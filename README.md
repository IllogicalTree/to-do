# CRUD todo app

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)


This is a basic to-do cli app written in golang which utilises a postgres database to store tasks. The main goals of creating this project were to experiment with a new language and to interact with a database.

## Example usage

```
todo.exe

Output: 

    Usage:
    todo.exe list
    todo.exe add task
    todo.exe update taskNo item
    todo.exe remove taskNo
```

### Adding and viewing items

```
todo.exe add 'Learn more about golang'
todo.exe add 'Explore more docker use cases'
todo.exe list 

Output:
1. Learn more about golang
2. Explore more docker use cases
```

### Updating and deleting items

```
todo.exe update 1 'Learn everyone about golang'
todo.exe delete 2
todon.exe list 

Output:
1. Learn everyone about golang
```

## Usage

A postgres database is required for this application, this can be quickly installed with docker or installed manually on your machine.

### Docker

With docker, a configured database can be spun up with just one command.

```
docker run --name postgres -e POSTGRES_DB=to-do -e POSTGRES_USER=todo-user -e POSTGRES_PASSWORD=todo-password -p 5432:5432 -d postgres
```

### Manually

You should follow the relevant instructions for your operating system at https://www.postgresql.org/download/.

You should also create a database named 'to-do' before continuing further or you will encounter errors.

Security was not seen as a concern for this project so the database credentials have been hard coded as 'todo-user' and 'todo-password' for the database user and password respectively. 

### Run in development mode

You should never blindly run executables, while this code isn't malicious you should review it before building anything.

```
cd app
go run . add 'Learn more about golang'
go run . add 'Explore more docker use cases'
go run . list

Output:
1. Learn more about golang
2. Explore more docker use cases
```

### Build binary

```
cd app
go build -o ../bin/todo.exe
cd ../bin

todo.exe add 'Learn more about golang'
todo.exe add 'Explore more docker use cases'
todo.exe list 

Output:
1. Learn more about golang
2. Explore more docker use cases
```

## Clean up

Remember to remove your postgres instance if you no longer need it, if you were following along with docker you may run the following command to first stop and then delete the container.

```
docker stop postgres && docker rm postgres
```

If you installed postgres manually I suggest you refer to the manual as the steps vary depending on your system.

## Issues

Should you discover any problems with the project you should raise an issue, if you feel confident in fixing the issue you are welcome to contribute.

## Contributing

If you feel the urge to contribute to this project you should clone the project and submit a pull request with your changes for review.