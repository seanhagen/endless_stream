import Phaser from "phaser";

class Character extends Phaser.GameObjects.Sprite {
  constructor(config) {
    super(config.scene, config.x, config.y, config.name)
    console.log("Character: ", this);
  }
  // preload(){}
  // create(){}
}

export class Fighter extends Character {
  constructor(config) {
    config.name = "Fighter"
    super(config);
    config.scene.add.existing(this);
  }
}
