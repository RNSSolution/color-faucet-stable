<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://color-platform.rnssol.com/"
      target="_blank"
      title="Color Wallet"
      subtitle="color-platform.rnssol.com")
    list-item(
      href="https://color-platform.rnssol.com:3000/validators"
      target="_blank"
      title="Validator Chat"
      subtitle="#color_validators")
    list-item(
      href="https://color-platform.rnssol.com:3000/"
      target="_blank"
      title="Testnet Explorer"
      subtitle="color-platform.rnssol.com:3000")
    list-item(
      href="https://color-platform.org/"
      target="_blank"
      title="Color Website"
      subtitle="color-platform.org")
</template>

<script>
import ListItem from "../components/NiListItem.vue";
export default {
  name: "faucet",
  components: {
    ListItem
  },
  data: () => ({}),
  methods: {}
};
</script>
