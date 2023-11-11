package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HangManData struct {
	Lettres_trouvees string // Word composed of '_', ex: H_ll_
	Solution         string // Final word chosen by the program at the beginning. It is the word to find
	Tentatives       int    // Number of attempts left
}

func LoadMainMenu() { //Titouan
	ClearConsole()
	texte := Colorize(Green+Bold, "  _    _                   __  __             \n | |  | |                 |  \\/  |            \n | |__| | __ _ _ __   __ _| \\  / | __ _ _ __  \n |  __  |/ _` | '_ \\ / _` | |\\/| |/ _` | '_ \\ \n | |  | | (_| | | | | (_| | |  | | (_| | | | |\n |_|  |_|\\__,_|_| |_|\\__, |_|  |_|\\__,_|_| |_|\n                      __/ |                   \n                     |___/                    \n\n\n")

	// Effacer le terminal avant d'afficher le texte
	ClearConsole()

	// Définir le délai entre chaque étape du serpent
	delay := 0.1

	// Animation en forme de serpent
	for i := 1; i <= len(texte); i += 15 {
		ClearConsole()
		fmt.Print(texte[:i])
		Timeout(delay)

	}
	for i := 9; i > -1; i-- {
		ClearConsole()
		fmt.Print(texte)
		fmt.Printf(Colorize(Gray, "\t#########################\n\t#                       #\n"))
		for i, line := range strings.Split(PositionPendu(i), "\n") {
			fmt.Println("\t" + Colorize(Gray, "#") + "\t" + line + "\t" + Colorize(Gray, "#"))
			if i > 5 {
				break
			}
		}
		fmt.Printf(Colorize(Gray, "\t#                       #\n\t#########################\n\n\n"))
		Timeout(delay * 2)
	}
	hangmanIn := HangManData{}
	switch len(os.Args) {
	case 1:
		fmt.Println(Colorize(LightYellow, "1") + " > Créer une nouvelle partie")
		fmt.Println(Colorize(Yellow, "2") + " > Charger une partie")
		fmt.Println(Colorize(Orange, "3") + " > Lancer le mode spectateur")
		inp := Input("\n" + Colorize(Purple, "->") + "")
		ClearConsole()
		switch inp {
		case "1":
			CreateParty(hangmanIn)
		case "2":
			LoadParty(hangmanIn)
		case "3":
			SpectateInit(hangmanIn)
		default:
			LoadMainMenu()
		}
	case 2:
		hangmanIn.Solution = Choose_word(os.Args[1])
		hangmanIn.Tentatives = 10
		hangmanIn.Lettres_trouvees = GiveRandomSolutionsLetters(hangmanIn.Solution)
		Game(hangmanIn)
	case 3:
		if os.Args[1] == "--startWith" {
			fileData, err := ioutil.ReadFile("saves/" + os.Args[2] + ".json")
			if err != nil {
				fmt.Print(Colorize(Red+Bold, "Nom de la sauvegarde invalide !"))
				return
			}
			json.Unmarshal(fileData, &hangmanIn)
		} else {
			fmt.Print(Colorize(Red+Bold, "Argument "+os.Args[1]+" invalide (veuillez taper --startWith)"))
		}
		Game(hangmanIn)
	}
}

