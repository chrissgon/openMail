<template>
  <article id="parallax" class="relative w-full">
    <div class="absolute w-full h-full" data-depth="0.6">
      <img class="path" src="@/assets/path.svg" />
    </div>
    <div class="absolute w-full h-full" data-depth="0.2">
      <img class="desktop" src="@/assets/desktop.svg" />
    </div>
    <div class="absolute w-full h-full" data-depth="0.3">
      <img class="mobile" src="@/assets/mobile.svg" />
    </div>
  </article>
</template>

<script>
import Parallax from "parallax-js";

export default {
  setup() {
    window.onload = () => {
      const parallax = document.getElementById("parallax");
      new Parallax(parallax);
    };

    return {};
  },
};
</script>

<style lang="scss" scoped>
img {
  &.path {
    max-width: 100%;
    transform: translate(10%, 5%);
  }
  &.desktop {
    max-width: 80%;
    transform: translate(5%, 15%);
  }
  &.mobile {
    max-width: 50%;
    transform: translate(110%, 50%);
  }
}
</style>
