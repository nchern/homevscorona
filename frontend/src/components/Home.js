import React, {Component} from "react";
import {Button, Container} from "react-bootstrap";
import ConnectionDay from "./ConnectionDay";
import "./Home.css";
import axios from "axios";

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
         axios.post('http://homevscorona.us.to/api/get_events')
        .then(res => res.json())
        .then((data) => {
          console.log(data);
          this.setState({ events: data['events'],
                                username: data['user_name']})
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
                {this.state.dates.map(element => <ConnectionDay key={element.toDateString()} date={element.toDateString()} events={
                    this.state.events.filter(event => event.time === (element.getTime() / 1000))}/>)}
            </Container>
        );
    }
}
export default Home;