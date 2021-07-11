<template>
  <div>
    <!-- <div>
      <p>Current State: {{ state }}</p>
    </div>
    <div>{{ animationQueue }}</div>
    <div>{{ currentAnimation }}</div> -->

    <div ref="switch">
      <div ref="switch-knob" @click="onClick"></div>
    </div>

    <div ref="bot"></div>
  </div>
</template>

<script>
function newAnimation(animationItems) {
  return {
    items: animationItems,
  };
}

function newAnimationItem(ref, preClasses, postClasses, steps) {
  return {
    ref,
    preClasses,
    postClasses,
    steps,
  };
}

const turnOnAnimation = newAnimation([
  newAnimationItem(
    'switch',
    ['switch', 'switch-off', 'switch-turn-on'],
    ['switch', 'switch-on'],
    1,
  ),
  newAnimationItem(
    'switch-knob',
    ['switch-knob', 'switch-knob-off', 'switch-knob-turn-on'],
    ['switch-knob', 'switch-knob-on'],
    1,
  ),
]);
const revertTurnOnAnimation = newAnimation([
  newAnimationItem(
    'switch',
    ['switch', 'switch-on', 'switch-revert-turn-on'],
    ['switch', 'switch-off'],
    1,
  ),
  newAnimationItem(
    'switch-knob',
    ['switch-knob', 'switch-knob-on', 'switch-knob-revert-turn-on'],
    ['switch-knob', 'switch-knob-off'],
    1,
  ),
]);
const turnOffAnimation = newAnimation([
  newAnimationItem('bot',
    ['bot', 'bot-hidden', 'bot-turn-off'],
    ['bot', 'bot-hidden'],
    3),
  newAnimationItem('switch',
    ['switch', 'switch-on', 'switch-turn-off'],
    ['switch', 'switch-off'],
    1),
  newAnimationItem('switch-knob',
    ['switch-knob', 'switch-knob-on', 'switch-knob-turn-off'],
    ['switch-knob', 'switch-knob-off'],
    1),
]);

// classes for refs if state is on or off
const initialClasses = {
  false: [
    {
      ref: 'switch',
      classNames: [
        'switch',
        'switch-off',
      ],
    },
    {
      ref: 'switch-knob',
      classNames: [
        'switch-knob',
        'switch-knob-off',
      ],
    },
    {
      ref: 'bot',
      classNames: [
        'bot',
        'bot-hidden',
      ],
    },
  ],
  true: [
    {
      ref: 'switch',
      classNames: [
        'switch',
        'switch-on',
      ],
    },
    {
      ref: 'switch-knob',
      classNames: [
        'switch-knob',
        'switch-knob-on',
      ],
    },
    {
      ref: 'bot',
      classNames: [
        'bot',
        'bot-hidden',
      ],
    },
  ],
};

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
      if (!oldState) {
        if (newState !== null) {
          // apply initial states
          for (const item of initialClasses[newState.on]) {
            this.$refs[item.ref].classList.add(...item.classNames);
          }
        }
        return;
      }

      const noChange = (
        oldState.on === newState.on
          && oldState.optimisticOn === newState.optimisticOn
      );
      if (noChange) {
        return;
      }

      const isOff = !oldState.on && !oldState.optimisticOn;
      const isOptimisticOn = !oldState.on && oldState.optimisticOn;
      const isOn = oldState.on && oldState.optimisticOn;
      if (oldState.on && !oldState.optimisticOn) {
        console.error('unexpected old state', oldState);
        return;
      }

      const willBeOff = !newState.on && !newState.optimisticOn;
      const willBeOptimisticOn = !newState.on && newState.optimisticOn;
      const willBeOn = newState.on && newState.optimisticOn;
      if (newState.on && !newState.optimisticOn) {
        console.error('unexpected new state', newState);
        return;
      }

      if (isOff && (willBeOn || willBeOptimisticOn)) {
        this.turnOn();
      } else if (isOptimisticOn && willBeOff) {
        this.revertTurnOn();
      } else if (isOptimisticOn && willBeOn) {
        // do nothing, optimism was justified
      } else if (isOn && willBeOptimisticOn) {
        console.error('unexpected state transition from off to optimistic on');
      } else if (isOn && willBeOff) {
        this.turnOff();
      } else {
        console.error('unexpected state transition', oldState, newState);
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

      const animationCounter = {};
      for (const [index, item] of this.currentAnimation.items.entries()) {
        const element = this.$refs[item.ref];
        element.classList.remove(...element.classList);
        for (const className of item.preClasses) {
          element.classList.add(className);
        }
        animationCounter[index] = item.steps;
      }

      for (const [index, item] of this.currentAnimation.items.entries()) {
        const element = this.$refs[item.ref];
        const animationEndListener = () => {
          animationCounter[index] -= 1;
          if (animationCounter[index] <= 0) {
            delete animationCounter[index];
          }

          if (!animationCounter[index]) {
            element.classList.remove(...element.classList);
            for (const className of item.postClasses) {
              element.classList.add(className);
            }
            element.removeEventListener('animationend', animationEndListener);
          }

          if (Object.keys(animationCounter).length === 0) {
            this.currentAnimation = null;
            this.advanceAnimation();
          }
        };
        element.addEventListener('animationend', animationEndListener);
      }
    },

    onClick() {
      if (this.currentAnimation !== null || this.animationQueue.length !== 0) {
        return;
      }
      if (this.state.on || this.state.optimisticOn) {
        return;
      }
      this.$emit('turnOn');
    },
  },
};
</script>

