<template>
  <div class="overlay" @click='close'>
    <div class="box" @click.stop>
      <div class="box-header">
        <h1>Welcome</h1>
        <button class="close-button" @click='close'><span>Close</span></button>
      </div>
      <div class="box-body">
        <p>
          This is a <a href="https://en.wikipedia.org/wiki/Useless_machine">useless machine</a>
          powered by Miner Extractable Value (MEV). Anyone can turn it on by clicking the switch.
          This will send a transaction to the useless machine smart contract and deposit a small
          bounty. The bounty is a form of MEV that bots can extract by turning the machine off
          again.
        </p>
        <p>
          The useless MEV machine lives on the <a href="https://blockscout.com/address">xDai
          blockchain</a>. It has been created by <a href="#">Maria Leonor Nogueira</a> and
          <a href="https://twitter.com/AnotherJannik">Jannik Luhn</a>. You can follow its very own
          <a href="">Twitter account</a> on which it tweets about the accounts turning it on and
          off.
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Info',
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
  padding-left: 1em;
  padding-right: 1em;
  padding-top: 1em;
  padding-bottom: 0.5em;
}

.box-body {
  padding-left: 1em;
  padding-right: 1em;
}

h1 {
  display: inline-block;
  margin: 0;
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
</style>
