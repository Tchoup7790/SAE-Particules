package config

// Config définit les champs qu'on peut trouver dans un fichier de config.json
type Config struct {
	Debug                    bool
	WindowTitle              string
	FullScreen				 bool
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	InitNumParticles         int
	SpawnRate				 float64
	SpawnType				 string
	SpawnX, SpawnY           float64
	RandomColor				 bool
	RandomOpacity			 bool
	ScaleX, ScaleY 			 float64	
	RandomSpeed				 bool
	Speed					 float64
	DeathOut				 bool
	Life 					 bool
	DeathScale				 bool
	LifeSpan				 float64
	Gravity				 	 float64
	Bounce					 bool
	BounceColor				 bool
	MouseBlower				 bool
	JailMouse				 bool
}

var General Config
 