import { expect } from "chai";
import { ethers, network } from "hardhat";
import { Lillybox } from "../typechain-types";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { address } from "hardhat/internal/core/config/config-validation";

describe("Lillybox", function () {
  let lillyboxContract: Lillybox;
  let signers: SignerWithAddress[];
  let signer: SignerWithAddress;

  before(async function () {
    const lillyboxFactory = await ethers.getContractFactory("Lillybox");
    const lillybox = await lillyboxFactory.deploy(10, 10);
    await lillybox.deployed();

    lillyboxContract = await ethers.getContractAt("Lillybox", lillybox.address);
    signers = await ethers.getSigners();
    signer = signers[0];
  });

  it("배포 완료 점검", async () => {
    const deployed = await lillyboxContract.deployed();
    expect(deployed.address).not.eq(undefined);
  });

  const cid =
    "https://play-lh.googleusercontent.com/VNN45wC_dB4JgERRX-DYtRplHeVB8A02vIvHO_KaSW-jnoiafoAQa98V4opXZzdWZQ=w480-h960-rw";

  it("동영상 NFT 민팅", async () => {
    const contractTransaction = await lillyboxContract.mintVod(cid);
    const receipt = await contractTransaction.wait(1);
    expect(receipt.status).eq(1);
  });

  it("동영상 NFT 조회", async () => {
    const s = await lillyboxContract.uri(2);
    await lillyboxContract.showVideos();
    expect(typeof s).eq("string");
  });

  it("광고 NFT 민팅", async () => {
    const data = await lillyboxContract.minimumAdsCost();

    const ret = await lillyboxContract.mintAd(cid, {
      value: data,
    });
    const receipt = await ret.wait(1);
    expect(receipt.status).eq(1);
  });

  it("광고 NFT 조회", async () => {
    const ads = await lillyboxContract.showAds();
    expect(ads.length).eq(1);
  });

  it("입금", async () => {
    await lillyboxContract.deposit({
      value: ethers.utils.parseEther("100").toString(),
    });
  });

  it("출금", async () => {
    const message = await lillyboxContract.showWallet(signer.address);
    await lillyboxContract.withdraw(ethers.utils.parseEther("99").toString());
    const message1 = await lillyboxContract.showWallet(signer.address);
  });

  it("만료된 광고 소각", async () => {
    const before = await lillyboxContract.showAds();
    const message = before[0];
    expect(before.length).eq(1);
    await lillyboxContract.vacuumAd(message.tokenId, 101);
    const after = await lillyboxContract.showAds();
    expect(after.length).eq(0);
  });

  it("리워드 조회", async () => {
    const block = await network.provider.send("eth_blockNumber");
    const reward = await lillyboxContract.getRewards(signer.address);
    const w = await lillyboxContract.showWallet(signer.address);

    expect(
      ethers.utils.formatEther(
        w.balance.mul(w.stakedAt.sub(block).abs()).mul("10").toString()
      )
    ).eq(ethers.utils.formatEther(reward.toString()));
  });

  it("리워드 받기", async () => {
    const reward = await lillyboxContract.getRewards(signer.address);
    await lillyboxContract.flushReward();
    console.log(await lillyboxContract.balanceOf(signer.address, 0));
    console.log(reward);
    // expect(
    //   ethers.utils
    //     .formatEther()
    //     .toString()
    // ).eq(ethers.utils.formatEther(reward).toString());
  });
});
