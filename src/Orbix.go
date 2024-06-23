package src

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"goCmd/cmdPress"
	"goCmd/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Absdir, _ = filepath.Abs("")
var DirUser, _ = filepath.Abs("")

func Orbix(commandInput string) {
	utils.SystemInformation()

	isWorking := true
	isPermission := true

	var promptText string

	isEmpty, err := isPasswordDirectoryEmpty()
	if err != nil {
		animatedPrint("Ошибка при проверке директории с паролями:" + err.Error() + "\n")
		return
	}
	username := ""
	if !isEmpty && commandInput == "" {
		dir, _ := os.Getwd()
		user := cmdPress.CmdUser(dir)
		nameuser, isSuccess := CheckUser(user)
		if !isSuccess {
			return
		}
		username = nameuser
	}

	for isWorking {
		dir, _ := os.Getwd()

		runFilePath := Absdir
		runFilePath += "\\isRun.txt"

		file, err := os.Open(runFilePath)
		if err != nil {
			red := color.New(color.FgRed).SprintFunc()
			fmt.Println(red("Запустите программу через run_main.bat либо если у вас Unix(Linux, MacOS) то запустите через main.sh"))
			break
		}
		defer file.Close()

		currentBranchGit, errGitBranch := GetCurrentGitBranch()
		if errGitBranch != nil {
			fmt.Println("Ошибка при получении текущей ветки Git")
		}

		// Ensure absolute path for activeUser.txt
		activeUserFilePath := DirUser
		activeUserFilePath += "\\activeUser.txt"

		fmt.Println("Путь к activeUser.txt:", activeUserFilePath)

		os.Create(activeUserFilePath)
		os.WriteFile(activeUserFilePath, []byte(username), 0644)

		cyan := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		magenta := color.New(color.FgMagenta).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		dirC := cmdPress.CmdDir(dir)
		user := cmdPress.CmdUser(dir)

		if username != "" {
			user = username
		}

		currentTime := time.Now().Format("15:04")
		location := os.Getenv("CITY")
		if location == "" {
			location = "Unknown City"
		}

		if promptText != "" {
			animatedPrint("\n" + promptText)
		} else {
			fmt.Print(fmt.Sprintf("\n%s%s%s%s%s%s%s%s %s%s%s%s%s%s%s%s%s\n",
				yellow("┌"), yellow("─"), yellow("("), cyan("Orbix@"+user), yellow(")"), yellow("─"), yellow("["),
				yellow(location), magenta(currentTime), yellow("]"), yellow("─"), yellow("["),
				cyan("~"), cyan(dirC), yellow("]"), yellow(" git:"), green("["+currentBranchGit+"]")))
			fmt.Print(fmt.Sprintf("%s%s%s %s",
				yellow("└"), yellow("─"), green("$"), green(commandInput)))
		}

		var commandLine string
		var commandParts []string
		var commandArgs []string
		var commandLower string
		var command string

		if commandInput != "" {
			isWorking = false
			isPermission = false
			commandLine = strings.TrimSpace(commandInput)
			commandParts = utils.SplitCommandLine(commandLine)
			if len(commandParts) == 0 {
				continue
			}

			command = commandParts[0]
			commandArgs = commandParts[1:]
			commandLower = strings.ToLower(command)
		} else {
			commandLine = prompt.Input("", autoComplete)
			commandLine = strings.TrimSpace(commandLine)
			commandParts = utils.SplitCommandLine(commandLine)

			if len(commandParts) == 0 {
				continue
			}

			command = commandParts[0]
			commandArgs = commandParts[1:]
			commandLower = strings.ToLower(command)

			commandHistory = append(commandHistory, commandLine)
		}

		animatedPrint("\n")

		if commandLower == "prompt" {
			handlePromptCommand(commandArgs, &promptText)
			continue
		}

		if commandLower == "help" {
			displayHelp(commandArgs, user, dir)
			continue
		}

		isValid := utils.ValidCommand(commandLower, commands)

		if !isValid {
			fullCommand := append([]string{command}, commandArgs...)
			err := utils.ExternalCommand(fullCommand)
			if err != nil {
				fullPath := filepath.Join(dir, command)
				fullCommand[0] = fullPath
				err = utils.ExternalCommand(fullCommand)
				if err != nil {
					suggestedCommand := suggestCommand(commandLower)
					fmt.Printf("Error executing command '%s': %v\n", commandLine, err)
					if suggestedCommand != "" {
						fmt.Printf("Did you mean: %s?\n", suggestedCommand)
					}
				}
			}
			continue
		}

		ExecuteCommand(commandLower, command, commandLine, dir, commands, commandArgs, &isWorking, isPermission)
	}
}
