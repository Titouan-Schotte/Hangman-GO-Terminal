package core

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func Read_words(path string) []string { //Jemima
	fichier, a := ioutil.ReadFile(path)
	if a != nil {
		fmt.Print(Colorize(Red+Bold, "Erreur de lecture du fichier !"))
		os.Exit(0)
	}
	mots := strings.Split(string(fichier), "\n")
	return mots
}

func Choose_word(path string) string { //Manel
	mots := Read_words(path)
	return mots[rand.Intn(len(mots))]
}
