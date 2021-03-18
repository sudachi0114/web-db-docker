package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/sudachi0114/web-db-docker/models"
	"github.com/sudachi0114/web-db-docker/repo"
)

// C(reate)
func CreateUser(c *gin.Context) {
	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db := repo.Connect(DBMS, CONNECTION)
	defer db.Close()

	var user models.User
	c.BindJSON(&user)

	log.Println("user", user)
	log.Println("user.Name", user.Name)
	log.Println("user.Email", user.Email)

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Accept Request",
	})
}

// R(ead)
func GetAllUser(c *gin.Context) {
	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db := repo.Connect(DBMS, CONNECTION)
	defer db.Close()

	var users []models.User
	db.Order("created_at asc").Find(&users)

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db := repo.Connect(DBMS, CONNECTION)
	defer db.Close()

	id := c.Param("id")
	log.Println("get user id = ", id)

	var user models.User
	db.Where("id = ?", id).First(&user) // SELECT * FROM users WHERE id = `id` LIMIT 1

	log.Println(user)
	c.JSON(http.StatusOK, user)
}

// U(pdate)

// D(elete)
func DeleteUser(c *gin.Context) {
	DBMS := "mysql"
	CONNECTION := "test:passw0rd@tcp(db:3306)/test_db?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db := repo.Connect(DBMS, CONNECTION)
	defer db.Close()

	id := c.Param("id")
	log.Println("delete user id = ", id)

	var user models.User
	db.First(&user, id)
	log.Println(user)

	// gorm.DB の Delete メソッドを呼ぶと
	//	gorm.Model の DeletedAt に日時が入り、これが入っていると削除として扱われる
	//	(論理削除になる。 ref:https://qiita.com/gold-kou/items/45a95d61d253184b0f33#delete)
	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Completed",
	})

}
