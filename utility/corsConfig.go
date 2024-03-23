package utility

//func CORS() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
//		c.Writer.Header().Add("Access-Control-Allow-Methods", "DELETE, OPTIONS, POST, GET, OPTIONS")
//		c.Writer.Header().Add("Access-Control-Allow-Headers", "*")
//
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//			return
//		}
//
//		c.Next()
//	}
//}
