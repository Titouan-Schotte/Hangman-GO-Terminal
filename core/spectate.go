package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Step struct {
	TentativesIn int
	ResponseIn   string
}

type Party struct {
	Solution string
	Steps    []Step
}

func (p *Party) InitSpectate(solution string) {
	p.Solution = solution
}
func (p *Party) AddStep(tentativesIn int, responseIn string) {
	p.Steps = append(p.Steps, Step{
		TentativesIn: tentativesIn,
		ResponseIn:   responseIn,
	})
}

func (p *Party) StoreSpectate(fileName string) { //Titouan
	file, _ := json.MarshalIndent(p, "", " ")
	ioutil.WriteFile("spectates/"+fileName+".json", file, 0644)
}

func Spectate(saveName string) { //Titouan
	var spectateIn Party
	fileData, err := ioutil.ReadFile("spectates/" + saveName + ".json")
	if err != nil {
		fmt.Print(Colorize(Red+Bold, "Nom de la sauvegarde invalide !"))
		return
	}
	json.Unmarshal(fileData, &spectateIn)
	for _, steps := range spectateIn.Steps {
		ClearConsole()
		fmt.Println(Colorize(LightRed+Bold, "###### SPECTATE MODE ######\n"))
		fmt.Println(steps.ResponseIn)
		Input(Colorize(Yellow, "(Touche Entrée pour continuer) \n=>"))
	}
	ClearConsole()
	fmt.Print(Colorize(Red+Bold, "FIN DE LA VISUALISATION"))
	for range [3]int{} {
		Timeout(0.7)
		fmt.Print(Colorize(Red+Bold, "."))
	}
	LoadMainMenu()
}
