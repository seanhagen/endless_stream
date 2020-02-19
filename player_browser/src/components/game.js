import React, { Component } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { Link, Route, withRouter } from "react-router-dom";
import { v4 } from "uuid";

import { gameState } from "../reducers/game";
import { GameClient } from "../grpc/endless_pb_service";
import { ClientType, Input, Register } from "../grpc/input_pb";

const gc = new GameClient(process.env.REACT_APP_GAME_SERVER);

export class Game extends Component {
  constructor(props) {
    super(props);
    this.state = {
      joining: false,
      error: null,
      code: props.match.params.code
    };
    this.checkStream = this.checkStream.bind(this);
  }

  componentDidMount() {
    // console.log("props: ", this.props.actions.connectToGame(this.state.code));
    const pid = v4();
    const reg = new Register();
    reg.setId(pid);
    reg.setCode(this.state.code);
    reg.setName("Testing JS Client");
    reg.setType(ClientType.CLIENTPLAYER);
    console.log("component mounted: ", reg.toObject());

    const inp = new Input();
    inp.setRegister(reg);

    const strm = gc.state();
    strm
      .on("data", msg => {
        console.log("stream data: ", msg.toObject());
        // dispatch(gameState.actions.receiveStreamMsg(msg.toObject()));
      })
      .on("status", msg => {
        console.log("stream status: ", msg);
      })
      .write(inp);

    this.stream = strm;
  }

  checkStream(ev) {
    ev.preventDefault();
    console.log("stream client: ", this.stream.client);
  }

  render() {
    const { joining, error } = this.state;
    const { code } = this.state;
    console.log("code: ", code);
    if (joining) {
      return (
        <div>
          <h1>Endless Stream</h1>
          <p>Joining Game: {code}</p>
        </div>
      );
    }

    if (error) {
      return (
        <div>
          <h1>Endless Stream</h1>
          <p>Error joining game: {error}</p>
        </div>
      );
    }

    return (
      <div>
        <h1>Endless Stream</h1>
        <p>Game Joined!</p>
        <p>Code: {code}</p>
        <button onClick={this.checkStream}>Check</button>
      </div>
    );
  }
}

const wr = withRouter(Game);

const mapStateToProps = state => {
  return { code: state.game.code };
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(gameState.actions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(wr);
