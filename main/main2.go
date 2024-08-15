package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/lib/pq"
// 	_ "github.com/lib/pq"
// )

// type skill struct {
// 	Key         string   `json:"key"`
// 	Name        string   `json:"name"`
// 	Description string   `json:"description"`
// 	Logo        string   `json:"logo"`
// 	Tags        []string `json:"tags"`
// }


// func main() {
// 	connStr := "postgres://dev:dev123@localhost:5432/postgres?sslmode=disable"
// 	var err error
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	// fmt.Print(db)
// 	if err = db.Ping();err != nil {
// 		log.Fatal(err)
// 	}

// 	// fmt.Println("Succesfully connected to PostgreSQL!")
// 	router := gin.Default()

// 	h := &handler{db:db}

// 	router.GET("/ping", getPing)
// 	router.GET("/api/v1/skills", h.getAllSkills)
// 	router.GET("/api/v1/skills/:key", h.getSkillById)
// 	router.POST("/api/v1/skills", h.postSkill)
// 	router.Run()
// }

// func getPing(context *gin.Context) {
// 	context.JSON(http.StatusOK, gin.H{"message": "pong"})
// }

// type handler struct{
// 	db *sql.DB
// }

// func (h *handler) postSkill(context *gin.Context) {
// 	var newSkill skill
// 	err := context.BindJSON(&newSkill)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, "Can not extract data from request body")
// 		return
// 	}
// 	stmt, err := h.db.Prepare("INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5)")
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Can not create statement to create new skill"})
// 		return
// 	}
// 	defer stmt.Close()
// 	if _, err := stmt.Exec(newSkill.Key, newSkill.Name,	newSkill.Description, newSkill.Logo, pq.Array(newSkill.Tags)); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Can not create new skill"})
// 		return
// 	}
// 	context.JSON(http.StatusOK, newSkill)
// }

// func (h *handler) getAllSkills(context *gin.Context) {
// 	rows, err := h.db.Query("SELECT key, name, description, logo, tags FROM skill;")
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Can not get data from database")})
// 		return
// 	}
// 	defer rows.Close()
// 	var skills []skill
// 	for rows.Next() {
// 		var s skill
// 		err := rows.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, pq.Array(&s.Tags))
// 		if err != nil {
// 			context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Can not get data from database")})
// 			return
// 		}
// 		skills = append(skills, s)
// 	}
// 	context.JSON(http.StatusOK, skills)
// }

// func (h *handler) getSkillById(context *gin.Context) {
// 	paramkey := context.Param("key")
// 	row := h.db.QueryRow(fmt.Sprintf("SELECT key, name, description, logo, tags from skill where key = '%v';", paramkey))
// 	var s skill
// 	err := row.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, pq.Array(&s.Tags))
// 	if err != nil {
// 		if err == sql.ErrNoRows{
// 			context.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Not found skill with the key %v", paramkey)})
// 			return
// 		}
// 		context.JSON(http.StatusBadRequest, nil)
// 		return
// 	}
// 	context.JSON(http.StatusOK, s)
// }

