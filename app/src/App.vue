<template>
  <section class="app flex justify-center items-center w-full h-screen">
    <section class="banner h-full w-full absolute lg:w-7/12 lg:flex lg:relative">
      <BannerMain />
    </section>
    <section class="form z-10 w-full flex justify-center items-center lg:w-5/12 lg:ml-10">
      <FormEmail />
    </section>
  </section>
</template>

<script>
import BannerMain from "@/components/BannerMain";
import FormEmail from "@/components/FormEmail";

export default {
  name: "App",
  components: { BannerMain, FormEmail },
};
</script>

<style lang="scss">
html,
body {
  overflow: hidden;
}
</style>
