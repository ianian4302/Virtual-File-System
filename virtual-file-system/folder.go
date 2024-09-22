package virtualfilesystem

import (
	"fmt"
	"sort"
	"time"
)

type Folder struct {
	Name        string
	Files       map[string]*File
	Time        time.Time
	Description string
}

func NewFolder(name string, time time.Time, description string) *Folder {
	return &Folder{
		Name:        name,
		Files:       make(map[string]*File),
		Time:        time,
		Description: description,
	}
}

func (f *Folder) AddFolder(folder *Folder) {
	f.Files[folder.Name] = NewFile(folder.Name, folder.Description)
}

func (f *Folder) DeleteFolder(foldername string) {
	delete(f.Files, foldername)
}

func ListFolders(user *User, sortMethod string, order string) {
	if sortMethod == "name" {
		//make a mapwith key = usename and value = other info
		list := make(map[string]Folder)
		for _, folder := range user.Folders {
			list[folder.Name] = *folder
		}
		//sort the map by key
		var keys []string
		for k := range list {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		if order == "desc" {
			for i := len(keys)/2 - 1; i >= 0; i-- {
				opp := len(keys) - 1 - i
				keys[i], keys[opp] = keys[opp], keys[i]
			}
		}
		//List all the folders within the [username] scope in following formats: [foldername]
		//[description] [created at] [username]
		for _, key := range keys {
			folder := list[key]
			//humanreadable date/time format.
			time := folder.Time.Format("2006-01-02 15:04:05")
			msg := folder.Name + " " + folder.Description + " " + time + " " + user.Username
			fmt.Println(msg)
		}
	} else if sortMethod == "created" {

	}
}

func IsValidateFolderName(foldername string) bool {
	//name can not start with number
	if foldername[0] >= '0' && foldername[0] <= '9' {
		return false
	}
	//name len should be greater than 1 chars and less than 20 chars
	if len(foldername) < 1 || len(foldername) > 20 {
		return false
	}
	//name only contains a-z, A-Z, 0-9
	for _, c := range foldername {
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return true
}
