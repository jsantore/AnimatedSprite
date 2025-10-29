package main

import (
	"embed"
	"fmt"
	"image"
	"log"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var EmbeddedAssets embed.FS

const (
	WINDOW_WIDTH     = 1000
	WINDOW_HEIGHT    = 1000
	COIN_DIMENSION   = 512.0
	GUY_FRAME_WIDTH  = 108
	GUY_HEIGHT       = 140
	FRAME_COUNT      = 4
	FRAMES_PER_SHEET = 8
)

type Direction int

const (
	RIGHT Direction = iota
	LEFT
	UP
	DOWN
)

type PlayerSprite struct {
	spriteSheet *ebiten.Image
	xLoc        int
	direction   Direction
	frame       int
	frameDelay  int
}

type AnimatedSpriteDemo3 struct {
	player PlayerSprite
}

func (demoGame *AnimatedSpriteDemo3) Update() error {
	getPlayerInput(demoGame)
	demoGame.player.frameDelay += 1
	if demoGame.player.frameDelay%FRAME_COUNT == 0 {
		demoGame.player.frame += 1
		if demoGame.player.frame >= FRAMES_PER_SHEET {
			demoGame.player.frame = 0
		}
		if demoGame.player.direction == LEFT {
			demoGame.player.xLoc -= 5
		} else if demoGame.player.direction == RIGHT {
			demoGame.player.xLoc += 5
		}

	}
	return nil
}

func getPlayerInput(game *AnimatedSpriteDemo3) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && game.player.xLoc > 0 {
		game.player.direction = LEFT
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) &&
		game.player.xLoc < WINDOW_WIDTH-GUY_FRAME_WIDTH {
		game.player.direction = RIGHT
	}
}

func (demoGame AnimatedSpriteDemo3) Draw(screen *ebiten.Image) {
	yLoc := 300.0 //when you add up and down move this to the player sprite and update it in update
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(float64(demoGame.player.xLoc), yLoc)
	screen.DrawImage(demoGame.player.spriteSheet.SubImage(image.Rect(demoGame.player.frame*GUY_FRAME_WIDTH,
		int(demoGame.player.direction)*GUY_HEIGHT,
		demoGame.player.frame*GUY_FRAME_WIDTH+GUY_FRAME_WIDTH,
		int(demoGame.player.direction)*GUY_HEIGHT+GUY_HEIGHT)).(*ebiten.Image), op)
}

func (demoGame AnimatedSpriteDemo3) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	animationGuy := LoadEmbeddedImage("", "scottpilgrim_multiple.png")
	myPlayer := PlayerSprite{spriteSheet: animationGuy,
		xLoc: WINDOW_WIDTH / 2}
	demo := AnimatedSpriteDemo3{
		player: myPlayer,
	}
	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Animated Sprite")
	ebiten.RunGame(&demo)
}

//type AnimatedSpriteDemo2 struct {
//	Raccoon    []*ebiten.Image
//	Frame      int
//	FrameDelay int
//}
//
//func (demo *AnimatedSpriteDemo2) Update() error {
//	demo.FrameDelay += 1
//	if demo.FrameDelay%5 == 0 {
//		demo.Frame += 1
//		if demo.Frame >= len(demo.Raccoon) {
//			demo.Frame = 0
//		}
//	}
//	return nil
//}
//
//func (demo AnimatedSpriteDemo2) Draw(screen *ebiten.Image) {
//	drawOps := ebiten.DrawImageOptions{}
//	drawOps.GeoM.Reset()
//	//drawOps.GeoM.Translate(float64(WINDOW_WIDTH/2), float64(WINDOW_HEIGHT/2))
//	screen.DrawImage(demo.Raccoon[demo.Frame], &drawOps)
//}
//
//func (demo AnimatedSpriteDemo2) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return outsideWidth, outsideHeight
//}
//
//func main() {
//	frames := LoadAllRaccoons()
//	ebiten.SetWindowSize(1000, 1000)
//	ebiten.SetWindowTitle("sliceOfImeages")
//	demo := AnimatedSpriteDemo2{
//		Raccoon:    frames,
//		Frame:      0,
//		FrameDelay: 0,
//	}
//	ebiten.RunGame(&demo)
//}
//
//func LoadAllRaccoons() []*ebiten.Image {
//	all_frames := make([]*ebiten.Image, 14, 20)
//	suffix_list := []string{"01", "03", "05", "07", "09", "11", "13", "15", "17", "19", "21", "23", "25", "27"}
//	for index, suffix := range suffix_list {
//		filename := fmt.Sprintf("victory-dance00%s.png", suffix)
//		frame_pict := LoadEmbeddedImage("victory-dance", filename)
//		all_frames[index] = frame_pict
//	}
//	return all_frames
//}

//type AnimatedSpriteDemo struct {
//	CoinImage  *ebiten.Image
//	xFrame     int
//	yFrame     int
//	FrameDelay int
//}
//
//func (demo *AnimatedSpriteDemo) Update() error {
//	demo.FrameDelay += 1
//	if demo.FrameDelay%5 == 0 {
//		demo.xFrame += 1
//		if demo.xFrame >= FRAME_COUNT {
//			demo.xFrame = 0
//			demo.yFrame += 1
//			if demo.yFrame >= FRAME_COUNT {
//				demo.yFrame = 0
//			}
//		}
//	}
//	return nil
//}
//
//// Draw Heavily based on the official ebiten animation demo
//// https://ebitengine.org/en/examples/animation.html#Code
//func (demo *AnimatedSpriteDemo) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	op.GeoM.Translate(COIN_DIMENSION/2, COIN_DIMENSION/2)
//	frameX := demo.xFrame * COIN_DIMENSION
//	frameY := demo.yFrame * COIN_DIMENSION
//	screen.DrawImage(demo.CoinImage.SubImage(image.Rect(frameX, frameY,
//		frameX+COIN_DIMENSION, frameY+COIN_DIMENSION)).(*ebiten.Image), op)
//}
//
//func (demo AnimatedSpriteDemo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return outsideWidth, outsideHeight
//}
//
//func main() {
//	coin := LoadEmbeddedImage("", "Coin_Spin_Animation_A.png")
//	demo := AnimatedSpriteDemo{CoinImage: coin} //xFrame and yFrame deliberately 0
//	ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
//	ebiten.SetWindowTitle("Sprite animation on Sprite Sheet")
//	err := ebiten.RunGame(&demo)
//	if err != nil {
//		fmt.Println("Error running game:", err)
//	}
//}

func LoadEmbeddedImage(folderName string, imageName string) *ebiten.Image {
	embeddedFile, err := EmbeddedAssets.Open(path.Join("assets", folderName, imageName))
	if err != nil {
		log.Fatal("failed to load embedded image ", imageName, err)
	}
	ebitenImage, _, err := ebitenutil.NewImageFromReader(embeddedFile)
	if err != nil {
		fmt.Println("Error loading tile image:", imageName, err)
	}
	return ebitenImage
}
