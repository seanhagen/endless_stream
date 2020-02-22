import Phaser from "phaser";

import innTiles from "../assets/worlds/inn/interior_1.png";
import innMap from "../assets/worlds/inn/inn.json";

import homeTiles from "../assets/worlds/inn/interior_1.png";
import homeMap from "../assets/worlds/inn/home.json";

import * as Chars from "../objects/characters/character";

import atlasIMG from "../assets/chars.png";
import atlasJSON from "../assets/chars_atlas.json";

import unknownPlayer from "../assets/unknown_char.png";

import f1Png from "../assets/fonts/f1.png";
import f1Fnt from "../assets/fonts/f1.xml";

import f2Png from "../assets/fonts/f2.png";
import f2Fnt from "../assets/fonts/f2.xml";

const width = 1280;
const height = 720;

export default class Home extends Phaser.Scene {
  constructor() {
    super({ key: "Home", active: true });
    this.handleData = this.handleData.bind(this);
  }

  handleData(data) {
    if (data.tick === undefined) {
      console.log("not tick (home): ", data);
    }
  }

  preload() {
    this.load.image("home-tiles", homeTiles);
    this.load.image("unknownPlayer", unknownPlayer);
    this.load.tilemapTiledJSON("home-map", homeMap);

    this.load.atlas("charAtlas", atlasIMG, atlasJSON);

    // console.log("load: ", this.load)

    this.load.bitmapFont("f1", f1Png, f1Fnt);

    this.cameras.main.setZoom(1).setPosition(0, 0);
  }

  create() {
    this.game.socket.on("data", this.handleData);
    this.events.on("transitionout", (ts, dur) => {
      console.log("home, transition out: ", ts, dur);
      this.game.socket.off("data", this.handleData);
    });
    this.input.once(
      "pointerdown",
      ev => {
        console.log("home pointerdown: ", ev);
        this.scene.start("Inn");
      },
      this
    );

    const map = this.make.tilemap({ key: "home-map" });

    // Parameters are the name you gave the tileset in Tiled and then the key of the tileset image in
    // Phaser's cache (i.e. the name you used in preload)
    const tileset = map.addTilesetImage("interior_1", "home-tiles");

    const homeBaseLayer = map.createStaticLayer("Base", tileset, 0, 0);
    const homeBGLayer = map.createStaticLayer("Background", tileset, 0, 0);
    const homeObjLayer = map.createStaticLayer("Objects", tileset, 0, 0);
    const homeObjFLayer = map.createStaticLayer(
      "Objects Forward",
      tileset,
      0,
      0
    );

    // const fighter = new Chars.Fighter({scene: this, x: 300, y: 500});
    // const cleric = new Chars.Cleric({scene: this, x: 550, y: 500});
    // const ranger = new Chars.Ranger({scene: this, x:800, y: 500});
    // const wizard = new Chars.Wizard({scene: this, x: 1050, y: 500});

    const unknown1 = new Chars.Unknown({ scene: this, x: 300, y: 500 });
    const unknown2 = new Chars.Unknown({ scene: this, x: 550, y: 500 });
    const unknown3 = new Chars.Unknown({ scene: this, x: 800, y: 500 });
    const unknown4 = new Chars.Unknown({ scene: this, x: 1050, y: 500 });

    // const label1 = this.add.bitmapFont(550, 300, 'f1', 'testing font', 128)
    const label1 = this.add.bitmapText(300, 400, "f1", "P1", 50);
    const label2 = this.add.bitmapText(550, 400, "f1", "P2", 50);
    const label3 = this.add.bitmapText(800, 400, "f1", "P3", 50);
    const label4 = this.add.bitmapText(1050, 400, "f1", "P4", 50);

    // const label2 = this.add.text(550, 250, 'testing regular text');
    // label2.style.color = "#000";
    // label2.style.fontFamily = "f1";

    // const label2 = this.make.text({
    //   x: 550,
    //   y: 150,
    //   text: 'text from config',
    //   style: {
    //     fontSize: '64px',
    //     fontFamily: 'Arial',
    //     color: '#000',
    //     align: 'center',
    //   }
    // })
  }
}
