import React, {Component} from "react";
import {Button, Col, Row} from "react-bootstrap";
import "./Home.css";

class Home extends Component {
    render() {
        return (
            <div>
                <h1>My Connections</h1>
                <div className="buttonGroup">
                    <Button size="lg" href="/checkin_person">Person</Button>
                    <Button size="lg" href="/checkin_location">Ort</Button>
                </div>
                <div className="connectionDay">
                    <p>Today</p>
                    <div className="container">
                        <div className="connectionItem">
                            <p>(IC)</p>
                            <div>
                                <p>Claudia, Gustav, + 3 others</p>
                                <p>13:05 * Glogauer Str. 9</p>
                            </div>
                            <p >(PF)</p>
                        </div>
                        <div className="connectionItem">
                            <p>(IC)</p>
                            <div>
                                <p>Claudia, Gustav, + 3 others</p>
                                <p>13:05 * Glogauer Str. 9</p>
                            </div>
                            <p >(PF)</p>
                        </div>
                    </div>
                </div>
                <div className="connectionDay">
                    <p>Yesterday</p>
                    <div className="container">
                    <div className="connectionItem">
                        <p>(IC)</p>
                        <div>
                            <p>Claudia, Gustav, + 3 others</p>
                            <p>13:05 * Glogauer Str. 9</p>
                        </div>
                        <p >(PF)</p>
                    </div>
                    <div className="connectionItem">
                        <p>(IC)</p>
                        <div>
                            <p>Claudia, Gustav, + 3 others</p>
                            <p>13:05 * Glogauer Str. 9</p>
                        </div>
                        <p >(PF)</p>
                    </div>
                    </div>
                </div>
                <div className="connectionDay">
                    <p>18. März 2020</p>
                    <div className="container">
                    <div className="connectionItem">
                        <p>(IC)</p>
                        <div>
                            <p>Claudia, Gustav, + 3 others</p>
                            <p>13:05 * Glogauer Str. 9</p>
                        </div>
                        <p >(PF)</p>
                    </div>
                    </div>
                </div>
            </div>
        );
    }
}
export default Home;