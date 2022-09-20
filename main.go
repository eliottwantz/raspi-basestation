package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"reflect"

	"app/db"

	_ "modernc.org/sqlite"
)

const (
	MAXLINE = 1024
)

func main() {

	err := DbStuff()
	if err != nil {
		panic(err)
	}

	// conn, err := net.ListenUDP("udp", &net.UDPAddr{
	// 	Port: 8080,
	// 	IP:   net.ParseIP("127.0.0.1"),
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// defer conn.Close()
	// fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	// for {
	// 	message := make([]byte, MAXLINE)
	// 	rlen, remote, err := conn.ReadFromUDP(message[:])
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	var sensorState pb.SensorState
	// 	err = proto.Unmarshal(message[:rlen], &sensorState)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Printf("[%s]: %v\n", remote, &sensorState)
	// }
}

//go:embed db/sql/schema.sql
var ddl string

func DbStuff() error {
	ctx := context.Background()

	pdb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return err
	}

	// create tables
	if _, err := pdb.ExecContext(ctx, ddl); err != nil {
		return err
	}

	queries := db.New(pdb)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return nil
}
