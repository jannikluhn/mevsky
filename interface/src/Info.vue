<template>
  <Box title="Welcome" @close="$emit('close')">
    <p>
      This is a <a href="https://en.wikipedia.org/wiki/Useless_machine">useless machine</a>
      running on a blockchain. It is powered by <i>miner extractable value</i> (MEV). Anyone can
      turn it on by clicking the switch. This will send a transaction to the useless machine smart
      contract and deposit a small bounty. The bounty is a form of MEV that bots can extract by
      turning the machine off again.
    </p>
    <p>
      The Useless MEV Machine is deployed on the
      <a
        href="https://blockscout.com/xdai/mainnet/address/0xeCbb670751ED019b021A062E30dfdf22Bc3fe8d9/"
      >
        xDai blockchain</a>. It has been created by
      <a href="http://marianono.superhi.com/">Maria Leonor Nogueira</a> and
      <a href="https://twitter.com/AnotherJannik">Jannik Luhn</a>. You can follow the machine's
      very own <a href="https://twitter.com/UselessMEV">Twitter account</a> on which it tweets
      about the accounts turning it on and off.
    </p>
  </Box>
</template>

<script>
import Box from './Box.vue';

export default {
  name: 'Info',
  components: {
    Box,
  },
};
</script>
