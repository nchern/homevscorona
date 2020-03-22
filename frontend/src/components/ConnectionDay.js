import React, {Component} from "react";
import ConnectionItem from "./ConnectionItem";

class ConnectionDay extends Component {

    render() {
        const no_events = this.props.events.length;
        if (no_events === 0) {
            return (<div className="connectionDay">
                <p>{this.props.date}</p>
                <ConnectionItem event={"noEvent"}/>
            </div>)
        } else {
        return (
            <div className="connectionDay">
                    <p>{this.props.date}</p>
                    <div className="container">
                    {this.props.events.map(element => <ConnectionItem event={element}/>)}
                    </div>
            </div>
        )
            }
    }
}

export default ConnectionDay;