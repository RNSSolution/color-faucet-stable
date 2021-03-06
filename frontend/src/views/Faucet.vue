<template lang="pug">
#faucet
  .section
    faucet-header
    form(v-on:submit.prevent='onSubmit', method='post')
      form-group(:error='$v.fields.response.$error'
        field-id='faucet-response' field-label='Captcha')
        vue-recaptcha#faucet-response(
          ref="recaptcha"
          @verify="onVerify"
          @expired="onExpired"
          :sitekey="config.recaptchaSiteKey")
        form-msg(name='Captcha' type='required' v-if='!$v.fields.response.required')
      form-group(:error='$v.fields.address.$error'
        field-id='faucet-address' field-label='Send To')
        field#faucet-address(
          type='text'
          v-model='fields.address'
          placeholder='Testnet address'
          size="lg")
        form-msg(name='Address' type='required' v-if='!$v.fields.address.required')
        form-msg(name='Address' type='bech32' :body="bech32error" v-else-if='!$v.fields.address.bech32Validate')
      form-group
        btn(v-if='sending' value='Sending...' disabled color="#15182d" size="lg")
        btn(v-else @click='onSubmit' value="Send me tokens" color="#15182d" size="lg" icon="send")
  section-join
  section-links
</template>

<script>
import axios from "axios";
import VueRecaptcha from "vue-recaptcha";
import { mapGetters } from "vuex";
import { required } from "vuelidate/lib/validators";
import b32 from "../scripts/b32";
import Btn from "@nylira/vue-button";
import Field from "@nylira/vue-field";
import FormGroup from "../components/NiFormGroup";
import FormMsg from "../components/NiFormMsg";
import FaucetHeader from "../components/FaucetHeader";
import SectionJoin from "../components/SectionJoin.vue";
import SectionLinks from "../components/SectionLinks.vue";
export default {
  name: "faucet",
  components: {
    Btn,
    Field,
    FormGroup,
    FaucetHeader,
    FormMsg,
    SectionJoin,
    SectionLinks,
    VueRecaptcha
  },
  computed: {
    ...mapGetters(["config"])
  },
  data() {
    return {
      fields: {
        address: "",
        response: ""
      },
      sending: false
    };
  },
  methods: {
    resetForm() {
      this.fields.address = "";
      this.fields.response = "";
      this.$refs.recaptcha.reset();
      this.$v.$reset();
    },
    onVerify(response) {
      this.fields.response = response;
    },
    onExpired() {
      this.$refs.recaptcha.reset();
    },
    async onSubmit() {
      this.$v.$touch();
      if (this.$v.$error) return;

      var datas = JSON.stringify({
        address: this.fields.address,
        response: this.fields.response
      });

      this.sending = true;
      axios
        .post(this.config.claimUrl, datas)
        .then(() => {
          this.sending = false;
          this.$store.commit("notify", {
            title: "Successfully Sent",
            body: `Sent tokens to ${this.fields.address}`
          });
          this.resetForm();
        })
        .catch(err => {
          this.sending = false;
          this.$store.commit("notifyError", {
            title: "Error Sending",
            body: `An error occurred while trying to send: "${err.message}"`
          });
        });
    },
    bech32Validate(param) {
      try {
        b32.decode(param);
        this.bech32error = null;
        return true;
      } catch (error) {
        this.bech32error = error.message;
        return false;
      }
    }
  },
  validations() {
    return {
      fields: {
        address: {
          required,
          bech32Validate: this.bech32Validate
        },
        response: { required }
      }
    };
  }
};
</script>

<style lang="stylus">
@import '~variables'

#faucet
  max-width 40rem
  width 100%
  margin 0 auto

.section
  margin 0.5rem
  padding 1rem
  background #fdfafaf7
  box-shadow 0 3px 1px -2px rgba(0, 0, 0, .2), 0 2px 2px 0 rgba(0, 0, 0, .14), 0 1px 5px 0 rgba(0, 0, 0, .12)
  position relative
  z-index 10
  label
    display none

  input:-webkit-autofill
    -webkit-text-fill-color black !important

  .section-main
    padding 0 1rem

@media screen and (min-width: 375px)
  .section
    padding 2rem 1rem

@media screen and (min-width: 768px)
  .section
    padding 3rem 2rem
    margin 1rem
</style>
