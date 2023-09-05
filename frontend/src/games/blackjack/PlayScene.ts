import path from 'path'
import Phaser from 'phaser'
import Text = Phaser.GameObjects.Text

export default class PlayScene extends Phaser.Scene {
  // TODO:　ここから下でBlackjackを作成する

  preload() {
    console.log('preload!!')
    this.load.atlas(
      'cards',
      '/assets/image/cards.png',
      '/assets/image/cards.json',
    )
  }

  create() {
    console.log('create!!')
    //  Create a stack of random cards

    const frames = this.textures.get('cards').getFrameNames()

    let x = 100
    let y = 100

    for (let i = 0; i < 64; i++) {
      this.add
        .image(x, y, 'cards', Phaser.Math.RND.pick(frames))
        .setInteractive()

      x += 4
      y += 4
    }

    // this.input.on('gameobjectdown', function ()
    // {

    //     //  Will contain the top-most Game Object (in the display list)
    //     this.tweens.add({
    //         x: { value: 1100, duration: 1500, ease: 'Power2' },
    //         y: { value: 500, duration: 500, ease: 'Bounce.easeOut', delay: 150 }
    //     });

    // }, this);
  }

  update() {
    console.log('update!!')
  }
}
