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
		util.ShowCommandMsg(handleCommand(input))

	}
}

func handleCommand(input string) virtualfilesystem.CommandType {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return virtualfilesystem.ErrCommandInvalid
	}
	command := strings.ToLower(parts[0])
	switch command {
	// common commands
	case "help":
		util.ShowHelp()
		return virtualfilesystem.Help
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
			return virtualfilesystem.ErrRegisterArgumentMissing
		} else if len(parts) > 2 {
			return virtualfilesystem.ErrRegisterArgumentInvalid
		}
		username := strings.ToLower(parts[1])
		if _, exist := users[username]; exist {
			msg := fmt.Sprintf("The %s has already existed. ", username)
			util.CallError(msg)
			return virtualfilesystem.ErrUserAlreadyExists
		} else if !virtualfilesystem.IsValidateUsername(username) {
			msg := fmt.Sprintf("The %s contain invalid chars. ", username)
			util.CallError(msg)
			return virtualfilesystem.ErrUserNameInvalid
		} else {
			users[username] = virtualfilesystem.NewUser(username)
			msg := fmt.Sprintf("Add %s successfully. ", username)
			util.CallSuccess(msg)
			return virtualfilesystem.Success
		}
	//folder commands
	case "create-folder":
		{
			//create-folder [username] [foldername] [description]?
			if len(parts) < 3 {
				return virtualfilesystem.ErrCreateFolderArgumentMissing
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
				return virtualfilesystem.ErrUserNotFound
			} else if _, exist := users[username].Folders[foldername]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderAlreadyExists
			} else if !virtualfilesystem.IsValidateFolderName(foldername) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNameInvalid
			} else {
				users[username].Folders[foldername] = virtualfilesystem.NewFolder(foldername, time, description)
				msg := fmt.Sprintf("Add %s successfully. ", foldername)
				util.CallSuccess(msg)
				return virtualfilesystem.Success
			}
		}
	case "delete-folder":
		{
			//delete-folder [username] [foldername]
			if len(parts) < 3 {
				return virtualfilesystem.ErrDeleteFolderArgumentMissing
			} else if len(parts) > 3 {
				return virtualfilesystem.ErrDeleteFolderArgumentInvalid
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return virtualfilesystem.ErrUserNotFound
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNotFound
			} else {
				delete(users[username].Folders, foldername)
				msg := fmt.Sprintf("Delete %s successfully. ", foldername)
				util.CallSuccess(msg)
				return virtualfilesystem.Success
			}
		}
	case "list-folders":
		{
			//list-folders [username] [--sort-name|--sort-created] [asc|desc]
			if len(parts) < 2 {
				return virtualfilesystem.ErrListFoldersArgumentMissing
			} else if len(parts) > 4 {
				return virtualfilesystem.ErrListFoldersArgumentInvalid
			}
			username := strings.ToLower(parts[1])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return virtualfilesystem.ErrUserNotFound
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
					return virtualfilesystem.ErrListFoldersArgumentInvalid
				}
			}
			if len(parts) > 3 {
				if parts[3] == "asc" {
					order = "asc"
				} else if parts[3] == "desc" {
					order = "desc"
				} else {
					return virtualfilesystem.ErrListFoldersArgumentInvalid
				}
			}
			virtualfilesystem.ListFolders(users[username], sort, order)
			return virtualfilesystem.Success
		}
	case "rename-folder":
		{
			//rename-folder [username] [foldername] [new-folder-name]
			if len(parts) < 4 {
				return virtualfilesystem.ErrRenameFolderArgumentMissing
			} else if len(parts) > 4 {
				return virtualfilesystem.ErrRenameFolderArgumentInvalid
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			newFolderName := strings.ToLower(parts[3])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return virtualfilesystem.ErrUserNotFound
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNotFound
			} else if _, exist := users[username].Folders[newFolderName]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", newFolderName)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderAlreadyExists
			} else if !virtualfilesystem.IsValidateFolderName(newFolderName) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", newFolderName)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNameInvalid
			} else {
				users[username].Folders[newFolderName] = users[username].Folders[foldername]
				delete(users[username].Folders, foldername)
				msg := fmt.Sprintf("Rename %s to %s successfully. ", foldername, newFolderName)
				util.CallSuccess(msg)
				return virtualfilesystem.Success
			}
		}
	//file commands
	case "create-file":
		{
			//create-file [username] [foldername] [filename] [description]?
			if len(parts) < 4 {
				return virtualfilesystem.ErrCreateFileArgumentMissing
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
				return virtualfilesystem.ErrUserNotFound
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNotFound
			} else if _, exist := users[username].Folders[foldername].Files[filename]; exist {
				msg := fmt.Sprintf("The %s has already existed. ", filename)
				util.CallError(msg)
				return virtualfilesystem.ErrFileAlreadyExists
			} else if !virtualfilesystem.IsValidateFileName(filename) {
				msg := fmt.Sprintf("The %s contain invalid chars. ", filename)
				util.CallError(msg)
				return virtualfilesystem.ErrFileNameInvalid
			} else {
				users[username].Folders[foldername].Files[filename] = virtualfilesystem.NewFile(filename, time, description)
				msg := fmt.Sprintf("Create %s in %s/%s successfully. ", filename, username, foldername)
				util.CallSuccess(msg)
				return virtualfilesystem.Success
			}
		}
	case "delete-file":
		{
			//delete-file [username] [foldername] [filename]
			if len(parts) < 4 {
				return virtualfilesystem.ErrDeleteFileArgumentMissing
			} else if len(parts) > 4 {
				return virtualfilesystem.ErrDeleteFileArgumentInvalid
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			filename := strings.ToLower(parts[3])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return virtualfilesystem.ErrUserNotFound
			} else if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNotFound
			}
			if _, exist := users[username].Folders[foldername].Files[filename]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", filename)
				util.CallError(msg)
				return virtualfilesystem.ErrFileNotFound
			} else {
				delete(users[username].Folders[foldername].Files, filename)
				msg := fmt.Sprintf("Delete %s in %s/%s successfully. ", filename, username, foldername)
				util.CallSuccess(msg)
				return virtualfilesystem.Success
			}
		}
	case "list-files":
		{
			//list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
			if len(parts) < 3 {
				return virtualfilesystem.ErrListFilesArgumentMissing
			} else if len(parts) > 5 {
				return virtualfilesystem.ErrListFilesArgumentInvalid
			}
			username := strings.ToLower(parts[1])
			foldername := strings.ToLower(parts[2])
			if _, exist := users[username]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", username)
				util.CallError(msg)
				return virtualfilesystem.ErrUserNotFound
			}
			if _, exist := users[username].Folders[foldername]; !exist {
				msg := fmt.Sprintf("The %s does not exist. ", foldername)
				util.CallError(msg)
				return virtualfilesystem.ErrFolderNotFound
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
					return virtualfilesystem.ErrListFilesArgumentInvalid
				}
			}
			if len(parts) > 4 {
				if parts[4] == "asc" {
					order = "asc"
				} else if parts[4] == "desc" {
					order = "desc"
				} else {
					return virtualfilesystem.ErrListFilesArgumentInvalid
				}
			}
			virtualfilesystem.ListFiles(users[username], foldername, sort, order)
			return virtualfilesystem.Success
		}
	//default (Invalid command)
	default:
		return virtualfilesystem.ErrCommandNotFound
	}
	return virtualfilesystem.Success
}
