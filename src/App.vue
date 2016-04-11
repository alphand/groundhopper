<template>
   <div id="app">
    <topbar></topbar>
    <router-view>
    </router-view>
  </div>

</template>

<script>
import Topbar from './components/topbar'
export default {
  components: {
    Topbar
  }
}
</script>

<style lang="scss">
@import "./app.scss";
</style>
