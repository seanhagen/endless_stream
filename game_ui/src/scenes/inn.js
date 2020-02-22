import Phaser from "phaser";

import innTiles from "../assets/worlds/inn/interior_1.png";
import innMap from "../assets/worlds/inn/inn.json";

const width = 1280;
const height = 720;

export default class Inn extends Phaser.Scene {
  constructor() {
    super({ key: "Inn" });
    console.log("created inn scene: ", this);
    this.handleData = this.handleData.bind(this);
  }

  init() {
    console.log("inn init");
  }

  handleData(data) {
    if (data.tick === undefined) {
      console.log("not tick: ", data);
    }
  }

  preload() {
    console.log("inn preload");
    this.load.image("inn-tiles", innTiles);
    this.load.tilemapTiledJSON("inn-map", innMap);

    this.cameras.main.setZoom(1).setPosition(0, 0);
  }

  create() {
    this.game.socket.on("data", this.handleData);
    this.events.on("transitionout", (ts, dur) => {
      console.log("inn, transition out: ", ts, dur);
      this.game.socket.off("data", this.handleData);
    });

    this.input.once(
      "pointerdown",
      ev => {
        this.scene.start("Forest");
        console.log("inn switch: ", ev);
      },
      this
    );
    const map = this.make.tilemap({ key: "inn-map" });

    // Parameters are the name you gave the tileset in Tiled and then the key of the tileset image in
    // Phaser's cache (i.e. the name you used in preload)
    const tileset = map.addTilesetImage("interior_1", "inn-tiles");

    const innBaseLayer = map.createStaticLayer("Base", tileset, 0, 0);
    const innBGLayer = map.createStaticLayer("Background", tileset, 0, 0);
    const innPrettyLayer = map.createStaticLayer("Pretty", tileset, 0, 0);
    const innSmallLayer = map.createStaticLayer("Small", tileset, 0, 0);
  }
}
