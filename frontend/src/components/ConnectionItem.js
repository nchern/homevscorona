import React, {Component} from "react";
import {FiChevronRight, FiUser, FiSmile, FiMapPin} from "react-icons/all";

export default class ConnectionItem extends Component {
    render() {
        if (this.props.event.type === "location") {
            return (<div className="connectionItem">
                <FiMapPin className="itemIcon"/>
                <div className="itemText">
                    <p><b>{this.props.event.details.name}</b></p>
                    <p>{this.props.event.time}・{this.props.event.details.adress}</p>
                </div>
                <div className="itemGoto"><FiChevronRight/></div>
            </div>)
        } else if (this.props.event.type === "person") {
            return (<div className="connectionItem">
                    <FiUser className="itemIcon"/>
                    <div className="itemText">
                        <p><b>Claudia, Gustav, + 3 others</b></p>
                        <p>13:05・Glogauer Str. 9</p>
                    </div>
                    <div className="itemGoto"><FiChevronRight/></div>
                </div>
            )
        } else {
            return (<div className="connectionItem">
                    <FiSmile className="itemIcon"/>
                    <div className="itemText">
                        <p><b>Du hast den ganzen Tag zuhause verbracht!</b></p>
                        <p></p>
                    </div>
                    <div className="itemGoto"><FiChevronRight/></div>
                </div>
            )
        }
    }
}