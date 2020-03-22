import React, {Component} from "react";
import {BrowserRouter as Router, Switch, Route, Link} from "react-router-dom";
import Home from "./Home";
import CheckInLocation from "./CheckInLocation";
import CheckInPerson from "./CheckInPerson";

export default class Main extends Component {
    render() {
    return (<Router>
                   <Link href="/" component={Home}/>
                   <Link href="/about" component={About}/>
                   <Link href="/checkin" component={CheckIn}/>
                   <Link href="/checkin_person" component={CheckInLocation}/>
                   <Link href="/checkin_location" component={CheckInPerson}/>
                   <Link href="*" component={() => "404 NOT FOUND"} />
              </Router> )
}
    }

        function About() {
  return <h1>About</h1>;
}

function CheckIn() {
    return <h1>Check-In</h1>
}

function NoMatch() {
    return <h1>No Match</h1>
}


