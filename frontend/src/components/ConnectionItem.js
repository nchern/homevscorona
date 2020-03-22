import React, {Component} from "react";
import {FiChevronRight, FiUser, FiSmile, FiMapPin} from "react-icons/all";

class ConnectionItem extends Component {
    render() {
        return (
            <div className="TMP">
            <div className="connectionItem">
                {/*if this.props.event.type == "person"*/}
                <FiUser className="itemIcon"/>
                <div className="itemText">
                    <p><b>Claudia, Gustav, + 3 others</b></p>
                    <p>13:05・Glogauer Str. 9</p>
                </div>
                <div className="itemGoto"><FiChevronRight/></div>
            </div>
            <div className="connectionItem">
                {/*if this.props.event.type == "location"*/}
                <FiMapPin className="itemIcon"/>
                <div className="itemText">
                    <p><b>Kater Blau</b></p>
                    <p>13:05・Glogauer Str. 9</p>
                </div>
                <div className="itemGoto"><FiChevronRight/></div>
            </div>
            <div className="connectionItem">
                {/*if this.props.event.type == "athome"*/}
                {/* {this.props.type}*/}
                {/*<p>(IC)</p>*/}
                {/*<div>*/}
                {/*    <p><b>Claudia, Gustav, + 3 others</b></p>*/}
                {/*    <p>13:05 * Glogauer Str. 9</p>*/}
                {/*</div>*/}
                <FiSmile className="itemIcon"/>
                <div className="itemText">
                    <p><b>Du hast den ganzen Tag zuhause verbracht!</b></p>
                    <p> </p>
                </div>
                <div className="itemGoto"><FiChevronRight/></div>
            </div>
            </div>
        )
    }
}

export default ConnectionItem;