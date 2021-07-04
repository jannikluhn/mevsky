import Vue from 'vue';
import { ethers } from 'ethers';
import App from './App.vue';
import * as config from './config';
import * as contractMetadata from './assets/MEVsky.json';

Vue.config.productionTip = false;

let provider;
let hasWalletConnection;
if (window.ethereum) {
  provider = new ethers.providers.Web3Provider(window.ethereum);
  window.ethereum.on('chainChanged', () => window.location.reload());
  hasWalletConnection = true;
} else {
  provider = new ethers.providers.JsonRpcProvider(config.rpcProviderUrl);
  hasWalletConnection = false;
}

const contract = new ethers.Contract(config.contractAddress, contractMetadata.abi, provider);

Vue.prototype.$provider = provider;
Vue.prototype.$contract = contract;
Vue.prototype.$hasWalletConnection = hasWalletConnection;

new Vue({
  render: (h) => h(App),
}).$mount('#app');
