<template>
   <div id="app">
    <topbar></topbar>
    <router-view>
    </router-view>
  </div>

</template>

<script>
import store from './store'
import Topbar from './components/topbar'
export default {
  components: {
    Topbar
  },
  store
}
</script>

<style lang="scss">
@import "./app.scss";
</style>
