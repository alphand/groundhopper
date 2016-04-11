<template>
  <div class="login-panel">
    <div class="row">
      <div class="medium-6 medium-centered large-4 large-centered columns">
        <form action="" class="row column log-in-form">
          <h4>Log in with your email!</h4>
          <div class="form-group">
            <label for="email">Email address:</label>
            <input type="text" id="email" name="email" />
          </div>
          <div class="form-group">
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" />
          </div>
          <div class="form-group">
            <input type="checkbox" id="show-password">
            <label for="show-password">show password</label>
          </div>
          <div class="form-actions">
            <button type="submit" class="button expanded">Submit</button>
          </div>
          <div class="form-group">
            <a href="">Forgot your password?</a>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style>
  .log-in-form {
    border: 1px solid #cacaca;
    padding: 1rem;
    border-radius: 3px;
  }
</style>
