/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 09/26/2023
 */

package routes

import (
	"fetch-rewards-receipt-processor-challenge/middlewares"

	"github.com/gin-gonic/gin"
)

// Initalizer for routes
func InitRouter(engine *gin.Engine) {
	engine.Use(middlewares.CORSMiddleware())
	RegisterPublicRoutes(engine)
}
