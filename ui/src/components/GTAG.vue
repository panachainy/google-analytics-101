<template>
  <div class="hello">
    <button id="gtm-test-btn-1" v-on:click="gtmSend">GTM TEST 1</button>
    <button id="gtm-test-btn-2">GTM TEST 2</button>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",

  methods: {
    gtmSend: function() {
      console.log("gtmSend ~");
      console.log("this.$gtm.enabled() ->", this.$gtm.enabled());

      this.$gtm.trackEvent({
        event: "event", // Event type [default = 'interaction'] (Optional)
        category: "Calculator",
        action: "click",
        label: "Home page SIP calculator",
        value: 5000,
        noninteraction: false, // Optional
      });
    },
  },
};
</script>
