<template>
  <div class="error">
    <!-- <h1>Error</h1>
    <p>{{ error.message }}</p> -->
  </div>
</template>

<script>
export default {
  name: 'Error',
  props: [
    'error',
  ],
};
</script>

<style scoped>
.error {
  /* position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  width: 500px;
  height: 300px;
  margin: auto;
  background-color: red; */
}
</style>
