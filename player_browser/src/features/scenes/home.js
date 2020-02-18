import React, {Component} from 'react';



export default class Home extends Component {
  constructor(props){
    super(props);

  }

  onClick(ev){
    ev.preventDefault();
    console.log("need to create game!")
  }

  render(){
    return (
      <div>
        <h1>Endless Stream</h1>
        <input type="text" />
        <button onClick={this.onClick}>Join Game!</button>
      </div>
    )
  }
}
