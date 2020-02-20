import React, { Component } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { withRouter } from "react-router-dom";
import { v4 } from "uuid";

import { gameState } from "../reducers/game";
import io from "socket.io-client";
import { ClientType, Input, Register } from "../grpc/input_pb";

export class Game extends Component {
  constructor(props) {
    super(props);
    this.state = {
      joining: false,
      error: null,
      code: props.match.params.code
    };
    this.checkStream = this.checkStream.bind(this);

    console.log("io: ", io);
    const s = io("http://localhost:3002");

    var pid = localStorage.getItem("player-uuid");
    console.log("pid from local storage: ", pid);
    if (pid === null) {
      pid = v4();
      localStorage.setItem("player-uuid", pid);
    }
    s.on("connect", () => {
      console.log("socket connected");

      const reg = new Register();
      reg.setId(pid);
      reg.setCode(this.state.code);
      reg.setName("Testing JS Client");
      reg.setType(ClientType.CLIENTPLAYER);
      console.log("component mounted: ", reg.toObject());

      const inp = new Input();
      inp.setRegister(reg);

      s.emit("input", inp.toObject());
    });
    s.on("data", data => {
      console.log("recieved data: ", data);
    });
    s.on("disconnect", () => {
      console.log("disconnected");
    });
    this.socket = s;
  }

  componentWillUnmount() {
    this.socket.close();
  }

  checkStream(ev) {
    ev.preventDefault();
    console.log("stream client: ", this.stream);
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
