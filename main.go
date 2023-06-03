package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
)

// main est la fonction principale du projet. Elle commence par lire le fichier
// de configuration, puis elle charge en mémoire l'image d'une particule. Elle
// initialise ensuite la fenêtre d'affichage, puis elle crée un système de
// particules encapsulé dans un "game" et appelle la fonction RunGame qui se
// charge de faire les mise-à-jour (Update) et affichages (Draw) de manière
// régulière. 
// FullScreen sezrt à faire apparaître la fen^tre en plein écran
// Si SpawnX (ou Y) est inférieur à 0, elle devient égal à la moitié de son axe
// cela sert à faire apparaitre des particules au milieu de l'écran
func main() {
	config.Get("config.json")
	assets.Get()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)

	if config.General.FullScreen {
		ebiten.SetFullscreen(true)
		config.General.WindowSizeX, config.General.WindowSizeY = ebiten.ScreenSizeInFullscreen()
	}

	if config.General.SpawnX < 0 {
		config.General.SpawnX = float64(config.General.WindowSizeX)/2
	}

	if config.General.SpawnY < 0 {
		config.General.SpawnY = float64(config.General.WindowSizeY)/2
	}

	g := game{system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
