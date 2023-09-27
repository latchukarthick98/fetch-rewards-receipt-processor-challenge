/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package models

// A single product in a receipt
type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"` // The short description provided for the item.
	Price            string `json:"price" validate:"required"`            // The price paid for this item.
}
