<template>
  <div>
    <div class="bar">
      <a href="https://blockscout.com"><img class="icon" src="./assets/eth-diamond-glyph.png"></a>
      <a href="https://github.com/jannikluhn/mevsky"><img class="icon" src="./assets/GitHub-Mark-64px.png"></a>
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
}
.icon {
  height: 64px;
  margin: 20px;
  opacity: 30%;
}
.icon:hover {
  opacity: 70%;
}
</style>
