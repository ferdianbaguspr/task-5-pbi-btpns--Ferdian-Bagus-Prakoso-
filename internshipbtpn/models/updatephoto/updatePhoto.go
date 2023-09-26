func updatePhoto(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	var photo Photo
	if err := db.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(404, gin.H{"message": "Photo not found"})
		return
	}

	c.BindJSON(&photo)

	if err := db.Save(&photo).Error; err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	} else {
		c.JSON(200, photo)
	}
}