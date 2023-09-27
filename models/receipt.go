/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package models

// Defines the strucure of a Receipt Object
type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"required"`
	Total        string `json:"total" validate:"required"`
}
