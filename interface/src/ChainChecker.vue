<template>
  <div>
    <p v-if="error">
      Failed to check chain.
    </p>
    <p v-else-if="!network">
      Chain check in progress.
    </p>
    <p v-else-if="correctChain">
      You are connected to the correct chain.
    </p>
    <p v-else-if="!correctChain">
      You are not connected to the correct chain.
    </p>
    <p v-else>
      This should never be shown.
    </p>
  </div>
</template>

<script>
import * as config from './config';

export default {
  name: 'ChainChecker',
  data() {
    return {
      network: null,
      error: null,
    };
  },
  computed: {
    correctChain() {
      if (!this.network) {
        return null;
      }
      return this.network.chainId === config.chainId;
    },
  },

  async mounted() {
    try {
      this.network = await this.$provider.getNetwork();
    } catch (e) {
      this.error = e;
      console.error(e);
    }
    this.$emit('chainChecked', this.correctChain);
  },
};
</script>
