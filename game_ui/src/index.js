import Phaser from "phaser";
import io from "socket.io-client";

// import logoImg from "./assets/logo.png";

import Home from "./scenes/home";
import Inn from "./scenes/inn";
import Forest from "./scenes/forest";

// import GameClient from "./grpc/endless_grpc_web_pb.js";
// import { ClientType, Register, Input } from "./grpc/input_pb";

const width = 1280;
const height = 720;

const config = {
  type: Phaser.AUTO,
  parent: "phaser-example",
  width: width,
  height: height,
  pixelArt: true,
  physics: {
    default: "arcade",
    arcade: {
      gravity: { y: 0 }
    }
  },
  scene: [Home, Inn, Forest]
};

const game = new Phaser.Game(config);
const s = io("http://localhost:8080");
s.on("connect", () => {
  console.log("connected to websocket");
  // const reg = new Register();
  // reg.setName("Game UI");
  // reg.setCode(getCodeFromURI());
  // reg.setType(ClientType.CLIENTDISPLAY);

  // const inp = new Input();
  // inp.setRegister(reg);
  // s.emit("input", inp.toObject());
});

s.on("data", data => {
  if (data.tick === undefined) {
    console.log("received data: ", data);
  }
});

s.on("error", err => {
  console.error("error from socket: ", err);
});

s.on("disconnect", () => {
  console.log("socket disconnected");
});

game.socket = s;

// game.scene.add("forest", Forest);
// game.scene.add("inn", Inn);
// game.scene.add("home", Home);
// game.scene.start("Home");

console.log("game: ", game);

const getCodeFromURI = () => {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get("code");
};
