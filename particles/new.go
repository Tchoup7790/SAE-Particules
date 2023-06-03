package particles

import (
	"container/list"
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math/rand"
	"math"
	"time"
)

// 'NewSystem()' est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// Elle créé un System et y ajoute le nombre de particule (grâce
// à la fonction 'NewListWithParticule()' (qu'on retrouve plus bas) donné avec 
// la variable 'InitNumParticles' présente dans le 'config.json'
func NewSystem() System {
	rand.Seed(time.Now().UnixNano())

	l := list.New()

	NewListWithParticule(l,int(config.General.InitNumParticles))

	return System{Content: l}
}

// La fonction 'NewListWithParticule()' sert à créé une particule dans une liste
// en prenant en compte les options du 'config.json' grâce à la fonction 
// 'VarInit()' créé dans le fichier 'random.go'.
// La variable 'a' est utilisé pour avoir une direction aléatoire lors de la creation de la particule.
// La variable 'RandomSpeedTracker' sert à garder une variable lors que l'option mouseTacker est activer
func NewListWithParticule(l *list.List, n int) {
	for i:=0; i!=n ; i++{
		a:=rand.Float64()* 2 * math.Pi

		p := Particle{
			PositionX: config.General.SpawnX,
			PositionY: config.General.SpawnY,
			SpeedX: math.Cos(a) * config.General.Speed , 
			SpeedY: math.Sin(a) * config.General.Speed ,
			RandomSpeedTracker: 10 * rand.Float64(),
			ScaleX: config.General.ScaleX, 
			ScaleY: config.General.ScaleY,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
			Life: config.General.LifeSpan * 10,
		}
		
		p.VarInit()

		l.PushFront(&p)
	}
}


// 'createSpawnRate()' permet de créer un nombre de particules donnés par le 'SpawnRate'
// présent dans le config.json 
func (s *System) createSpawnRate() {
	s.Spawn += config.General.SpawnRate
	if s.Spawn >= 1 {
		l := list.New()

		NewListWithParticule(l, int(s.Spawn))

		s.Content.PushFrontList(l)

		_, s.Spawn = math.Modf(s.Spawn)
	}

	l := list.New()

	NewListWithParticule(l, int(config.General.SpawnRate))

	s.Content.PushFrontList(l)
}


// 'createMouseRate()' permet de créer un nombre de particules donnés par le 'SpawnRate'
// présent dans le config.json à la position de la souris
func (s *System) createMouseRate() {
	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight) > 0 {
		mouseX, mouseY := ebiten.CursorPosition()
			
		config.General.SpawnX, config.General.SpawnY = float64(mouseX), float64(mouseY)

		s.createSpawnRate()
	}
}

// 'createTrackerRate()' permet de créer un nombre de particules donnés par le 'SpawnRate'
// présent dans le config.json et de le déplacer grâce au flèche du clavier
func (s *System) createTrackerRate() {
	s.createSpawnRate()
	if config.General.SpawnX >= 10 {
		if inpututil.KeyPressDuration(ebiten.KeyLeft) > 0 {
			config.General.SpawnX -= 7
		}
	}
	if  config.General.SpawnX < float64(config.General.WindowSizeX)-10 { 
		if inpututil.KeyPressDuration(ebiten.KeyRight) > 0 {
			config.General.SpawnX += 7
		}
	}
	if config.General.SpawnY >= 10 {
		if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
			config.General.SpawnY -= 7
		}
	}
	if config.General.SpawnY < float64(config.General.WindowSizeY)-10 {
		if inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0 {
			config.General.SpawnY += 7
		}
	}
}


// 'createStack()' permet de créer un paquet de particules
// à l'endroit de la souris
func (s *System) createStack() {
	config.General.InitNumParticles = 0

	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonRight) > 0 {
		mouseX, mouseY := ebiten.CursorPosition()
			
		config.General.SpawnX, config.General.SpawnY = float64(mouseX), float64(mouseY)

		for i:=0; i!=5; i++{
			l := list.New()

			NewListWithParticule(l, i*10)

			s.Content.PushFrontList(l)
		}
	}
}