package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	certPath string
	keyPath  string
	address  string

	correctAnswers int
	wrongAsnwers   int
)

func init() {
	flag.StringVar(&certPath, "cert", "", "path to SSL/TLS certificate file")
	flag.StringVar(&keyPath, "key", "", "path to SSL/TLS private key file")
	flag.StringVar(&address, "a", ":7000", "address to use")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

	r.GET("/ws/:id", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		for {
			// Read message from client
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				// panic(err)
				log.Printf("%s, error while reading message\n", err.Error())
				c.AbortWithError(http.StatusInternalServerError, err)
				break
			}

			f, err := os.OpenFile("./code/answer.cpp", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
			if err != nil {
				log.Printf(err.Error())
				break
			}
			_, err = f.Write(p)
			if err != nil {
				log.Printf(err.Error())
				break
			}
			var res []byte

			cmd := exec.Command("./code/comp.sh", "./answer.cpp", "./bb")
			compOut, err := cmd.CombinedOutput()
			if err != nil {
				res = compOut
				err = conn.WriteMessage(messageType, res)
				if err != nil {
					// panic(err)
					log.Printf("%s, error while writing message\n", err.Error())
					c.AbortWithError(http.StatusInternalServerError, err)
					break
				}
			} else {
				exe := exec.Command("/usr/bin/bash", "-c", "./code/run.sh")

				res, _ = exe.Output()
			}

			b := true

			exe := exec.Command("/usr/bin/bash", "-c", "./code/test.sh")

			_, err = exe.Output()
			if err != nil {
				b = false
			}

			// Echo message back to client
			err = conn.WriteMessage(messageType, []byte(fmt.Sprintln(b)))
			if err != nil {
				// panic(err)
				log.Printf("%s, error while writing message\n", err.Error())
				c.AbortWithError(http.StatusInternalServerError, err)
				break
			}
		}
	})

	if certPath == "" || keyPath == "" {
		log.Println("Warning: SSL/TLS certificate and/or private key file not provided. Running server unsecured.")
		err := r.Run(address)
		if err != nil {
			panic(err)
		}
	} else {
		err := r.RunTLS(address, certPath, keyPath)
		if err != nil {
			panic(err)
		}
	}
}

func compareFiles(file1, file2 string) (bool, error) {
	f1, err := os.Open(file1)
	if err != nil {
		return false, err
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false, err
	}
	defer f2.Close()

	f1scanner := bufio.NewScanner(f1)
	f2scanner := bufio.NewScanner(f2)

	for f1scanner.Scan() {
		for f2scanner.Scan() {
			if f1scanner.Text() != f2scanner.Text() {
				log.Print("File1: ", f1scanner.Text())
				log.Print("File2: ", f2scanner.Text())
				return false, nil
			}
		}
	}

	return true, nil
}