func SpectateInit(hangmanIn HangManData) { //Titouan
	ClearConsole()
	files, _ := ioutil.ReadDir("./spectates/")
	var fileNames []string
	if len(files) == 0 {
		fmt.Println("Vous avez actuellement aucune partie en mode spectateur à charger !\nNous allons donc créer votre première partie !")
		CreateParty(hangmanIn)
		return
	}
	fmt.Println(Colorize(Underline+Bold, "Quelle partie en mode spectateur voulez vous charger :\n"))
	for i, v := range files {
		fileName := strings.Split(v.Name(), ".")[0]
		fmt.Printf("%v > %v", Colorize(Yellow, i+1), fileName)
		fmt.Println()
		fileNames = append(fileNames, fileName)
	}
	for inp := Input("" + Colorize(Purple, "->") + ""); true; inp = Input("" + Colorize(Purple, "->") + "") {
		if len(inp) == 0 {
			fmt.Println("Veuillez réessayer !")
			continue
		}
		inpInt, err := strconv.Atoi(string(inp[0]))
		if err == nil {
			if len(fileNames) >= inpInt && inpInt > 0 {
				ClearConsole()
				fmt.Print(Colorize(Red, "CHARGEMENT MODE SPECTATEUR"))
				Timeout(0.5)
				fmt.Print(Colorize(Red, "."))
				Timeout(0.5)
				fmt.Print(Colorize(Red, "."))
				Timeout(0.5)
				fmt.Print(Colorize(Red, "."))
				Spectate(fileNames[inpInt-1])
			}
		}
	}
}
func LoadParty(hangmanIn HangManData) { //Titouan
	ClearConsole()
	files, _ := ioutil.ReadDir("./saves/")
	var fileNames []string
	if len(files) == 0 {
		fmt.Println("Vous avez actuellement aucune partie à charger !\nNous allons donc créer votre première partie !")
		CreateParty(hangmanIn)
		return
	}
	fmt.Println(Colorize(Underline+Bold, "Quelle partie voulez vous charger :\n"))
	for i, v := range files {
		fileName := strings.Split(v.Name(), ".")[0]
		fmt.Printf("%v > %v", Colorize(Yellow, i+1), fileName)
		fmt.Println()
		fileNames = append(fileNames, fileName)
	}
	for inp := Input("" + Colorize(Purple, "->") + ""); true; inp = Input("" + Colorize(Purple, "->") + "") {
		if len(inp) == 0 {
			fmt.Println("Veuillez réessayer !")
			continue
		}
		inpInt, err := strconv.Atoi(string(inp[0]))
		if err == nil {
			if len(fileNames) >= inpInt && inpInt > 0 {
				ClearConsole()
				fmt.Print(Colorize(Yellow, "CHARGEMENT EN COURS"))
				Timeout(0.5)
				fmt.Print(Colorize(Yellow, "."))
				Timeout(0.5)
				fmt.Print(Colorize(Yellow, "."))
				Timeout(0.5)
				fmt.Print(Colorize(Yellow, "."))
				Timeout(0.7)
				ClearConsole()
				fileData, erre := ioutil.ReadFile("saves/" + fileNames[inpInt-1] + ".json")
				if erre != nil {
					fmt.Print(Colorize(Red+Bold, "Nom de la sauvegarde invalide !"))
					return
				}
				json.Unmarshal(fileData, &hangmanIn)
				Game(hangmanIn)
			}
		}
	}
}
func CreateParty(hangmanIn HangManData) { //Titouan
	switch Input("Choisissez votre langage (" + Colorize(Blue, "f") + Colorize(Red, "r") + " / " + Colorize(LightRed, "e") + Colorize(LightBlue, "n") + ")\n" + Colorize(Purple, "->") + "") {
	case "fr":
		hangmanIn.Solution = GetFrenchWord()
	case "en":
		hangmanIn.Solution = GetEnglishWord()
	default:
		ClearConsole()
		CreateParty(hangmanIn)
		return
	}
	ClearConsole()
	hangmanIn.Tentatives = 10
	hangmanIn.Lettres_trouvees = GiveRandomSolutionsLetters(hangmanIn.Solution)
	Game(hangmanIn)
}
func Game(hangmanIn HangManData) { //Titouan & Lisa & Manel & Jemima
	//Transformation du mot en affichage crypté pour le joueur (afin qu'il ne sache pas quel est le mot)
	affichage := GenerateCryptedSolution(hangmanIn.Solution, hangmanIn.Lettres_trouvees)
	affichageAscii := GetAsciiCryptendSolution(affichage)
	spectate := Party{}
	spectate.InitSpectate(hangmanIn.Solution)
	//Liste des lettres déjà utilisées
	alreadyUsed := strings.Split(hangmanIn.Lettres_trouvees, "")
	for hangmanIn.Tentatives > 0 {
		response := ""
		response += fmt.Sprintf("Il vous reste %s tentatives pour réussir !\n", Colorize(Purple, hangmanIn.Tentatives))
		response += fmt.Sprint(Colorize(Bold, "Mot à deviner : \n") + affichageAscii + "\n")
		//Affichage des lettres déjà proposées
		if len(alreadyUsed) > 0 {
			response += fmt.Sprintln(Colorize(UnderlineBold, "\n\nLettres déjà proposées :"))
			for i, letter := range alreadyUsed {
				if i%4 == 0 && i != 0 {
					response += "\n"
				}
				response += fmt.Sprint(Colorize(Pink, string(letter)) + "\t")
			}
		}
		response += "\n\nProposez une lettre : "
		fmt.Println(response)
		var proposition, buff string
		isContinuing := true
		//Obtention d'une entrée utilisateur pour la lettre concernée correcte
		for isContinuing {
			proposition = Input(" " + Colorize(Purple, "->") + " ")
			buff = "\n" + Colorize(Purple, "->") + "" + Colorize(Pink, proposition) + "\n\n"
			if proposition == "" {
				continue
			}
			//Stop hangman
			if proposition == "stop" {
				ClearConsole()
				file, _ := json.MarshalIndent(hangmanIn, "", " ")
				_ = os.WriteFile("saves/"+Input("Donnez le nom de votre sauvegarde :\n"+Colorize(Purple, "->")+"")+".json", file, 0644)
				fmt.Print(Colorize(Blue+Bold, "Fin du jeu, votre partie a bien était sauvegardée !"))
				Timeout(3)
				LoadMainMenu()
			}
			//Vérifier si c'est un mot qui est donné
			if len(proposition) >= 2 {
				if proposition == hangmanIn.Solution {
					Win(hangmanIn.Solution, spectate)
				} else {
					hangmanIn.Tentatives--
				}
			}
			//Permet de traiter l'entrée utilisateur et de soit mettre un message d'erreur (car les pré-conditions ne sont pas réspectées) ou bien de sortir de la boucle quand l'input est correct.
			if !strings.ContainsAny(proposition, "abcdefghijklmnopqrstuvwxyz") {
				fmt.Print(Colorize(Red, "Veuillez entrer une lettre uniquement !\n"))
			} else {
				//Ajout de la proposition à la liste des lettres proposées par l'utilisateur
				if slices.Contains(alreadyUsed, proposition) {
					fmt.Println(Colorize(LightBlue, "Vous avez déjà proposé cette lettre"))
					continue
				}
				alreadyUsed = append(alreadyUsed, proposition)
				isContinuing = false
			}
		}
		ClearConsole()
		//Ajout de la dernière proposition faite (qui correspond à une solution qu'on considère comme correcte)
		response += buff
		//On reset la variable pour pouvoir la réutiliser
		buff = ""
		//Comparaison à la solution
		if strings.Contains(hangmanIn.Solution, proposition) {
			hangmanIn.Lettres_trouvees += proposition
			buff += fmt.Sprintf(""+Colorize(Purple, "->")+" Bravo vous avez trouvé la lettre \"%s\"\n\n", Colorize(Yellow, proposition))
		} else {
			hangmanIn.Tentatives -= 1
			buff += fmt.Sprintln(Colorize(White+RedBackground, "Nope !\n"))
			buff += PositionPendu(hangmanIn.Tentatives)
		}
		response += buff
		fmt.Println(buff)
		spectate.AddStep(hangmanIn.Tentatives, response)
		//GAMEOVER
		if hangmanIn.Tentatives == 0 {
			GameOver(hangmanIn.Solution, spectate)
		}
		affichage = GenerateCryptedSolution(hangmanIn.Solution, hangmanIn.Lettres_trouvees)
		affichageAscii = GetAsciiCryptendSolution(affichage)
		if !strings.Contains(affichage, "_") {
			Win(hangmanIn.Solution, spectate)
		}
	}
}
func New_game(spectate Party) { //Titouan
	ClearConsole()
	if len(os.Args) > 2 {
		os.Exit(0)
	}
	if Input("Voulez vous enregistrer la partie ("+Colorize(Green, "y")+"/"+Colorize(Red, "n")+") \n "+Colorize(Purple, "->")+" ") == "y" {
		spectate.StoreSpectate(Input("Nom du fichier de sauvegarde\n" + Colorize(Purple, "->") + ""))
	}
	ClearConsole()
	proposition := Input("Nouvelle partie ??? (" + Colorize(Green, "y") + "/" + Colorize(Red, "n") + ") \n " + Colorize(Purple, "->") + " ")
	if proposition == "y" {
		ClearConsole()
		fmt.Print(Colorize(Yellow+Bold, "NOUVELLE PARTIE"))
		for range [3]int{} {
			Timeout(0.7)
			fmt.Print(Colorize(Yellow+Bold, "."))
		}
		Timeout(0.7)
		ClearConsole()
		CreateParty(HangManData{})
		return
	} else if proposition == "n" {
		LoadMainMenu()
	} else {
		New_game(spectate)
	}
}
