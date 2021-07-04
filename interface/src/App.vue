<template>
  <div id="app">
    <ChainChecker
      v-if="!correctChain"
      @chainChecked="correctChain = $event"
    />
    <Machine v-if="correctChain" />
  </div>
</template>

<script>
import ChainChecker from './ChainChecker.vue';
import Machine from './Machine.vue';

export default {
  name: 'App',
  components: {
    ChainChecker,
    Machine,
  },

  data() {
    return {
      correctChain: null,
    };
  },
};
</script>
