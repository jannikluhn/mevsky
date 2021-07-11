<template>
  <div>
    <div>
      <p>Current State: {{ state }}</p>
    </div>
    <div>{{ animationQueue }}</div>
    <div>{{ currentAnimation }}</div>

    <div class="switch switch-off" ref="switch"></div>

    <div class="bot bot-hidden" ref="bot"></div>
  </div>
</template>

<script>
function newAnimation(animationItems) {
  return {
    items: animationItems,
  };
}

function newAnimationItem(ref, preClasses, postClasses) {
  return {
    ref,
    preClasses,
    postClasses,
  };
}

const turnOnAnimation = newAnimation([
  newAnimationItem('switch', ['switch', 'switch-off', 'switch-turn-on'], ['switch', 'switch-on']),
]);
const revertTurnOnAnimation = newAnimation([
  newAnimationItem('switch', ['switch', 'switch-on', 'switch-revert-turn-on'], ['switch', 'switch-off']),
]);
const turnOffAnimation = newAnimation([
  newAnimationItem('bot', ['bot', 'bot-hidden', 'bot-turn-off'], ['bot', 'bot-hidden']),
  newAnimationItem('switch', ['switch', 'switch-on', 'switch-turn-off'], ['switch', 'switch-off']),
]);

export default {
  name: 'MachineAnimation',
  props: [
    'state',
  ],
  data() {
    return {
      animationQueue: [],
      currentAnimation: null,
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
      this.animationQueue.push(turnOnAnimation);
      this.advanceAnimation();
    },
    turnOff() {
      this.animationQueue.push(turnOffAnimation);
      this.advanceAnimation();
    },
    revertTurnOn() {
      this.animationQueue.push(revertTurnOnAnimation);
      this.advanceAnimation();
    },

    advanceAnimation() {
      if (this.currentAnimation !== null) {
        return;
      }
      if (this.animationQueue.length === 0) {
        this.currentAnimation = null;
        return;
      }
      this.currentAnimation = this.animationQueue.shift();

      const startedAnimations = new Set();
      for (const [index, item] of this.currentAnimation.items.entries()) {
        const element = this.$refs[item.ref];
        element.classList.remove(...element.classList);
        for (const className of item.preClasses) {
          element.classList.add(className);
        }
        startedAnimations.add(index);
      }

      for (const [index, item] of this.currentAnimation.items.entries()) {
        const element = this.$refs[item.ref];
        element.addEventListener('animationend', () => {
          element.classList.remove(...element.classList);
          for (const className of item.postClasses) {
            element.classList.add(className);
          }
          startedAnimations.delete(index);

          if (startedAnimations.size === 0) {
            this.currentAnimation = null;
            this.advanceAnimation();
          }
        },
        {
          once: true,
        });
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.switch {
  position: relative;
  width: 100px;
  height: 100px;
}

.switch-off {
  background-color: red;
}

.switch-on {
  background-color: green;
}

.bot {
  position: relative;
  width: 100px;
  height: 100px;
}

.bot-hidden {
  background-color: grey;
}

.switch-turn-on {
  animation-name: switch-turn-on;
  animation-duration: 5s;
}

.switch-revert-turn-on {
  animation-name: switch-revert-turn-on;
  animation-duration: 1s;
}

.bot-turn-off {
  animation-name: bot-turn-off;
  animation-duration: 5s;
}

.switch-turn-off {
  animation-name: switch-turn-off;
  animation-duration: 5s;
}

@keyframes switch-turn-on {
  from {
    background-color: red;
  }
  to {
    background-color: green;
  }
}

@keyframes switch-revert-turn-on {
  from {
    background-color: green;
  }
  to {
    background-color: red;
  }
}

@keyframes bot-turn-off {
  from {
    background-color: grey;
  }
  50% {
    background-color: blue;
  }
  to {
    background-color: grey;
  }
}

@keyframes switch-turn-off {
  from {
    transform: rotateY(0);
    background-color: green;
  }
  to {
    transform: rotateY(720deg);
    background-color: red;
  }
}
</style>
