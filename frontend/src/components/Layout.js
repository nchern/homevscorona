import React from "react";
import {Container} from "react-bootstrap";
export const Layout = (props) => (
    <Container style={{paddingTop:"30px"}}>
        {props.children}
    </Container>
);
