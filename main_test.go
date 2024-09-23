package main

import (
	"reflect"
	"testing"

	virtualfilesystem "github.com/ianian4302/Virtual-File-System/virtual-file-system"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_handleCommand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want virtualfilesystem.CommandType
	}{
		//init test
		{
			name: "Init Test 1",
			args: args{
				input: "register name1",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 2",
			args: args{
				input: "register name2",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 3",
			args: args{
				input: "create-folder name1 folder1",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 4",
			args: args{
				input: "create-folder name1 folder2 description2",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 5",
			args: args{
				input: "create-folder name2 folder1",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 6",
			args: args{
				input: "create-folder name2 folder2 forDelete",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 7",
			args: args{
				input: "create-folder name2 folder3 forRename",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 8",
			args: args{
				input: "create-file name2 folder1 file1",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Init Test 9",
			args: args{
				input: "create-file name2 folder1 file2 forDelete",
			},
			want: virtualfilesystem.Success,
		},
		//test for help
		{
			name: "Help Test 1",
			args: args{
				input: "help",
			},
			want: virtualfilesystem.Help,
		},
		//register test
		{
			name: "Register Test 1",
			args: args{
				input: "register name1",
			},
			want: virtualfilesystem.ErrUserAlreadyExists,
		},
		{
			name: "Register Test 2",
			args: args{
				input: "register",
			},
			want: virtualfilesystem.ErrRegisterArgumentMissing,
		},
		{
			name: "Register Test 3",
			args: args{
				input: "register name1 name2",
			},
			want: virtualfilesystem.ErrRegisterArgumentInvalid,
		},
		//create-folder test
		{
			name: "Create-Folder Test 1",
			args: args{
				input: "create-folder name1 folder1",
			},
			want: virtualfilesystem.ErrFolderAlreadyExists,
		},
		{
			name: "Create-Folder Test 2",
			args: args{
				input: "create-folder name1",
			},
			want: virtualfilesystem.ErrCreateFolderArgumentMissing,
		},
		{
			name: "Create-Folder Test 3",
			args: args{
				input: "create-folder name99 folder1",
			},
			want: virtualfilesystem.ErrUserNotFound,
		},
		//delete-folder test
		{
			name: "Delete-Folder Test 1",
			args: args{
				input: "delete-folder name2 folder2",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Delete-Folder Test 2",
			args: args{
				input: "delete-folder name2 folder2",
			},
			want: virtualfilesystem.ErrFolderNotFound,
		},
		{
			name: "Delete-Folder Test 3",
			args: args{
				input: "delete-folder name2",
			},
			want: virtualfilesystem.ErrDeleteFolderArgumentMissing,
		},
		{
			name: "Delete-Folder Test 4",
			args: args{
				input: "delete-folder name2 folder2 folder3",
			},
			want: virtualfilesystem.ErrDeleteFolderArgumentInvalid,
		},
		{
			name: "Delete-Folder Test 5",
			args: args{
				input: "delete-folder name99 folder1",
			},
			want: virtualfilesystem.ErrUserNotFound,
		},
		//list-folders test
		{
			name: "List-Folders Test 1",
			args: args{
				input: "list-folders name1",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 2",
			args: args{
				input: "list-folders name1 --sort-name",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 3",
			args: args{
				input: "list-folders name1 --sort-name asc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 4",
			args: args{
				input: "list-folders name1 --sort-name desc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 5",
			args: args{
				input: "list-folders name1 --sort-created",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 6",
			args: args{
				input: "list-folders name1 --sort-created asc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 7",
			args: args{
				input: "list-folders name1 --sort-created desc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Folders Test 8",
			args: args{
				input: "list-folders name1 --sort-name asc desc",
			},
			want: virtualfilesystem.ErrListFoldersArgumentInvalid,
		},
		{
			name: "List-Folders Test 9",
			args: args{
				input: "list-folders name1 --sort-nam123 desc asc",
			},
			want: virtualfilesystem.ErrListFoldersArgumentInvalid,
		},
		//rename-folder test
		{
			name: "Rename-Folder Test 1",
			args: args{
				input: "rename-folder name2 folder3 folder3",
			},
			want: virtualfilesystem.ErrFolderAlreadyExists,
		},
		{
			name: "Rename-Folder Test 2",
			args: args{
				input: "rename-folder name2 folder3 folder4",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Rename-Folder Test 3",
			args: args{
				input: "rename-folder name2 folder3 folder99",
			},
			want: virtualfilesystem.ErrFolderNotFound,
		},
		{
			name: "Rename-Folder Test 4",
			args: args{
				input: "rename-folder name2 folder3",
			},
			want: virtualfilesystem.ErrRenameFolderArgumentMissing,
		},
		{
			name: "Rename-Folder Test 5",
			args: args{
				input: "rename-folder name2 folder3 folder4 folder5",
			},
			want: virtualfilesystem.ErrRenameFolderArgumentInvalid,
		},
		{
			name: "Rename-Folder Test 6",
			args: args{
				input: "rename-folder name99 folder1 folder2",
			},
			want: virtualfilesystem.ErrUserNotFound,
		},
		//create-file test
		{
			name: "Create-File Test 1",
			args: args{
				input: "create-file name2 folder1",
			},
			want: virtualfilesystem.ErrCreateFileArgumentMissing,
		},
		{
			//user not found
			name: "Create-File Test 2",
			args: args{
				input: "create-file name99 folder1 file1",
			},
			want: virtualfilesystem.ErrUserNotFound,
		},
		{
			//folder not found
			name: "Create-File Test 3",
			args: args{
				input: "create-file name2 folder99 file1",
			},
			want: virtualfilesystem.ErrFolderNotFound,
		},
		{
			//duplicate file
			name: "Create-File Test 4",
			args: args{
				input: "create-file name2 folder1 file1",
			},
			want: virtualfilesystem.ErrFileAlreadyExists,
		},
		//delete-file test
		{
			name: "Delete-File Test 1",
			args: args{
				input: "delete-file name2 folder1 file2",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "Delete-File Test 2",
			args: args{
				input: "delete-file name2 folder1 file2",
			},
			want: virtualfilesystem.ErrFileNotFound,
		},
		{
			name: "Delete-File Test 3",
			args: args{
				input: "delete-file name2 folder1",
			},
			want: virtualfilesystem.ErrDeleteFileArgumentMissing,
		},
		{
			name: "Delete-File Test 4",
			args: args{
				input: "delete-file name2 folder1 file1 file2",
			},
			want: virtualfilesystem.ErrDeleteFileArgumentInvalid,
		},
		{
			//user not found
			name: "Delete-File Test 5",
			args: args{
				input: "delete-file name99 folder1 file1",
			},
			want: virtualfilesystem.ErrUserNotFound,
		},
		{
			//folder not found
			name: "Delete-File Test 6",
			args: args{
				input: "delete-file name2 folder99 file1",
			},
			want: virtualfilesystem.ErrFolderNotFound,
		},
		//list-files test
		{
			name: "List-Files Test 1",
			args: args{
				input: "list-files name2 folder1",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 2",
			args: args{
				input: "list-files name2 folder99",
			},
			want: virtualfilesystem.ErrFolderNotFound,
		},
		{
			name: "List-Files Test 3",
			args: args{
				input: "list-files name2 folder1 --sort-name",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 4",
			args: args{
				input: "list-files name2 folder1 --sort-name asc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 5",
			args: args{
				input: "list-files name2 folder1 --sort-name desc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 6",
			args: args{
				input: "list-files name2 folder1 --sort-created",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 7",
			args: args{
				input: "list-files name2 folder1 --sort-created asc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 8",
			args: args{
				input: "list-files name2 folder1 --sort-created desc",
			},
			want: virtualfilesystem.Success,
		},
		{
			name: "List-Files Test 9",
			args: args{
				input: "list-files name2 folder1 --sort-name asc desc",
			},
			want: virtualfilesystem.ErrListFilesArgumentInvalid,
		},
		{
			name: "List-Files Test 10",
			args: args{
				input: "list-files name2 folder1 --sort-nam123 desc asc",
			},
			want: virtualfilesystem.ErrListFilesArgumentInvalid,
		},
		//default test
		{
			name: "Default Test 1",
			args: args{
				input: "default",
			},
			want: virtualfilesystem.ErrCommandNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleCommand(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
