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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You can use cammand 'help' to get help")
	fmt.Println("You can use cammand 'exit' or 'q' to exit the program")
	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()
		handleCommand(input)
	}
}

func handleCommand(input string) bool {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return false
	}
	command := strings.ToLower(parts[0])
	switch command {
	// common commands
	case "help":
		util.ShowHelp()
		return true
	case "exit":
		fmt.Println("Exit the program. ")
		os.Exit(0)
	case "q":
		fmt.Println("Exit the program. ")
		os.Exit(0)
	case "quit":
		fmt.Println("Exit the program. ")
		os.Exit(0)
	// user commands
	case "register":
		//register [username]
		if len(parts) < 2 {
			util.CallError("Please provide username. ")
			return false
		} else if len(parts) > 2 {
			util.CallError("Please provide only one username. ")
			return false
		}
		username := strings.ToLower(parts[1])
		if _, exist := users[username]; exist {
			msg := fmt.Sprintf("The %s has already existed. ", username)
			util.CallError(msg)
			return false
		} else if !virtualfilesystem.IsValidateUsername(username) {
			msg := fmt.Sprintf("The %s contain invalid chars. ", username)
			util.CallError(msg)
			return false
		} else {
			users[username] = virtualfilesystem.NewUser(username)
			msg := fmt.Sprintf("Add %s successfully. ", username)
			util.CallSuccess(msg)
			return true
		}
	//folder commands
	case "create-folder":
		{
			//create-folder [username] [foldername] [description]?
			if len(parts) < 3 {
				util.CallError("Please provide username and foldername. ")
				return false
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
				return false
			} else if _, exist := users[username].Folders[foldername]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", foldername)
				util.CallError(msg)
				return false
			} else if !virtualfilesystem.IsValidateFolderName(foldername) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", foldername)
				util.CallError(msg)
				return false
			} else {
				users[username].Folders[foldername] = virtualfilesystem.NewFolder(foldername, time, description)
				msg := fmt.Sprintf("Add %s successfully. ", foldername)
				util.CallSuccess(msg)
				return true
			}
		}
	case "delete-folder":
		{
			//delete-folder [username] [foldername]
			if len(parts) < 3 {
				util.CallError("Please provide username and foldername. ")
				return false
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return false
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return false
			} else {
				delete(users[username].Folders, foldername)
				msg := fmt.Sprintf("Delete %s successfully. ", foldername)
				util.CallSuccess(msg)
				return true
			}
		}
	case "list-folders":
		{
			//list-folders [username] [--sort-name|--sort-created] [asc|desc]
			if len(parts) < 2 {
				util.CallError("Please provide username. ")
				return false
			}
			username := strings.ToLower(parts[1])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return false
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
					return false
				}
			}
			if len(parts) > 3 {
				if parts[3] == "asc" {
					order = "asc"
				} else if parts[3] == "desc" {
					order = "desc"
				} else {
					util.CallError("Invalid order option. ")
					return false
				}
			}
			virtualfilesystem.ListFolders(users[username], sort, order)
			return true
		}
	case "rename-folder":
		{
			//rename-folder [username] [foldername] [new-folder-name]
			if len(parts) < 4 {
				util.CallError("Please provide username, foldername and new-folder-name. ")
				return false
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			newFolderName := strings.ToLower(parts[3])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return false
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return false
			} else if _, exist := users[username].Folders[newFolderName]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", newFolderName)
				util.CallError(msg)
				return false
			} else if !virtualfilesystem.IsValidateFolderName(newFolderName) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", newFolderName)
				util.CallError(msg)
				return false
			} else {
				users[username].Folders[newFolderName] = users[username].Folders[foldername]
				delete(users[username].Folders, foldername)
				msg := fmt.Sprintf("Rename %s to %s successfully. ", foldername, newFolderName)
				util.CallSuccess(msg)
				return true
			}
		}
	//file commands
	case "create-file":
		{
			//create-file [username] [foldername] [filename] [description]?
			if len(parts) < 4 {
				util.CallError("Please provide username, foldername and filename. ")
				return false
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			filename := strings.ToLower(parts[3])
			time := util.GetCurrentTime()
			//default description is empty
			description := ""
			if len(parts) > 4 {
				description = strings.Join(parts[4:], " ")
			}
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return false
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return false
			} else if _, exist := users[username].Folders[foldername].Files[filename]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", filename)
				util.CallError(msg)
				return false
			} else if !virtualfilesystem.IsValidateFileName(filename) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", filename)
				util.CallError(msg)
				return false
			} else {
				users[username].Folders[foldername].Files[filename] = virtualfilesystem.NewFile(filename, time, description)
				msg := fmt.Sprintf("Create %s in %s/%s successfully. ", filename, username, foldername)
				util.CallSuccess(msg)
				return true
			}
		}
	case "delete-file":
		{
			//delete-file [username] [foldername] [filename]
			if len(parts) < 4 {
				util.CallError("Please provide username, foldername and filename. ")
				return false
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			filename := strings.ToLower(parts[3])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return false
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return false
			}
			if _, exist := users[username].Folders[foldername].Files[filename]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", filename)
				util.CallError(msg)
				return false
			} else {
				delete(users[username].Folders[foldername].Files, filename)
				msg := fmt.Sprintf("Delete %s in %s/%s successfully. ", filename, username, foldername)
				util.CallSuccess(msg)
				return true
			}
		}
	case "list-files":
		{
			//list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
			if len(parts) < 3 {
				util.CallError("Please provide username and foldername. ")
				return false
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return false
			}
			if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return false
			}
			// [--sort-name|--sort-created]
			sort := "name"
			// [asc|desc]
			order := "asc"
			if len(parts) > 3 {
				if parts[3] == "--sort-name" {
					sort = "name"
				} else if parts[3] == "--sort-created" {
					sort = "created"
				} else {
					util.CallError("Invalid sort option. ")
					return false
				}
			}
			if len(parts) > 4 {
				if parts[4] == "asc" {
					order = "asc"
				} else if parts[4] == "desc" {
					order = "desc"
				} else {
					util.CallError("Invalid order option. ")
					return false
				}
			}
			virtualfilesystem.ListFiles(users[username], foldername, sort, order)
			return true
		}
	//default (Invalid command)
	default:
		util.CallError("Invalid command. ")
		return false
	}
	return true
}
