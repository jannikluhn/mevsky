<template>
  <div id="app">
    <Machine
      v-if="provider && contract"
      :provider="provider"
      :contract="contract"
      :network="network"
      @error="error = $event"
      @showInfo="infoShown = true"
    />
    <transition name="box-transition">
      <Error v-if="error" :error="error" @close="error = null" />
    </transition>
    <transition name="box-transition">
      <Info v-if="infoShown" @close="infoShown = false" />
    </transition>
    <Bar />
  </div>
</template>

<script>
import { ethers } from 'ethers';

import * as config from './config';
import * as contractMetadata from './assets/MEVsky.json';
import Machine from './Machine.vue';
import Bar from './Bar.vue';
import Error from './Error.vue';
import Info from './Info.vue';

export default {
  name: 'App',
  components: {
    Machine,
    Error,
    Info,
    Bar,
  },

  data() {
    return {
      network: null,
      provider: null,
      contract: null,
      error: null,

      infoShown: false,
    };
  },

  async mounted() {
    // Use wallet provider, if user has a wallet and it is connected to the right network.
    // Otherwise, use backup provider for now. We will ask them to switch when they try to click
    // the machine's switch.
    let provider;
    if (this.$hasWalletConnection) {
      try {
        this.network = await this.$provider.getNetwork();
        if (this.network.chainId === config.chainId) {
          provider = this.$provider;
        } else {
          provider = this.$backupProvider;
        }
      } catch (e) {
        this.error = {
          error: e,
          message: 'Failed to query network',
        };
        provider = this.$backupProvider;
      }
    } else {
      provider = this.$backupProvider;
    }

    this.provider = provider;
    this.contract = new ethers.Contract(config.contractAddress, contractMetadata.abi, provider);
  },
};
</script>

<style>
body {
  background-color: #F8D2FF;
  font-family: Arial, Helvetica, sans-serif;
}
</style>
