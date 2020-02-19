import React, { Component, Fragment } from "react";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";
import { Link, Route, withRouter } from "react-router-dom";

import { gameState } from "../reducers/game";

const defaultState = { code: "" };

class HomeChoice extends Component {
  render() {
    return (
      <Fragment>
        <ul>
          <li>
            <Link to="/game/create">Create A Game</Link>
          </li>
          <li>
            <Link to="/game/join">Join A Game</Link>
          </li>
          <li>
            <Link to="/game/rejoin">Re-join A Game</Link>
          </li>
          <li>
            <Link to="/game/audience">Join As Audience</Link>
          </li>
        </ul>
      </Fragment>
    );
  }
}

export class HomeCreateGame extends Component {
  componentWillReceiveProps(nextProps) {
    if (nextProps.code !== "undefined") {
      this.props.history.push("/player/" + nextProps.code);
    }
  }

  render() {
    const { onClick, code } = this.props;

    return (
      <div>
        <p>HomeCreateGame, code: {code}</p>
        <button onClick={onClick}>Create Game!</button>
      </div>
    );
  }
}

export class HomeJoinGame extends Component {
  constructor(props) {
    super(props);
    this.inputRef = React.createRef();
    this.state = {
      code: props.code,
      error: null
    };
    this.onClick = this.onClick.bind(this);
    this.onChange = this.onChange.bind(this);
  }

  onClick(ev) {
    const { code } = this.state;
    if (code.length === 4) {
      this.props.onClick(ev, code);
    } else {
      this.setState({ error: "Code must be 4 characters" });
    }
  }

  onChange(ev) {
    ev.preventDefault();
    var val = ev.target.value;
    val = val.toUpperCase();
    this.setState({ code: val, error: null });
  }

  render() {
    const { code, error } = this.state;
    return (
      <Fragment>
        {error ? <p>{error}</p> : <Fragment />}
        <input
          type="text"
          ref={this.inputRef}
          onChange={this.onChange}
          value={code}
        />
        <button onClick={this.onClick}>Join Game</button>
      </Fragment>
    );
  }
}

export class Home extends Component {
  constructor(props) {
    super(props);
    this.state = defaultState;
    this.createGame = this.createGame.bind(this);
    this.joinGame = this.joinGame.bind(this);
  }

  createGame(ev) {
    ev.preventDefault();
    console.log("[Home] need to create game!");
    //joinGame({code: 'AAAA'});
    console.log("[Home] props: ", this.props);
    console.log("[Home] state: ", this.state);
    this.props.actions.createGame();
    // createGame();
  }

  joinGame(ev, code) {
    ev.preventDefault();
    console.log("need to join game!", this.props.actions);
    var ret = this.props.actions.joinGame(code);
    console.log("joinGame action call: ", ret, this.props);
    this.props.history.push("/player/" + code);
  }

  render() {
    const { code } = this.props;
    return (
      <div>
        <h1>Endless Stream</h1>
        <p>Code: {code}</p>
        <Route exact path="/game" render={props => <HomeChoice {...props} />} />
        <Route
          exact
          path="/game/create"
          render={props => (
            <HomeCreateGame {...props} onClick={this.createGame} code={code} />
          )}
        />
        <Route
          exact
          path="/game/join"
          render={props => (
            <HomeJoinGame {...props} onClick={this.joinGame} code={code} />
          )}
        />
      </div>
    );
  }
}

const wr = withRouter(Home);

const mapStateToProps = state => {
  console.log("[Home] state updated: ", state.game.code);
  return { code: state.game.code };
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(gameState.actions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(wr);
