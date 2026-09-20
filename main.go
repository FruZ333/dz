package main

import (
	"context"
	"dz/postgres-sql/connection"
	"fmt"
)

func main() {
	ctx := context.Background()

	conn, err := connection.ConnectionDB(ctx)
	if err != nil {
		fmt.Println("error:", err)
	} 
}
