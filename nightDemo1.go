package main

//
//import (
//	"embed"
//	"fmt"
//	"image"
//	"log"
//	"path"
//
//	"github.com/hajimehoshi/ebiten/v2"
//	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
//)
//
////go:embed assets/*
//var EmbeddedAssets embed.FS
//
//const (
//	WINDOW_WIDTH   = 1000
//	WINDOW_HEIGHT  = 1000
//	COIN_DIMENSION = 512.0
//	FRAME_COUNT    = 4
//)
//
//type AnimatedSpriteDemo struct {
//	CoinImage  *ebiten.Image
//	xFrame     int
//	yFrame     int
//	FrameDelay int
//}
//
//func (demo AnimatedSpriteDemo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//	return outsideWidth, outsideHeight
//}
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
//func (demo *AnimatedSpriteDemo) Update() error {
//	demo.xFrame += 1
//	if demo.xFrame >= FRAME_COUNT {
//		demo.xFrame = 0
//		demo.yFrame += 1
//		if demo.yFrame >= FRAME_COUNT {
//			demo.yFrame = 0
//		}
//	}
//	return nil
//}
//
//func (demo *AnimatedSpriteDemo) Draw(screen *ebiten.Image) {
//	op := &ebiten.DrawImageOptions{}
//	op.GeoM.Translate(COIN_DIMENSION/2, COIN_DIMENSION/2)
//	frameX := demo.xFrame * COIN_DIMENSION
//	frameY := demo.yFrame * COIN_DIMENSION
//	screen.DrawImage(demo.CoinImage.SubImage(image.Rect(frameX, frameY,
//		frameX+COIN_DIMENSION, frameY+COIN_DIMENSION)).(*ebiten.Image), op)
//}
//
//func LoadEmbeddedImage(folderName string, imageName string) *ebiten.Image {
//	embeddedFile, err := EmbeddedAssets.Open(path.Join("assets", folderName, imageName))
//	if err != nil {
//		log.Fatal("failed to load embedded image ", imageName, err)
//	}
//	ebitenImage, _, err := ebitenutil.NewImageFromReader(embeddedFile)
//	if err != nil {
//		fmt.Println("Error loading tile image:", imageName, err)
//	}
//	return ebitenImage
//}
