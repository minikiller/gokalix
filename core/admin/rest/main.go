package main

import (
	_ "github.com/lib/pq"
	"github.com/gocraft/dbr"
	"encoding/json"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"
	"kalix.com/core/api/persistence"
	"kalix.com/core/api/dao"
	"fmt"
	"log"
)

const (
	DB_USER = "postgres"
	DB_PASSWORD = "1234"
	DB_NAME = "kalix"
)

type User struct {
	Id   string `json:"id"  binding:"required"`
	Name string `json:"name"   binding:"required"`
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", postUser)
	router.PUT("/users", putUser)
	router.DELETE("/users/:id", deleteUser)
	router.Run()
}

func postUser(c *gin.Context) {
	var json User
	if c.Bind(&json) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(json.Name)
		log.Println(json.Id)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	}

	/*if c.BindWith(&json,binding.JSON) == nil {
		if json.Name == "admin" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}*/

}

func putUser(c *gin.Context) {
	var json User
	if c.Bind(&json) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(json.Name)
		log.Println(json.Id)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	}

}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Hello %s", id)
}

func getUsers(c *gin.Context) {
	_page := c.Query("page")
	if _page =="" {
		c.String(http.StatusOK, "{\"status\":\"error:page is required! \"}")
		return
	}
	page, err := strconv.ParseInt(_page, 10, 64)
	if err != nil {
		c.String(http.StatusOK, "{\"status\":\"error:no valid number of page\"}")
		return
	}

	_limit := c.Query("limit")
	limit, err := strconv.ParseInt(_limit, 10, 64)
	if err != nil {
		c.String(http.StatusOK, "{\"status\":\"error:no valid number of page\"}")
		return
	}
	c.String(http.StatusOK, getJson(page, limit))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getJson(page, perPage int64) string {

	jsonData := persistence.JsonData{}
	/*dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	       DB_USER, DB_PASSWORD, DB_NAME)*/
	/*db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()*/


	// create a connection (e.g. "postgres", "mysql", or "sqlite3")
	conn, err := dao.GetConn()
	// create a connection (e.g. "postgres", "mysql", or "sqlite3")
	//conn, err := dao.GetConn()
	checkErr(err)
	//defer conn.Close()
	// create a session for each business unit of execution (e.g. a web request or goworkers job)
	sess := conn.NewSession(nil)

	users := []User{}
	// get a record
	count, err := sess.Select("id", "name").From("sys_user").Paginate(uint64(page), uint64(perPage)).Load(&users)
	fmt.Println("total is " + strconv.Itoa(count))
	checkErr(err)
	jsonData.Data = users

	//result, _ := json.Marshal(&users)

	sess.Select("count(id)").From(
		dbr.Select("*").From("sys_user").As("count"),
	).Load(&jsonData.TotalCount)


	// JSON-ready, with dbr.Null* types serialized like you want
	//fmt.Println(string(result))

	//jsonData.Data =make([]interface{}, len(suggestion))

	/*for i,v:=range suggestion{
	       jsonData.Data[i]=v
	}*/

	v, _ := json.Marshal(jsonData)
	return string(v)

}