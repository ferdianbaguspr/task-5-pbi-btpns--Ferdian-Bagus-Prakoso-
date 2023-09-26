func deleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	var user User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(400, gin.H{"message": "Failed to delete user"})
	} else {
		c.JSON(200, gin.H{"message": "User deleted"})
	}
}
