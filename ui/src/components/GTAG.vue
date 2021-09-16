<template>
  <div class="hello">
    <button id="gtm-test-btn" v-on:click="gtmSend">GTM TEST</button>
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
        event: 'event', // Event type [default = 'interaction'] (Optional)
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
