# R1.01.SAE.eq06_EVENO-Jules_JULIO-Baptiste

Projet Particules, permettant de créé des particules et de les modifier.


## Options modifiables grâce au config.json

- **Debug**: permet de voir le nombre de Particule 'Vivant' et le TPS 
- **WindowTitle**: correspond au nom de la fenêtre 
- **FullScreen** : permet d'avoir la fénêtre en plein écran ou non
- **WindowSizeX**, **WindowSizeY**: permet de changer la taille de la fenêtre lorsque le **FullScreen** est désactivé 
- **ParticleImage**: correspond au nom du fichier utilisé comme particule
- **InitNumParticles**: sert à choisir le nombre de particule de départ
- **SpawnRate**: nombre de particule crée (exploité de différente manière selon le **SpawnType**)
- **SpawnType**: change certaines options selon le mode de Spawn (mode disponible : "Stack", "Tracker", "Random", "Mouse") (*Si ne marche pas vérifier que la majuscule est bien présente*)
        - "Stack" : fait apparaître des particules par paquet quand on appuye sur le clique droit (**SpawnRate** n'est pas utilisé) (créé les particules de départ à **SpawnX**, **SpawnY**)
        - "Tracker" : fait apparaître régulierement un nombre de particules donné par **SpawnRate**, le point d'apparition peut bouger grâce au flèche du clavier (créé les particules de départ à **SpawnX**, **SpawnY**) (désactive la vitesse aléatoire et met **Speed** à 12)
        - "Random" : fait apparaître les particules de manière aléatoire
        - "Mouse" : fait apparaître un nombre de particules donné par **SpawnRate** quand on appuye sur le clique droit (créé les particules de départ à **SpawnX**, **SpawnY**)
- **SpawnX**, **SpawnY**: définie les coordonnées où apparait la particule (n'est pas utilisé si **SpawnType** est *Random*) (si X (ou Y) est inférieur à 0, la coordonnée est mis au centre de son axe)
- **RandomColor**: décide si la couleur de la particule est aléatoire (blanc si l'option est *false*)
- **RandomOpacity**: permet d'activer l'opacité aléatoire (est égal à un si désactivé)
- **ScaleX**, **ScaleY**: permet de définir la largeur et longueur de la particule
- **RandomSpeed**: décide si la vitesse de la particule est aléatoire
- **Speed**: choisir la vitesse ( fonctionne si **RandomSpeed** est *false*)
- **DeathOut**: si actif, porovque la "mort" de la particule lorsqu'elle sort de l'écran
- **Life**: permet d'activer la mort ou non de la particle
- **DeathScale**: active la diminution de la taille de la particule avant sa mort (fonctionne si **Life** est sur *true*)
- **LifeSpan**: permet de définir la durée de vie
- **Gravity**: permet de choisir l'intensité de la gravité (désactivé si égal à 0)
- **Bounce**: active le rebond de la particule sur le rebord de la fenêtre
- **BounceColor**: sert à activer le changement de couleur lors du rebond (fonctionne si **Bounce** est *true*)
- **MouseBlower**: une fois activé, les particules sont comme "aspiré" par la souris lorsque le clique gauche est maintnenue et "dispersé" lorsque le bouton est relaché (met **Gravity** à 0)
- **MouseJail**: active une prison dans un rayon de 100 pixels avec la bar espace (désactive **RandomColor**) (les particules en "prison" seront orange)


## ***Description des fichiers modifiés***

## main.go

Le fichier main.go a été légerement modifié afin que le champ **FullScreen**, présent dans le *config.json* fonctionne. Cela permet l'affichage de la fenêtre en pleine écran et la création des particule au centre de l'écran.
Aussi, il permet de modifié SpawnX et SpawnY si ils sont inférieur à 0 pour les mettre au centre de leurs axe respectif.

## draw.go

Le fichier draw.go a été légerement modifié afin que, lorsque que le mode debug est activé, on affiche le TPS arrondis et le nombre de particule du systeme.
Il permet aussi d'afficher les instructions selon le **SpawnType** choisi.


## config.json

Ce fichier a reçus des modification afin d'ajouter ou utiliser des options tels que le mode d'apparition de la particule ou bien le nombre de particule au départ.


## config/type.go

Ce fichier sert à définir les variables présentent dans le config.json et qui seront utilisées dans tout le projet.

**Les prochains fichiers se trouvent dans le dossier** ***'particles/'***

## new.go

Ce fichier contient trois méthodes appelé ou non selon le **SpawnType** :
  - **NewSystem()**
  - **NewListWithParticule()** 
  - **createSpawnRate()**
  - **createMouseRate()**
  - **createTrackerRate()**
  - **createStack()**

### NewSystem() 

Cette fonction sert tout d'abord à charger une nouvelle seed lorsqu'on utilisera le package rand. Autrement dit, cela permet que chaque nombre généré aléatoirement soit à nouveau défini lorsqu'on relance le projet.

Cette fonction sert aussi à créé un systeme au quel sera ajouté une liste et où sera ajouté des paticules grâce à la fonction **NewListWithParticule()** dont on parlera plus tard.

### createSpawnRate()

Cette méthode permet de créé le nombre de particules choisi dans un système défini par **SpawnRate** pésent dans le fichier *config.json* en prenant en compte les décimals si il y en a.

### createMouseRate()

Cette méthode permet de créé, à l'endroit de la souris, le nombre de particules choisi dans un système défini par **SpawnRate** pésent dans le fichier *config.json* en prenant en compte les décimals si il y en a.

### createTrackerRate()

Cette méthode permet de créé le nombre de particules choisi dans un système défini par **SpawnRate** pésent dans le fichier *config.json* en prenant en compte les décimals si il y en a.
Le point d'apparition peut être changer en tant réel grâce au flèche du clavier.

### createStack()

Cette méthode permet de créé des paquets de particules à l'endroit de la souris.


### NewListWithParticule()

Cette méthode permet la création d'une particule dans un système en la définissant. Elle appelle aussi la fonction **VarInit()** dont on parlera plus tard.


## random.go

Ce fichier contient deux méthodes :
  - **VarInit()**
  - **VarUpdate()**

### check_random()

Cette méthode sert a vérifier si une des options du *config.json* concerne la création de la particule et si c'est le cas applique cette option. Par exemple, la méthode vérifie si **RandomSpawn** présent de le *config.json* est *true* et si c'est le cas le position de départ de la particule est aléatoire.

### VarUpdate()

Cette méthode sert a vérifier si une des options du *config.json* concerne la mise à jour de la particule et si c'est le cas applique cette option. Par exemple, elle ajoute de la vitesse quand la gravité est supérieur à 0 et, si l'option **MouseJail** est sur *true*, vérifie si la bar espace est pressé pour créé une prison où les particules ne peuvent ni sortir ni rentrer.


## update.go

Ce fichier ne contient qu'une méthode, **Update()**.

**Update()** sert a, selon le**SpawnType**, créé des particules à l'endroit demandé en appelant une des mathodes de *new.go*.
La méthode vérifie grâce à une routine la valeur **Life** et si elle est égal à 0 supprime cette particule.


## type.go

Ce fichier contient deux structure :
  -**System**
  -**Particle**

### System

Cette structure contient deux chose, **Content** qui est la liste de particle contenu par le System et **Spawn** qui contiendra aidera lorsque **SpawnRate** est utilisé.

### Particle

Cette structure est la plus importane du projet, elle désigne la particle dans son ensemble avec toutes ses variables qui sont indispensables au projet.

## particles_test.go

Ce fichier est le fichier de test du projet.
Il contient 4 fonctions :
  - init()
  - TestAddParticuleInSys()
  - TestGoodLen()
  - TestGoodSpeed()

### init()

Cette fonction sert a implémenter les valeurs du json pour tout les tests.

### TestAddParticuleInSys()

Cette fonction teste que la fonction **NewListWithParticule()** ajoute bien une particule au systeme.

### TestGoodLen()

Cette fonction teste que le nombre de particule demandé au départ est bien le nombre créé.

### TestGoodSpeed()

Cette fonction teste que les particules ont bien la vitesse demandé si le **RandomSpeed** du *config.json* est désactivé.

### TestGoodDeath()

Cette fonction teste que les particules meurt bien avec la fonction update quand **Life** est égal à 0.



## ***Installation et Exécution***

Pour installer le projet, placez-vous dans le dossier sur un terminal et exécutez

```bash
  go install
```

Ensuite, pour chaque lancement, il suffit d'exécuter

```bash
  go build
  ./project-particles
```


## ***Running Tests***

Pour lancer les tests, rendez-vous dans le dossier *'particles'* et exécutez

```bash
  cd particles
  go test
```


## ***Support***

Pour nous contacter, envoyé un mail à baptiste.julio@etu.univ-nantes.fr ou jules.eveno-gallen@etu.univ-nantes.fr


