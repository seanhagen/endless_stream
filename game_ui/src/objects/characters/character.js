import Phaser from "phaser";


class Character extends Phaser.GameObjects.Sprite {
  constructor(config) {
    if (config.atlas) {
      super(config.scene, config.x, config.y, config.atlas, config.name);
    } else {
      super(config.scene, config.x, config.y, config.name);
    }

    config.scene.add.existing(this, config.x, config.y);
    console.log("character ", config.name, this.scale);
    this.scale = 2.5
  }
}

export class Fighter extends Character {
  constructor(config) {
    config.name = "fighter_01"
    config.atlas = "charAtlas"
    super(config);
  }
}

export class Ranger extends Character {
  constructor(config) {
    config.name = "ranger_01"
    config.atlas = "charAtlas"
    super(config);
  }
}

export class Cleric extends Character {
  constructor(config) {
    config.name = "cleric_01";
    config.atlas = "charAtlas"
    super(config)
  }
}

export class Wizard extends Character {
  constructor(config) {
    config.name ="wizard_01";
    config.atlas = "charAtlas"
    super(config);
  }
}

export class Unknown extends Character {
  constructor(config){
    config.name = "unknownPlayer";
    super(config);
  }
}
