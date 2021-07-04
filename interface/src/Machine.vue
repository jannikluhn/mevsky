<template>
  <div>
    <p v-if="error">{{ error.message }}</p>
    <p v-else-if="loading">
      Loading...
    </p>
    <div>
      <button
        :disabled="buttonDisabled"
        @click="onClickButton"
      >
        Turn on
      </button>
      <MachineAnimation
        :state="state"
      />
      <!-- <button @click="state = {on: false, optimisticOn: true}">Turn on button</button>
      <button @click="state = {on: true, optimisticOn: true}">Turn on chain</button>
      <button @click="state = {on: false, optimisticOn: false}">Turn off</button>
      <button @click="state = {on: false, optimisticOn: false}">revert</button> -->
    </div>
  </div>
</template>

<script>
import MachineAnimation from './MachineAnimation.vue';

export default {
  name: 'Machine',
  components: {
    MachineAnimation,
  },
  data() {
    return {
      error: null,
      state: null,
      minBounty: null,

      txInFlight: false,
    };
  },
  computed: {
    loading() {
      return this.state === null;
    },
    buttonDisabled() {
      return (
        !this.$hasWalletConnection
        || this.loading
        || this.state.on
        || this.txInFlight
      );
    },
  },

  async mounted() {
    this.$contract.on('TurnedOn', this.onTurnedOn);
    this.$contract.on('TurnedOff', this.onTurnedOff);

    this.$contract.functions.on()
      .then(([on]) => {
        this.state = {
          on,
          optimisticOn: on,
        };
      })
      .catch((e) => {
        this.error = {
          error: e,
          message: 'Failed to query machine status',
        };
      });

    this.$contract.minBounty()
      .then((minBounty) => {
        this.minBounty = minBounty;
      })
      .catch((e) => {
        this.error = {
          error: e,
          message: 'Failed to query minimum bounty',
        };
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
      if (this.buttonDisabled) {
        return;
      }
      this.txInFlight = true;

      try {
        await window.ethereum.request({ method: 'eth_requestAccounts', params: [] });
      } catch (e) {
        this.error = {
          error: e,
          message: 'Failed to request accounts.',
        };
        this.txInFlight = false;
        console.error(e);
        return;
      }

      const signer = this.$provider.getSigner();
      const signingContract = this.$contract.connect(signer);

      let tx;
      try {
        tx = await signingContract.turnOn({ value: this.minBounty });
      } catch (e) {
        this.error = {
          error: e,
          message: 'Failed to send transaction.',
        };
        this.txInFlight = false;
        console.error(e);
        return;
      }

      this.state = {
        on: this.state.on,
        optimisticOn: true, // we assume the tx will go through
      };

      try {
        await tx.wait();
      } catch (e) {
        this.error = {
          error: e,
          message: 'Transaction failed.',
        };
        this.state = {
          on: this.state.on,
          optimisticOn: this.state.on, // we're realistic now
        };
        this.txInFlight = false;
        console.error(e);
        return;
      }

      this.txInFlight = false;
    },
  },
};
</script>
