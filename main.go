package main

import (
	"context"
	"fmt"
	"fossil-api/internal/repository"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	repo := repository.New(conn)

	fmt.Println("getting users data...")
	data, err := repo.GetUsers(ctx)
	fmt.Println(data)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)
}
