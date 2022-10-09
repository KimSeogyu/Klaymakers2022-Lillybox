import { ethers } from "hardhat";

async function main() {
  const lillyboxFactory = await ethers.getContractFactory("Lillybox");
  const lillybox = await lillyboxFactory.deploy();
  await lillybox.deployed();

  console.log(`Ownership deployed to ${lillybox.address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
