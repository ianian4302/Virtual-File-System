package virtualfilesystem

//enum error type
type CommandType int

const (
	//success
	Success CommandType = iota
	//help
	Help
	//error type for folder
	ErrFolderNotFound
	ErrFolderAlreadyExists
	ErrFolderNameInvalid
	//error type for file
	ErrFileNotFound
	ErrFileAlreadyExists
	ErrFileNameInvalid
	//error type for user
	ErrUserNotFound
	ErrUserAlreadyExists
	ErrUserNameInvalid
	//error type for command
	ErrCommandNotFound
	ErrCommandInvalid
	//error type for argument
	// ErrArgumentInvalid
	// ErrArgumentMissing
	//argument
	ErrRegisterArgumentInvalid
	ErrRegisterArgumentMissing
	ErrCreateFolderArgumentInvalid
	ErrCreateFolderArgumentMissing
	ErrDeleteFolderArgumentInvalid
	ErrDeleteFolderArgumentMissing
	ErrListFoldersArgumentInvalid
	ErrListFoldersArgumentMissing
	ErrRenameFolderArgumentInvalid
	ErrRenameFolderArgumentMissing
	ErrCreateFileArgumentInvalid
	ErrCreateFileArgumentMissing
	ErrDeleteFileArgumentInvalid
	ErrDeleteFileArgumentMissing
	ErrListFilesArgumentInvalid
	ErrListFilesArgumentMissing
)
