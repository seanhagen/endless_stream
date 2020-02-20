import React, { Component, Fragment } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { withRouter } from "react-router-dom";
import { v4 } from "uuid";

import { gameState } from "../reducers/game";
import io from "socket.io-client";
import { ClientType, CharSelect, Register, Input } from "../grpc/input_pb";
import { Display, Class, ClassType } from "../grpc/util_pb";

console.log("class type: ", ClassType);

const emitInput = "input";

export class Game extends Component {
  constructor(props) {
    super(props);
    var pid = localStorage.getItem("player-uuid");
    if (pid === null) {
      pid = v4();
      localStorage.setItem("player-uuid", pid);
    }
    var name = localStorage.getItem("player-name");
    if (name === null) {
      name = "testing name";
      localStorage.setItem("player-name", name);
    }

    this.state = {
      pid: pid,
      name: name,
      joining: false,
      error: null,
      code: props.match.params.code,
      game: null
    };

    this.processData = this.processData.bind(this);
    this.selectCharacter = this.selectCharacter.bind(this);

    const s = io(process.env.REACT_APP_GAME_SERVER);
    s.on("connect", () => {
      const reg = new Register();
      reg.setId(pid);
      reg.setName(name);
      reg.setCode(this.state.code);
      reg.setType(ClientType.CLIENTPLAYER);

      const inp = new Input();
      inp.setRegister(reg);
      s.emit(emitInput, inp.toObject());
    });
    s.on("data", this.processData);
    s.on("error", e => {
      console.error("error from socket: ", e);
    });
    s.on("disconnect", () => {
      console.log("disconnected");
    });
    this.socket = s;
  }

  processData(msg) {
    switch (msg.data) {
      case "tick":
        break;
      case "state":
        this.setState({ game: msg.state });
        break;
      case "joined":
        console.log("somone joined: ", msg);
        break;
      default:
        console.log("don't know how to handle this message type: ", msg.data);
    }
  }

  componentWillUnmount() {
    this.socket.close();
  }

  renderJoining(code) {
    return (
      <div>
        <h1>Endless Stream</h1>
        <p>Joining Game: {code}</p>
      </div>
    );
  }

  renderError(error) {
    return (
      <div>
        <h1>Endless Stream</h1>
        <p>Error joining game: {error}</p>
      </div>
    );
  }

  renderState() {
    const { game } = this.state;
    if (game === null) {
      return <p>Joining game!</p>;
    }

    switch (game.display) {
      case Display.SCREENCHARSELECT:
        return this.renderCharSelect();
    }

    return <p>Game Joined!</p>;
  }

  selectCharacter(ev) {
    ev.preventDefault();
    console.log("character select: ", ev.target.dataset);

    const c = new Class();
    c.setClass(ClassType[ev.target.dataset.class.toUpperCase()]);

    const cs = new CharSelect();
    cs.setPlayerId(this.state.pid);
    cs.setChoice(c);

    const inp = new Input();
    inp.setPlayerId(this.state.pid);
    inp.setCharSelect(cs);

    console.log("selecting character: ", inp.toObject());

    this.socket.emit(emitInput, inp.toObject());
  }

  makeCharSelectButton(cl) {
    const name = cl.charAt(0).toUpperCase() + cl.slice(1);
    return (
      <button data-class={cl} onClick={this.selectCharacter}>
        {name}
      </button>
    );
  }

  renderCharSelect() {
    const { game } = this.state;
    console.log("selected: ", game.selected);
    return (
      <Fragment>
        <h3>Character Select!</h3>
        {this.makeCharSelectButton("fighter")}
        {this.makeCharSelectButton("ranger")}
        {this.makeCharSelectButton("cleric")}
        {this.makeCharSelectButton("wizard")}
      </Fragment>
    );
  }

  render() {
    const { joining, error } = this.state;
    const { code } = this.state;
    if (joining) {
      return this.renderJoining(code);
    }

    if (error) {
      return this.renderError(error);
    }

    return (
      <div>
        <h1>Endless Stream</h1>
        {this.renderState()}
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
