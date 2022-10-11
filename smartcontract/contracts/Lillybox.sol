// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.7;

import "@klaytn/contracts/KIP/token/KIP37/extensions/KIP37URIStorage.sol";
import "@klaytn/contracts/access/Ownable.sol";
import "@klaytn/contracts/utils/Counters.sol";

// 릴리박스 서비스 컨트랙트
contract Lillybox is KIP37URIStorage, Ownable {
    // 클레이 출금 이벤트
    event WithdrawKlay(
        address indexed user,
        uint256 amount,
        uint256 usedLIL,
        uint256 blocknumber
    );
    // 새로운 스테이킹 추가 이벤트
    event NewStaking(
        address indexed user,
        uint256 amount,
        uint256 blocknumber,
        uint256 index
    );
    // 스테이킹 해제 이벤트
    event UnStaking(address indexed user, uint256 amount, uint256 blocknumber);
    // 리워드 회수 이벤트
    event FlushReward(
        address indexed user,
        uint256 amount,
        uint256 blocknumber
    );
    // 새로운 도네이션 발생 이벤트
    event NewDonation(
        address indexed from,
        address indexed to,
        uint256 amount,
        uint256 usedLIL,
        uint256 blocknumber
    );

    // 전체 NFT를 인덱싱할 토큰 ID 카운터
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    // 거버넌스 토큰
    uint256 public constant LIL = 0;

    // 예치금액을 스테이킹으로 전환했을 때 교환해주는 토큰, 일종의 sKLAY
    uint256 public constant LKLAY = 1;

    // LIL decimal
    uint256 public constant LIL_DECIMALS = 10**18;

    // 광고 10회 노출당 비용
    uint256 public AdsCostPer10Times;

    // 블록당 보상 수량
    uint256 public RewardPerBlock;

    // 광고 수익 총계
    uint256 private AdsRevenues;

    // 수수료 수익 총계
    uint256 private FeeRevenues;

    // 언스테이킹 처리 소요 시간
    uint256 private UnstakeProcessTime;

    // 지갑 정보
    struct Wallet {
        // 예치되어 있는 KLAY 잔고
        uint256 klayBalance;
        // 동시성처리를 위한 Mutex lock
        bool lock;
    }

    enum StakeStatus {
        // 스테이킹 완료
        Staked,
        // 언스테이킹 진행
        UnstakePending,
        // 언스테이킹 완료
        Unstaked,
        // Unexpected Error
        Failed
    }

    struct StakeRequest {
        // 수량
        uint256 amount;
        // 스테이킹 신청한 블록
        uint256 stakedBlock;
        // 언스테이킹 신청한 블록
        uint256 unstakeRequestBlock;
        // 스테이킹 상태
        StakeStatus status;
    }

    struct StakeWallet {
        // 유저의 스테이킹/언스테이킹 신청 배열
        StakeRequest[] request;
        // 뮤텍스 락
        bool lock;
    }

    mapping(address => StakeWallet) Table;

    // 주소로 매핑한 지갑 정보
    mapping(address => Wallet) Wallets;

    // 닉네임으로 매핑한 주소 정보
    mapping(string => address) private UserNicknames;

    constructor(
        // 광고 10회 노출 비용: 5KLAY
        uint256 _adsCostPer10Times,
        // 블록당 리워드 수량: 0.0001 LIL
        uint256 _rewardPerBlock,
        // 언스테이킹 처리 소요 시간
        uint256 _unstakeProcessTime
    ) KIP37("https://lillybox.xyz/metadata/{id}.json") {
        // 광고 10회 노출 비용: 5KLAY
        AdsCostPer10Times = _adsCostPer10Times;
        // 블록당 리워드 수량: 0.0001 LIL
        RewardPerBlock = _rewardPerBlock;
        // 언스테이킹 처리 소요 시간
        UnstakeProcessTime = _unstakeProcessTime;

        AdsRevenues = 0;
        FeeRevenues = 0;
        _tokenIds.increment();
        _tokenIds.increment();
        _mint(msg.sender, LIL, 21000000 * LIL_DECIMALS, "");
    }

    // Transaction Lock 구현체
    modifier walletMutex() {
        // 호춣한 계정의 지갑이 unlock 상태인지 확인
        require(Wallets[msg.sender].lock == false);
        // 지갑을 lock 상태로 업데이트
        Wallets[msg.sender].lock = true;

        _;

        // 지갑을 unlock 상태로 업데이트
        Wallets[msg.sender].lock = false;
    }

    // 컨트랙트가 호출자에게 승인받았는지 확인
    modifier isApproved() {
        require(
            isApprovedForAll(msg.sender, address(this)) == true,
            "Caller have to approve this contract address managing asset"
        );

        _;
    }

    // 지갑 소유자 또는 컨트랙트 관리자인지 확인
    modifier walletOwnerOrAdmin(address user) {
        require(
            msg.sender == user || msg.sender == owner(),
            "Only owner of contract or owner of wallet can see wallet information"
        );

        _;
    }

    // 최소 금액 확인, 에러메시지 기본 값
    modifier minValueCheck(uint256 _minAmount) {
        require(
            msg.value >= _minAmount,
            string(
                abi.encodePacked(
                    "msg.value should be higher than",
                    Strings.toString(_minAmount)
                )
            )
        );

        _;
    }

    // Transaction Lock 구현체
    modifier StakeMutex() {
        require(Table[msg.sender].lock == false);

        Table[msg.sender].lock = true;

        _;

        Table[msg.sender].lock = false;
    }

    // 유저의 모든 요청 기록 조회
    function showTable()
        public
        view
        isApproved
        returns (StakeRequest[] memory)
    {
        return Table[msg.sender].request;
    }

    // Stake
    function stake() public payable StakeMutex isApproved returns (bool) {
        // mapping 테이블에 넣어주기
        Table[msg.sender].request.push(
            StakeRequest(msg.value, block.number, 0, StakeStatus.Staked)
        );
        // LKLAY mint
        _mint(msg.sender, LKLAY, msg.value, "");
        // 이벤트
        emit NewStaking(
            msg.sender,
            msg.value,
            block.number,
            Table[msg.sender].request.length - 1
        );
        return true;
    }

    function unstake(uint256[] memory index)
        public
        StakeMutex
        isApproved
        returns (bool)
    {
        for (uint256 i = 0; i < index.length; i++) {
            // Index가 올바른지 확인
            require(
                Table[msg.sender].request.length > index[i],
                "You accessed an element that does not exist"
            );
            // 해당 Request가 Staked 된 상태인지 확인
            require(
                Table[msg.sender].request[index[i]].status ==
                    StakeStatus.Staked,
                "You can unstake transaction only you have already staked"
            );
            // unstakeRequestBlock을 현재 블록으로 등록
            Table[msg.sender].request[index[i]].unstakeRequestBlock = block
                .number;
            // UnstakePending 상태로 변경
            Table[msg.sender].request[index[i]].status = StakeStatus
                .UnstakePending;
        }
        return true;
    }

    function flushUnstakePendingBalance(uint256[] memory index)
        public
        payable
        StakeMutex
        isApproved
        returns (bool)
    {
        for (uint256 i = 0; i < index.length; i++) {
            // Index가 올바른지 확인
            require(
                Table[msg.sender].request.length > index[i],
                "You accessed an element that does not exist"
            );
            // 해당 Request가 UnstakePending 상태인지 확인
            require(
                Table[msg.sender].request[index[i]].status ==
                    StakeStatus.UnstakePending,
                "You have to make a request unstake first!"
            );
            // 21일이 지났는지 확인
            require(
                Table[msg.sender].request[index[i]].unstakeRequestBlock +
                    UnstakeProcessTime <
                    block.number,
                "Unstaking process usually takes 21 days"
            );
            // burn
            _burn(
                msg.sender,
                LKLAY,
                Table[msg.sender].request[index[i]].amount
            );
            // 해당 Request는 Unstaked 된 상태로 변경해준다.
            Table[msg.sender].request[index[i]].status = StakeStatus.Unstaked;
            // 유저 지갑에 해당 amount 만큼 더해준다.
            payable(msg.sender).transfer(
                Table[msg.sender].request[index[i]].amount
            );
        }
        return true;
    }

    // 수수료 수익 조회 (컨트랙트 주인만 볼 수 있음)
    function showRevenuesFromFee() public view onlyOwner returns (uint256) {
        return FeeRevenues;
    }

    // 광고 수익 조회 (컨트랙트 주인만 볼 수 있음)
    function showRevenuesFromAds() public view onlyOwner returns (uint256) {
        return AdsRevenues;
    }

    // 지갑 정보 조회
    function showWallet(address user)
        public
        view
        walletOwnerOrAdmin(user)
        isApproved
        returns (Wallet memory)
    {
        return Wallets[user];
    }

    // 비디오 NFT 민팅
    function mintVod(string memory tokenURI)
        public
        payable
        isApproved
        returns (uint256)
    {
        uint256 id = mintNFT(tokenURI);
        return id;
    }

    // 최소 광고 비용 조회
    function minimumAdsCost() public view isApproved returns (uint256) {
        return AdsCostPer10Times * 100;
    }

    // 광고 NFT 민팅
    function mintAd(string memory tokenURI)
        public
        payable
        minValueCheck(minimumAdsCost())
        isApproved
        returns (uint256)
    {
        uint256 id = mintNFT(tokenURI);
        AdsRevenues += msg.value;
        return id;
    }

    // NFT 민팅
    function mintNFT(string memory tokenURI) private returns (uint256) {
        _tokenIds.increment();
        uint256 newTokenId = _tokenIds.current();
        _mint(msg.sender, newTokenId, 1, "");
        _setURI(newTokenId, tokenURI);
        setApprovalForAll(address(this), true);

        return newTokenId;
    }

    function vacuumAd(address _owner, uint256 _tokenId)
        public
        onlyOwner
        returns (bool)
    {
        _burn(_owner, _tokenId, 1);

        return true;
    }

    function _burn(
        address from,
        uint256 id,
        uint256 amount
    ) internal override walletOwnerOrAdmin(msg.sender) {
        super._burn(from, id, amount);
    }

    function flushKlayBalance(uint256 _amount, uint256 _lilAmount)
        public
        payable
        walletMutex
        isApproved
        returns (bool)
    {
        require(balanceOf(msg.sender, LIL) >= _lilAmount, "Not enough LIL");

        require(_lilAmount <= 1000, "Max amount is 1000 LIL");

        require(
            Wallets[msg.sender].klayBalance >= _amount,
            "Not enough deposit"
        );

        uint256 _fee = withdrawFee(_amount, _lilAmount);
        Wallets[msg.sender].klayBalance -= _amount;
        FeeRevenues += _fee;
        // Send Klay
        payable(msg.sender).transfer(_amount - _fee);
        _burn(msg.sender, LIL, _lilAmount);

        emit WithdrawKlay(msg.sender, _amount, _lilAmount, block.number);

        return true;
    }

    function withdrawFee(uint256 v, uint256 _lilAmount)
        private
        pure
        returns (uint256)
    {
        // 1 LIL = -0.01%
        // 10 LIL = -0.1%
        // 100 LIL = -1%
        // 1000 LIL = -10%
        // withdraw fee = 10%
        return (v / 10000) * (1000 - _lilAmount);
    }

    function donationFee(uint256 v, uint256 _lilAmount)
        private
        pure
        returns (uint256)
    {
        // 1 LIL = -0.01%
        // 10 LIL = -0.1%
        // 100 LIL = -1%
        // 1000 LIL = -10%
        // 2500 LIL = -25%
        // donation fee = 25%
        return (v / 10000) * (2500 - _lilAmount);
    }

    function flushReward(uint256[] memory index)
        public
        isApproved
        StakeMutex
        returns (uint256)
    {
        uint256 _total = 0;

        for (uint256 i = 0; i < index.length; i++) {
            require(
                Table[msg.sender].request[index[i]].status ==
                    StakeStatus.Staked,
                "You can only get reward from staked block"
            );

            require(
                Table[msg.sender].request[index[i]].stakedBlock != 0,
                "Your staked block is missing"
            );

            _total += estimateRewards(msg.sender, index[i]);
            Table[msg.sender].request[i].stakedBlock = block.number;
        }
        _mint(msg.sender, LIL, _total, "");

        emit FlushReward(msg.sender, _total, block.number);
        return _total;
    }

    function estimateRewards(address user, uint256 index)
        public
        view
        walletOwnerOrAdmin(user)
        returns (uint256)
    {
        if (
            Table[user].request[index].stakedBlock == 0 ||
            Table[user].request[index].status != StakeStatus.Staked
        ) {
            return 0;
        }

        uint256 passedBlocks = block.number -
            Table[user].request[index].stakedBlock;
        return balanceOf(user, LKLAY) * passedBlocks * RewardPerBlock;
    }

    function donation(address destUser, uint256 _lilAmount)
        public
        payable
        walletMutex
        isApproved
        returns (bool)
    {
        require(balanceOf(msg.sender, LIL) >= _lilAmount, "Not enough LIL");

        require(_lilAmount <= 2500, "Max amount is 2500 LIL");

        uint256 _fee = donationFee(msg.value, _lilAmount);

        Wallets[destUser].klayBalance += msg.value - _fee;

        if (_lilAmount > 0) {
            _burn(msg.sender, LIL, _lilAmount);
        }
        FeeRevenues += _fee;

        emit NewDonation(
            msg.sender,
            destUser,
            msg.value,
            _lilAmount,
            block.number
        );

        return true;
    }

    function withdrawRevenue() public onlyOwner {
        uint256 _amount = FeeRevenues + AdsRevenues;
        FeeRevenues = 0;
        AdsRevenues = 0;
        payable(owner()).transfer(_amount);
    }
}
