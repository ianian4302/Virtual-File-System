package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ianian4302/Virtual-File-System/util"
	virtualfilesystem "github.com/ianian4302/Virtual-File-System/virtual-file-system"
)

var users = make(map[string]*virtualfilesystem.User)
var folders = make(map[string]*virtualfilesystem.Folder)
var files = make(map[string]*virtualfilesystem.File)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You can use cammand 'help' to get help")
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		handleCommand(input)
	}
}

func handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}
	command := strings.ToLower(parts[0])
	switch command {
	// common commands
	case "help":
		fmt.Println("Help, Not supported. ")
	case "exit":
		fmt.Println("Exit the program. ")
		os.Exit(0)
	case "q":
		fmt.Println("Exit the program. ")
		os.Exit(0)
	// user commands
	case "register":
		//register [username]
		if len(parts) < 2 {
			util.CallError("Please provide username. ")
			return
		}
		username := strings.ToLower(parts[1])
		if _, exist := users[username]; exist {
			msg := fmt.Sprintf("The %s has already existed. ", username)
			util.CallError(msg)
		} else if !virtualfilesystem.IsValidateUsername(username) {
			msg := fmt.Sprintf("The %s contain invalid chars. ", username)
			util.CallError(msg)
		} else {
			users[username] = virtualfilesystem.NewUser(username)
			msg := fmt.Sprintf("Add %s successfully. ", username)
			util.CallSuccess(msg)
		}
	//folder commands
	case "create-folder":
		{
			//create-folder [username] [foldername] [description]?
			if len(parts) < 3 {
				util.CallError("Please provide username and foldername. ")
				return
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			time := util.GetCurrentTime()
			//default description is empty
			description := ""
			if len(parts) > 3 {
				description = strings.Join(parts[3:], " ")
			}
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return
			} else if _, exist := users[username].Folders[foldername]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", foldername)
				util.CallError(msg)
				return
			} else if !virtualfilesystem.IsValidateFolderName(foldername) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", foldername)
				util.CallError(msg)
			} else {
				users[username].Folders[foldername] = virtualfilesystem.NewFolder(foldername, time, description)
				msg := fmt.Sprintf("Add %s successfully. ", foldername)
				util.CallSuccess(msg)
			}
		}
	case "delete-folder":
		{
			//delete-folder [username] [foldername]
			if len(parts) < 3 {
				util.CallError("Please provide username and foldername. ")
				return
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return
			} else {
				delete(users[username].Folders, foldername)
				msg := fmt.Sprintf("Delete %s successfully. ", foldername)
				util.CallSuccess(msg)
			}
		}
	case "list-folders":
		{
			//list-folders [username] [--sort-name|--sort-created] [asc|desc]
			if len(parts) < 2 {
				util.CallError("Please provide username. ")
				return
			}
			username := strings.ToLower(parts[1])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return
			}
			// [--sort-name|--sort-created]
			sort := "name"
			// [asc|desc]
			order := "asc"
			if len(parts) > 2 {
				if parts[2] == "--sort-name" {
					sort = "name"
				} else if parts[2] == "--sort-created" {
					sort = "created"
				} else {
					util.CallError("Invalid sort option. ")
					return
				}
			}
			if len(parts) > 3 {
				if parts[3] == "asc" {
					order = "asc"
				} else if parts[3] == "desc" {
					order = "desc"
				} else {
					util.CallError("Invalid order option. ")
					return
				}
			}
			virtualfilesystem.ListFolders(users[username], sort, order)
		}
	case "rename-folder":
		{

		}
	//file commands
	case "create-file":
		{
		}
	case "delete-file":
		{
		}
	case "list-files":
		{
		}
	//default (Invalid command)
	default:
		util.CallError("Invalid command. ")
	}
}
