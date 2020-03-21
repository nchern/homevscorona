import React, {Component} from "react";
import {FaChevronCircleRight, FaUser, FaSmile, FaMap} from "react-icons/all";

class ConnectionItem extends Component {
    render() {
        return (
            <div className="TMP">
            <div className="connectionItem">
                {/*if this.props.event.type == "person"*/}
                <FaUser/>
                <div className="itemText">
                    <p><b>Claudia, Gustav, + 3 others</b></p>
                    <p>13:05 * Glogauer Str. 9</p>
                </div>
                <div className="itemGoto"><FaChevronCircleRight/></div>
            </div>
            <div className="connectionItem">
                {/*if this.props.event.type == "location"*/}
                <FaMap/>
                <div className="itemText">
                    <p><b>Kater Blau</b></p>
                    <p>13:05 * Glogauer Str. 9</p>
                </div>
                <div className="itemGoto"><FaChevronCircleRight/></div>
            </div>
            <div className="connectionItem">
                {/*if this.props.event.type == "athome"*/}
                {/* {this.props.type}*/}
                {/*<p>(IC)</p>*/}
                {/*<div>*/}
                {/*    <p><b>Claudia, Gustav, + 3 others</b></p>*/}
                {/*    <p>13:05 * Glogauer Str. 9</p>*/}
                {/*</div>*/}
                <FaSmile/>
                <div className="itemText">
                    <p><b>Du hast den ganzen Tag zuhause verbracht!</b></p>
                    <p> </p>
                </div>
                <div className="itemGoto"><FaChevronCircleRight/></div>
            </div>
            </div>
        )
    }
}

export default ConnectionItem;