require('dotenv').config()
require("@nomiclabs/hardhat-waffle");

module.exports = {
  solidity: {
    version: "0.8.6",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      },
    },
  },

  networks: {
    goerli: {
      url: "https://goerli.infura.io/v3/cb47771bf3324acc895994de6752654b",
      accounts: [
        process.env.DEPLOY_KEY_GOERLI,
      ],
    },
    xdai: {
      url: "https://rpc.xdaichain.com/",
      accounts: [
        process.env.DEPLOY_KEY_XDAI,
      ],
    }
  },
};

task("deploy", "Deploys the contract")
  .setAction(async () => {
    const factory = await ethers.getContractFactory("MEVsky");
    const contract = await factory.deploy();
    console.log("deploy tx:", contract.deployTransaction.hash);
    const receipt = await contract.deployTransaction.wait();
    if (receipt.status == 1) {
      console.log("contract deployed at", receipt.contractAddress);
    } else {
      console.log("deploy transaction failed");
      console.log(receipt);
    }
  });
