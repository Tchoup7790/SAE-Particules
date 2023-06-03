package particles

import (
	"testing"
	"project-particles/config"
	"fmt"
)

// Cette fonction permet de récupérer les informations
// du fichier config.json pour les prochains test
func init() {
	config.Get("../config.json")
}

// Ce test permet de vérifier qu'avec la fonction 
// 'NewListWithParticule()' on ajoute bien 
// une nouvelle particule dans le systeme
func TestAddParticuleInSys(t *testing.T) {
	s := NewSystem()

	l := s.Content.Len()

	NewListWithParticule(s.Content, 1) 

	if l+1 != s.Content.Len() {
		t.Error("\nLe NewListWithParticule() ne fonctionne pas\n ")
		t.Log("\n nombre avant l'appel de la fonction : ", l, "nombre après l'appel de la fonction : ", s.Content.Len())
	}
	fmt.Println("\nTest Add Particule In Sys : OK\nLe NewListWithParticule() fonctionne\n ")
}


// Ce test permet de vérifier que le nombre de particule
// demandé au départ est bien le nombre créé
func TestGoodLen(t *testing.T) {
	s := NewSystem()

	if s.Content.Len() != config.General.InitNumParticles{
		t.Error("Le nombre de particules de départ n'est pas celui demandé")
		t.Log("\n nombre demandé : ", config.General.InitNumParticles, "\n nombre créé : ", s.Content.Len(), "\n ")
	}
	fmt.Println("Test Good Len : OK\nLe nombre de particules de départ est celui demandé\n ")
}

// Ce test permet de vérifier que les particules ont la vitesse demandé
// si RandomSpeed est désactivé
func TestGoodSpeed(t *testing.T) {
		config.General.InitNumParticles = 2

		s := NewSystem()

		p := s.Content.Front()

		p1 := p.Value.(*Particle)

		p.Next()

		p2 := p.Value.(*Particle)

		if !config.General.RandomSpeed {
			if p1.SpeedX * p1.SpeedY != p2.SpeedX * p2.SpeedY{
				t.Error("Les vitesses sont aléatoire\n ")
				t.Log("\n particule 1 Vitesse X : ", p1.SpeedX, "particule 1 Vitesse Y : ", p1.SpeedY,"\n particule 2 Vitesse X : ", p2.SpeedX, "particule 2 Vitesse Y : ", p2.SpeedY)
			}
			fmt.Println("Test Good Speed : OK\nLes vitesses ne sont pas aléatoire\n ")
		}
}

// Ce test permet de vérifier que les particules mortent 
// soient bien supprimé
func TestGoodDeath(t *testing.T) {
	config.General.InitNumParticles = 1

	config.General.Life = true

	config.General.LifeSpan = 1

	s := NewSystem()

	part := s.Content.Front()

	p := part.Value.(*Particle)

	len := s.Content.Len()

	s.Update()

	if len > s.Content.Len() {
		t.Error("La particule n'est pas supprimé")
		t.Log("\n particule Life : ", p.Life, "\n nombre de particule : ", s.Content.Len())
	}
	fmt.Println("Test Good Death : OK \nLa particule est bien supprimé")
}

