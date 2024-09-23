#   Backend Assignment - Virtual File System

##  Desiner
Wen Yen Chen / 陳文彥 

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
## Environment
*   IDE : Vscode
*   Extension : Go for Visual Studio Code
*   go version : go version go1.23.1 windows/amd64

## Dependence
Use the following commands or shortcut keys to confirm the extensions and tools version. 

1. Ctrl+Shift+P
2. Go: Install/Update Tools
3. Select all and dowload.

## Run Code
Please use the following command to execute the program to prevent unknown errors. 
```java
go run main.go
go run .
```
If you have extension "Code Runner", you can use the button "Run Code"(Ctrl+Alt+N) with my setting.json. 

##  Run Test
```java
go test ./... -cover
```

|Folder|Coverage|
|-|-|
|./|82.1%|
|./virtual-file-system|33.3%|
|./util|0.0%|
##  Command
### User
```java
//register
register [username]
```

|Field|Required|Default|
|-|-|-|
|username|Yes||
### Folder
```java
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
|new-folder-name|Yes||
|description|No|""|
|--sort-name --sort-created|Auto|--sort-name|
|asc desc|Auto|asc|
### File
```java
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
### Universal
1.  It is not case sensitive. "FOLDER1" and "folder1" have the same name. Please make the name unique and cannot be repeated.
2.  Can not start with number. 
3.  Only contains a-z, A-Z, 0-9. 
4.  Cannot contain special characters whitespace. 

### User
####    Name
1.  Can not start with number. 
2.  Len should be greater than 3 chars and less than 20 chars. 
3.  Only contains a-z, A-Z, 0-9. 

### Folder
####    Name
1.  Can not start with number. 
2.  Len should be greater than 1 chars and less than 20 chars
3.  Only contains a-z, A-Z, 0-9
### File
####    
1.  Can not start with number. 
2.  Len should be greater than 1 chars and less than 20 chars
3.  Only contains a-z, A-Z, 0-9

##  Design concept
In this project, I designed the following architecture: 

First, the user, file, folder structure and the relationship between them are defined.

Using this structure allows me to clearly feel the relationship between them when reading the information.

Next, I use **handleCommand function** to wrap the command processing part to make the logic in the command clearer and more readable.

In order to make **handleCommand** look more concise, I packaged some functions into **package virtualfilesystem** and **package util** Let handleCommand focus on reading the commands. 

I defined the error type for handlecommand to distinguish the error types that were originally mixed together.

At this point, the entire architecture is clear and leaves room for new functions in the future.

I just need to focus on setting the rules and paying attention to edge cases, and follow the already established framework to complete the system. 

##  Unit Test Design
|Folder|Coverage|
|-|-|
|./|82.1%|
|./virtual-file-system|33.3%|
|./util|0.0%|

Even though I got what seemed like good coverage, I think my unit test design was pretty poor.

I did take most of the situations into consideration, but I almost adopted a brute force approach in designing the unit test.


