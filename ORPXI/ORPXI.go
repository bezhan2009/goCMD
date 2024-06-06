package ORPXI

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"goCmd/cmdPress"
	"goCmd/commands/commandsWithSignaiture/Create"
	"goCmd/commands/commandsWithSignaiture/Edit"
	"goCmd/commands/commandsWithSignaiture/Read"
	"goCmd/commands/commandsWithSignaiture/Remove"
	"goCmd/commands/commandsWithSignaiture/Rename"
	"goCmd/commands/commandsWithSignaiture/Write"
	"goCmd/commands/commandsWithSignaiture/newShablon"
	"goCmd/commands/commandsWithoutSignature/CD"
	"goCmd/commands/commandsWithoutSignature/Clean"
	"goCmd/debug"
	"goCmd/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CMD() {
	utils.SystemInformation()
	isWorking := true
	counterORPXI := 1
	reader := bufio.NewReader(os.Stdin)

	prompt := ""

	for isWorking {
		if utils.IsHidden() {
			fmt.Println("You are BLOCKED!!!")
			return
		}

		cyan := color.New(color.FgCyan).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		dir, _ := os.Getwd()
		dirC := cmdPress.CmdDir(dir)
		user := cmdPress.CmdUser(dir)

		if prompt != "" {
			fmt.Printf("\n%s", prompt)
		} else {
			fmt.Printf("\n┌─(%s)-[%s%s]\n", cyan("ORPXI "+user), cyan("~"), cyan(dirC))
			fmt.Printf("└─$ %s", green(""))
		}

		commandLine, _ := reader.ReadString('\n')
		commandLine = strings.TrimSpace(commandLine)
		commandParts := strings.Fields(commandLine)

		if len(commandParts) == 0 {
			continue
		}

		command := commandParts[0]
		commandArgs := commandParts[1:]
		commandLower := strings.ToLower(command)

		if commandLower == "prompt" {
			if len(commandArgs) < 1 {
				fmt.Println("prompt <name_prompt>")
				fmt.Println("to delete prompt enter:")
				fmt.Println("prompt delete")
				continue
			}

			namePrompt := commandArgs[0]

			if namePrompt != "delete" {
				namePrompt = strings.TrimSpace(namePrompt)
				prompt = namePrompt
				fmt.Printf("Prompt set to: %s\n", prompt)
			} else {
				prompt, _ = os.Getwd()
				fmt.Printf("Prompt set to: %s\n", prompt)
				prompt = ""
			}

			continue
		}

		if commandLower == "help" {
			fmt.Println("Для получения сведений об командах наберите ORPXIHELP")
			fmt.Println("CREATE             создает новый файл")
			fmt.Println("CLEAN              очистка экрана")
			fmt.Println("CD                 смена текущего каталога")
			fmt.Println("NEWSHABLON         создает новый шаблон комманд для выполнения")
			fmt.Println("REMOVE             удаляет файл")
			fmt.Println("READ               выводит на экран содержимое файла")
			fmt.Println("PROMPT             Изменяет ORPXI.")
			fmt.Println("PASSWORD           пароль для ORPXI.")
			fmt.Println("ORPXI              запускает ещё одну ORPXI")
			fmt.Println("SHABLON            выполняет определенный шаблон комманд")
			fmt.Println("SYSTEMGOCMD        вывод информации о ORPXI")
			fmt.Println("SYSTEMINFO         вывод информации о системе")
			fmt.Println("TREE               Графически отображает структуру каталогов диска или пути.")
			fmt.Println("WRITE              записывает данные в файл")
			fmt.Println("EDIT               редактирует файл")
			fmt.Println("EXIT               Выход")
			fmt.Println("LS                 выводит содержимое каталога")

			errDebug := debug.Commands(command, true)
			if errDebug != nil {
				fmt.Println(errDebug)
			}
			continue
		}

		if commandLower == "help" {
			continue
		}

		commands := []string{"newshablon", "shablon", "password", "promptSet", "systemgocmd", "rename", "remove", "read", "write", "create", "exit", "orpxi", "clean", "cd", "edit", "ls"}

		isValid := utils.ValidCommand(commandLower, commands)

		if !isValid {
			fullCommand := append([]string{command}, commandArgs...)
			err := utils.ExternalCommand(fullCommand)
			if commandLower == "help" {
				continue
			}
			if err != nil {
				fullPath := filepath.Join(dir, command)
				fullCommand[0] = fullPath
				err = utils.ExternalCommand(fullCommand)
				if err != nil {
					fmt.Printf("Ошибка при запуске команды '%s': %v\n", commandLine, err)
				}
			}
			continue
		}

		switch commandLower {
		case "newshablon":
			newShablon.Make()
		case "password":
			Password()
		case "systemgocmd":
			utils.SystemInformation()
		case "orpxi":
			counterORPXI += 1
			CMD()
		case "exit":
			isWorking = false
		case "create":
			name, err := Create.File(commandArgs)
			if err != nil {
				fmt.Println(err)
				debug.Commands(command, false)
			} else if name != "" {
				fmt.Printf("Файл %s успешно создан!!!\n", name)
				fmt.Printf("Директория нового файла: %s\n", filepath.Join(dir, name))
				debug.Commands(command, true)
			}
		case "write":
			Write.File(commandLower, commandArgs)
		case "read":
			Read.File(commandLower, commandArgs)
		case "remove":
			name, err := Remove.File(commandArgs)
			if err != nil {
				debug.Commands(command, false)
				fmt.Println(err)
			} else {
				debug.Commands(command, true)
				fmt.Printf("Файл %s успешно удален!!!\n", name)
			}
		case "rename":
			errRename := Rename.Rename(commandArgs)
			if errRename != nil {
				debug.Commands(command, false)
				fmt.Println(errRename)
			} else {
				debug.Commands(command, true)
			}
		case "clean":
			Clean.Screen()
		case "cd":
			if len(commandArgs) == 0 {
				dir, _ := os.Getwd()
				fmt.Println(dir)
			} else {
				err := CD.ChangeDirectory(commandArgs[0])
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
		case "edit":
			if len(commandArgs) < 1 {
				fmt.Println("Использование: edit <файл>")
				continue
			}
			filename := commandArgs[0]
			err := Edit.File(filename)
			if err != nil {
				fmt.Println(err)
			}
		case "ls":
			printLS()
		default:
			validCommand := utils.ValidCommand(commandLower, commands)
			if !validCommand {
				fmt.Printf("'%s' не является внутренней или внешней командой,\nисполняемой программой или пакетным файлом.\n", commandLine)
			}
		}
	}
}

func printLS() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("%s/\t", file.Name())
		} else {
			fmt.Print(file.Name(), "\t")
		}
	}
}
