package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type WordFrenchResponse struct {
	Name string `json:"name"`
}

func GetFrenchWord() string { //Titouan
	response, err := http.Get("https://trouve-mot.fr/api/random")
	if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
		os.Exit(0)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		os.Exit(0)
	}

	var wordResponse []WordFrenchResponse

	if err := json.Unmarshal(body, &wordResponse); err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		os.Exit(0)
	}

	if len(wordResponse) > 0 {
		// Afficher le mot
		return removeAccents(wordResponse[0].Name)
	} else {
		fmt.Println("Aucun élément trouvé dans la réponse JSON")
		return ""
	}
}

func GetEnglishWord() string { //Titouan
	response, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		fmt.Println("Erreur lors de la requête GET:", err)
		os.Exit(0)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse:", err)
		os.Exit(0)
	}

	var wordResponse []string

	if err := json.Unmarshal(body, &wordResponse); err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		os.Exit(0)
	}

	if len(wordResponse) > 0 {
		// Return the first element of the array
		return strings.ReplaceAll(removeAccents(wordResponse[0]), " ", "")
	} else {
		fmt.Println("Aucun élément trouvé dans la réponse JSON")
		return ""
	}
}

func removeAccents(input string) string { //Lisa
	accentMap := map[rune]rune{
		'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ä': 'a', 'å': 'a',
		'ç': 'c',
		'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e',
		'ì': 'i', 'í': 'i', 'î': 'i', 'ï': 'i',
		'ñ': 'n',
		'ò': 'o', 'ó': 'o', 'ô': 'o', 'õ': 'o', 'ö': 'o',
		'ù': 'u', 'ú': 'u', 'û': 'u', 'ü': 'u',
		'ý': 'y',
		'ÿ': 'y',
	}

	inputRunes := []rune(input)

	for i, r := range inputRunes {
		if replacement, exists := accentMap[r]; exists {
			inputRunes[i] = replacement
		}
	}

	result := string(inputRunes)

	return result
}
