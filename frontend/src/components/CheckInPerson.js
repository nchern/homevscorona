import React, {Component} from "react";
import {Form} from "react-bootstrap";

class CheckInPerson extends Component {
    render() {
        return (
            <div>
                <h1>Mit wem bist du unterwegs?</h1>
                <Form>
                    <Form.Control placeholder="Suche nach Fruenden..."/>
                </Form>
                <p>Zuletzt</p>

            </div>
        );
    }
}
export default CheckInPerson;