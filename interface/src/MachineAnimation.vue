<template>
  <div>
    <div>
      <p>Current State: {{ state }}</p>
    </div>
    <div v-if="animationQueue.length != 0">
      <p>Current Animation: {{ animationQueue[0].type }}</p>
      <p>Progress: {{ animationQueue[0].progress }}</p>
    </div>
    <div v-else>
      <p>No animation in progress</p>
    </div>
  </div>
</template>

<script>
function newAnimation(type, duration) {
  return {
    type,
    duration,
    startTime: null,
    progress: null,
  };
}

function newTurnOffAnimation() {
  return newAnimation('turnOff', 2);
}

function newTurnOnAnimation() {
  return newAnimation('turnOn', 2);
}

function newRevertTurnOnAnimation() {
  return newAnimation('revertTurnOn', 2);
}

export default {
  name: 'MachineAnimation',
  props: [
    'state',
  ],
  data() {
    return {
      animationQueue: [],
      animationRunning: false,
    };
  },

  watch: {
    state(newState, oldState) {
      if (oldState === null) {
        return;
      }
      if (oldState.on === newState.on && oldState.optimisticOn === newState.optimisticOn) {
        return;
      }

      if (!oldState.on && !oldState.optimisticOn && !newState.on && newState.optimisticOn) {
        this.turnOn();
      } else if (!oldState.on && !oldState.optimisticOn && newState.on && newState.optimisticOn) {
        this.turnOn();
      } else if (!oldState.on && oldState.optimisticOn && newState.on && newState.optimisticOn) {
        // do nothing
      } else if (!oldState.on && oldState.optimisticOn && !newState.on && !newState.optimisticOn) {
        this.revertTurnOn();
      } else if (oldState.on && oldState.optimisticOn && !newState.on && !newState.optimisticOn) {
        this.turnOff();
      } else {
        console.error('unexpected state transition from', oldState, 'to', newState);
      }
    },
  },

  methods: {
    turnOn() {
      this.animationQueue.push(newTurnOnAnimation());
      this.animate();
    },
    turnOff() {
      this.animationQueue.push(newTurnOffAnimation());
      this.animate();
    },
    revertTurnOn() {
      this.animationQueue.push(newRevertTurnOnAnimation());
      this.animate();
    },

    animate() {
      if (this.animationRunning) {
        return;
      }
      this.animationRunning = true;

      const step = (timestamp) => {
        this.animationStep(timestamp);
        if (this.animationQueue.length > 0) {
          window.requestAnimationFrame(step);
        } else {
          this.animationRunning = false;
        }
      };
      window.requestAnimationFrame(step);
    },

    animationStep(timestamp) {
      if (this.animationQueue.length === 0) {
        return;
      }
      const animation = this.animationQueue[0];

      if (animation.startTime === null) {
        animation.startTime = timestamp;
      }
      animation.progress = (timestamp - animation.startTime) / 1000 / animation.duration;

      if (animation.progress > 1) {
        this.animationQueue.shift();
      }
    },
  },
};
</script>
