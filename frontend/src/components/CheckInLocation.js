import React, {Component} from "react";
import {Container, Form} from "react-bootstrap";
import {Button} from "react-bootstrap";

class CheckInLocation extends Component {
    render() {
        return (
            <Container>
                <h1>Wo bist du?</h1>
                <Form>
                  <Form.Group >
                    <Form.Label>Name des Ortes</Form.Label>
                    <Form.Control placeholder="Rewe an der Hauptwache" />
                  </Form.Group>
                  <Form.Group >
                    <Form.Label>Adresse</Form.Label>
                    <Form.Control type="adress" placeholder="Zeil 106-110, 60313 Frankfurt am Main" />
                  </Form.Group>
                  <Form.Group controlId="formBasicCheckbox">
                    <Form.Check type="checkbox" label="Check me out" />
                  </Form.Group>
                  <Button variant="primary" type="submit">
                    Eintragen
                  </Button>
                </Form>
                <br/>
                <p>Deine letzten Orte</p>

            </Container>
        );
    }
}
export default CheckInLocation;