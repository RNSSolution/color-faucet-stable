<template lang="pug">
.section
  .section-header
    h1 Don't have a testnet address?
    p There's two ways to get one. The first is by using Color Wallet, the crypto wallet for Color. If you know command-line, you can also generate an address with the Color SDK.
  .section-main
    btn(type="anchor" href="https://rnssolution.github.io/colorprojectdocs/joinTestNet/" value="Join the latest testnet" size="lg" target="_blank" icon="wifi_tethering")
</template>

<script>
import Btn from "@nylira/vue-button";
export default {
  name: "section-join",
  components: {
    Btn
  }
};
</script>
