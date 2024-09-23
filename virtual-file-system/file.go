package virtualfilesystem

import (
	"fmt"
	"sort"
	"time"
)

type File struct {
	Name        string
	Time        time.Time
	Description string
}

//create-file [username] [foldername] [filename] [description]?

func NewFile(name string, time time.Time, description string) *File {
	return &File{
		Name:        name,
		Time:        time,
		Description: description,
	}
}

func ListFiles(user *User, foldername string, sortMethod string, order string) {
	if sortMethod == "name" {
		//make a mapwith key = usename and value = other info
		list := make(map[string]File)
		for _, file := range user.Folders[foldername].Files {
			list[file.Name] = *file
		}
		//sort the map by key
		var keys []string
		for k := range list {
			keys = append(keys, k)
		}
		if order == "asc" {
			sort.Strings(keys)
		} else {
			for i := len(keys)/2 - 1; i >= 0; i-- {
				opp := len(keys) - 1 - i
				keys[i], keys[opp] = keys[opp], keys[i]
			}
		}
		//title
		fmt.Println("| filename | description | created_at | foldername |")
		//List all the files within the [username] scope in following formats
		//format: [filename] [description] [created at] [foldername]
		for _, key := range keys {
			file := list[key]
			//humanreadable date/time format.
			time := file.Time.Format("2006-01-02 15:04:05")
			msg := "| " + file.Name + " | " + file.Description + " | " + time + " | " + foldername + " |"
			fmt.Println(msg)
		}
	} else if sortMethod == "created" {
		//make a mapwith key = usename and value = other info
		list := make(map[time.Time]File)
		for _, file := range user.Folders[foldername].Files {
			list[file.Time] = *file
		}
		//sort the map by key
		var keys []time.Time
		for k := range list {
			keys = append(keys, k)
		}
		if order == "asc" {
			sort.Slice(keys, func(i, j int) bool {
				return keys[i].Before(keys[j])
			})
		} else {
			sort.Slice(keys, func(i, j int) bool {
				return keys[i].After(keys[j])
			})
		}
		//title
		fmt.Println("| filename | description | created_at | foldername |")
		//List all the files within the [username] scope in following formats
		//format: [filename] [description] [created at] [foldername]
		for _, key := range keys {
			file := list[key]
			//humanreadable date/time format.
			time := file.Time.Format("2006-01-02 15:04:05")
			msg := "| " + file.Name + " | " + file.Description + " | " + time + " | " + foldername + " |"
			fmt.Println(msg)
		}
	}
}

func IsValidateFileName(filename string) bool {
	//name can not start with number
	if filename[0] >= '0' && filename[0] <= '9' {
		return false
	}
	//name len should be greater than 1 chars and less than 20 chars
	if len(filename) < 1 || len(filename) > 20 {
		return false
	}
	//name only contains a-z, A-Z, 0-9
	for _, c := range filename {
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
