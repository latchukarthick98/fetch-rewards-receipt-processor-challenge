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

type pointsResultItem struct {
	Points int `json:"points"`
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

// Returns 25 points if the amout is a multiple of quarter (0.25 cents)
func calculateQuarterPoints(amount string) int {
	if helpers.IsMultipleOfQuarter(amount) {
		return 25
	}
	return 0
}

// Calculates the points for each pair of items (5 points each) in the list
func calculatePointsForPairs(items []models.Item) int {
	return helpers.CountPairs(items) * 5
}

/* Calculates the points based on Description of each item in the list by multiplying
*  the price by 0.2 and rounding it up to nearest integer
 */
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

// Calculates 6 points when the day is odd
func calculateOddDayPoints(date string) int {
	if helpers.IsDayOdd(date) {
		return 6
	}
	return 0
}

// Calculates 10 points when the purchase time is between 2:00 pm and 4:00 pm
func calculatePointsForTimeRange(time string) int {
	if helpers.IsTimeInRange(time, "14:00", "16:00") {
		return 10
	}

	return 0
}

// Cumulatively calculates the points based on the rules defined
func calculatePoints(receipt models.Receipt) int {
	res := calculateRetailerPoints(receipt.Retailer)
	res += calculateRoundAmountPoints(receipt.Total)
	res += calculateQuarterPoints(receipt.Total)
	res += calculatePointsForPairs(receipt.Items)
	res += calculatePointsForItemDesc(receipt.Items)
	res += calculateOddDayPoints(receipt.PurchaseDate)
	res += calculatePointsForTimeRange(receipt.PurchaseTime)
	return res
}

// Controller for the receipt processing POST endpoint
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

	// ds := datastore.Points
	datastore.Points[u.String()] = points

	fmt.Println(datastore.Points)
	result := resultItem{
		ID: u.String(),
	}

	c.IndentedJSON(http.StatusOK, result)

}

// Controller for getting the points based on id
func GetPoints(c *gin.Context) {
	id := c.Param("id")

	// ds := datastore.Points
	value, ok := datastore.Points[id]
	if !ok {
		c.IndentedJSON(http.StatusNotFound, "No receipt found for that id")
		return
	}
	points := value

	fmt.Println(datastore.Points)
	result := pointsResultItem{
		Points: points,
	}

	c.IndentedJSON(http.StatusOK, result)

}
