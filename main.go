package main

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	actorFrameSpriteAnimateX = 0
	actorFrameSpriteAnimateY = 0
	screenWidth              = 640
	screenHeight             = 480
	actorSpriteWidth         = 32
	actorSpriteHeigh         = 32
	actorFrameNum            = 8
	sampleRate               = 48000
	introLengthInSecond      = 1
	loopLengthInSecond       = 90
)

var (
	logoImage, runnerImage *ebiten.Image
	copyRight              *ebiten.Image
	pushStartKey           *ebiten.Image
)

func init() {
	var err error
	runnerImage, _, err = ebitenutil.NewImageFromFile("character_animation.png")
	if err != nil {
		log.Fatal(err)
	}
	logoImage, _, err = ebitenutil.NewImageFromFile("logo.png")
	if err != nil {
		log.Fatal(err)
	}
	copyRight, _, err = ebitenutil.NewImageFromFile("copy_right.png")
	if err != nil {
		log.Fatal(err)
	}
	pushStartKey, _, err = ebitenutil.NewImageFromFile("push_start_key.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	player       *audio.Player
	audioContext *audio.Context
	count        int
}

func (g *Game) Update() error {
	g.count++
	if g.player != nil {
		return nil
	}
	if g.audioContext == nil {
		g.audioContext = audio.NewContext(sampleRate)
	}
	data, err := ioutil.ReadFile("01.mp3")
	if err != nil {
		log.Fatal(err)
	}
	mp3, err := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	s := audio.NewInfiniteLoopWithIntro(mp3, introLengthInSecond*4*sampleRate, loopLengthInSecond*4*sampleRate)
	g.player, err = audio.NewPlayer(g.audioContext, s)
	if err != nil {
		return err
	}
	g.player.Play()
	return nil
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func (g *Game) Draw(screen *ebiten.Image) {
	logoOp := &ebiten.DrawImageOptions{}
	logoOp.ColorM.Scale(1, 1, 1, 0.5)
	w, _ := logoImage.Size()
	wposition := float64(w / 2)
	logoOp.GeoM.Translate((screenWidth/2)-wposition, (screenHeight/2)-130)
	screen.DrawImage(
		logoImage,
		logoOp,
	)

	actorOp := &ebiten.DrawImageOptions{}
	actorOp.ColorM.Scale(1, 1, 1, 0.5)
	actorOp.GeoM.Translate(screenWidth/4, screenHeight/2)
	i := (g.count / 5) % actorFrameNum
	x, y := actorFrameSpriteAnimateX+i*actorSpriteWidth, actorFrameSpriteAnimateY
	screen.DrawImage(
		runnerImage.SubImage(image.Rect(x, y, x+actorSpriteWidth, y+actorSpriteHeigh)).(*ebiten.Image),
		actorOp,
	)

	pushStartOp := &ebiten.DrawImageOptions{}
	pushStartOp.ColorM.Scale(1, 1, 1, 0.5)
	w, _ = pushStartKey.Size()
	wposition = float64(w / 2)
	pushStartOp.GeoM.Translate((screenWidth/2)-wposition, (screenHeight/2)+5)
	screen.DrawImage(
		pushStartKey,
		pushStartOp,
	)

	copyOp := &ebiten.DrawImageOptions{}
	copyOp.ColorM.Scale(1, 1, 1, 0.5)
	w, _ = copyRight.Size()
	wposition = float64(w / 2)
	copyOp.GeoM.Translate((screenWidth/2)-wposition, (screenHeight/2)+40)
	screen.DrawImage(
		copyRight,
		copyOp,
	)
}
func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Hacker Adventure")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
