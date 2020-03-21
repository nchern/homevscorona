import React, {Component} from "react";
import {Form} from "react-bootstrap";

class CheckInLocation extends Component {
    render() {
        return (
            <div>
                <h1>Wer ist bei dir?</h1>
                <Form>
                    <Form.Control placeholder="Suche nach Personen..."/>
                </Form>
                <p>Favoriten</p>

            </div>
        );
    }
}
export default CheckInLocation;