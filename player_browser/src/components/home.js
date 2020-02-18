import React, {Component} from 'react';
import { connect } from 'react-redux'
import { bindActionCreators } from 'redux'
import { withRouter } from "react-router-dom";

import {gameState} from '../reducers/game';

const defaultState = {code: ''};

export class Home extends Component {
  constructor(props){
    super(props);
    this.state = defaultState;
    this.onClick = this.onClick.bind(this);
    this.onChange = this.onChange.bind(this);
  }

  onClick(ev){
    ev.preventDefault();
    console.log("need to create game!")
    //joinGame({code: 'AAAA'});
    console.log("props: ", this.props);
    console.log("state: ", this.state);
    var ret = this.props.actions.joinGame();
    console.log("action call: ", ret);
  }

  onChange(ev) {
    ev.preventDefault();
  }

  render(){
    const {code} = this.state;

    return (
        <div>
        <h1>Endless Stream</h1>
        <input type="text" value={code} onChange={this.onChange}/>
        <button onClick={this.onClick}>Join Game!</button>
        </div>
    )
  }
}

const wr = withRouter(Home);

const mapStateToProps = state => {
  return {code: state.game.code};
}


const mapDispatchToProps = dispatch => ({
    actions: bindActionCreators(gameState.actions, dispatch)
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(wr)
