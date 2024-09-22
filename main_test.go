package main

import "testing"

func Test_handleCommand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//init field
		{
			name: "Init Test case 1",
			args: args{
				input: "register userName1",
			},
			want: true,
		},
		{
			name: "Init Test case 2",
			args: args{
				input: "register userName2",
			},
			want: true,
		},
		{
			name: "Init Test case 3",
			args: args{
				input: "create-folder userName1 folderName1",
			},
			want: true,
		},
		{
			name: "Init Test case 4",
			args: args{
				input: "create-folder userName1 folderName2",
			},
			want: true,
		},
		{
			name: "Init Test case 5",
			args: args{
				input: "create-folder userName2 folderName1",
			},
			want: true,
		},
		{
			name: "Init Test case 6",
			args: args{
				input: "create-folder userName2 folderName2",
			},
			want: true,
		},
		{
			name: "Init Test case 7",
			args: args{
				input: "create-file userName1 folderName1 fileName1",
			},
			want: true,
		},
		{
			name: "Init Test case 8",
			args: args{
				input: "create-file userName1 folderName1 fileName2",
			},
			want: true,
		},
		//help field
		{
			name: "Help",
			args: args{
				input: "help",
			},
			want: true,
		},
		//register field
		{
			name: "Register Test case 1",
			args: args{
				input: "register",
			},
			want: false,
		},
		{
			name: "Register Test case 2",
			args: args{
				input: "register userName1 userName2",
			},
			want: false,
		},
		{
			name: "Register Test case 3",
			args: args{
				input: "register a",
			},
			want: false,
		},
		{
			//duplicate username
			name: "Register Test case 4",
			args: args{
				input: "register userName1",
			},
			want: false,
		},
		//---------------------------------------------------------------------------------------
		//folder field
		//create-folder field
		{
			name: "Create-folder Test case 1",
			args: args{
				input: "create-folder",
			},
			want: false,
		},
		{
			name: "Create-folder Test case 2",
			args: args{
				input: "create-folder userName1",
			},
			want: false,
		},
		{
			//duplicate folder name
			name: "Create-folder Test case 3",
			args: args{
				input: "create-folder userName1 folderName1",
			},
			want: false,
		},
		{
			//user not found
			name: "Create-folder Test case 4",
			args: args{
				input: "create-folder userName99 folderName1",
			},
			want: false,
		},
		{
			//folder name length is greater than 20
			name: "Create-folder Test case 5",
			args: args{
				input: "create-folder userName1 folderName123456789012345678901",
			},
			want: false,
		},
		//delete-folder field
		{
			name: "Delete-folder Test case 1",
			args: args{
				input: "delete-folder",
			},
			want: false,
		},
		{
			name: "Delete-folder Test case 2",
			args: args{
				input: "delete-folder userName2",
			},
			want: false,
		},
		{
			name: "Delete-folder Test case 3",
			args: args{
				input: "delete-folder userName2 folderName2",
			},
			want: true,
		},
		{
			//folder not found
			name: "Delete-folder Test case 4",
			args: args{
				input: "delete-folder userName2 folderName2",
			},
			want: false,
		},
		{
			//user not found
			name: "Delete-folder Test case 5",
			args: args{
				input: "delete-folder userName99 folderName3",
			},
			want: false,
		},
		//list-folder field
		{
			name: "List-folders Test case 1",
			args: args{
				input: "list-folders",
			},
			want: false,
		},
		{
			name: "List-folders Test case 2",
			args: args{
				input: "list-folders userName1",
			},
			want: true,
		},
		{
			//user not found
			name: "List-folders Test case 3",
			args: args{
				input: "list-folders userName99",
			},
			want: false,
		},
		{
			//sort method is not valid
			name: "List-folders Test case 4",
			args: args{
				input: "list-folders userName1 sortMethod",
			},
			want: false,
		},
		{
			name: "List-folders Test case 5",
			args: args{
				input: "list-folders userName1 --sort-name",
			},
			want: true,
		},
		{
			name: "List-folders Test case 6",
			args: args{
				input: "list-folders userName1 --sort-name desc",
			},
			want: true,
		},
		{
			name: "List-folders Test case 7",
			args: args{
				input: "list-folders userName1 --sort-created",
			},
			want: true,
		},
		{
			name: "List-folders Test case 8",
			args: args{
				input: "list-folders userName1 --sort-created desc",
			},
			want: true,
		},
		{
			//invalid sort method
			name: "List-folders Test case 9",
			args: args{
				input: "list-folders userName1 --sort-created desc1",
			},
			want: false,
		},
		//rename-folder field
		{
			name: "Rename-folder Test case 1",
			args: args{
				input: "rename-folder",
			},
			want: false,
		},
		{
			name: "Rename-folder Test case 2",
			args: args{
				input: "rename-folder userName2",
			},
			want: false,
		},
		{
			name: "Rename-folder Test case 3",
			args: args{
				input: "rename-folder userName2 folderName1 folderName2",
			},
			want: true,
		},
		{
			//folder new name already exists
			name: "Rename-folder Test case 4",
			args: args{
				input: "rename-folder userName2 folderName1 folderName2",
			},
			want: false,
		},
		{
			//folder not found
			name: "Rename-folder Test case 5",
			args: args{
				input: "rename-folder userName2 folderName1 folderName3",
			},
			want: false,
		},
		{
			name: "Rename-folder Test case 3",
			args: args{
				input: "rename-folder userName2 folderName2 folderName1",
			},
			want: true,
		},
		{
			//user not found
			name: "Rename-folder Test case 5",
			args: args{
				input: "rename-folder userName99 folderName3 folderName4",
			},
			want: false,
		},
		{
			//folder name length is greater than 20
			name: "Rename-folder Test case 6",
			args: args{
				input: "rename-folder userName2 folderName1 folderName123456789012345678901",
			},
			want: false,
		},
		//---------------------------------------------------------------------------------------
		//file field
		//create-file field
		{
			name: "Create-file Test case 1",
			args: args{
				input: "create-file",
			},
			want: false,
		},
		{
			name: "Create-file Test case 2",
			args: args{
				input: "create-file userName1",
			},
			want: false,
		},
		{
			name: "Create-file Test case 3",
			args: args{
				input: "create-file userName1 folderName1",
			},
			want: false,
		},
		{
			//duplicate file name
			name: "Create-file Test case 4",
			args: args{
				input: "create-file userName1 folderName1 fileName1",
			},
			want: false,
		},
		{
			//folder not found
			name: "Create-file Test case 5",
			args: args{
				input: "create-file userName1 folderName99 fileName1",
			},
			want: false,
		},
		{
			//user not found
			name: "Create-file Test case 6",
			args: args{
				input: "create-file userName99 folderName1 fileName1",
			},
			want: false,
		},
		{
			//file name length is greater than 20
			name: "Create-file Test case 7",
			args: args{
				input: "create-file userName1 folderName1 fileName123456789012345678901",
			},
			want: false,
		},
		//delete-file field
		{
			name: "Delete-file Test case 1",
			args: args{
				input: "delete-file",
			},
			want: false,
		},
		{
			name: "Delete-file Test case 2",
			args: args{
				input: "delete-file userName1",
			},
			want: false,
		},
		{
			name: "Delete-file Test case 3",
			args: args{
				input: "delete-file userName1 folderName1",
			},
			want: false,
		},
		{
			name: "Delete-file Test case 4",
			args: args{
				input: "delete-file userName1 folderName1 fileName1",
			},
			want: true,
		},
		{
			//file not found
			name: "Delete-file Test case 5",
			args: args{
				input: "delete-file userName1 folderName1 fileName1",
			},
			want: false,
		},
		{
			//folder not found
			name: "Delete-file Test case 6",
			args: args{
				input: "delete-file userName1 folderName99 fileName1",
			},
			want: false,
		},
		{
			//user not found
			name: "Delete-file Test case 7",
			args: args{
				input: "delete-file userName99 folderName1 fileName1",
			},
			want: false,
		},
		//list-file field
		{
			name: "List-files Test case 1",
			args: args{
				input: "list-files",
			},
			want: false,
		},
		{
			name: "List-files Test case 2",
			args: args{
				input: "list-files userName1",
			},
			want: false,
		},
		{
			//user not found
			name: "List-files Test case 3",
			args: args{
				input: "list-files userName99",
			},
			want: false,
		},
		{
			//folder not found
			name: "List-files Test case 4",
			args: args{
				input: "list-files userName1 folderName99",
			},
			want: false,
		},
		{
			//sort method is not valid
			name: "List-files Test case 5",
			args: args{
				input: "list-files userName1 folderName1 sortMethod",
			},
			want: false,
		},
		{
			name: "List-files Test case 6",
			args: args{
				input: "list-files userName1 folderName1 --sort-name",
			},
			want: true,
		},
		{
			name: "List-files Test case 7",
			args: args{
				input: "list-files userName1 folderName1 --sort-name desc",
			},
			want: true,
		},
		{
			name: "List-files Test case 8",
			args: args{
				input: "list-files userName1 folderName1 --sort-created",
			},
			want: true,
		},
		{
			name: "List-files Test case 9",
			args: args{
				input: "list-files userName1 folderName1 --sort-created desc",
			},
			want: true,
		},
		{
			//invalid sort method
			name: "List-files Test case 10",
			args: args{
				input: "list-files userName1 folderName1 --sort-created desc1",
			},
			want: false,
		},
		//default
		{
			name: "Default Test case 1",
			args: args{
				input: "default",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleCommand(tt.args.input); got != tt.want {
				t.Errorf("handleCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
