<template lang="pug">
#app
  tm-cookie-consent
  .top-bar
    .left
      img(src="~assets/ColorWhiteplatform1.png" height="30")
      .key testnet:
      .value {{ config.chain }}
    .right
      a(href="https://antliafinalexplorer.herokuapp.com/" target="_blank") Explorer #[i.material-icons assessment]
  router-view
  notifications(:notifications='notifications' theme='color' )
  #bottom
      h3 All rights  reserved &copy; 2019
        p Color Network
</template>


<script>
import { mapGetters } from "vuex";
import { TmCookieConsent } from "@tendermint/ui";
import Notifications from "@nylira/vue-notifications";

export default {
  name: "app",
  components: {
    Notifications,
    TmCookieConsent
  },
  computed: {
    ...mapGetters(["notifications", "config"])
  },
  mounted() {}
};
</script>

<style lang="stylus" src="./styles/app.styl"></style>

<style lang="stylus">
@import '~variables'

img 
  margin 7px !important

.top-bar
  background-image url("header.jpg")
  background-size 750px
  position fixed
  top 0
  right 0
  text-align center
  height 4rem
  width 100%
  z-index 100

  display flex
  justify-content space-between
  line-height 3rem

  .left
    display flex
    padding 0 1rem
  .key
    text-transform: lowercase
    margin-right 0.5rem
    margin-left 30px
  .value
    font-weight bold
    color var(--bright)
  .right
    a
      display block
      background #32c2fbad
      padding 0 1rem
      color var(--bright)
      display flex
      align-items center
      i
        margin-left 0.5rem
#bottom
  background rgb(21, 24, 45)
  color white
  font-size 0.75rem
  padding 1rem 0

  h3
    padding 0 1rem
    p
      float right
      display flex
      padding 0 1rem
      align-items center
</style>
