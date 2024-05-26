package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	address string

	correctAnswers int
	wrongAsnwers   int
)

func main() {
	flag.Parse()

	r := gin.Default()

	// Serve HTML page to trigger connection
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	r.GET("/correct", func(c *gin.Context) { c.JSON(http.StatusOK, correctAnswers) })
	r.GET("/wrong", func(c *gin.Context) { c.JSON(http.StatusOK, wrongAsnwers) })
	r.POST("/setcorrect/:amount", func(c *gin.Context) {
		amount := c.Param("amount")
		num, err := strconv.Atoi(amount)
		if err != nil {
			return
		}
		correctAnswers = num
	})

	r.POST("/setwrong/:amount", func(c *gin.Context) {
		amount := c.Param("amount")
		num, err := strconv.Atoi(amount)
		if err != nil {
			return
		}
		wrongAsnwers = num
	})

	// Handle WebSocket connections
	r.POST("/code", func(c *gin.Context) {
		body := c.Request.Body
		codeBytes, err := io.ReadAll(body)
		if err != nil {
			log.Printf(err.Error())
		}
		f, err := os.OpenFile("./code/answer.cpp", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		_, err = f.Write(codeBytes)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		var res []byte

		cmd := exec.Command("./comp.sh", "./code/answer.cpp", "./code/bb")
		compOut, err := cmd.CombinedOutput()
		if err != nil {
			res = compOut
			c.JSON(http.StatusBadRequest, string(res))
			if err != nil {
				// panic(err)
				log.Printf("%s, error while writing message\n", err.Error())
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		} else {
			exe := exec.Command("/usr/bin/bash", "-c", "./code/run.sh")

			res, err = exe.Output()
			if err != nil {
				c.JSON(http.StatusBadRequest, string(res))
				return
			}
			log.Println(string(res))
			c.JSON(http.StatusOK, string(res))
			return
		}

	})

	r.GET("/code/check", func(c *gin.Context) {
		b := true

		exe := exec.Command("/usr/bin/bash", "-c", "./code/test.sh")

		_, err := exe.Output()
		log.Println(err)
		if err != nil {
			b = false
		}

		// Echo message back to client
		if b {
			c.JSON(http.StatusOK, "Yay!")
			correctAnswers += 1
		} else {
			c.JSON(http.StatusBadRequest, "Nooo!")
			wrongAsnwers += 1
		}
	})

	r.Run(":7000")
}
