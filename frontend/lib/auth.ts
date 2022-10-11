import axios from "axios";
import type { IUserInfo } from "./types";

const axiosInstance = axios.create({
  baseURL: process.env.AUTH_SERVER_ENDPOINT,
});

const getLoginRequestID = async (account: string): Promise<string> => {
  return axiosInstance
    .post("/request_id", { account })
    .then((res) => res.data.request_id)
    .catch((err) => null)
};

const insertUser = async (account: string, nickname: string): Promise<string> => {
  return axiosInstance
    .post("/sign", { account, nickname })
    .then((res) => res.data.request_id)
    .catch((err) => null)
};

const checkNickname = async (nickname: string): Promise<boolean> => {
  try {
    const CommonRegExp = /[^a-z|^A-Z|^0-9]/gi;
    return CommonRegExp.test(nickname) == false ? (nickname != '' ? true : false) : false;
  } catch (err) {
    console.error(err);
    return false;
  }
};

const login = async (account: string, signature: string): Promise<boolean> => {
  // return true;
  return axiosInstance
    .post("/login", { account, signature })
    .then((res) => res.data.success);
};

const getUserInfo = async (account: string): Promise<IUserInfo> => {
  return axiosInstance
    .post("/user", {account})
    .then((res) => res.data.result)
    .catch((err) => null)
};

const AuthAPI = {
  getLoginRequestID,
  login,
  insertUser,
  checkNickname,
  getUserInfo,
};

export default AuthAPI;
