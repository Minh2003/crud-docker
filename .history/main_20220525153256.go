package main

import (
  "fmt"
  "strconv"
  "time"

  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  gorm.Model
  Name string
  Email string
}

func getAllUser(ctx *gin.Context) {
  db := sqlConnect()
  var users []User
  db.Order("Name asc").Find(&users)
  defer db.Close()

  if len(users) > 0 {
    fmt.Print(&users[0].CreatedAt, &users[0].Name)
  }


  ctx.HTML(200, "index.html", gin.H{
    "users": users,
  })
}

func createUser(ctx *gin.Context) {
  db := sqlConnect()
  name := ctx.PostForm("name")
  email := ctx.PostForm("email")
  db.Create(&User{Name: name, Email: email})
  defer db.Close()

  ctx.Redirect(302, "/")
}

func updateUser(ctx *gin.Context) {
  db := sqlConnect()
  id := 
}

func deleteUser(ctx *gin.Context) {
  db := sqlConnect()
    n := ctx.Param("id")
    id, err := strconv.Atoi(n)
    if err != nil {
      panic("id is not a number")
    }
    var user User
    db.First(&user, id)
    db.Delete(&user)
    defer db.Close()

    ctx.Redirect(302, "/")
}

func main() {
  db := sqlConnect()
  db.AutoMigrate(&User{})
  defer db.Close()

  router := gin.Default()
  router.LoadHTMLGlob("templates/*.html")

  router.GET("/", getAllUser)

  router.POST("/new", createUser)

  router.POST("/update/:id", deleteUser)

  router.POST("/delete/:id", deleteUser)
}

func sqlConnect() (database *gorm.DB) {
  DBMS := "mysql"
  USER := "go_test"
  PASS := "password"
  PROTOCOL := "tcp(db:3306)"
  DBNAME := "go_database"

  CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true"
  
  count := 0
  db, err := gorm.Open(DBMS, CONNECT)
  if err != nil {
    for {
      if err == nil {
        fmt.Println("")
        break
      }
      fmt.Print(".")
      time.Sleep(time.Second)
      count++
      if count > 180 {
        fmt.Println("")
        panic(err)
      }
      db, err = gorm.Open(DBMS, CONNECT)
    }
  }

  return db
}