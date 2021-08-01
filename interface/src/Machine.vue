<template>
  <div>
    <div>
      <MachineAnimation
        :state="state"
        @turnOn="onClickButton"
        @showInfo="$emit('showInfo')"
      />
    </div>
  </div>
</template>

<script>
import { ethers } from 'ethers';

import * as config from './config';
import MachineAnimation from './MachineAnimation.vue';

export default {
  name: 'Machine',
  props: [
    'provider',
    'contract',
    'network',
  ],
  components: {
    MachineAnimation,
  },
  data() {
    return {
      state: null,
      minBounty: null,

      txInFlight: false,
    };
  },
  computed: {
    loading() {
      return this.state === null;
    },
  },

  async mounted() {
    this.contract.on('TurnedOn', this.onTurnedOn);
    this.contract.on('TurnedOff', this.onTurnedOff);

    this.contract.functions.on()
      .then(([on]) => {
        this.state = {
          on,
          optimisticOn: on,
        };
      })
      .catch((e) => {
        this.$emit('error', {
          error: e,
          message: 'Failed to query machine status.',
        });
      });

    this.contract.minBounty()
      .then((minBounty) => {
        this.minBounty = minBounty;
      })
      .catch((e) => {
        this.$emit('error', {
          error: e,
          message: 'Failed to query minimum bounty amount.',
        });
      });
  },

  methods: {
    async onTurnedOn() {
      this.state = {
        on: true,
        optimisticOn: true,
      };
    },

    async onTurnedOff() {
      this.state = {
        on: false,
        optimisticOn: false,
      };
    },

    async onClickButton() {
      if (this.txInFlight || !this.state || this.state.on || this.state.optimisticOn) {
        return;
      }
      this.txInFlight = true;

      if (!this.$hasWalletConnection) {
        this.$emit('error', {
          error: null,
          message: 'To turn on the machine, you need a browser wallet such as Metamask. Please '
          + 'install one and try again.',
        });
        this.txInFlight = false;
        return;
      }

      // If the user's wallet is connected to the wrong chain, ask them to switch.
      if (!this.network || this.network.chainId !== config.chainId) {
        try {
          await window.ethereum.request({
            method: 'wallet_switchEthereumChain',
            params: [{ chainId: ethers.BigNumber.from(config.chainId).toHexString() }],
          });
        } catch (e) {
          this.$emit('error', {
            error: e,
            message: 'Failed to change network.',
          });
          this.txInFlight = false;
          return;
        }
      }

      // Request accounts
      try {
        await window.ethereum.request({ method: 'eth_requestAccounts', params: [] });
      } catch (e) {
        this.$emit('error', {
          error: e,
          message: 'Failed to request accounts.',
        });
        this.txInFlight = false;
        return;
      }

      const signer = this.provider.getSigner();
      const signingContract = this.contract.connect(signer);

      let tx;
      try {
        tx = await signingContract.turnOn({ value: this.minBounty });
      } catch (e) {
        this.$emit('error', {
          error: e,
          message: 'Failed to send transaction.',
        });
        this.txInFlight = false;
        return;
      }

      this.state = {
        on: this.state.on,
        optimisticOn: true, // we assume the tx will go through
      };

      try {
        await tx.wait();
      } catch (e) {
        this.$emit('error', {
          error: e,
          message: 'Transaction failed.',
        });
        this.state = {
          on: this.state.on,
          optimisticOn: this.state.on, // we're realistic now
        };
        this.txInFlight = false;
        return;
      }

      this.txInFlight = false;
    },
  },
};
</script>
