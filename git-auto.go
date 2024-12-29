package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage : git-auto commit <message>")
		return
	}


	subCommand := os.Args[1]
	message := os.Args[2] 

	if subCommand != "commit" {
		fmt.Println("Seule la sous-commande 'commit' est supportée pour le moment.")
		return
	}

	err := launchCommand(message)
	if err != nil {
		log.Fatalf("Erreur lors de l'exécution de la commande : %v", err)
	}

	fmt.Println("Commandes Git exécutées avec succès !")
}

func launchCommand(message string) error {
	commands := fmt.Sprintf("git add . && git commit -S -m '%s' && git push origin main", message)

	cmd := exec.Command("bash", "-c", commands)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}