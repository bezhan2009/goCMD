package Create

import (
	"fmt"
	"goCmd/commands/commandsWithSignaiture/Create/utils"
)

func File(commandArgs []string) (string, error) {
	if len(commandArgs) < 1 {
		fmt.Println("Использования: create <файл>")
		return "", nil
	}

	var name string

	name = commandArgs[0]

	if name == "debug.txt" {
		fmt.Println("PermissionDenied: You cannot write, delete or create a debug.txt file")
		return name, nil
	}

	errExisting := utils.IsExists(name)

	if errExisting == nil {
		return name, errExisting
	}

	name, err := utils.CreateFile(name)

	if err != nil {
		return name, err
	} else {
		return name, err
	}
}
