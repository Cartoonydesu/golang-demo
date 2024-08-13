package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type skill struct {
// 	Key         string   `json:"key"`
// 	Name        string   `json:"name"`
// 	Description string   `json:"description"`
// 	Logo        string   `json:"logo"`
// 	Tags        []string `json:"tags"`
// }

// var skills = []skill{
// 	{
// 		Key:         "go",
// 		Name:        "Go",
// 		Description: "Go is a statically typed, compiled programming language designed at Google.",
// 		Logo:        "https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg",
// 		Tags:        []string{"programming language", "system"},
// 	},
// 	{
// 		Key:         "nodejs",
// 		Name:        "Node.js",
// 		Description: "Node.js is an open-source, cross-platform, JavaScript runtime environment that executes JavaScript code outside of a browser.",
// 		Logo:        "https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg",
// 		Tags:        []string{"runtime", "javascript"},
// 	},
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/api/v1/skills", getSkills)
// 	router.GET("/api/v1/skills/:key", getSkillByKey)
// 	router.POST("api/v1/skills", postSkills)
// 	router.Run()
// }

// func getSkills(context *gin.Context) {
// 	context.Header("Content-Type", "application/json")
// 	fmt.Println(skills)
// 	context.IndentedJSON(http.StatusOK, skills)
// }

// func postSkills(context *gin.Context) {
// 	var newSkill skill
// 	err := context.BindJSON(&newSkill)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	skills = append(skills, newSkill)
// 	context.IndentedJSON(http.StatusCreated, newSkill)
// }

// func getSkillByKey(context *gin.Context) {
// 	key := context.Param("key")
// 	for _, skill := range skills {
// 		if skill.Key == key {
// 			context.IndentedJSON(http.StatusOK, skill)
// 			return
// 		}
// 	}
// 	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Skill not found."})
// }
