func getPhotos(c *gin.Context) {
	var photos []Photo
	if err := db.Find(&photos).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, photos)
	}
}