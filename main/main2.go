package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

var db *sql.DB

func main() {
	// connStr := "localhost://postgres:postgres@localhost:5432/demo?sslmode=disable"
	connStr := "user=dev password=dev123 dbname=postgres sslmode=disable port=5432 host=127.0.0.1"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("First")
		log.Fatal(err)
	}
	// fmt.Print(db)
	err = db.Ping()
	if err != nil {
		fmt.Print("Second")
		log.Fatal(err)
	}
	fmt.Println("Succesfully connected to PostgreSQL!")
	router := gin.Default()
	router.GET("/ping", getPing)
	router.GET("api/v1/skills", getAllSkills)
	router.GET("/api/v1/skills/:key", getSkillById)

	router.Run()
}

func getPing(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}

func getAllSkills(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	// rows, err := db.Query(fmt.Sprintf("SELECT * FROM %q;", "skill"))
	rows, err := db.Query("SELECT key, name, description, logo, tags FROM skill;")
	// rows, err := db.Query(fmt.Sprintf("SELECT key, name, description, logo, tags FROM %q;", "skill"))
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	defer rows.Close()
	context.JSON(200, gin.H{"message": "pong"})
}

func getSkillById(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	paramkey := context.Param("key")
	// rows, err := db.Query(fmt.Sprintf("SELECT * from skill where key = %v", paramkey))
	// if err != nil {
	// 	log.Panic("Error get skill by id!")
	// }
	// defer rows.Close()

	// err = rows.Scan(&rows.key, &rows.name, )
	context.JSON(200, gin.H{
		"message": paramkey,
	})
}
