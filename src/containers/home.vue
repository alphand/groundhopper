<template>
  <div class="hp-header">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        msg: 'this is from homepage'
      }
    }
  }
</script>

<style scoped>
  .hp-header {
    font-size:20px;
  }
</style>
