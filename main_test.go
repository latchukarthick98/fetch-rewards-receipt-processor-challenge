/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package main

import (
	"bytes"
	"encoding/json"
	"fetch-rewards-receipt-processor-challenge/controllers"
	"fetch-rewards-receipt-processor-challenge/datastore"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type proccessResultBody struct {
	ID string `json:"id"`
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestProcessReceipt(t *testing.T) {
	mockResponse := `{"points":28}`
	t.Cleanup(datastore.Cleanup)
	r := SetUpRouter()
	r.POST("/receipts/process", controllers.ProcessReceipt)
	r.GET("/receipts/:id/points", controllers.GetPoints)

	data := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
		  {
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
		  },{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
		  },{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
		  },{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
		  },{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
		  }
		],
		"total": "35.35"
	  }`

	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(data)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var res proccessResultBody
	err := json.Unmarshal(w.Body.Bytes(), &res)
	// fmt.Printf("ID: %s\n", res.ID)
	// fmt.Printf("Result body: %s\n", w.Body.String())
	if err != nil {
		fmt.Println(err)
		return
	}
	secondReqURL := fmt.Sprintf("/receipts/%s/points", res.ID)
	req, _ = http.NewRequest("GET", secondReqURL, nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	// fmt.Println(datastore.Points)
	// fmt.Printf("Second Response: %s", responseData)
	require.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}