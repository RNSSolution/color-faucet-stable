const dev = process.env.NODE_ENV === `development`;
const VUE_APP_CHAIN =
  process.env.VUE_APP_CHAIN || (dev ? `rns-test-01` : `colors-test-01`);

const VUE_APP_CLAIM_URL =
  process.env.VUE_APP_CLAIM_URL ||
  (dev
    ? `https://proxy.testnet.color-platform.rnssol.com:8000/claim`
    : `https://proxy.testnet.color-platform.org:9000/claim`);

const state = {
  chain: VUE_APP_CHAIN,
  claimUrl: VUE_APP_CLAIM_URL,
  recaptchaSiteKey: process.env.VUE_APP_RECAPTCHA_SITE_KEY
};

const mutations = {};

export default {
  state,
  mutations
};
