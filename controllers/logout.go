package controllers

import (
	"net/http"

	"dinsos_kuburaya/config"
	"dinsos_kuburaya/models"

	"github.com/gin-gonic/gin"
)

// Struktur request logout
type LogoutRequest struct {
	TokenID string `json:"token_id" binding:"required"`
}

// Logout menghapus token dari tabel secret_tokens
func Logout(c *gin.Context) {
	var input LogoutRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Input tidak valid"})
		return
	}

	db := config.DB

	// Cari token berdasarkan token_id
	var secretToken models.SecretToken
	if err := db.Where("id = ?", input.TokenID).First(&secretToken).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Token tidak ditemukan"})
		return
	}

	// Hapus token dari database (logout)
	if err := db.Delete(&secretToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout berhasil",
	})
}
