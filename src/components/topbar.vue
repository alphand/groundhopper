<template>
  <div class="top-bar">

    <div class="top-bar-left">
      <ul class="menu">
        <li class="menu-text">
          <a v-link="{path:'/'}">
            <img class="logo" src="../assets/logo.png" alt="GlassPilot">
          </a>
          GlassPilot Gate
        </li>
      </ul>
    </div>

    <div class="top-bar-right">
      <ul class="menu menu-right">
        <li><a v-link="{path:'/accounts/register'}">Register</a></li>
        <li><a v-link="{path:'/accounts/login'}">Login</a></li>
      </ul>
    </div>
  </div>
</template>

<style lang="scss">

  .logo {
    height:50px;
    width:auto;
  }

  .menu-right{
    justify-content:flex-end;
  }
</style>
