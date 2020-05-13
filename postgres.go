package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/jackc/pgx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "miz"
	password = "admin"
	dbname   = "super_awesome_application"
)

func main() {
	// Connect to postgres
	conn, err := pgx.Connect(context.Background(), "postgresql://miz:admin@localhost/super_awesome_application")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Query Posts
	rows, err := conn.Query(context.Background(), "select id, title, content from posts")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	type Post struct {
		id      int8
		title   string
		content string
	}

	posts := []Post{}

	for rows.Next() {
		var (
			id      int8
			title   string
			content string
		)

		err := rows.Scan(&id, &title, &content)

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		}

		posts = append(posts, Post{
			title:   title,
			content: content,
			id:      id,
		})
	}

	for i, post := range posts {
		fmt.Println(i, post.title, post.content, post.id)
		fmt.Println(reflect.TypeOf(post.id))
	}

	// Check Data Type reflect.TypeOf
	fmt.Println(reflect.TypeOf(posts))

	fmt.Printf("Connected")
}
