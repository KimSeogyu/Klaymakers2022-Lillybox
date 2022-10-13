# Lillybox Smart Contract

<div align=center>
   <img width=200 height=200 src="https://upload.wikimedia.org/wikipedia/commons/9/98/Solidity_logo.svg" />
	<img width=200 height=200 src="https://cryptologos.cc/logos/klaytn-klay-logo.png?v=023" />
</div>

## Events

### WidthdrawKlay

```solidity
event WithdrawKlay(
        address indexed user,
        uint256 amount,
        uint256 usedLIL,
        uint256 blocknumber
    );
```

### NewStaking

```solidity
event NewStaking(
        address indexed user,
        uint256 amount,
        uint256 blocknumber,
        uint256 index
    );    
```

### UnStaking

```solidity
event UnStaking(address indexed user, uint256 amount, uint256 blocknumber);
```

### FlushReward

```solidity
event FlushReward(
        address indexed user,
        uint256 amount,
        uint256 blocknumber
    );
```

### NewDonation

```solidity
event NewDonation(
        address indexed from,
        address indexed to,
        uint256 amount,
        uint256 usedLIL,
        uint256 blocknumber
    );
```

## Data Structure

### StakeStatus

- Type : Enum
- Desc : Define the status of each stake request.

| Name           | Description                             |
| -------------- | --------------------------------------- |
| Staked         | Status of staked request                |
| UnstakePending | Status of unstake request pending       |
| Unstaked       | Status of unstaked block                |
| Failed         | Status that occurs something went wrong |


### StakeRequest

- Type : Struct
- Desc : Struct that includes information about each stake request.

| Name                | Type          | Description                                    |
| ------------------- | ------------- | ---------------------------------------------- |
| amount              | `uint256`     | Amount of user staked                          |
| stakedBlock         | `uint256`     | Block number of user staked                    |
| unstakeRequestBlock | `uint256`     | Block number of when user made unstake request |
| status              | `StakeStatus` | Define the status of each stake request        |


### StakeWallet 

- Type : Struct
- Desc : Struct that includes `StakeRequest` and `mutex lock status`.

| Name    | Type             | Description             |
| ------- | ---------------- | ----------------------- |
| request | `StakeRequest[]` | Array of `StakeRequest` |
| lock    | `bool`           | Status of mutex lock    |


### Wallet

- Type : Struct 
- Desc : Struct that includes user's donation reward balance and status of mutex lock.

| Name           | Type      | Description            |
| -------------- | --------- | ---------------------- |
| donationReward | `uint256` | KLAY balance deposited |
| lock           | `bool`    | Status of mutex lock   |

### Table

- Type : Mapping
- Desc : Table mapped with address and StakeWallet

### Wallets

- Type : Mapping
- Desc : Table mapped with address and Wallet

### UserNicknames

- Type : Mapping
- Desc : Table mapped with string and address


## Functions

### showTable

Query the history of all the user's steak/unstake requests.

```solidity
function showTable()
        public
        view
        isApproved
    returns (StakeRequest[] memory)
    {
        return Table[msg.sender].request;
    }
```

### stake


Stake as much as `msg.value`. The steak request is added to `Table[msg.sender].request`

Mint LKAY in proportion to the amount of KLAY and generate a New Staking event.

```solidity
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

```

### unstake

The index list is received as a parameter. Change all the stake requests corresponding to the index to an unstake pending state.

```solidity
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
```

### flushUnstakePendingBalance

Among the requests in the Unstake pending state, unstake work is performed only for requests that have passed 21 days.

And the amount of the request is remitted to the user.

```solidity
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
```

### showWallet

Wallet Information Inquiry

```solidity
unction showWallet(address user)
        public
        view
        walletOwnerOrAdmin(user)
        isApproved
        returns (Wallet memory)
    {
        return Wallets[user];
    }

```

### mintVod

Mint video NFT

```solidity
function mintVod(string memory tokenURI)
        public
        payable
        isApproved
        returns (uint256)
    {
        uint256 id = mintNFT(tokenURI);
        return id;
    }
```

### mintNFT

Mint NFT

```solidity
function mintNFT(string memory tokenURI) private returns (uint256) {
        _tokenIds.increment();
        uint256 newTokenId = _tokenIds.current();
        _mint(msg.sender, newTokenId, 1, "");
        _setURI(newTokenId, tokenURI);
        setApprovalForAll(address(this), true);

        return newTokenId;
    }
```

### flushDonationReward

Transfer money to the user's wallet as much as user`s donation revenue.

```solidity
function flushDonationReward(uint256 _amount, uint256 _lilAmount)
        public
        payable
        walletMutex
        isApproved
        returns (bool)
    {
        require(balanceOf(msg.sender, LIL) >= _lilAmount, "Not enough LIL");

        require(_lilAmount <= 1000, "Max amount is 1000 LIL");

        require(
            Wallets[msg.sender].donationReward >= _amount,
            "Not enough deposit"
        );

        uint256 _fee = withdrawFee(_amount, _lilAmount);
        Wallets[msg.sender].donationReward -= _amount;
        FeeRevenues += _fee;
        // Send Klay
        payable(msg.sender).transfer(_amount - _fee);
        _burn(msg.sender, LIL, _lilAmount);

        emit WithdrawKlay(msg.sender, _amount, _lilAmount, block.number);

        return true;
    }
```

### flushReward 

Receive compensation from Staked request. Rewards are given in LIL tokens.

```solidity
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
```

### donation 

Send as much money as msg.value sent by msg.sender to donationReward in destUser. 

At this time, fee will be reduced by the LIL entered by the msg.sender.

```solidity
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

        Wallets[destUser].donationReward += msg.value - _fee;

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
```
