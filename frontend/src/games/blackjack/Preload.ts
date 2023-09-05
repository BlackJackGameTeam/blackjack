import Phaser from "phaser";
import Main from "./PlayScene";
import { useEffect } from "react";

export default function Index() {
  useEffect(() => {
    loadGame();
    console.log("load game");
  }, []);

  const loadGame = async () => {
    if (typeof window !== "object") {
      return;
    }

    var config = {
      type: Phaser.AUTO,
      width: 960, // window.innerWidth,
      height: 540, // window.innerHeight,
      // width: window.innerWidth * window.devicePixelRatio,
      // height: window.innerHeight * window.devicePixelRatio,
      backgroundColor: "#eef",
      physics: {
        default: "arcade",
        arcade: {
          gravity: { y: 200 },
        },
      },

      parent: "game-content",
      fps: { 
        forceSetTimeOut: true, 
        target: 30
      },

      scale: {
        mode: Phaser.Scale.WIDTH_CONTROLS_HEIGHT,
        autoCenter: Phaser.Scale.CENTER_BOTH,
      },

      render: {
        pixelArt: false,
      }
    };

    var game = new Phaser.Game(config)  
    game.scene.add("main", Main);
    game.scene.start("main");
  };

  return null;
}
