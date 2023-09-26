func updateUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	var user User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}

	c.BindJSON(&user)

	if err := db.Save(&user).Error; err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}