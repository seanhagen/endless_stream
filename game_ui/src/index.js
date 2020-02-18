import Phaser from "phaser";

import logoImg from "./assets/logo.png";

import Home from "./scenes/home";
import Inn from "./scenes/inn";
import Forest from "./scenes/forest";

import GameClient from "./grpc/endless_grpc_web_pb.js"



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
  // scenes: [Forest]
};

const game = new Phaser.Game(config);
game.scene.add('forest', Forest);
game.scene.add('inn', Inn)
game.scene.add('home',Home)
game.scene.start('home')
