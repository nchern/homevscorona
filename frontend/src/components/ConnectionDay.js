import React, {Component} from "react";
import ConnectionItem from "./ConnectionItem";

class ConnectionDay extends Component {
    render() {
        return (
            <div className="connectionDay">
                    <p>{this.props.date}</p>
                    <div className="container">
                    {/*{this.props.events.map(element => <ConnectionItem event={element}/>)}*/}
                    {/*if this.props.events is empty: event: {"type":"athome"}*/}
                    <ConnectionItem event={"TestEvent"}/>
                    </div>
            </div>
        )
    }
}

export default ConnectionDay;