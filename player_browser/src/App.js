import React, {Component} from 'react';
import { compose } from 'redux';
import { Route, withRouter, BrowserRouter as Router } from 'react-router-dom';
import logo from './logo.svg';
import { Counter } from './features/counter/Counter';
import './App.css';

import Home from './components/home';

// import {GameClient} from './grpc/endless_pb_service';
// import {CreateGame} from './grpc/endless_pb';

// const gc = new GameClient('http://localhost:8080')

// const createGame = new CreateGame();
// gc.create(createGame, (err, msg) => {
//   if (err){
//     var x = err.toObject();
//     console.error("error creating game: ", x)
//     return
//   }

//   var resp = msg.toObject();
//   console.log("response: ", resp);
// })
// console.log("create game: ", createGame);

// gc.create()
// var stateClient = gc.state()
// console.log("state client: ", stateClient);

// stateClient.on('data', (msg) => {
//   console.log("state message: ", msg)
// })

class App extends Component {
  render(){
    return (
        <div className="App">
        <Route exact path="/" component={Home} />
        </div>
    );
  }
}

export default compose(withRouter)(App);
