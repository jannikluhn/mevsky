<template>
  <div>
    <div class="bar">
      <a href="https://blockscout.com/xdai/mainnet/address/0xeCbb670751ED019b021A062E30dfdf22Bc3fe8d9/">
        <img class="icon" src="./assets/ethereum-100.png">
      </a>
      <a href="https://github.com/jannikluhn/mevsky"><img class="icon" src="./assets/octocat-100.png"></a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Bar',
};
</script>

<style scoped>
.bar {
  position: fixed;
  bottom: 0;
  right: 25px;
  z-index: -1;
}
.icon {
  height: 54px;
  margin: 15px 5px;
  opacity: 0.3;
}
.icon:hover {
  opacity: 0.7;
}
</style>
