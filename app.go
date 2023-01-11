package main

import (
	"database/sql"
	"fmt"
	"log"

	"tt/gen/model"
	. "tt/gen/table"

	. "github.com/go-jet/jet/v2/sqlite"
	_ "modernc.org/sqlite"
)

func main() {

	// db db!
	db, err := sql.Open("sqlite", "./db/database.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	stm := SELECT(Posts.AllColumns, Comments.AllColumns).
		FROM(Posts.LEFT_JOIN(Comments, Comments.PostID.EQ(Posts.ID)))

	var dist []struct {
		model.Posts
		Comments []model.Comments
	}

	err = stm.Query(db, &dist)
	if err != nil {
		log.Fatalln(err)
	}

	for _, p := range dist {
		fmt.Printf("Post ID: %v, Title: %v, CreatedAt: %v\n", *p.ID, p.Title, p.CreatedAt)
		for _, c := range p.Comments {
			fmt.Printf("Comment ID: %v, Text: %v, CreatedAt: %v, PostID: %v\n", *c.ID, c.Text, c.CreatedAt, c.PostID)
		}
		fmt.Println("-------------")
	}

}
