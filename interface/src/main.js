import Vue from 'vue';
import { ethers } from 'ethers';
import App from './App.vue';
import * as config from './config';

Vue.config.productionTip = false;

let provider;
let hasWalletConnection;
const backupProvider = new ethers.providers.JsonRpcProvider(
  config.backupRPCProviderURL,
);

if (window.ethereum) {
  provider = new ethers.providers.Web3Provider(window.ethereum);
  window.ethereum.on('chainChanged', () => window.location.reload());
  hasWalletConnection = true;
} else {
  provider = backupProvider;
  hasWalletConnection = false;
}

Vue.prototype.$hasWalletConnection = hasWalletConnection;
Vue.prototype.$provider = provider;
Vue.prototype.$backupProvider = backupProvider;

new Vue({
  render: (h) => h(App),
}).$mount('#app');
