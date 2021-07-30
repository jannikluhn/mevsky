<template>
  <Box title="Error" @close="$emit('close')">
    <p v-if="error">
      {{ error.message }}
    </p>
  </Box>
</template>

<script>
import Box from './Box.vue';

export default {
  name: 'Error',
  props: [
    'error',
  ],
  components: {
    Box,
  },
};
</script>
