func createPhoto(c *gin.Context) {
	var photo Photo
	c.BindJSON(&photo)

	if err := db.Create(&photo).Error; err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, photo)
	}
}