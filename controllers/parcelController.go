package controllers

import (
	"net/http"
	"strconv"
	"sistem-tracking/config"
	"sistem-tracking/models"

	"github.com/gin-gonic/gin"
)

// GetParcels - Mengambil semua data pengiriman
func GetParcels(c *gin.Context) {
	parcels, err := models.GetAllParcels(config.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pengiriman"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": parcels})
}

// GetParcelByID - Mengambil data pengiriman berdasarkan ID
func GetParcelByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	parcel, err := models.GetParcelByID(config.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data pengiriman tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": parcel})
}

// CreateParcel - Menambahkan data pengiriman baru
func CreateParcel(c *gin.Context) {
	var parcel models.Parcel
	if err := c.ShouldBindJSON(&parcel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if err := models.CreateParcel(config.DB, &parcel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan data pengiriman"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data pengiriman berhasil ditambahkan", "data": parcel})
}

// UpdateParcel - Mengupdate data pengiriman
func UpdateParcel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var updatedParcel models.Parcel
	if err := c.ShouldBindJSON(&updatedParcel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if err := models.UpdateParcel(config.DB, uint(id), updatedParcel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data pengiriman"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pengiriman berhasil diupdate"})
}

// DeleteParcel - Menghapus data pengiriman dan tracking status terkait
func DeleteParcel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Hapus parcel dan tracking status terkait
	if err := models.DeleteParcel(config.DB, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data pengiriman"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pengiriman dan tracking status berhasil dihapus"})
}

// AddTrackingStatus - Menambahkan tracking status baru untuk parcel
func AddTrackingStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	if err := models.AddTrackingStatus(config.DB, uint(id), input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan tracking status"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tracking status berhasil ditambahkan"})
}

// GetTrackingStatus - Mengambil semua riwayat tracking status untuk parcel tertentu
func GetTrackingStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	statuses, err := models.GetTrackingStatus(config.DB, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil tracking status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statuses})
}
