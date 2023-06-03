package particles

import (
	"project-particles/config"
)

// 'Update()' mets à jour l'état du système de particules à chaque pas de temps. 
// Elle est appellée exactement 60 fois par seconde (de manière régulière) 
// par la fonction principale du projet.
// Elle change la position de la Particle grâce à la vitesse et créé le nombre 
// de particule demandé suivant le "SpawnType" demandé.
// Elle vérifie aussi que le temps de vie restant à la particule est à 0 pour 
// la supprimée et prend en compte les potentiels changement à appliquer
// sur la particule grâce à 'VarUpdate()'
func (s *System) Update() {

	if config.General.SpawnType == "Mouse" {
		s.createMouseRate()
	}
	if config.General.SpawnType == "Tracker" {
		s.createTrackerRate()
	}
	if config.General.SpawnType == "Stack" {
		
		config.General.RandomSpeed = false

		config.General.Speed = 12

		s.createStack()
	}


	for e := s.Content.Front(); e != nil; {
		p, ok := e.Value.(*Particle)

		if !ok {
			continue
		}
		
		p.VarUpdate()

		p.PositionX +=  p.SpeedX
		p.PositionY +=  p.SpeedY

		next := e.Next()

		if p.Life <=0 {
			s.Content.Remove(e)
		}

		e = next
	}
}
