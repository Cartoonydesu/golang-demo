package main

import (
	_ "github.com/lib/pq"
)

// type skill struct {
// 	key         string
// 	name        string
// 	description string
// 	logo        string
// 	tags        []string
// }

// var db *sql.DB

// // var (
// // 	UNAMEDB string = "dev"
// // 	PASSDB  string = "dev123"
// // 	HOSTDB  string = "postgres"
// // 	DBNAME  string = "demo"
// // )

// func main_ex() {
// 	var err error
// 	// connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
// 	// connStr := "postgres://dev:dev123@localhost:5432/demo?sslmode=disable"
// 	connStr := "user=postgres password=1234 dbname=demo sslmode=disable host=localhost port=5432"
// 	// connStr := "user=dev password=dev123 dbname=demo sslmode=disable host=postgres port=5432"
// 	// db, err = sql.Open("postgres", "port=5432 dbname=demo user=dev password=dev123 sslmode=disable")
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		fmt.Print(err)
// 		log.Fatalf("Can't connect to database: %v", err)
// 	}
// 	defer db.Close()

// 	router := gin.Default()
// 	router.GET("/api/v1/skills", getAllSkill)
// 	router.GET("/api/v1/skills/:key", getSkillByKey)
// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	router.Run()
// }

// func getAllSkill(context *gin.Context) {
// 	context.Header("Content-type", "application/json")
// 	rows, err := db.Query("SELECT * from skill")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// var skills []skill
// 	// for rows.Next() {
// 	// 	var s skill
// 	// 	err := rows.Scan(&s.key, &s.name, &s.description, &s.logo, &s.tags)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	skills = append(skills, s)
// 	// }
// 	// err = rows.Err()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// context.IndentedJSON(http.StatusOK, skills)
// 	context.JSON(200, gin.H{
// 		"message": "all skill",
// 	})
// }

// func getSkillByKey(context *gin.Context) {
// 	context.Header("Content-Type", "application/json")
// 	paramkey := context.Param("key")
// 	rows, err := db.Query(fmt.Sprintf("SELECT * from skill where key = %q;", paramkey))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// var s skill
// 	// errRows := rows.Scan(&s.key, &s.name, &s.description, &s.logo, &s.tags)
// 	// if errRows != nil {
// 	// 	log.Fatal(errRows)
// 	// }
// 	// context.JSON(http.StatusOK, s)

// 	context.JSON(200, gin.H{
// 		"message": "skill by id",
// 	})
// }
