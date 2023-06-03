package particles

import "container/list"

// System définit un système de particules.
// Spawn est utilisé pour servir de variable lorsque
// le SpawnRate est utilisé
type System struct {
	Content *list.List
	Spawn 	float64
}

// Particle définit une particule.
// Elle possède une position, une rotation, une vitesse, une taille, une couleur, une
// opacité et une durée de vie. 
// "RandomSpeed" sert à garder une valeur lorsque les particules sont aspirées (avec "MouseBlower")
// et que la vitesse est aléatoire
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	SpeedX, SpeedY					float64
	ScaleX, ScaleY                  float64
	RandomSpeedTracker				float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	Life							float64
}

