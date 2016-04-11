<template>
  <div>
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
  import topbar from '../components/topbar'
  export default {
    components: {
      topbar: topbar
    },

    data () {
      return {
        msg: 'this is from homepage'
      }
    }
  }
</script>
