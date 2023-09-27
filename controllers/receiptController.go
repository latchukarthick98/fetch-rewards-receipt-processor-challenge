/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package controllers

import (
	"fetch-rewards-receipt-processor-challenge/datastore"
	"math"
	"strconv"

	"fetch-rewards-receipt-processor-challenge/models"

	"fetch-rewards-receipt-processor-challenge/helpers"

	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type resultItem struct {
	ID string `json:"id"`
}

// Calculates points based on the no. of Alpha numeric characters in the Retailer's name
func calculateRetailerPoints(name string) int {
	return helpers.CountAlphaNumeric(name)
}

// Calculate points if the Total amount is rounded with no cents
func calculateRoundAmountPoints(amount string) int {
	if helpers.IsRounded(amount) {
		return 50
	}
	return 0
}

func calculateQuarterPoints(amount string) int {
	if helpers.IsMultipleOfQuarter(amount) {
		return 25
	}
	return 0
}

func calculatePointsForPairs(items []models.Item) int {
	return helpers.CountPairs(items) * 5
}

func calculatePointsForItemDesc(items []models.Item) int {
	points := 0
	for _, item := range items {
		if helpers.TrimmedLength(item.ShortDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return 0
			}
			point := int(math.Ceil(price * 0.2))
			fmt.Printf("Point for %s : %d\n", item.ShortDescription, point)
			points += point
		}
	}
	fmt.Printf("Points for Item Description: %d \n", points)
	return points
}

func calculatePoints(receipt models.Receipt) int {
	res := calculateRetailerPoints(receipt.Retailer)
	res += calculateRoundAmountPoints(receipt.Total)
	res += calculateQuarterPoints(receipt.Total)
	res += calculatePointsForPairs(receipt.Items)
	res += calculatePointsForItemDesc(receipt.Items)
	return res
}

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	// Generate a random UUID (version 4)
	u, error := uuid.NewRandom()
	if error != nil {
		fmt.Println("Failed to generate UUID:", error)
		return
	}

	// Parse request body
	if err := c.BindJSON((&receipt)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Invalid Input",
			"error": err,
		})
		return
	}

	// Validate request body
	validate := validator.New()
	err := validate.Struct(receipt)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(400, gin.H{
			"error": validationErrors.Error(),
		})
		return
	}

	points := calculatePoints(receipt)

	ds := datastore.Points
	ds[u.String()] = points

	fmt.Println(ds)
	result := resultItem{
		ID: u.String(),
	}

	c.IndentedJSON(http.StatusCreated, result)

}
