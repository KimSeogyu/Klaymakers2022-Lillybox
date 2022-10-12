/* eslint-disable react-hooks/exhaustive-deps */
import {
  BsBell,
  BsCameraVideo,
  BsMicFill,
  BsSearch,
} from "react-icons/bs";
import Avatar from "react-avatar";
import ReactTooltip from "react-tooltip";
import { useEffect, useState } from "react";
import { useRouter } from "next/router";
import Link from "next/link";
import Image from "next/image";
import Caver from "caver-js";
import AuthAPI from "../lib/auth";
import { userInfoStore } from "../lib/store";
import { callShowWallet, setApprovalForAll } from "../lib/contract";
import type { IUserInfo } from "../lib/types";
import { ThemeSwitch } from "../lib/utils";

function Header({ toggleSidebar, isMounted }: any) {
  const [InputValue, setInputValue] = useState("");
  const [notification, setNotification] = useState(2);
  const [connectedAccount, setConnectedAccount] = useState("");
  const [userNickname, setUserNickname] = useState("");
  const { setUserInfo } = userInfoStore();
  const router = useRouter();
  useEffect(() => {
    setNotification(Math.floor(Math.random() * 10));
    let loginNickname = sessionStorage.getItem("loginNickname");
    if (loginNickname) {
      AccountHandler();
    }
    window.klaytn.on("accountsChanged", (newAccount: string) => {
      setConnectedAccount(newAccount.toString());
      UserInfoHandler(newAccount.toString());
    });
  }, []);
  const AccountHandler = async () => {
    if (window.klaytn) {
      const [account] = await window.klaytn.enable();
      setConnectedAccount(account);
      UserInfoHandler(account.toString());
    }
  }
  const UserInfoHandler = async (account:string) => {
    const userInfo:IUserInfo = await AuthAPI.getUserInfo(account);
    if (userInfo === null) {
      LoginHandler();
      return ;
    }
    console.log()
    setUserInfo(account.toString(), userInfo.nickname.toString())
    setUserNickname(userInfo.nickname.toString());
    sessionStorage.setItem("loginNickname", userInfo.nickname.toString());
  }
  const LoginHandler = async () => {
    if (window.klaytn === undefined) {
      alert(
        "Please install kaikas!\nURL: https://chrome.google.com/webstore/detail/kaikas/jblndlipeogpafnldhgmapagcccfchpi"
      );
    } else {
      try {
        const [account] = await window.klaytn.enable();
        const caver = new Caver(window.klaytn);
        let message = await AuthAPI.getLoginRequestID(account.toString());
        while (message === null) {
          const nickname: any = prompt("Please enter your nickname")
          if (nickname === null)
            break ;
          const flag: boolean = await AuthAPI.checkNickname(nickname);
          if (!!!flag) {
            alert("Please enter valid nickname");
            return;
          }
          const result = await AuthAPI.insertUser(account.toString(), nickname);
          if (result === null) {
            alert("Already exists")
            continue;
          }
          message = result;
		  await setApprovalForAll();
        }
        if (message === null)
         return ;
        console.log("account, message : ", account, message);
        const signResult = await caver.rpc.klay.sign(account, message);
        const v = "0x" + signResult.substring(2).substring(128, 130);
        const r = "0x" + signResult.substring(2).substring(0, 64);
        const s = "0x" + signResult.substring(2).substring(64, 128);
        const signature = [v, r, s];
        const validate = await caver.validator.validateSignedMessage(
          message,
          signature,
          account
        );
        setConnectedAccount(account);
        UserInfoHandler(account.toString());
        console.log(`validate signed message: ${validate}`);
      } catch (error) {
        console.log("LoginHandler Error",error);
      }
    }
  };
  const showWallet = async () => {
    try {
      const caver = new Caver();
      const Wallet = await callShowWallet();
      console.log(Wallet);
      alert(`
      Klay Balance : ${caver.utils.fromPeb(Wallet.klayBalance)}
      `);
    } catch (error) {
      console.log("showWallet Error",error);
    }
  };;
  return (
      <div className="bg-white dark:bg-black flex items-center justify-between border-b-2 fixed top-0 left-0 right-0 z-20">
        {/* tailwind css hamburger menu */}
        <div
          className="flex cursor-pointer justify-between h-[1.2rem] w-[1.5rem] flex-col ml-2 group md:ml-7"
          onClick={() => {
            toggleSidebar();
          }}
        >
          <div className="border-r-1 h-[2px] bg-gray-700 group-hover:bg-gray-500"></div>
          <div className="border-r-1 h-[2px] bg-gray-700 group-hover:bg-gray-500"></div>
          <div className="border-r-1 h-[2px] bg-gray-700 group-hover:bg-gray-500"></div>
        </div>

        <div className="flex items-center justify-center relative ml-1 pt-1 hover:text-gray-800">
          <Link href="/">
            <Image
              src="/lilly-plain.png"
              width={80}
              height={50}
              className="text-4xl text-red-600 pr-1"
              alt=""
            />
          </Link>
        </div>
        {/* Query field */}
        <div className="flex items-center justify-center focus:outline-none focus:bg-gray-300">
          <div className="flex items-center justify-center bg-gray-200 rounded-lg mr-2 my-2">
            <input
              type="text"
              placeholder="Search"
              className="bg-gray-200 rounded-lg px-4 py-2 w-20 text-gray-800 md:w-[20em] focus:outline-none lg:w-[35em]"
              value={InputValue}
              onChange={(e) => {
                setInputValue(e.target.value);
              }}
              onKeyDown={(e) => {
                if (e.key === "Enter") {
                  router.push(`/?query=${InputValue.toLowerCase()}`);
                }
              }}
            />
            <BsSearch className="mr-3" />
          </div>
          <BsMicFill className="text-2xl mr-3 hidden md:block" />
        </div>
        {/* Avatar */}
        <div className="flex items-center justify-between">
          <Link href="/upload">
            <BsCameraVideo
              className="hidden md:block text-3xl mr-7"
              data-tip="Add a new Video"
            />
          </Link>
          <div className="relative mr-3 md:mr-7">
            <BsBell className="text-2xl md:text-3xl" data-tip="Notifications" />
            <span className="absolute -top-3 -right-1 w-4 h-4 text-center flex items-center justify-center rounded-full bg-red-600 p-3 text-[12px] text-white ">
              {notification}
            </span>
          </div>
          <ThemeSwitch />
          <div className="mr-5 md:mr-7 " data-tip="Your account">
              {connectedAccount && userNickname ? (
                <button onClick={() => {showWallet();}}>
                  <Avatar name="Lilly0" size={"35"} round={true} />
                </button>
              ) : (
                <button onClickCapture={LoginHandler}>
                  <strong>Login</strong>
                </button>
              )}
          </div>
        </div>
        {isMounted && <ReactTooltip backgroundColor="#403e3f" isCapture />}
      </div>
  );
}

export default Header;
