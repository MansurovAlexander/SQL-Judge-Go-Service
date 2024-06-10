package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPWhiteList(whitelist map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !whitelist[c.ClientIP()] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			return
		}
	}
}

func getClientIPByHeaders(req *http.Request) (ip string, err error) {
	// Client could be behid a Proxy, so Try Request Headers (X-Forwarder)
	ipSlice := []string{}
	ipSlice = append(ipSlice, req.Header.Get("X-Forwarded-For"))
	ipSlice = append(ipSlice, req.Header.Get("x-forwarded-for"))
	ipSlice = append(ipSlice, req.Header.Get("X-FORWARDED-FOR"))
	for _, v := range ipSlice {
		log.Printf("debug: client request header check gives ip: %v", v)
		if v != "" {
			return v, nil
		}
	}
	err = errors.New("error: Could not find clients IP address from the Request Headers")
	return "", err
}
