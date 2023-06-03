package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"project-particles/config"
)

// ParticleImage est une variable globale pour stocker l'image d'une particule
var ParticleImage *ebiten.Image

// Get charge en mémoire l'image de la particule. (particle.png)
func Get() {
	var err error
	ParticleImage, _, err = ebitenutil.NewImageFromFile(config.General.ParticleImage)
	if err != nil {
		log.Fatal("Problem while loading particle image: ", err)
	}
}
