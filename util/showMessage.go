package util

import "fmt"

const (
	Reset = "\033[0m"
	Red   = "\033[31;1m"
	Green = "\033[32;1m"
)

func CallError(str string) {
	fmt.Println(Red, "Error: ", str, Reset)
}

func CallSuccess(str string) {
	fmt.Println(Green, str, Reset)
}

func ShowHelp() {
	helpMsg := (`Commands:
	- register [username]
	- create-folder [username] [foldername] [description]?
	- delete-folder [username] [foldername]
	- list-folders [username] [--sort-name|--sort-created] [asc|desc]
	- rename-folder [username] [foldername] [new-folder-name]
	- create-file [username] [foldername] [filename] [description]?
	- delete-file [username] [foldername] [filename]
	- list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
	- exit | quit | q
	- help
	`)
	fmt.Println(helpMsg)
}
