package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
	"math"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. 
// Le Debug a été modifié pour faire apparaître un nombre frame arrondis et 
// le nombre de particules présent de le System
// Des "if" ont été ajouté afin d'afficher le "SpawnType" et de donné des instructions
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint("Frames : ",math.Round( ebiten.ActualTPS())," , " ,"Nombre de Particules : ", g.system.Content.Len()))
	}

	if config.General.MouseBlower{
		ebitenutil.DebugPrintAt(screen, "Left clic for Blower", config.General.WindowSizeX - 160, 20)
	} 


	if config.General.JailMouse{
		ebitenutil.DebugPrintAt(screen, "SpaceBar for Jail", config.General.WindowSizeX - 160, 40)
	
	}

	if config.General.SpawnType == "Tracker"{
		ebitenutil.DebugPrintAt(screen, "Spawn Type: Tracker", config.General.WindowSizeX - 160, 0)
		ebitenutil.DebugPrintAt(screen, "Left for SpawnX Left", config.General.WindowSizeX - 160, 60)
		ebitenutil.DebugPrintAt(screen, "Right for SpawnX Right", config.General.WindowSizeX - 160, 80)
		ebitenutil.DebugPrintAt(screen, "Top for SpawnY Up", config.General.WindowSizeX - 160, 100)
		ebitenutil.DebugPrintAt(screen, "Bottom for SpawnX Down", config.General.WindowSizeX - 160, 120)
	
	}

	if config.General.SpawnType == "Mouse"{
		ebitenutil.DebugPrintAt(screen, "Spawn Type:  Mouse", config.General.WindowSizeX - 160, 0)
		ebitenutil.DebugPrintAt(screen, "Right clic for Create", config.General.WindowSizeX - 160, 60)
	
	}

	if config.General.SpawnType == "Stack"{
		ebitenutil.DebugPrintAt(screen, "Spawn Type: Stack", config.General.WindowSizeX - 160, 0)
		ebitenutil.DebugPrintAt(screen, "Right clic for Create", config.General.WindowSizeX - 160, 60)
	}
	if config.General.SpawnType == "Random"{
		ebitenutil.DebugPrintAt(screen, "Spawn Type: Random", config.General.WindowSizeX - 160, 0)
	}
}
