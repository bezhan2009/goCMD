package utils

import (
	"fmt"
	"github.com/fatih/color"
	"goCmd/editCMD"
)

func SystemInformation() {
	editCMD.StartEditing()
	magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()
	fmt.Printf("%s\n", magenta("Orbix [Версия 0.74]"))
	fmt.Printf("%s\n", magenta("(c) Orbix Software, 2024. Все права защищены."))
}
