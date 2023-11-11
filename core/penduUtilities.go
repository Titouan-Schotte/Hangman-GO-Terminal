package core

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

func GenerateCryptedSolution(solution string, lettres_trouvees string) string { // Jemima & Manel
	affichage := ""
	for _, l := range solution {
		if strings.Contains(lettres_trouvees, string(l)) {
			affichage += string(l) + " "
		} else {
			if string(l) == " " {
				affichage += " "
			} else if string(l) == "-" {
				affichage += "- "
			} else {
				affichage += "_ "
			}
		}
	}
	return affichage
}

func GiveRandomSolutionsLetters(solution string) string { //Manel
	if len(solution) < 3 {
		return ""
	}
	revealedLetters := ""
	for i := 0; i < len(solution)/2-1; i++ {
		for true {
			choosenLetter := string(solution[rand.Intn(len(solution))])
			if !strings.Contains(revealedLetters, choosenLetter) && strings.Count(solution, choosenLetter) == 1 {
				revealedLetters += choosenLetter
				break
			}
		}

	}
	return revealedLetters
}

func LettreAscii(lettreByte int32, lettresAscii []string) string { //Lisa
	if int(lettreByte)-32 >= 0 && int(lettreByte)-32 < len(lettresAscii) {
		return lettresAscii[lettreByte-32]
	}
	// Gérer le cas où l'index est hors limites
	return ""
}

func GetAsciiCryptendSolution(affichage string) string { //Lisa
	//Charger et mettre les donnés du fichier dans une chaine
	affichageascii := ""
	fileData, _ := ioutil.ReadFile("standard.txt")
	lettresAscii := strings.Split(string(fileData), "\n\r\n")
	var listLettresAsciiChoisies []string
	for _, lettreByte := range affichage {
		listLettresAsciiChoisies = append(listLettresAsciiChoisies, LettreAscii(lettreByte, lettresAscii))
	}
	//Initialiser une liste pour stocker le résultat finale
	var output []string
	for i, letter := range listLettresAsciiChoisies {
		letter = strings.ReplaceAll(letter, "\r", "") // Pour supprimer le \r du fichier et éviter les bugs
		temp := strings.Split(letter, "\n")           // séparer les lignes de chaque lettre de ma liste précedente
		if i == 0 {
			//initialiser ma liste finale avec les lignes de ma premiere lettre pour eviter les append
			if affichage[i] == '_' {
				for i, _ := range temp {
					temp[i] = Colorize(Gray, temp[i])
				}
			} else {
				for i, _ := range temp {
					temp[i] = Colorize(Blue, temp[i])
				}
			}

			output = temp
		} else {
			for k := range output {
				//rajouter les lignes des caractéres suivant avec un espace avant
				if k == 0 {
					output[k] += "   "
				}
				if string(affichage[i]) == "_" {
					output[k] += Colorize(Gray, temp[k])
				} else {
					output[k] += Colorize(Blue, temp[k])
				}

			}
		}
	}
	// le résultat finale
	for _, line := range output {
		affichageascii += line + "\n"

	}
	return affichageascii
}

func PositionPendu(tentatives int) string { //Lisa
	var finalPendu string
	hangmanPositions, _ := ioutil.ReadFile("hangman.txt")
	if strings.ContainsAny(string(hangmanPositions), "\r") {
		hangmanPositions = []byte(strings.ReplaceAll(string(hangmanPositions), "\r", "\n"))
		hangmanPositions = []byte(strings.ReplaceAll(string(hangmanPositions), "\n\n", "\n"))
	}
	hangman := strings.Split(string(hangmanPositions), "\n\n")[10-tentatives-1]
	for _, l := range hangman {
		toPrint := string(l)
		switch toPrint {
		case "=":
			toPrint = Colorize(Red+Bold, toPrint)
		case "|", "/", "\\", "":
			toPrint = Colorize(Yellow+Bold, toPrint)
		case "-":
			toPrint = Colorize(Yellow+Bold, toPrint)
		case "+":
			toPrint = Colorize(Red, toPrint)
		}
		finalPendu += toPrint
	}
	finalPendu += "\n\n"
	return finalPendu
}

func Win(solution string, spectate Party) { //Titouan
	ClearConsole()
	buff := ""
	buff += fmt.Sprintln(Colorize(Green+Purple, "\t  ====================="))
	buff += fmt.Sprintln("\t       * " + Colorize(Yellow, "YOU WIN") + " *    ")
	buff += fmt.Sprintln(Colorize(Green+Purple, "\t  =====================\n"))
	buff += fmt.Sprintln(Colorize(LightYellow, "\t            o \n\t         o^/|\\^o\n\t      o_^|\\/*\\/|^_o\n\t     o\\*`'.\\|/.'`*/o\n\t      \\\\\\\\\\\\|//////\n\t       {><><@><><}\n\t       `\"\"\"\"\"\"\"\"\"`\n"))
	buff += fmt.Sprintln("Le mot était : " + Colorize(Blue+Bold, solution))
	fmt.Print(buff)
	spectate.AddStep(-1, buff)
	Input(Colorize(Purple, "-> "))
	New_game(spectate)
}

func GameOver(solution string, spectate Party) { //Titouan
	ClearConsole()
	buff := ""
	buff += fmt.Sprintln(Colorize(Bold+Purple, "\t  ====================="))
	buff += fmt.Sprintln("\t       * " + Colorize(Red+Bold, "GAME OVER") + " *    ")
	buff += fmt.Sprintln(Colorize(Bold+Purple, "\t  =====================\n"))
	buff += fmt.Sprintln(Colorize(Gray+Bold, "\t#########################\n\t#                       #"))
	for i, line := range strings.Split(PositionPendu(0), "\n") {
		buff += fmt.Sprintln("\t" + Colorize(Gray, "#") + "\t" + line + "\t" + Colorize(Gray, "#"))
		if i > 5 {
			break
		}
	}
	buff += fmt.Sprint(Colorize(Gray+Bold, "\t#                       #\n\t#########################\n\n\n"))

	buff += fmt.Sprintln("Le mot était : " + Colorize(Blue+Bold, solution))
	fmt.Print(buff)
	spectate.AddStep(-1, buff)
	Input(Colorize(Purple, ""+Colorize(Purple, "->")+" "))
	New_game(spectate)
}
