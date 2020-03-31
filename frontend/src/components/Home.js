import React, {Component} from "react";
import {Button, Container, Switch} from "react-bootstrap";
import ConnectionDay from "./ConnectionDay";
import "./Home.css";
import axios from "axios";
import {BrowserRouter, Route, Redirect} from "react-router-dom";
import CheckInPerson from "./CheckInPerson";
import CheckInLocation from "./CheckInLocation";
import Auth from "./Auth";
import Config from "./Config";


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

class Home extends Component {
    constructor(props) {
    super(props);
    this.handler = this.handler.bind(this);
    this.state = {
        loginRequired: false,

        events: [{
                  "type": "location",
                  "time": 1584910435,
                  "details": {
                    "location_id": "location-id-1",
                    "name": "Rewe",
                    "adress": "Berlin "
                  }
                },
                {
                  "type": "location",
                  "time": 1584914267,
                  "details": {
                    "location_id": "location-id-1",
                    "name": "Aldi",
                    "adress": "Frankfurt"
                  }
                },
                {
                  "type": "person",
                  "time": 1584754202,
                  "details": {
                    "users": [
                      {
                        "user_id": "user-id",
                        "user_name": "User regitered name",
                        "name": "Gustav"
                },
                ]}},
                {
                  "type": "person",
                  "time": 1584900000,
                  "details": {
                    "users": [
                      {
                        "user_id": "user-id",
                        "user_name": "User regitered name",
                        "name": "Martin"
                },
                ]}
                }],
        username: 'Angie',
        dates: getDateArray()
      }
    }

    handler(param) {
        this.setState({
            events: this.state.events.concat([param])}
        );
    }

     componentDidMount() {
        const headers = {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + Auth.getToken()
        };
        const data = {};
        const ep = Config.getApiHost() + '/api/get_events';
        axios.post(ep, data, {headers:headers})
        .then(res => res.data)
        .then((data) => {
          this.setState({ events: data['events'], username: data['user_name']});
        })
        .catch((ex) => {
            console.log(ex.response)
            const status = ex.response.status;
            if ( status === 401 || status === 400) {
                this.setState({ loginRequired: true });
            }
        })
      }

    render() {
        if ( this.state.loginRequired ) {
            return <Redirect to="/login" />
        }

        return (
            <Container>
                <BrowserRouter>
                    <Switch/>
                    <PrivateRoute exact path="/checkin_person">
                      <CheckInPerson handler={this.handler}/>
                  </PrivateRoute>
                  <PrivateRoute exact path="/checkin_location">
                      <CheckInLocation handler={this.handler}/>
                  </PrivateRoute>
                  <Route exact path="/">
                      <div>
                        <h1>Hallo {this.state.username}</h1>
                        <p>Hier kannst du deine letzten Kontakte und Orte eintragen.</p>
                        <div className="buttonGroup">
                            <Button size="lg" href="/checkin_person">Person</Button>
                            <Button size="lg" href="/checkin_location">Ort</Button>
                        </div>
                        {this.state.dates.map(element => <ConnectionDay key={element.toDateString()} date={element.toDateString()}
                           events={
                               this.state.events.filter(event => (new Date(event.time * 1000).getDate() === element.getDate()))}/>)}
                      </div>
                  </Route>
                <Switch/>
                </BrowserRouter>
            </Container>
        );
    }
}
export default Home;
