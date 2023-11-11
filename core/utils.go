package core

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ClearConsole() { //Titouan
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Timeout(seconds float64) { //Titouan
	duration := time.Duration(seconds * float64(time.Second))
	time.Sleep(duration)
}
func Input(prompt string) string { //Titouan
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return strings.ToLower(input)
}
