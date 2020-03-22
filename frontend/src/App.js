import React, {Component} from 'react';
import {Layout} from './components/Layout';
import Login from './components/Login';
import Home from "./components/Home";
import {BrowserRouter as Router, Switch, Route, Redirect} from "react-router-dom";
import './App.css';
import Auth from "./components/Auth";

function PrivateRoute({ children, ...rest }) {
  return (
    <Route
      {...rest}
      render={({ location }) =>
        Auth.isAuthenticated() ? (
          children
        ) : (
          <Redirect
            to={{
              pathname: "/login",
              state: { from: location }
            }}
          />
        )
      }
    />
  );
}

export default class App extends Component {
    render() {
    return (
        <div className="App">
            <Router>
            <Layout>
                <Switch>
                  <Route path="/login">
                  <Login />
                  </Route>
                  <PrivateRoute path="/">
                      <Home/>
                  </PrivateRoute>
                  <Route path="*" component={NoMatch}/>
                </Switch>
            </Layout>
            </Router>
        </div>
      );
    }

}
function NoMatch() {
    return <h2>Seite nicht gefunden.</h2>

}