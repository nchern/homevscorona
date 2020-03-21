import React, {Component} from "react";
import {Form} from "react-bootstrap";

class CheckInPerson extends Component {
    render() {
        return (
            <div>
                <h1>Wo bist du?</h1>
                <Form>
                    <Form.Control placeholder="Suche nach Orten..."/>
                </Form>
                <p>Zuletzt</p>

            </div>
        );
    }
}
export default CheckInPerson;