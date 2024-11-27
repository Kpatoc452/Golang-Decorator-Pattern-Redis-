package main

import (
	"context"
	"fmt"

	// "github.com/jackc/pgx/v5"
	// "github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6380",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}

// func ExampleClientPSQL(){
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// 	dbURL := os.Getenv("DATABASE_URL")
	
// 	conn, err := pgx.Connect(ctx, dbURL)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(ctx)


// 	_, err = conn.Exec(ctx, "INSERT INTO persons (firsName, lastName, city) VALUES($1, $2, $3)", "denis", "gatilov", "orel")
// 	if err != nil{
// 		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
// 		os.Exit(1)
// 	}	
// }

func main(){
	ExampleClient()
	// ExampleClientPSQL()
}
