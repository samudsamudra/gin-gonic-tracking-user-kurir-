package controllers

import (
	"net/http"
	"sistem-tracking/models"
	"strconv" 

	"github.com/gin-gonic/gin"
)

// GetParcels - Mengambil semua data pengiriman
func GetParcels(c *gin.Context) {
	parcels, err := models.GetAllParcels()
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

	parcel, err := models.GetParcelByID(id)
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

	err := models.CreateParcel(parcel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan data pengiriman"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data pengiriman berhasil ditambahkan"})
}

// UpdateParcel - Mengupdate data pengiriman
func UpdateParcel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var parcel models.Parcel
	if err := c.ShouldBindJSON(&parcel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	err = models.UpdateParcel(id, parcel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data pengiriman"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data pengiriman berhasil diupdate"})
}

// DeleteParcel - Menghapus data pengiriman
func DeleteParcel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	// Hapus data parcel
	err = models.DeleteParcel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data pengiriman"})
		return
	}

	// Hapus tracking status yang terkait
	err = models.DeleteTrackingStatusByParcelID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus tracking status yang terkait"})
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

	var request struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	err = models.AddTrackingStatus(id, request.Status)
	if err != nil {
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

	statuses, err := models.GetTrackingStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil tracking status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": statuses})
}
