<template>
  <div id="app">
    <ChainChecker
      v-if="!correctChain"
      @chainChecked="correctChain = $event"
    />
    <Machine v-if="correctChain" />
    <Bar />
  </div>
</template>

<script>
import ChainChecker from './ChainChecker.vue';
import Machine from './Machine.vue';
import Bar from './Bar.vue';

export default {
  name: 'App',
  components: {
    ChainChecker,
    Machine,
    Bar,
  },

  data() {
    return {
      correctChain: null,
    };
  },
};
</script>

<style>
body {
  background-color: #F8D2FF;
}
</style>
