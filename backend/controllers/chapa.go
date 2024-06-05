package controllers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitializeTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {

		url := "https://api.chapa.co/v1/transaction/initialize"
		method := "POST"

		payload := strings.NewReader(`{
		"amount": "100",
		"currency": "ETB",
		"email": "abebech_bekele@gmail.com",
		"first_name": "Bilen",
		"last_name": "Gizachew",
		"phone_number": "0912345678",
		"tx_ref": "chewatatest-6668",
		"callback_url": "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
		"return_url": "https://www.google.com/",
		"customization[title]": "Payment for my favourite merchant",
		"customization[description]": "I love online payments"
	}`)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		req.Header.Add("Authorization", "Bearer CHASECK_TEST-0zTPIQ2wFsVRrogvJ4Lc1oTPsImdJ26o")
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(string(body))
		c.JSON(http.StatusOK, gin.H{"response": string(body)})
	}
}
