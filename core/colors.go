package core

import "fmt"

var (
	Reset = "\033[0m"

	Bold      = "\033[1m"
	Underline = "\033[4m"

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	BlackBackground  = "\033[40m"
	RedBackground    = "\033[41m"
	GreenBackground  = "\033[42m"
	YellowBackground = "\033[43m"
	BlueBackground   = "\033[44m"
	PurpleBackground = "\033[45m"
	CyanBackground   = "\033[46m"
	GrayBackground   = "\033[47m"
	WhiteBackground  = "\033[107m"
	// Couleurs du texte
	Orange      = "\033[38;5;208m"
	Pink        = "\033[38;5;200m"
	LightBlue   = "\033[38;5;39m"
	LightGreen  = "\033[38;5;46m"
	LightRed    = "\033[38;5;196m"
	LightYellow = "\033[38;5;226m"

	// Couleurs de fond
	OrangeBackground      = "\033[48;5;208m"
	PinkBackground        = "\033[48;5;200m"
	LightBlueBackground   = "\033[48;5;39m"
	LightGreenBackground  = "\033[48;5;46m"
	LightRedBackground    = "\033[48;5;196m"
	LightYellowBackground = "\033[48;5;226m"

	// Styles de texte
	Italic        = "\033[3m"
	Strikethrough = "\033[9m"

	// Combinaison de styles
	BoldItalic    = "\033[1;3m"
	UnderlineBold = "\033[4;1m"
)

func Colorize(color string, s any) string { //Titouan
	switch s := s.(type) {
	case string:
		return color + s + Reset
	default:
		return color + fmt.Sprint(s) + Reset
	}
}
