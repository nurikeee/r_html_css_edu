package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Congratulations! You reached a server!")
	})

	caCertPool := x509.NewCertPool()
	caCertFile, err := ioutil.ReadFile("./root.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool.AppendCertsFromPEM(caCertFile)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert, //<-- this is the key
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", 11080),
		TLSConfig: tlsConfig,
	}

	server.Handler = r
	err = server.ListenAndServeTLS("root.pem", "rootCA.key")
	if err != nil {
		log.Fatal(err)
	}
}
