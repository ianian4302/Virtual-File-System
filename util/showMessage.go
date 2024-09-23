package util

import (
	"fmt"
	"os"

	virtualfilesystem "github.com/ianian4302/Virtual-File-System/virtual-file-system"
)

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

func ShowCommandMsg(command virtualfilesystem.CommandType) {
	if command == virtualfilesystem.Success {
		CallSuccess("Success")
	} else if command == virtualfilesystem.Help {
		ShowHelp()
	} else if command == virtualfilesystem.ErrFolderNotFound {
		CallError("Folder not found")
	} else if command == virtualfilesystem.ErrFolderAlreadyExists {
		CallError("Folder already exists")
	} else if command == virtualfilesystem.ErrFolderNameInvalid {
		CallError("Folder name is invalid")
	} else if command == virtualfilesystem.ErrFileNotFound {
		CallError("File not found")
	} else if command == virtualfilesystem.ErrFileAlreadyExists {
		CallError("File already exists")
	} else if command == virtualfilesystem.ErrFileNameInvalid {
		CallError("File name is invalid")
	} else if command == virtualfilesystem.ErrUserNotFound {
		CallError("User not found")
	} else if command == virtualfilesystem.ErrUserAlreadyExists {
		CallError("User already exists")
	} else if command == virtualfilesystem.ErrUserNameInvalid {
		CallError("User name is invalid")
	} else if command == virtualfilesystem.ErrCommandNotFound {
		CallError("Command not found")
	} else if command == virtualfilesystem.ErrCommandInvalid {
		CallError("Command is invalid")
	} else if command >= virtualfilesystem.ErrRegisterArgumentInvalid && command < virtualfilesystem.ErrRegisterArgumentInvalid {
		if command == virtualfilesystem.ErrRegisterArgumentInvalid {
			CallError("Register argument is invalid")
			fmt.Fprintf(os.Stderr, "register [username]\n")
		} else if command == virtualfilesystem.ErrRegisterArgumentMissing {
			CallError("Register argument is missing")
			fmt.Fprintf(os.Stderr, "register [username]\n")
		} else if command == virtualfilesystem.ErrCreateFolderArgumentInvalid {
			CallError("Create folder argument is invalid")
			fmt.Fprintf(os.Stderr, "create-folder [username] [foldername] [description]?\n")
		} else if command == virtualfilesystem.ErrCreateFolderArgumentMissing {
			CallError("Create folder argument is missing")
			fmt.Fprintf(os.Stderr, "create-folder [username] [foldername] [description]?\n")
		} else if command == virtualfilesystem.ErrDeleteFolderArgumentInvalid {
			CallError("Delete folder argument is invalid")
			fmt.Fprintf(os.Stderr, "delete-folder [username] [foldername]\n")
		} else if command == virtualfilesystem.ErrDeleteFolderArgumentMissing {
			CallError("Delete folder argument is missing")
			fmt.Fprintf(os.Stderr, "delete-folder [username] [foldername]\n")
		} else if command == virtualfilesystem.ErrListFoldersArgumentInvalid {
			CallError("List folders argument is invalid")
			fmt.Fprintf(os.Stderr, "list-folders [username] [--sort-name|--sort-created] [asc|desc]\n")
		} else if command == virtualfilesystem.ErrListFoldersArgumentMissing {
			CallError("List folders argument is missing")
			fmt.Fprintf(os.Stderr, "list-folders [username] [--sort-name|--sort-created] [asc|desc]\n")
		} else if command == virtualfilesystem.ErrRenameFolderArgumentInvalid {
			CallError("Rename folder argument is invalid")
			fmt.Fprintf(os.Stderr, "rename-folder [username] [foldername] [new-folder-name]\n")
		} else if command == virtualfilesystem.ErrRenameFolderArgumentMissing {
			CallError("Rename folder argument is missing")
			fmt.Fprintf(os.Stderr, "rename-folder [username] [foldername] [new-folder-name]\n")
		} else if command == virtualfilesystem.ErrCreateFileArgumentInvalid {
			CallError("Create file argument is invalid")
			fmt.Fprintf(os.Stderr, "create-file [username] [foldername] [filename] [description]?\n")
		} else if command == virtualfilesystem.ErrCreateFileArgumentMissing {
			CallError("Create file argument is missing")
			fmt.Fprintf(os.Stderr, "create-file [username] [foldername] [filename] [description]?\n")
		} else if command == virtualfilesystem.ErrDeleteFileArgumentInvalid {
			CallError("Delete file argument is invalid")
			fmt.Fprintf(os.Stderr, "delete-file [username] [foldername] [filename]\n")
		} else if command == virtualfilesystem.ErrDeleteFileArgumentMissing {
			CallError("Delete file argument is missing")
			fmt.Fprintf(os.Stderr, "delete-file [username] [foldername] [filename]\n")
		} else if command == virtualfilesystem.ErrListFilesArgumentInvalid {
			CallError("List files argument is invalid")
			fmt.Fprintf(os.Stderr, "list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
		} else if command == virtualfilesystem.ErrListFilesArgumentMissing {
			CallError("List files argument is missing")
			fmt.Fprintf(os.Stderr, "list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
		}
	}
}
