package middlewares

import (
	"mygram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}

// func SosmedAuthorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		db := database.GetDB()
// 		sosmedId, err := strconv.Atoi(c.Param("sosmedId"))
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
// 				"error":    "Bad Request",
// 				"messsage": "Invalid Parameter",
// 			})
// 			return
// 		}
// 		userData := c.MustGet("userData").(jwt.MapClaims)
// 		UserID := uint(userData["id"].(float64))
// 		SocialMed := models.SocialMed{}

// 		err = db.Select("user_id").First(&SocialMed, uint(sosmedId)).Error

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"error":   "Data Not Found",
// 				"message": "Data doest'n exist",
// 			})
// 			return
// 		}

// 		if SocialMed.ID != UserID {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"error":   "Unauthorized",
// 				"message": "You are not allowed to access this data",
// 			})
// 		}
// 		c.Next()
// 	}
// }
