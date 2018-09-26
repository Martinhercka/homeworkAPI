package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:test@/world_x")

	type City struct {
		Code string
	}
	router := gin.Default()

	router.GET("/city/:name", func(c *gin.Context) {
		var (
			city   City
			result gin.H
		)
		Name := c.Param("name")
		row := db.QueryRow("select countrycode from city where name = ?;", Name)
		err = row.Scan(&city.Code)
		if err != nil {

			result = gin.H{
				"result": nil,
			}
		} else {
			result = gin.H{
				"result": city,
			}
		}
		c.JSON(http.StatusOK, result)

	})
	type Struct struct {
		Word string `json:"city"`
	}

	router.POST("/getCode", func(c *gin.Context) {
		var json Struct

		if err = c.ShouldBindJSON(&json); err == nil {

			rows, err := db.Query("SELECT countrycode from city where name=? LIMIT 1", json.Word)

			if err != nil {

				panic(err)
			}
			defer rows.Close()
			for rows.Next() {

				var cCode string
				err = rows.Scan(&cCode)

				if err != nil {

					panic(err)
				}

				c.JSON(http.StatusOK, gin.H{"Code": cCode})
			}

			err = rows.Err()
			if err != nil {
				panic(err)
			}

		}
	})

	type Struct2 struct {
		Word2 string `json:"code"`
	}

	router.POST("/getCity", func(c *gin.Context) {
		var json Struct2

		if err = c.ShouldBindJSON(&json); err == nil {

			rows, err := db.Query("SELECT name from city where countrycode=?", json.Word2)

			if err != nil {

				panic(err)
			}
			defer rows.Close()
			var cName string
			for rows.Next() {

				err = rows.Scan(&cName)

				if err != nil {

					panic(err)
				}

				c.JSON(http.StatusOK, gin.H{"City": cName})
			}

		}
	})

	router.Run(":3000")
}
