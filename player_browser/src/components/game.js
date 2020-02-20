import React, { Component, Fragment } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { withRouter } from "react-router-dom";
import { v4 } from "uuid";
import * as _ from "lodash";

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
      game: null,
      selectedCharacter: 0,
      amVip: false
    };

    this.processData = this.processData.bind(this);
    this.selectCharacter = this.selectCharacter.bind(this);
    this.unselectCharacter = this.unselectCharacter.bind(this);
    this.handleJoin = this.handleJoin.bind(this);

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
    const { game } = this.state;
    switch (msg.data) {
      case "tick":
        break;
      case "state":
        if (!_.isEqual(game, msg.state)) {
          console.log("new state: ", msg.state);
          this.setState({ game: msg.state });
        }

        break;
      case "joined":
        this.handleJoin(msg.joined);
        break;
      default:
        console.log("don't know how to handle this message type: ", msg.data);
    }
  }

  handleJoin(data) {
    const { pid } = this.state;
    if (pid == data.id) {
      if (data.isVip) {
        this.setState({ amVip: true });
      }
      if (data.asAudience) {
        console.log("this user is part of the audience!");
      }
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

    const s = ClassType[ev.target.dataset.class.toUpperCase()];
    this.setState({ selectedCharacter: s });

    const c = new Class();
    c.setClass(s);

    const cs = new CharSelect();
    cs.setPlayerId(this.state.pid);
    cs.setChoice(c);

    const inp = new Input();
    inp.setPlayerId(this.state.pid);
    inp.setCharSelect(cs);

    const { pid, game } = this.state;
    var { selected } = game;
    selected.selected[pid] = s;
    game.selected = selected;
    this.setState({ game: game });
    this.socket.emit(emitInput, inp.toObject());
  }

  unselectCharacter(ev) {
    ev.preventDefault();
    const cs = new CharSelect();
    cs.setPlayerId(this.state.pid);

    const inp = new Input();
    inp.setPlayerId(this.state.pid);
    inp.setCharSelect(cs);

    const { pid, game } = this.state;
    var { selected } = game;
    delete selected.selected[pid];
    game.selected = selected;
    this.setState({ game: game });
    this.socket.emit(emitInput, inp.toObject());
  }

  makeCharSelectButton(cl, selected) {
    const name = cl.charAt(0).toUpperCase() + cl.slice(1);
    const clVal = ClassType[cl.toUpperCase()];
    const alreadySelected = Object.values(selected);
    var taken = alreadySelected.includes(clVal);

    if (taken) {
      return <button disabled={true}>{name}</button>;
    }

    return (
      <button data-class={cl} onClick={this.selectCharacter}>
        {name}
      </button>
    );
  }

  renderSelected() {
    const { pid, game } = this.state;
    const selected = game.selected.selected[pid];
    switch (selected) {
      case ClassType.FIGHTER:
        return (
          <Fragment>
            <h3>Fighter</h3>
            <button onClick={this.unselectCharacter}>Unselect</button>
          </Fragment>
        );
      case ClassType.RANGER:
        return (
          <Fragment>
            <h3>Ranger</h3>
            <button onClick={this.unselectCharacter}>Unselect</button>
          </Fragment>
        );
      case ClassType.CLERIC:
        return (
          <Fragment>
            <h3>Cleric</h3>
            <button onClick={this.unselectCharacter}>Unselect</button>
          </Fragment>
        );
      case ClassType.WIZARD:
        return (
          <Fragment>
            <h3>Wizard</h3>
            <button onClick={this.unselectCharacter}>Unselect</button>
          </Fragment>
        );
      default:
        return (
          <Fragment>
            <h3>Unknown</h3>
            <button onClick={this.unselectCharacter}>Unselect</button>
          </Fragment>
        );
    }
  }

  startGame(ev) {
    ev.preventDefault();
    console.log("need to start the game");
  }

  renderCharSelect() {
    const { amVip, pid, game } = this.state;
    const { selected } = game.selected;
    return (
      <Fragment>
        <h3>Character Select!</h3>
        {selected[pid] === undefined ? (
          <Fragment>
            {this.makeCharSelectButton("fighter", selected)}
            {this.makeCharSelectButton("ranger", selected)}
            {this.makeCharSelectButton("cleric", selected)}
            {this.makeCharSelectButton("wizard", selected)}
          </Fragment>
        ) : (
            <Fragment>{this.renderSelected()}</Fragment>
          )}
        {amVip ? (
          <Fragment>
            <br />
            <br />
            <button onClick={this.startGame}>Start Game</button>
          </Fragment>
        ) : (
            <Fragment />
          )}
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
