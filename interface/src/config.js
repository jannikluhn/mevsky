// localhost
// module.exports = {
//   backupRPCProviderURL: 'https://rpc.xdaichain.com/',
//   contractAddress: '0x5FbDB2315678afecb367f032d93F642f64180aa3',
//   chainId: 31337,
// };

// goerli
// module.exports = {
//   backupRPCProviderURL: 'https://goerli.infura.io/v3/cb47771bf3324acc895994de6752654b',
//   contractAddress: '0x284f6A78BB277716B6397fEd2Efe9D2FEfb8731e',
//   chainId: 5,
// };

// xdai
module.exports = {
  backupRPCProviderURL: 'https://rpc.xdaichain.com/',
  contractAddress: '0xeCbb670751ED019b021A062E30dfdf22Bc3fe8d9',
  chainId: 100,
  addEthereumChainSettings: {
    chainId: '0x64',
    chainName: 'xDai',
    nativeCurrency: {
      name: 'xDai',
      symbol: 'xDai',
      decimals: 18,
    },
    rpcUrls: [
      'https://rpc.xdaichain.com/',
    ],
    blockExplorerUrls: [
      'https://blockscout.com/xdai/mainnet',
    ],
  },
};
