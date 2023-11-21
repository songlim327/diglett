package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Record struct {
	Address string `json:"address" binding:"required"`
	Port    string `json:"port" binding:"required"`
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.POST("/check", checkPortHandler)
	r.GET("/ipv4", getIpv4Handler)

	r.Run()
}

// checkPortHandler check if port is open or closed on given IP address
func checkPortHandler(c *gin.Context) {
	// validate payload format
	record := &Record{}
	if err := c.BindJSON(record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	// validate address input
	if !isIp(record.Address) {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": fmt.Sprintf("%v is not a valid IP address", record.Address)})
		return
	}

	// validate port input
	if !isPort(record.Port) {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": fmt.Sprintf("%v is not a valid port", record.Port)})
		return
	}

	err := tcpCheck(record.Address, record.Port)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("port %v closed on %v", record.Port, record.Address)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("port %v open on %v", record.Port, record.Address)})
	}
}

// getIpv4Handler returns the public IPv4 address of the request
func getIpv4Handler(c *gin.Context) {
	c.String(http.StatusOK, "%v", c.ClientIP())
}

// tcpCheck run tcp port check and return if given port is open or closed
func tcpCheck(addr, port string) error {
	timeout := time.Second
	_, err := net.DialTimeout("tcp4", net.JoinHostPort(addr, port), timeout)
	if err != nil {
		return err
	}
	return nil
}

// isIp validate given IP address
func isIp(addr string) bool {
	return net.ParseIP(addr) != nil
}

// isPort validate given port number
func isPort(port string) bool {
	p, err := strconv.ParseInt(port, 10, 0)
	if err != nil {
		return false
	}

	if p <= 0 || p >= 65535 {
		return false
	}

	return true
}
