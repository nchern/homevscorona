import React, {Component} from "react";
import {Button, Container, Form} from "react-bootstrap";
import "./CheckIn.css";
import {Redirect} from "react-router-dom";

class CheckInPerson extends Component {
    state = {
        redirect: false
    };
    render() {
        const now = new Date();
        if (this.state.redirect) {
            return <Redirect to="/"/>
        } else {
        return (
            <Container>
                <h1>Wer ist bei dir?</h1>
                <Form>
                  <Form.Group controlId="formName">
                    <Form.Label>Name</Form.Label>
                    <Form.Control placeholder="Max Mustermann" />
                  </Form.Group>
                    <Form.Group controlId="formBasicEmail">
                        <Form.Label>Email-Adresse</Form.Label>
                        <Form.Control type="email" placeholder="max@muster.de"/>
                        <Form.Text className="text-muted">
                          Falls dein Kontakt auch ein Corona-Scout ist, können wir im das Treffen auch zuweisen.
                        </Form.Text>
                      </Form.Group>
                    <Form.Group controlId="formDateTime">
                        <Form.Label>Zeitpunkt</Form.Label>
                        <Form.Control defaultValue={now.toUTCString()} />
                        <Form.Text className="text-muted">
                          Falls dein Kontakt auch Corona-Scout ist, gib bitte hier seine Email ein.
                        </Form.Text>
                      </Form.Group>
                    <div className="buttonGroup">
                        <Button variant="primary" type="submit" onClick={() => this.setState({redirect:true})}>
                         Eintragen </Button>
                        <Button variant="secondary" onClick={() => this.setState({redirect:true})}>Zurück</Button>
                    </div>
                </Form>

            </Container>
        );
    }
    }
}
export default CheckInPerson;