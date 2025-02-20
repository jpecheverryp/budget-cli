package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"budget-cli.jpech.dev/internal/repository"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
    ctx := context.Background()

    conn, err := pgx.Connect(context.Background(), os.Getenv("BUDGET_CLI_DB_URL"))
    if err!=nil{
        log.Fatal(err)
        os.Exit(1)
    }
    defer conn.Close(ctx)

    fmt.Println("Budget-CLI Application")
    repo := repository.New(conn)
    _, err = repo.CreateAccount(ctx, "Chase")
    if err!=nil{
        log.Fatal(err)
    }

    accounts, err := repo.ReadAllAccounts(ctx)
    if err!=nil{
        log.Fatal(err)
    }
    for _, account := range accounts {
        log.Println(account.ID, account.Name, account.CreatedAt)
    }
}
