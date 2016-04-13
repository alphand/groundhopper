<template>
  <div class="login-panel">
    <div class="row align-center">
      <div class="medium-6 large-4 column">
        <validator name="loginValidator">
          <form class="log-in-form"
            v-on:submit.prevent="onSubmit">
            <h4>Log in with your email!</h4>
            <div class="form-group">
              <label for="email">
                Email address:
                <input type="text" id="email" name="email"
                  v-model="email"
                  v-validate:username="{required:true}"/>
              </label>
            </div>
            <div class="form-group">
              <label for="password" >
                Password:
                <input type="password" id="password" name="password"
                  v-model="password"
                  v-bind:type="(isPassShown)? 'text':'password'"
                  v-validate:password="{required:true,minLength:6}"/>
              </label>
              <p class="help-text">
                Your password must have at least 6 characters
              </p>
            </div>
            <div class="form-group">
              <input type="checkbox" id="show-password" v-model="isPassShown" />
              <label for="show-password">
                show password
              </label>
            </div>
            <div class="form-actions">
              <div class="errors">
                <p class="form-error"
                  v-if="$loginValidator.username.dirty &&
                  $loginValidator.username.required">
                  Email is required
                </p>
              </div>
              <button type="submit" class="button expanded">Submit</button>
            </div>
            <div class="form-group">
              <a href="">Forgot your password?</a>
            </div>
          </form>
        </validator>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      email: '',
      password: '',
      isPassShown: false
    }
  },

  methods: {
    onSubmit () {
      console.log('submitting-data', this.email, this.password, this.$loginValidator)
    }
  }
}
</script>

<style>
  .log-in-form {
    margin-top: 30px;
    border: 1px solid #cacaca;
    padding: 1rem;
    border-radius: 3px;
  }
</style>
