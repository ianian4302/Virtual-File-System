#   Backend Assignment - Virtual File System

##  Desiner
陳文彥

#  Start
##  Workspace
Please confirm that you are executing the file in the following path.  
```
../Virtual-File-System/
```
You should be in the same workspace as the main.go file. 
```
../Virtual-File-System/main.go
```
## Run Code
Please use the following command to execute the program to prevent unknown errors. 
```
go run .
```
If you have extension "Code Runner", you can use the button "Run Code"(Ctrl+Alt+N) with my setting.json. 

## Environment

*   IDE : vscode
*   go version : go version go1.23.1 windows/amd64

## dependence
Use the following commands or shortcut keys to confirm the extensions and tools version. 

1. Ctrl+Shift+P
2. Go: Install/Update Tools
3. Select all and dowload.

##  Command
### User
```go
//register
register [username]
```
|Field|Required|Default|
|-|-|-|
|username|Yes||
### Folder
```go
//create
create-folder [username] [foldername] [description]?

//delete
delete-folder [username] [foldername]

//list
list-folders [username] [--sort-name|--sort-created] [asc|desc]
//list example
//The user is ian, sort based on foldername, and follow the asc order.
list-folders ian --sort-name asc

//rename
rename-folder [username] [foldername] [new-folder-name]
```
|Field|Required|Default|
|-|-|-|
|username|Yes||
|foldername|Yes||
|description|No|""|
|--sort-name --sort-created|Auto|--sort-name|
|asc desc|Auto|asc|
### File
```go
//create
create-file [username] [foldername] [filename] [description]?

//delete
delete-file [username] [foldername] [filename]

//list
list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
```
|Field|Required|Default|
|-|-|-|
|username|Yes||
|foldername|Yes||
|filename|Yes||
|description|No|""|
|--sort-name --sort-created|Auto|--sort-name|
|asc desc|Auto|asc|

##  Input Validation and Restrictions
### User
####    Name
1.  can not start with number. 
2.  len should be greater than 3 chars and less than 20 chars. 
3.  only contains a-z, A-Z, 0-9. 

### Folder
####    Name
1.  can not start with number. 
2.  len should be greater than 1 chars and less than 20 chars
3.  only contains a-z, A-Z, 0-9
### File
