import React, {Component} from "react";
import GoogleLogin from "react-google-login";
import {Redirect} from "react-router-dom";

export default class Login extends Component {
    state = {
        redirect: false
    };
    responseGoogle = response => {
        localStorage.setItem('ggToken', response.tokenObj.id_token);
        this.setState({redirect:true});
  };
  render () {
      const redirect = this.state.redirect;
      if (redirect === true) {
          return <Redirect to="/"/>
      } else {
  return (
    <div className="Login">
        <h2>Bitte melde dich an!</h2>
        <GoogleLogin
                clientId="328361320618-g6bo1k25hqnnbngu6145u17lkaj383fd.apps.googleusercontent.com"
                buttonText="Login"
                onSuccess={this.responseGoogle}
                onFailure={this.responseGoogle}
                cookiePolicy={'single_host_origin'}
            />
    </div>
        )
  }}}