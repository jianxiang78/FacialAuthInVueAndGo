<template>
  <section class="login_content">
    <form @submit.prevent="userLogin" id="form-login">
      <div class="head">
        <h1>Web APP Login</h1>
      </div>
      <div class="content">
        <input type="text" class="form-control"
               placeholder="LoginName" required
               name="loginName" id="loginName" />
      </div>
      <div class="content">
        <input type="password" class="form-control"
               placeholder="Password" required
               name="passWord" id="passWord"
        />
      </div>
      <div class="footer">
        <button type="submit" class="btn btn-primary" style="width: 100%">Name And Password Login</button>
      </div>
      <div class="footer">
        <button type="button" class="btn btn-primary" style="width: 100%" @click="faceLogin">Face Recognition Login</button>
      </div>

      <div class="footer">
        <router-link to="/dashboard">
          <button type="button" class="btn btn-primary" style="width: 100%" >Go Directly To The Dashboard</button>
        </router-link>
      </div>
      <div class="footer">
        Facial Auth Demo In Vue.Js & Golang
      </div>
    </form>
  </section>
</template>

<script>

  export default {
    name: 'userLogin',
    methods: {
      userLogin: function() {
        const that=this
        let data={}
        data.loginName=document.getElementById("loginName").value
        data.passWord=document.getElementById("passWord").value
        this.AjaxPost("login",data,loginResult)
        return false;
        function loginResult(Result){
          if(Result.Code===1){
            sessionStorage.setItem("myAccessToken",Result.Result.accessToken)
            that.$router.push('/dashboard')
          }else {
            alert("login failed")
          }
        }
      },
      faceLogin: function() {
        const that=this
        function loginResult(Result) {
          if(Result.Code===1){
            sessionStorage.setItem("myAccessToken",Result.Result.accessToken)
            that.$router.push('/dashboard')
          }else {
            alert("login failed")
          }
        }
        function checkFaceLogin(userInfo) {
            console.log(userInfo)
            let data={}
            data.FacialId=userInfo.facialId
            that.AjaxPost("loginWithFace",data,loginResult)
          }
          window.myFaceIO.authenticate({
              "locale": "auto"
          }).then(userInfo => {
              checkFaceLogin(userInfo);
          }).catch(errCode => {
              console.log(errCode);
              window.myFaceIO.restartSession();
          })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .login_content {
    margin: 0px auto;
    max-width: 400px;
    background-color: #efefef
  }
  .login_content .head{
    padding-top: 10px;
    background-color: #efefef;
    text-align: center;
    margin: 10px;
  }
  .login_content .content{
    margin: 20px;
  }
  .login_content .footer{
    margin: 20px;
    text-align: center;
    padding-bottom: 10px;
  }
</style>
