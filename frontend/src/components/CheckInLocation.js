import React, {Component} from "react";
import {Container, Form} from "react-bootstrap";
import {Button} from "react-bootstrap";
import {Redirect} from "react-router-dom";
import "./CheckIn.css";

class CheckInLocation extends Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            adress: '',
            datetime: new Date().toUTCString(),
            redirect: false};
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
      }

    handleChange(event) {
        const fieldName = event.target.id;
        var obj = {};
        obj[fieldName] = event.target.value;
        this.setState(obj);
    }

    handleSubmit(event) {
        var obj = {type: "location"};
        obj["time"] = Date.parse(this.state.datetime);
        obj["details"] = {"location_id": new Date().getTime(),
            "name": this.state.name,
            "adress": this.state.adress

        };
        this.setState({redirect:true});
        this.props.handler(obj);
        event.preventDefault();
    }
    render() {
        const now = new Date();
        if (this.state.redirect) {
            return <Redirect to="/"/>
        } else {
        return (
            <Container>
                <h1>Wo bist du?</h1>
                <Form onSubmit={this.handleSubmit}>
                  <Form.Group controlId="name">
                    <Form.Label>Name des Ortes</Form.Label>
                    <Form.Control onChange={this.handleChange} placeholder="Rewe an der Hauptwache" />
                  </Form.Group>
                  <Form.Group controlId="adress">
                    <Form.Label>Adresse</Form.Label>
                    <Form.Control onChange={this.handleChange}  placeholder="Zeil 106-110, 60313 Frankfurt am Main" />
                  </Form.Group>
                  <Form.Group controlId="datetime">
                        <Form.Label>Zeitpunkt</Form.Label>
                        <Form.Control onChange={this.handleChange} defaultValue={now.toUTCString()} />
                        <Form.Text className="text-muted">
                          Falls dein Kontakt auch Corona-Scout ist, gib bitte hier seine Email ein.
                        </Form.Text>
                  </Form.Group>
                    <div className="buttonGroup">
                  <Button variant="primary" type="submit">
                    Eintragen
                  </Button>
                   <Button variant="secondary" onClick={() => this.setState({redirect:true})}>Zurück</Button>
                    </div>
                </Form>
            </Container>
        );
    }
    }
}
export default CheckInLocation;