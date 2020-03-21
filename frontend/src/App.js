import React from 'react';
import {BrowserRouter as Router, Switch, Route } from "react-router-dom";
import {Nav, Navbar} from "react-bootstrap";
import {Layout} from './components/Layout';
import Login from './components/Login';
import Home from "./components/Home";
import CheckInLocation from "./components/CheckInLocation";
import CheckInPerson from "./components/CheckInPerson";

import './App.css';

function App() {
  return (
    <div className="App">
    <Layout>
    <Router>
          <div>
            <Navbar>
                <Nav.Link href="/">Home</Nav.Link>
                <Nav.Link href="/about">About</Nav.Link>
                <Nav.Link href="/login">Login</Nav.Link>
                <Nav.Link href="/checkin">CheckIn</Nav.Link>
            </Navbar>
            <br/>
            <Switch>
              <Route exact path="/about" component={About}/>
              <Route exact path="/login" component={Login}/>
              <Route exact path="/checkin" component={CheckIn}/>
              <Route exact path="/checkin_person" component={CheckInLocation}/>
              <Route exact path="/checkin_location" component={CheckInPerson}/>
              <Route exact path="/" component={Home}/>
              <Route path="*" component={NoMatch}/>
            </Switch>
          </div>
        </Router>
    </Layout>
    </div>
  );
}

export default App;


function About() {
  return <h1>About</h1>;
}

function CheckIn() {
    return <h1>Check-In</h1>
}

function NoMatch() {
    return <h1>No Match</h1>
}