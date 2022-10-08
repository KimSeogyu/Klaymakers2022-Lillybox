// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.7;

import "@klaytn/contracts/KIP/token/KIP37/extensions/KIP37URIStorage.sol";
import "@klaytn/contracts/access/Ownable.sol";
import "@klaytn/contracts/utils/Counters.sol";

// 릴리박스 서비스 컨트랙트
contract Lillybox is KIP37URIStorage, Ownable {
    // 클레이 입금 이벤트
    event DepositKlay(address indexed user, uint256 amount, uint256 blocknumber);
    // 클레이 출금 이벤트
    event WithdrawKlay(address indexed user, uint256 amount, uint256 usedLIL, uint256 blocknumber);
    // 새로운 스테이킹 추가 이벤트
    event NewStaking(address indexed user, uint256 amount, uint256 blocknumber);
    // 스테이킹 해제 이벤트
    event UnStaking(address indexed user, uint256 amount, uint256 blocknumber);
    // 리워드 회수 이벤트
    event FlushReward(address indexed user, uint256 amount, uint256 blocknumber);
    // 새로운 도네이션 발생 이벤트
    event NewDonation(address indexed from, address indexed to, uint256 amount, uint256 usedLIL, uint256 blocknumber);

    // 전체 NFT를 인덱싱할 토큰 ID 카운터
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    // 거버넌스 토큰
    uint256 public constant LIL = 0;

    // 예치금액을 스테이킹으로 전환했을 때 교환해주는 토큰, 일종의 sKLAY
    uint256 public constant LKLAY = 1;

    // LIL decimal
    uint256 public constant LIL_DECIMALS = 10 ** 18;

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
        // 스테이킹 되어있는 KLAY 잔고
        uint256 stakedBalance;
        // 스테이킹 시작한 블록
        uint256 stakedBlock;
        // 스테이킹 해제 신청한 블록
        uint256 unstakeRequestBlock;
        // 스테이킹 해제 대기중인 수량
        uint256 unstakePendingBalance;

        // 동시성처리를 위한 Mutex lock
        bool lock;
    }

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
    ) KIP37("https://lillybox.xyz/metadata/{id}.json")
    {
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
    modifier walletMutex()
    {
        // 호춣한 계정의 지갑이 unlock 상태인지 확인
        require(Wallets[msg.sender].lock == false);
        // 지갑을 lock 상태로 업데이트
        Wallets[msg.sender].lock = true;

        _;

        // 지갑을 unlock 상태로 업데이트
        Wallets[msg.sender].lock = false;
    }

    // 컨트랙트가 호출자에게 승인받았는지 확인
    modifier isApproved()
    {
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
            msg.value > _minAmount,
            string(abi.encodePacked(
                "msg.value should be higher than",
                Strings.toString(_minAmount)
            ))
        );

        _;
    }

    // 수수료 수익 조회 (컨트랙트 주인만 볼 수 있음)
    function showRevenuesFromFee()
    public
    view
    onlyOwner
    returns (uint256)
    {
        return FeeRevenues;
    }

    // 광고 수익 조회 (컨트랙트 주인만 볼 수 있음)
    function showRevenuesFromAds()
    public
    view
    onlyOwner
    returns (uint256)
    {
        return AdsRevenues;
    }

    // 지갑 정보 조회
    function showWallet(address user)
    public
    view
    walletOwnerOrAdmin(user)
    returns (Wallet memory)
    {
        return Wallets[user];
    }

    // 비디오 NFT 민팅
    function mintVod(string memory tokenURI)
    public
    payable
    returns (uint256)
    {
        uint256 id = mintNFT(tokenURI);
        return id;
    }

    // 최소 광고 비용 조회
    function minimumAdsCost()
    public
    view
    returns (uint256)
    {
        return AdsCostPer10Times * 100;
    }


    // 광고 NFT 민팅
    function mintAd(string memory tokenURI)
    public
    payable
        // 호출한 계정은 반드시 최소 광고 비용 이상을 보내야 한다
    minValueCheck(minimumAdsCost())
    returns (uint256)
    {
        uint256 id = mintNFT(tokenURI);
        AdsRevenues += msg.value;
        return id;
    }

    // NFT 민팅
    function mintNFT(string memory tokenURI)
    private
    returns (uint256)
    {
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

        return false;
    }

    function _burn(
        address from,
        uint256 id,
        uint256 amount
    )
    internal
    override
    onlyOwner {
        super._burn(from, id, amount);
    }

    function deposit()
    public
    payable
    walletMutex
    minValueCheck(10 * LIL_DECIMALS)
    returns (bool)
    {
        // KLAY 예치금액에 더 해준다
        // 입금은 수수료가 없음
        Wallets[msg.sender].klayBalance += msg.value;

        emit DepositKlay(msg.sender, msg.value, block.number);

        return true;
    }

    function withdraw(uint256 _amount, uint256 _lilAmount)
    public
    walletMutex
    returns (bool)
    {
        require(
            balanceOf(msg.sender, LIL) > _lilAmount,
            "Not enough LIL"
        );

        require(
            _lilAmount <= 1000,
            "Max amount is 1000 LIL"
        );

        require(
            Wallets[msg.sender].klayBalance >= _amount,
            "Not enough deposit"
        );

        uint256 _fee = withdrawFee(_amount, _lilAmount);
        Wallets[msg.sender].klayBalance = Wallets[msg.sender].klayBalance - _amount;
        FeeRevenues += _fee;
        // Send Klay
        payable(msg.sender).transfer(_amount - _fee);
        _burn(msg.sender, LIL, _lilAmount);

        emit WithdrawKlay(msg.sender, _amount, _lilAmount, block.number);

        return true;
    }

    function fee(uint256 v)
    private
    pure
    returns (uint256)
    {
        // 0.05%
        return v / 2000;
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
        return v / 10000 * (1000 - _lilAmount);
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
        return v / 10000 * (2500 - _lilAmount);
    }

    function flushReward()
    public
    isApproved
    walletMutex
    returns (uint256)
    {
        uint256 _reward = getRewards(msg.sender);
        require(_reward > 0, "You have no rewards yet!");
        Wallets[msg.sender].stakedBlock = block.number;
        _mint(msg.sender, LIL, _reward, "");
        emit FlushReward(msg.sender, _reward, block.number);
        return _reward;
    }

    function getRewards(address user)
    public
    view
    walletOwnerOrAdmin(user)
    returns (uint256)
    {
        if (Wallets[user].stakedBlock == 0) {
            return 0;
        }
        uint256 passedBlocks = block.number - Wallets[user].stakedBlock;
        // LKLAY 수량 * 스테이킹을 유지한 블록 수 * 블록 당 리워드 수량
        return balanceOf(user, LKLAY) * passedBlocks * RewardPerBlock;
    }

    function donation(string memory toNickname, uint256 _donationAmount, uint256 _lilAmount)
    public
    walletMutex
    isApproved
    returns (bool)
    {
        require(
            balanceOf(msg.sender, LIL) > _lilAmount,
            "Not enough LIL"
        );

        require(
            _lilAmount <= 2500,
            "Max amount is 2500 LIL"
        );

        require(
            Wallets[msg.sender].klayBalance < _donationAmount,
            "Not enough KLAY"
        );

        if (isApprovedForAll(msg.sender, address(this)) == false) {
            setApprovalForAll(address(this), true);
        }

        Wallets[msg.sender].klayBalance -= _donationAmount;

        uint256 _fee = donationFee(_donationAmount, _lilAmount);

        address destUser = UserNicknames[toNickname];
        Wallets[destUser].klayBalance += _donationAmount - _fee;

        if (_lilAmount > 0) {
            _burn(msg.sender, LIL, _lilAmount);
        }
        FeeRevenues += _fee;

        emit NewDonation(msg.sender, destUser, _donationAmount, _lilAmount, block.number);

        return true;
    }

    function withdrawRevenue()
    public
    onlyOwner
    {
        uint256 _amount = FeeRevenues + AdsRevenues;
        FeeRevenues = 0;
        AdsRevenues = 0;
        payable(owner()).transfer(_amount);
    }

    function stake(uint256 _stakeAmount)
    public
    walletMutex
    isApproved
    returns (uint256)
    {
        require(
            Wallets[msg.sender].klayBalance < _stakeAmount,
            "Not enough KLAY in your Wallet"
        );

        Wallets[msg.sender].klayBalance -= _stakeAmount;
        Wallets[msg.sender].stakedBlock = block.number;
        _mint(msg.sender, LKLAY, _stakeAmount, "");

        emit NewStaking(msg.sender, _stakeAmount, block.number);

        return _stakeAmount;
    }

    function unstake(uint256 _amount)
    public
    walletMutex
    isApproved
    returns (bool)
    {
        require(
            balanceOf(msg.sender, LKLAY) < _amount,
            "Not enough LKLAY in your Wallet"
        );

        // 언스테이킹 신청한 적이 없는 경우 지금 신청했다고 친다
        if (Wallets[msg.sender].unstakeRequestBlock == 0) {
            Wallets[msg.sender].unstakeRequestBlock = block.number;
            return true;
        }

        require(
            Wallets[msg.sender].unstakeRequestBlock + UnstakeProcessTime < block.number,
            "Unstaking process usually takes 21 days"
        );

        Wallets[msg.sender].unstakeRequestBlock = 0;
        _burn(msg.sender, LKLAY, _amount);

        emit UnStaking(msg.sender, _amount, block.number);

        return true;
    }
}
