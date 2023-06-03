package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"project-particles/config"
	"math/rand"
	"math"
)

// la fonction 'VarInit()' sert à changer la particule si
// la condition est vérifié et si l'option est activé dans le 
// fichier config.json au moment où elle est appelée
func (p *Particle) VarInit(){
	if config.General.SpawnType == "Random" {
		p.PositionX, p.PositionY = rand.Float64() * (float64(config.General.WindowSizeX)),rand.Float64() * (float64(config.General.WindowSizeY))
	}
	if config.General.RandomSpeed {
		p.SpeedX *= rand.Float64() 
		p.SpeedY *= rand.Float64()
	}
	if config.General.RandomColor{
		p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
	}
	if config.General.RandomOpacity{
		p.Opacity = rand.Float64()
	}
}

// la fonction 'VarUpdate()' sert à changer la particule si
// la condition est vérifié et si l'option est activé dans le 
// fichier config.json à chaque appel
// Elle permet, par exemple, de vérifier si la particule sort 
// de l'écran ou d'aspirer la particule si l'option est activé
// et que le click gauche est maintenu
func (p *Particle) VarUpdate(){
	if config.General.Bounce{
		if p.PositionX <= 0 || p.PositionX >= float64(config.General.WindowSizeX)-p.ScaleX*10 {
			p.SpeedX *= -1
			if config.General.BounceColor{
				p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
			}
		}
		if p.PositionY <= 0 || p.PositionY >= float64(config.General.WindowSizeY)-p.ScaleY*10 {
			p.SpeedY *= -1
			if config.General.BounceColor{
				p.ColorRed, p.ColorGreen, p.ColorBlue = rand.Float64(), rand.Float64(), rand.Float64()
			}
		}
	}	

	if config.General.DeathOut{
		if p.PositionX < -p.ScaleX*10  || p.PositionX > float64(config.General.WindowSizeX) || p.PositionY < -p.ScaleY*10 || p.PositionY > float64(config.General.WindowSizeY) {
			p.Life = 0
		}
	}

	if config.General.MouseBlower{

		config.General.Gravity = 0

		if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0{

			mouseX, mouseY := ebiten.CursorPosition()
		   
		    distanceX := float64(mouseX) - p.PositionX
			distanceY := float64(mouseY) - p.PositionY
			    
			distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
			distanceX /= distance
			distanceY /= distance
			    
			p.SpeedX = distanceX * config.General.Speed
			p.SpeedY = distanceY * config.General.Speed 	

			if config.General.RandomSpeed {
				p.SpeedX = distanceX * math.Abs(p.RandomSpeedTracker)
				p.SpeedY = distanceY * math.Abs(p.RandomSpeedTracker)
			}else{
				p.SpeedX = distanceX * config.General.Speed
				p.SpeedY = distanceY * config.General.Speed 
			}

			if p.PositionX == float64(mouseX) {
				p.SpeedX = 0
			}
			if p.PositionY == float64(mouseY) {
				p.SpeedY = 0
			}
		}
	}

	if config.General.JailMouse{
		p.ColorRed, p.ColorGreen, p.ColorBlue = 1, 1, 1

		config.General.RandomColor = false

		if inpututil.KeyPressDuration(ebiten.KeySpace) > 0{

			mouseX, mouseY := ebiten.CursorPosition()
		   
		    distanceX := float64(mouseX) - p.PositionX
			distanceY := float64(mouseY) - p.PositionY
			    
			distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
			    
			if distance < 105 && distance >= 95{
				if p.PositionX <= float64(mouseX)+100 || p.PositionX >= float64(mouseX)-100-p.ScaleX*10 {
					p.SpeedX *= -1
				}
				if p.PositionY <= float64(mouseX)+100 || p.PositionY >= float64(mouseX)-100-p.ScaleY*10 {
					p.SpeedY *= -1
				}
			} 
			if distance <= 100 {
				p.ColorRed, p.ColorGreen, p.ColorBlue = 0.95, 0.49, 0.24
			}
		}
	}
	if config.General.Life {
		if config.General.DeathScale {
			if p.Life <= 10 && p.Life >= 1{
				p.ScaleX -= 0.1
				p.ScaleY -= 0.1
			}
		}
		p.Life-- 
	}

	p.SpeedY += config.General.Gravity * 1/6
}
