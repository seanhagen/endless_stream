import Phaser from "phaser";

import logoImg from "./assets/logo.png";

import Forest from "./scenes/forest";


const config = {
  type: Phaser.AUTO,
  parent: "phaser-example",
  width: width,
  height: height,
  // scene: {
  //   preload: preload,
  //   create: create,
  //   update: update
  // },
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
game.scene.start('forest')
