import Phaser from "phaser";

import forestTiles from "../assets/worlds/forest/forest.png";
import forestMap from "../assets/worlds/forest/forest-map.json";

const width = 1280;
const height = 720;

export default class Forest extends Phaser.Scene {
  constructor() {
    super({key: 'Forest', active: true});
  }

  preload(){
    this.load.image("forest-tiles", forestTiles);
    this.load.tilemapTiledJSON("forest-map", forestMap);
    this.cameras.main
      .setZoom(0.5)
      .setPosition(-650, -500)
      .setSize(width/0.5, height/0.5);
  }

  create(){
    const map = this.make.tilemap({ key: "forest-map" });

    // Parameters are the name you gave the tileset in Tiled and then the key of the tileset image in
    // Phaser's cache (i.e. the name you used in preload)
    const tileset = map.addTilesetImage("forest", "forest-tiles");

    // Parameters: layer name (or index) from Tiled, tileset, x, y
    const forestBaseLayer = map.createStaticLayer("Base", tileset, 0, 0);
    const forestBackgroundLayer = map.createStaticLayer("Background", tileset, 0,0);
    const forestTreesLayer = map.createStaticLayer("Trees", tileset, 0,0);
    const forestTreeFixessLayer = map.createStaticLayer("Tree Fixes", tileset, 0,0);
    const forestRoadLayer = map.createStaticLayer("Road", tileset, 0, 0);
    const forestPrettyLayer = map.createStaticLayer("Pretty", tileset, 0, 0);
  }
}
