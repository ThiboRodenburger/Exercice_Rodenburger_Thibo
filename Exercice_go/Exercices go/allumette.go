package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var allumettes, choix, tour int
	rand.Seed(time.Now().UnixNano())
	for allumettes = 20; allumettes > 0; {
		tour = (tour + 1) % 2
		fmt.Printf("Il reste %d allumettes\n", allumettes)
		if tour == 0 {
			fmt.Println("C'est votre tour ! Combien d'allumettes voulez-vous retirer ?")
			fmt.Scanln(&choix)
			if choix > 3 || choix < 1 || choix > allumettes {
				fmt.Println("Nombre invalide d'allumettes ! Veuillez choisir entre 1 et 3 allumettes.")
				continue
			}
		} else {
			choix = rand.Intn(3) + 1
			if choix > allumettes {
				choix = allumettes
			}
			fmt.Printf("L'ordinateur retire %d allumettes\n", choix)
		}
		allumettes -= choix
		if allumettes == 0 {
			if tour == 0 {
				fmt.Println("Vous avez perdu !")
			} else {
				fmt.Println("L'ordinateur a perdu !")
			}
		}
	}
}