<style lang="scss" scoped>
$switch-width: 380px;
$switch-height: 150px;
$switch-knob-padding: 12px;
$switch-border-radius: 20px;
$switch-border-width: 8px;
$switch-color: blue;

$switch-knob-size: $switch-height - 2 * $switch-knob-padding;

$switch-knob-left-off: $switch-knob-padding;
$switch-knob-left-on: $switch-width - $switch-height + $switch-border-width;

$bot-width: 150px;
$bot-height: 100px;

$bot-margin-left-home: $switch-width / 2 + $switch-border-width + 150px;

.switch {
  position: fixed;
  top: 50%;
  left: 50%;
  margin-top: -$switch-height / 2 - $switch-border-width;
  margin-left: -$switch-width / 2 - $switch-border-width;
  width: $switch-width;
  height: $switch-height;

  border-style: solid;
  border-radius: $switch-border-radius;
  border-width: $switch-border-width;
  border-color: $switch-color;
}

.switch-on {}

.switch-off {}

.switch-knob {
  position: absolute;
  top: $switch-knob-padding;
  width: $switch-knob-size;
  height: $switch-knob-size;
  background-color: $switch-color;
  border-radius: $switch-border-radius;
}

.switch-knob-off {
  left: $switch-knob-left-off;
  cursor: pointer;
}

.switch-knob-on {
  left: $switch-knob-left-on;
  cursor: not-allowed;
}

.bot {
  position: fixed;
  top: 50%;
  left: 50%;
  margin-top: -$bot-height / 2;
  margin-left: $bot-margin-left-home;
  width: $bot-width;
  height: $bot-height;
  background-color: red;
}

.bot-hidden {
  opacity: 0;
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
  animation-fill-mode: forwards;
  animation-name: bot-appear, bot-hit, bot-disappear;
  animation-duration: 0.5s, 0.2s, 0.5s;
  animation-delay: 0s, 0.5s, 0.85s;
}

.switch-turn-off {
  animation-name: switch-turn-off;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
}

.switch-knob-turn-on {
  animation-name: switch-knob-turn-on;
  animation-duration: 1s;
  cursor: not-allowed;
}

.switch-knob-revert-turn-on {
  animation-name: switch-knob-revert-turn-on;
  animation-duration: 0.3s;
  cursor: not-allowed;
}

.switch-knob-turn-off {
  animation-name: switch-knob-turn-off;
  animation-duration: 0.2s;
  animation-delay: 0.6s;
  cursor: not-allowed;
}

@keyframes switch-turn-on {

}

@keyframes switch-revert-turn-on {

}

@keyframes bot-appear {
  from {
    opacity: 0%;
  }
  to {
    opacity: 100%;
  }
}

@keyframes bot-disappear {
  from {
    opacity: 100%;
  }
  to {
    opacity: 0%;
  }
}

@keyframes bot-hit {
  0% {
    margin-left: $bot-margin-left-home;
  }
  50% {
    margin-left: $bot-margin-left-home - 160px;
  }
  100% {
    margin-left: $bot-margin-left-home;
  }
}

@keyframes switch-turn-off {
  // from {
  //   transform: rotateY(0);
  // }
  // to {
  //   transform: rotateY(900deg);
  // }
}

@keyframes switch-knob-turn-on {
  from {
    left: $switch-knob-left-off;
  }
  to {
    left: $switch-knob-left-on;
  }
}

@keyframes switch-knob-revert-turn-on {
  from {
    left: $switch-knob-left-on;
  }
  to {
    left: $switch-knob-left-off;
  }
}

@keyframes switch-knob-turn-off {
  from {
    left: $switch-knob-left-on;
  }
  to {
    left: $switch-knob-left-off;
  }
}
</style>
