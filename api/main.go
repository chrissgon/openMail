package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	r.GET(":subject/:email/:message", func(c *gin.Context) {
		subject := c.Param("subject")
		email := c.Param("email")
		message := c.Param("message")

		if err := send(subject, email, message); err != nil {
			fmt.Println(err)
			c.Status(500)
			return
		}

		c.Status(200)
	})

	r.Run(":3000")
}

func send(subject, email, message string) error {
	var (
		host     = os.Getenv("HOST")
		port, _  = strconv.Atoi(os.Getenv("PORT"))
		username = os.Getenv("USERNAME")
		password = os.Getenv("PASSWORD")
	)

	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html; charset=\"utf-8\"", fmt.Sprintf(`
	<meta charset="utf-8">

	<b>Olá</b>, esse é um email enviado pela <i>OpenMail</i>! <br>
	
	<h4>%s</h4> <br><br>
	
	<i>OpenMail, rápido e seguro.</i>
	`, message))

	d := gomail.NewDialer(host, port, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	return d.DialAndSend(m)
}
