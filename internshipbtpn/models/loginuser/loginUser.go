func loginUser(c *gin.Context) {
	email := c.DefaultQuery("email", "")
	password := c.DefaultQuery("password", "")

	var user User
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"message": "Login failed. Invalid email or password"})
	} else {
		c.JSON(200, user)
	}
}