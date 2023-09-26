func deletePhoto(c *gin.Context) {
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	var photo Photo
	if err := db.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(404, gin.H{"message": "Photo not found"})
		return
	}

	if err := db.Delete(&photo).Error; err != nil {
		c.JSON(400, gin.H{"message": "Failed to delete photo"})
	} else {
		c.JSON(200, gin.H{"message": "Photo deleted"})
	}
}