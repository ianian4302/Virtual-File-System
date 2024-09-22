package virtualfilesystem

type User struct {
	Username string
	Folders  map[string]*Folder
}

func NewUser(username string) *User {
	return &User{
		Username: username,
		Folders:  make(map[string]*Folder),
	}
}

func IsValidateUsername(username string) bool {
	//name can not start with number
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}
	//name len should be greater than 3 chars and less than 20 chars
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	//name only contains a-z, A-Z, 0-9
	for _, c := range username {
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
