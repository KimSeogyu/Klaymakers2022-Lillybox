/* eslint-disable react-hooks/exhaustive-deps */
import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";
import { userInfoStore } from "../lib/store";
import Avatar from "react-avatar";
import { isInteger } from "../lib/utils";
import type { IModal, IStakeRequest } from "../lib/types";
import {
  callShowWallet,
  callStake,
  callUnStake,
  callShowTable,
  callFlushReward,
  callShowLil,
  callFlushUnstakePendingBalance,
  callFlushKlayBalance,
} from "../lib/contract";
import {
  WalletWrapper,
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
  const [myKlayBalance, setmyKlayBalance] = useState("");
  const [myLil, setMyLil] = useState("");
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

  const myWallet = async () => {
    try {
      const caver = new Caver();
      const Wallet = await callShowWallet();
      const lil = await callShowLil();
      setmyKlayBalance(`${caver.utils.fromPeb(Wallet.klayBalance)}`);
      setMyLil(`${caver.utils.fromPeb(lil)}`);
    } catch (error) {
      console.log("myWallet error\n", error);
    }
  };
  const myTable = async () => {
    try {
      const Table: IStakeRequest[] = await callShowTable();
      setStakeRequest(Table);
    } catch (error) {
      console.log("myTable error\n", error);
    }
  };
  const stake = async () => {
    try {
      if (!(amount > 0)) {
        alert("Please enter a number greater than 0.");
        throw new SyntaxError("Incomplete data: A number less than 1");
      }
      if (!isInteger(amount)) {
        alert("Please enter an integer");
        throw new SyntaxError("Incomplete data: Not an integer");
      }
      const receipt = await callStake(amount);
      console.log("stake", receipt);
      // alert(JSON.stringify(receipt, null, 2));
      alert("Success");
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
      if (!checkSelectBlock()) {
        throw new SyntaxError("Incomplete data: no select");
      }
      const receipt = await callUnStake(index);
      console.log("unStake", receipt);
      alert("Success");
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
      if (!checkSelectBlock()) {
        throw new SyntaxError("Incomplete data: no select");
      }
      const caver = new Caver();
      const receipt = await callFlushReward(index);
      const lil = await callShowLil();
      setMyLil(`${caver.utils.fromPeb(lil)}`);
      alert(`Success`);
      console.log("flushReward", receipt);
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
      if (!checkSelectBlock()) {
        throw new SyntaxError("Incomplete data: no select");
      }
      const caver = new Caver();
      const receipt = await callFlushUnstakePendingBalance(index);
      alert(`Success`);
      const Wallet = await callShowWallet();
      setmyKlayBalance(`${caver.utils.fromPeb(Wallet.klayBalance)}`);
      console.log("flushUnstakePending", receipt);
    } catch (error) {
      console.log("flushUnstakePendingBalance error\n", error);
      alert(`Transaction Failed`);
    }
    setIndex([]);
    myWallet();
    myTable();
  };
  const flushKlayBalance = async () => {
    try {
      const caver = new Caver();
      const receipt = await callFlushKlayBalance(amount, lilAmount);
      alert(`Success`);
      const Wallet = await callShowWallet();
      setmyKlayBalance(`${caver.utils.fromPeb(Wallet.klayBalance)}`);
      const lil = await callShowLil();
      setMyLil(`${caver.utils.fromPeb(lil)}`);
      console.log("flushUnstakePending", receipt);
    } catch (error) {
      console.log("flushKlayBalance error\n", error);
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
    if (index.length === 0) {
      alert("Please select your stake block");
      return false;
    }
    return true;
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
                  <ModalButton onClick={() => callFunction !== 0 ? (callFunction === 1 ? stake() : flushKlayBalance()) : null}>Confirm</ModalButton>
                </ButtonGroup>
              </ModalView>
            </ModalBackdrop>
          ) : null}
          <WalletInfo>
            <Text>{myAccount}</Text>
            <Text>
              <Avatar name="Lilly0" size={"50"} round={true} />
            </Text>
            <Text>Nickname: {myNickname}</Text>
            <Text>{myKlayBalance} KLAY</Text>
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
                  title: "Flush KLAY Balance",
                  description1: "KLAY Amount",
                  description2: "LIL Amount",
                });
              }}
              color="#A2B5BB"
            >
              flush klay balance
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
