<template>
  <div class="contanier-title">
    <dotlottie-player
      class="farm cow"
      src="https://lottie.host/01f4bcf3-6257-481a-aca6-7741d9a79d2a/6ahqUaaz6Y.json"
      background="transparent"
      speed="1"
      style="width: 300px; height: 300px"
      loop
      autoplay
    ></dotlottie-player>
    <dotlottie-player
      class="farm farmer"
      src="https://lottie.host/e843036e-5cfc-43ec-bdb1-8bf90ef8014a/FhcFVVmOqN.json"
      background="transparent"
      speed="1"
      style="width: 700px; height: 700px"
      loop
      autoplay
    ></dotlottie-player>
    <dotlottie-player
      class="farm"
      src="https://lottie.host/3e7040cd-15a2-4e35-8cd0-cccb93b01d6d/VSKlwF7xBH.json"
      background="transparent"
      speed="1"
      style="width: 700px; height: 700px"
      loop
      autoplay
    ></dotlottie-player>
  </div>
</template>
<script>
export default {
  name: "Icon-view",
};
</script>
<style scoped>
.farm {
  z-index: 10;
  position: absolute;
  top: -14rem;
}
.farmer {
  z-index: 13;
}
.cow {
  transform: scaleX(-1);
  z-index: 12;
  top: 0.92rem;
  left: 25rem;
}
</style>
