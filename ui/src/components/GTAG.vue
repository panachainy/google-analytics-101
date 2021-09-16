<template>
  <div class="hello">
    <button v-on:click="gtmSend">GTM TEST</button>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",

  methods: {
    gtmSend: function() {
      console.log("gtmSend ~");
    },
  },
};
</script>
