import React from "react";
import {Container} from "react-bootstrap";
export const Layout = (props) => (
    <Container style={{paddingTop:"10px"}}>
        {props.children}
    </Container>
);
