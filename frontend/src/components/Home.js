import React, {Component, useState} from "react";
import {Button, Container, Modal} from "react-bootstrap";
import ConnectionDay from "./ConnectionDay";
import "./Home.css";

var getDateArray = function() {
    const end = new Date();
    const start = new Date(Date.now() - 14 * 24 * 60 * 60 * 1000);
  var
    arr = new Array(),
    dt = new Date(start);

  while (dt <= end) {
    arr.push(new Date(dt));
    dt.setDate(dt.getDate() + 1);
  }

  return arr.reverse();

};

class Home extends Component {
    state = {
        events: [],
        username: '',
        dates: getDateArray()
      };

     componentDidMount() {
        fetch('https://6ca70b10-5803-40af-af91-41d37a0d9133.mock.pstmn.io/events')
        .then(res => res.json())
        .then((data) => {
          this.setState({ events: data['events'],
                                username: data['username']})
        })
        .catch(console.log)
      }

    render() {
        return (
            <Container>
                <h1>Hallo {this.state.username}</h1>
                <p>Hier kannst du deine letzten Kontakte und Orte eintragen.</p>
                <div className="buttonGroup">
                    <Button size="lg" href="/checkin_person">Person</Button>
                    <Button size="lg" href="/checkin_location">Ort</Button>
                </div>
                {this.state.dates.map(element => <ConnectionDay date={element.toDateString()} events={"TestEvents"}/>)}
            </Container>
        );
    }
}
export default Home;