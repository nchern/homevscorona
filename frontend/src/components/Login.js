import React, {Component} from "react";
import GoogleLogin from "react-google-login";
import {Redirect} from "react-router-dom";
import axios from "axios";

export default class Login extends Component {
    state = {
        redirect: false
    };
    responseGoogle = response => {
        console.log(response);
        localStorage.setItem('ggToken', response.tokenObj.id_token);
        this.setState({redirect:true});
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer' + response.tokenObj.id_token
        };
        const data = {
            'provider': 'google',
            'name': response.profileObj.name
        };
        axios.post("http://homevscorona.us.to/api/signup", data,
            {headers:headers})
        .then(res => res.json())
        .then((data) => {
          console.log(data);
        })
        .catch(console.log)
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