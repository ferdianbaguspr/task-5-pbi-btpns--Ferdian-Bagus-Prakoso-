func createUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if err := db.Create(&user).Error; err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}