package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

func createTable() {
	conn.Exec(context.Background(), "create table tasks(id serial primary key, description text not null)")
}

func listTasks() error {
	rows, _ := conn.Query(context.Background(), "select * from tasks")

	for rows.Next() {
		var id int32
		var description string
		err := rows.Scan(&id, &description)
		if err != nil {
			return err
		}
		fmt.Printf("%d. %s\n", id, description)
	}

	return rows.Err()
}

func showUsage() {
	var usage = `
	Usage:
	todo.exe list
	todo.exe add task
	todo.exe update taskNo item
	todo.exe remove taskNo`
	fmt.Print(usage)
}

func main() {

	var err error

	conn, err = pgx.Connect(context.Background(), "postgresql://todo-user:todo-password@localhost:5432/to-do")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

	createTable()

	if len(os.Args) == 1 {
		showUsage()
		os.Exit(0)
	}

	switch os.Args[1] {

	case "list":
		err = listTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to list tasks: %v\n", err)
			os.Exit(1)
		}

	case "add":
		_, err := conn.Exec(context.Background(), "insert into tasks(description) values($1)", os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to add task: %v\n", err)
			os.Exit(1)
		}

	case "update":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int32: %v\n", err)
			os.Exit(1)
		}
		_, err = conn.Exec(context.Background(), "update tasks set description=$1 where id=$2", os.Args[3], int32(n))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to update task: %v\n", err)
			os.Exit(1)
		}

	case "remove":
		n, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable convert task_num into int32: %v\n", err)
			os.Exit(1)
		}
		_, err = conn.Exec(context.Background(), "delete from tasks where id=$1", int32(n))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to remove task: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
		showUsage()
		os.Exit(1)
	}
}
