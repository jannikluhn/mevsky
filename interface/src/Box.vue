<template>
  <div class="overlay" @click='close'>
    <div class="box" @click.stop>
      <div class="box-header">
        <h1>{{ title }}</h1>
        <button class="close-button" @click='close'><span>Close</span></button>
      </div>
      <div class="box-body">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Info',
  props: [
    'title',
  ],
  methods: {
    close() {
      this.$emit('close');
    },
  },
};
</script>

<style scoped>
.overlay {
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background-color: rgba(0, 0, 0, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
}

.box {
  width: 500px;
  background-color: #F8D2FF;
  border-style: solid;
  border-color: black;
  border-width: 5px;
  border-radius: 20px;
  overflow: hidden;
  box-shadow: 0px 0px 70px 40px darkgrey;
}

.box-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-left: 1em;
  padding-right: 1em;
  padding-top: 0.5em;
  padding-bottom: 0.5em;
}

.box-body {
  padding: 0 1.5em;
}

h1 {
  display: inline-block;
  margin: 0 0 0 0.2em;
}

a {
  color: #5d30dd;
}

.close-button {
  float: right;
  background-image: url(./assets/cancel.png);
  background-color: transparent;
  background-position: 50% 50%;
  width: 30px;
  height: 30px;
  background-size: 30px;
  border: 0;
  opacity: 100%;
  cursor: pointer;
}

.close-button span{
  display: none;
}

.box-transition-enter-active {
  animation: fade-in 0.1s;
}
.box-transition-leave-active {
  animation: fade-in 0.1s reverse;
}
@keyframes fade-in {
  0% {
    opacity: 0%;
  }
  100% {
    opacity: 100%;
  }
}
</style>
