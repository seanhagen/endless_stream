import React, { Component, Fragment } from "react";
import { compose } from "redux";
import { Link, Route, withRouter } from "react-router-dom";
import { connect } from "react-redux";
import { bindActionCreators } from "redux";

import { gameState } from "./reducers/game";

import Home from "./components/home";
import Game from "./components/game";

import "./App.css";

class Welcome extends Component {
  render() {
    return (
      <Fragment>
        <h1>Welcome To Endless Stream!</h1>
        <Link to="/game">Start or Join A Game!</Link>
      </Fragment>
    );
  }
}

class App extends Component {
  render() {
    return (
      <div className="App">
        <Route exact path="/" component={Welcome} />
        <Route path="/game" component={Home} />
        <Route path="/player/:code" component={Game} />
      </div>
    );
  }
}

const mapStateToProps = state => {
  return { code: state.game.code };
};

const mapDispatchToProps = dispatch => ({
  actions: bindActionCreators(gameState.actions, dispatch)
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(compose(withRouter)(App));
