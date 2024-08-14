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


func main() {
	connStr := "postgres://dev:dev123@0.0.0.0:5432/postgres?sslmode=disable"
	// connStr := "user=dev password=dev123 dbname=postgres sslmode=disable port=5432 host=127.0.0.1"
	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(db)
	// if err = db.Ping();err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("Succesfully connected to PostgreSQL!")
	router := gin.Default()
	router.GET("/ping", getPing)

	h := &handler{db:db}
	router.GET("api/v1/skills", h.getAllSkills)
	router.GET("/api/v1/skills/:key", getSkillById)

	router.Run()
}

func getPing(context *gin.Context) {
	context.JSON(200, gin.H{"message": "pong"})
}

type handler struct{
	db *sql.DB
}

func  (h *handler)getAllSkills(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	// rows, err := db.Query(fmt.Sprintf("SELECT * FROM %q;", "skill"))
	// rows, err := db.Query("SELECT * FROM pg_catalog.pg_tables where schemaname = 'postgres';")
	rows, err := h.db.Query("SELECT * FROM skill;")
	// rows, err := db.Query("SELECT key, name, description, logo, tags FROM skill")
	// rows, err := db.Query(fmt.Sprintf("SELECT key, name, description, logo, tags FROM %q;", "skill"))
	
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	// var count = 0
	// for rows.Next(){
	// 	count++
	// }
	var skills []skill
	for rows.Next() {
		var s skills
		err := rows.Scan(&s.key, &s.name, &s.description, &s.logo, &s.tags)
		if err != nil {
			log.Fatal(err)
		}
		skills.append(s)
	}
	
	defer rows.Close()
	context.JSON(200, skill)
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
