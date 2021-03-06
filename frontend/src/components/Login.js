import React, {Component} from "react";
import GoogleLogin from "react-google-login";
import {Redirect} from "react-router-dom";
import {Container, Alert} from "react-bootstrap";
import axios from "axios";
import Config from "./Config";

export default class Login extends Component {
    state = {
        redirect: false,
        showAlert: false
    };
    responseGoogle = response => {
        localStorage.setItem('ggToken', response.tokenObj.id_token);
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + response.tokenObj.id_token
        };
        const data = {
            'provider': 'google',
            'name': response.profileObj.name
        };
        const ep = Config.getApiHost() + "/api/signup";
        axios.post(ep, data, {headers:headers})
        .then(res => res.data)
        .then((data) => {
            console.log(data);
            this.setState({redirect:true});
        })
        .catch((ex) => {
            console.log(ex);
            this.setState({showAlert: true});
        })
      };

  render () {
      const redirect = this.state.redirect;
      if (redirect === true) {
          return <Redirect to="/"/>
      } else {
  return (
      <Container>
        <Alert variant="danger" show={ this.state.showAlert }>
            Auth failed!
        </Alert>
        <div className="Login">
            <h2>Bitte melde dich an!</h2>
            <p>Hi, wir freuen uns, dass du Interesse daran hast, auch ein Corona-Scout zu werden.</p>
            <p>Bitte melde dich über mit deinem Google-Konto an. Danach kann es sofort losgehen.</p>
            <p>Da dies ein Prototyp ist, der während des #WirVsVirus-Hackathon entstanden ist,
               können leider noch nicht alle Einträge gespeichert werden.  </p>
            <p>Lass uns wissen, wenn dir unser Konzept gefällt, wir arbeiten mit Hochdruck am fertigen Produkt!</p>
            <p>Vielen Dank für euer Verständnis und viel Spaß beim Ausprobieren!</p>
            <GoogleLogin
                    clientId="328361320618-g6bo1k25hqnnbngu6145u17lkaj383fd.apps.googleusercontent.com"
                    buttonText="Login"
                    onSuccess={this.responseGoogle}
                    onFailure={this.responseGoogle}
                    cookiePolicy={'single_host_origin'}
                />
        </div>
      </Container>
        )
  }}}
