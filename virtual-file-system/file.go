package virtualfilesystem

type File struct {
	Name        string
	Description string
}

//create-file [username] [foldername] [filename] [description]?

func NewFile(name, description string) *File {
	return &File{
		Name:        name,
		Description: description,
	}
}
