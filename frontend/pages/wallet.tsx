/* eslint-disable react-hooks/exhaustive-deps */
import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { userInfoStore } from "../lib/store";
import Avatar from "react-avatar";
import { isPositiveInteger } from "../lib/utils";
import type { IModal, IStakeRequest } from "../lib/types";
import {
  callShowWallet,
  callStake,
  callUnStake,
  callShowTable,
  callFlushReward,
  callShowLil,
  callFlushUnstakePendingBalance,
  callFlushDonationReward,
  isApprovedForAll,
  guard,
} from "../lib/contract";
import {
  WalletWrapper,
  TextTitle,
  Text,
  WalletInfo,
  Info,
  Blocks,
  Block,
  Button,
  Hr,
  BlockInfo,
} from "../styles/wallet.styled";
import {
  ButtonGroup,
  ModalBackdrop,
  ModalView,
  InputField,
  ModalButton,
} from "../styles/modal.styled";

const Caver = require("caver-js");
const StakeStatusHandler = (status: string) => {
  switch (status) {
    case "0":
      return "Staked";
    case "1":
      return "Unstake Pending";
    case "2":
      return "Unstaked";
    case "3":
      return "Unexpected Error";
  }
};
const StakeStatusColorHandler = (status: string) => {
  switch (status) {
    case "0":
      return "#F5EDDC";
    case "1":
      return "#CFD2CF";
    case "2":
      return "#CFD2CF";
    case "3":
      return "#EB1D36";
  }
};

