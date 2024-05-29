package server

import "github.com/gin-gonic/gin"

type Server struct {
	address string
}

func NewServer(address string) *Server {

	return &Server{
		address: address,
	}
}

func (s *Server) Serve() error {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"methode": "Get",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		var user struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"methode": "Post"})
	})
	err := r.Run(s.address)
	if err != nil {
		return err
	}
	return nil
}
