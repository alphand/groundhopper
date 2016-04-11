<template>
  <div>
    <ul>
      <li>
        <a v-link="{path:'/login'}">Go to Login</a>
      </li>
      <li>
        <a v-link="{path:'/register'}">Go to Register</a>
      </li>
      <li>
        <a v-link="{path:'/home'}">Go to Home</a>
      </li>
    </ul>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        msg: 'testing'
      }
    }
  }
</script>