export default function Wallet() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const { myAccount, myNickname } = userInfoStore();
  const [myDonationReward, setMyDonationReward] = useState("0");
  const [myLil, setMyLil] = useState("0");
  const [myStakeBalance, setMyStakeBalance] = useState(0);
  const [modal, setModal] = useState<IModal>({
    title: "",
    description1: "",
    description2: "",
  });
  const [modalOpen, setModalOpen] = useState(false);
  const [callFunction, setCallFunction] = useState(0);
  const [amount, setAmount] = useState(0);
  const [lilAmount, setLilAmount] = useState(0);
  const [index, setIndex] = useState<number[]>([]);
  const [stakeRequest, setStakeRequest] = useState<IStakeRequest[]>([]);
  useEffect(() => {
    let loginNickname = sessionStorage.getItem("loginNickname");
    if (loginNickname) {
      myWallet();
      myTable();
    } else {
      console.log("error: Please use it after Login.");
      alert("Please use it after Login.");
      router.push("/");
    }
    setIsLoading(true);
  }, [myNickname, ]);
  const checkStakBalance = (stakeTable: IStakeRequest[]) => {
    let sumStakeBalance = 0
    stakeTable.forEach(function(value) {
      if(StakeStatusHandler(value.status) === "Staked") {
        sumStakeBalance += Number(value.amount);
      };
    })
    setMyStakeBalance(sumStakeBalance);
  };
  const myWallet = async () => {
    try {
      const caver = new Caver();
      const flag = await guard();
      if (flag) {
        const Wallet = await callShowWallet();
        const lil = await callShowLil();
        setMyDonationReward(`${caver.utils.fromPeb(Wallet.donationReward)}`);
        setMyLil(`${caver.utils.fromPeb(lil)}`);
      }
    } catch (error) {
      console.log("myWallet error\n", error);
    }
  };
  const myTable = async () => {
    try {
      const Table: IStakeRequest[] = await callShowTable();
      setStakeRequest(Table);
      checkStakBalance(Table);
    } catch (error) {
      console.log("myTable error\n", error);
    }
  };
  const stake = async () => {
    try {
      if (isPositiveInteger(amount)) {
        const receipt = await callStake(amount);
        console.log("stake", receipt);
        // alert(JSON.stringify(receipt, null, 2));
        alert("Success");
      }
    } catch (error) {
      console.log("stake error\n", error);
      alert(`Transaction Failed`);
    }
    setModalOpen(false);
    setCallFunction(0);
    setIndex([]);
    setAmount(0);
    myWallet();
    myTable();
  };
  const unStake = async () => {
    try {
      if (checkSelectBlock()) {
        const receipt = await callUnStake(index);
        console.log("unStake", receipt);
        alert("Success");
      }
    } catch (error) {
      console.log("unStake error\n", error);
      alert(`Transaction Failed`);
    }
    setIndex([]);
    myWallet();
    myTable();
  };
  const flushReward = async () => {
    try {
      if (checkSelectBlock()) {
        const caver = new Caver();
        const receipt = await callFlushReward(index);
        const lil = await callShowLil();
        setMyLil(`${caver.utils.fromPeb(lil)}`);
        alert(`Success`);
        console.log("flushReward", receipt);
      }
    } catch (error) {
      console.log("flushReward error\n", error);
      alert(`Transaction Failed`);
    }
    setIndex([]);
    myWallet();
    myTable();
  };
  const flushUnstakePendingBalance = async () => {
    try {
      if (checkSelectBlock()) {
        const caver = new Caver();
        const receipt = await callFlushUnstakePendingBalance(index);
        alert(`Success`);
        const Wallet = await callShowWallet();
        setMyDonationReward(`${caver.utils.fromPeb(Wallet.donationReward)}`);
        console.log("flushUnstakePending", receipt);
      }
    } catch (error) {
      console.log("flushUnstakePendingBalance error\n", error);
      alert(`Transaction Failed`);
    }
    setIndex([]);
    myWallet();
    myTable();
  };
  const flushDonationReward = async () => {
    try {
      const caver = new Caver();
      const receipt = await callFlushDonationReward(amount, lilAmount);
      const Wallet = await callShowWallet();
      setMyDonationReward(`${caver.utils.fromPeb(Wallet.donationReward)}`);
      const lil = await callShowLil();
      setMyLil(`${caver.utils.fromPeb(lil)}`);
      alert(`Success`);
      console.log("flushUnstakePending", receipt);
    } catch (error) {
      console.log("flushDonationReward error\n", error);
      alert(`Transaction Failed`);
    }
    setIndex([]);
    setModalOpen(false);
    setCallFunction(0);
    setAmount(0);
    setLilAmount(0);
    myWallet();
    myTable();
  };
  const checkSelectBlock = () => {
    try {
      if (index.length === 0) {
        alert("Please select your stake block");
  			throw new SyntaxError("Incomplete data: A number less than 1");
      }
      return true;
    } catch (error) {
        console.log("checkSelectBlock: ", error);
        return false;
    }
  };
  const changeHandler = (checked: boolean, id: number) => {
    if (checked) {
      setIndex((prev) => [...prev, id]);
    } else {
      setIndex(index.filter((el) => el !== id));
    }
  };
  return (
    <>
      {isLoading && myNickname && myAccount ? (
        <WalletWrapper>
          {modalOpen ? (
            <ModalBackdrop>
              <ModalView>
                <h3>{modal.title}</h3>
                <p>{modal.description1}</p>
                <InputField
                  type="text"
                  placeholder="amount"
                  onChange={(e) => {
                    setAmount(Number(e.target.value));
                  }}
                />
                {modal.description2 ? (
                  <>
                    <p>{modal.description2}</p>
                    <InputField
                      type="text"
                      placeholder="amount"
                      onChange={(e) => {
                        setLilAmount(Number(e.target.value));
                      }}
                    />
                  </>
                ) : null}
                <ButtonGroup>
                  <ModalButton onClick={() => setModalOpen(false)}>
                    Cancle
                  </ModalButton>
                  <ModalButton onClick={() => callFunction !== 0 ? (callFunction === 1 ? stake() : flushDonationReward()) : null}>Confirm</ModalButton>
                </ButtonGroup>
              </ModalView>
            </ModalBackdrop>
          ) : null}
          <WalletInfo>
            <Text>{myAccount}</Text>
            <TextTitle>
              <Avatar name="Lilly0" size={"50"} round={true} />
            </TextTitle>
            <TextTitle><b>Nickname</b></TextTitle>
            <Text>{myNickname}</Text>
            <TextTitle><b>Balance</b></TextTitle>
            <Text>Donation reward: {myDonationReward} KLAY</Text>
            <Text>Staked: {Caver.utils.fromPeb(myStakeBalance)} KLAY</Text>
            <Text>{myLil} LIL</Text>
          </WalletInfo>
          <Hr />
          <Info>
            <Button
              onClick={() => {
                setCallFunction(1);
                setModalOpen(true);
                setModal({
                  title: "New Stake",
                  description1: "KLAY Amount",
                  description2: ""
                });
              }}
              color="#A2B5BB"
            >
              new stake
            </Button>
            <Button
              onClick={() => {
                unStake();
              }}
              color="#A2B5BB"
            >
              unstake
            </Button>
            <Button
              onClick={() => {
                flushReward();
              }}
              color="#A2B5BB"
            >
              flush reward
            </Button>
            <Button
              onClick={() => {
                flushUnstakePendingBalance();
              }}
              color="#A2B5BB"
            >
              flush unstake pending balance
            </Button>
            <Button
              onClick={() => {
                setCallFunction(2);
                setModalOpen(true);
                setModal({
                  title: "Flush donation reward",
                  description1: "KLAY Amount",
                  description2: "LIL Amount",
                });
              }}
              color="#A2B5BB"
            >
              flush donation reward balance
            </Button>
          </Info>
          <Hr />
          <div>
            <Text>
              <p>Show History</p>
            </Text>
            <Blocks>
              {stakeRequest &&
                stakeRequest.map((item, idx) => {
                  return (
                    <Block
                      key={item.stakedBlock}
                      color={StakeStatusColorHandler(item.status)}
                    >
                      <input
                        type={"checkbox"}
                        disabled={item.status === "0" || "1" ? false : true}
                        onChange={(e) => {
                          changeHandler(e.currentTarget.checked, idx);
                        }}
                        checked={index.includes(idx) ? true : false}
                      />
                      <BlockInfo>{StakeStatusHandler(item.status)}</BlockInfo>
                      <BlockInfo>
                        {Caver.utils.fromPeb(item.amount)} KLAY
                      </BlockInfo>
                    </Block>
                  );
                })}
            </Blocks>
          </div>
        </WalletWrapper>
      ) : (
        <div>No Data</div>
      )}
    </>
  );
}
