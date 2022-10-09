import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
require("hardhat-abi-exporter");

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.7",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  networks: {
    klaytn: {
      url: "https://api.baobab.klaytn.net:8651",
      gasPrice: 250000000000,
      accounts: [
        "0xfd22e1eba8d59ffb9f9b36d291259c415abd57982ccf61a4a7167e43cfc51673",
      ],
    },
  },
};

export default config;
