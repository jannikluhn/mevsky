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
};
