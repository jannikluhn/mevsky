<template>
  <div>
    <div class="machine"></div>
    <div ref="switch"></div>
    <div ref="arm"></div>
    <button class="switch-button" @click="onClick"></button>
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
]);
const revertTurnOnAnimation = newAnimation([
  newAnimationItem(
    'switch',
    ['switch', 'switch-on', 'switch-revert-turn-on'],
    ['switch', 'switch-off'],
    1,
  ),
]);
const turnOffAnimation = newAnimation([
  newAnimationItem('switch',
    ['switch', 'switch-on', 'switch-turn-off'],
    ['switch', 'switch-off'],
    1),
  newAnimationItem('arm',
    ['arm', 'arm-off', 'arm-turn-off'],
    ['arm', 'arm-off'],
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
      ref: 'arm',
      classNames: [
        'arm',
        'arm-off',
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
      ref: 'arm',
      classNames: [
        'arm',
        'arm-on',
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

.part {
  background-repeat: no-repeat;
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
}

.switch-button {
  position: fixed;
  left: calc(50% - 60px);
  top: calc(50% - 150px);
  width: 60px;
  height: 100px;
  background-color: red;
  opacity: 0;
  cursor: pointer;
}

//
// Machine
//

.machine {
  @extend .part;
  background-image: url("./assets/machine.png");
  background-position: 50% 50%;
  z-index: 0;
}

//
// Switch
//

$switch-angle-off: -23deg;
$switch-angle-on: 0deg;

.switch {
  @extend .part;
  background-image: url("./assets/switch.png");
  background-position: calc(50% - 18px) calc(50% - 100px);
  z-index: -2;
  transform-origin: calc(50% - 25px) calc(50% - 68px);

  // width: 10px;
  // height: 10px;
  // background-color: red;
  // position: absolute;
  // left: calc(50% - 25px);
  // top: calc(50% - 68px);
  // z-index: 100;
}

.switch-off {
  transform: rotate($switch-angle-off);
}

.switch-on {
  transform: rotate($switch-angle-on);
}

.switch-turn-on {
  animation-name: switch-turn-on;
  animation-duration: 0.5s;

  @keyframes switch-turn-on {
    from {transform: rotate($switch-angle-off)};
    to {transform: rotate($switch-angle-on)};
  }
}

.switch-revert-turn-on {
  animation-name: switch-revert-turn-on;
  animation-duration: 0.5s;

  @keyframes switch-revert-turn-on {
    from {transform: rotate($switch-angle-on)};
    to {transform: rotate($switch-angle-off)};
  }
}

.switch-turn-off {
  animation-name: switch-turn-off;
  animation-duration: 0.2s;
  animation-delay: 0.35s;

  @keyframes switch-turn-off {
    from {transform: rotate($switch-angle-on)};
    to {transform: rotate($switch-angle-off)};
  }
}

//
// Arm
//

$arm-angle-off: 65deg;
$arm-angle-on: 0deg;

.arm {
  @extend .part;
  background-image: url("./assets/arm.png");
  background-position: calc(50% - 30px) calc(50% - 25px);
  transform-origin: calc(50% - 30px) calc(50% - 25px);
  z-index: -1;
}

.arm-off {
  transform: rotate($arm-angle-off);
}

.arm-on {
  transform: rotate($arm-angle-on);
}

.arm-turn-off {
  animation-name: arm-turn-off;
  animation-duration: 1s;

  @keyframes arm-turn-off {
    0% {transform: rotate($arm-angle-off)};
    50% {transform: rotate($arm-angle-on)};
    100% {transform: rotate($arm-angle-off)};
  }
}
</style>
