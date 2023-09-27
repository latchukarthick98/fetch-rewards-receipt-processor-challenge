/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/03/2022
 */

package routes

import (
	"fetch-rewards-receipt-processor-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine) {
	r.POST("/receipts/process", controllers.ProcessReceipt)
}
