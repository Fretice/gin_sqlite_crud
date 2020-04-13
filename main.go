package main

import (
	"fmt"
	"gin_sqlite_crud_demo/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/ini.v1"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB

const pagesize = 10

func init() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println("Can Not Load App.ini file, Please Check app.ini file is exist")
		os.Exit(1)
	}
	dbPath := cfg.Section("database").Key("PATH").String()
	fmt.Println(dbPath)
	db, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Cant Not Open db file, Please Check db file or db config")
		os.Exit(1)
	}

	dbExist := db.HasTable("students")
	if !dbExist {
		db.CreateTable(&models.Student{})
	}

}

func main() {

	defer db.Close()

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		var model = models.Student{}
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		datas := model.GetList(page,pagesize, db)
		total_count := model.GetListTotalCount(db)
		var page_List []int
		page_count := int(math.Ceil(float64(total_count / pagesize)))
		if total_count%pagesize != 0 {
			page_count += 1
		}
		for i := 1; i <= page_count; i++ {
			page_List = append(page_List, i)
		}
		c.HTML(200, "index/home", gin.H{
			"Datas":      datas,
			"CurPage":    page,
			"PageList":   page_List,
		})
	})

	router.GET("/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/add", gin.H{})

	})

	router.GET("/detail", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		var model models.Student
		model = model.GetDetail(id, db)
		c.HTML(200, "index/edit", gin.H{
			"Detail": model,
		})
	})

	router.POST("/save", func(c *gin.Context) {
		age, _ := strconv.Atoi(c.PostForm("age"))
		gender, _ := strconv.Atoi(c.PostForm("gender"))
		timeLayout := "2006-01-02" //转化所需模板
		loc, _ := time.LoadLocation("Local")
		barth := c.PostForm("birthday")
		id, _ := strconv.Atoi(c.PostForm("id"))
		birthDay, _ := time.ParseInLocation(timeLayout, barth, loc)
		var student models.Student
		if id != 0 {
			student = student.GetDetail(id, db)
			student.BirthDay = birthDay
			student.Name = c.PostForm("name")
			student.Email = c.PostForm("email")
			student.Phone = c.PostForm("phone")
			student.Gender = gender
			student.Address = c.PostForm("address")
			student.Update(&student, db)
		} else {
			student = models.Student{Name: c.PostForm("name"), Age: age, Address: c.PostForm("address"), Gender: gender, Phone: c.PostForm("phone"), Email: c.PostForm("email"), BirthDay: birthDay}

			student.Create(&student, db)
		}

		c.Redirect(302, "/")
	})

	router.GET("/search", func(c *gin.Context) {
		searchKey := c.Query("searchkey")
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		var model = models.Student{}
		students := model.Search(searchKey,page,pagesize, db)
		total_count:=  model.SearchResultCount(searchKey, db)
		var page_List []int
		page_count := int(math.Ceil(float64(total_count / pagesize)))
		if total_count%pagesize != 0 {
			page_count += 1
		}
		for i := 1; i <= page_count; i++ {
			page_List = append(page_List, i)
		}
		c.HTML(200, "index/home", gin.H{
			"Datas":      students,
			"CurPage":    page,
			"PageList":   page_List,
			"SearchKey":  searchKey,
		})

	})

	router.DELETE("/delete", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		var model models.Student
		model.Delete(id, db)
	})

	var _ = router.Run(":8080")
}
