package models

import (
	"errors"
	"gorm.io/gorm"
)

// Parcel - Struktur data untuk tabel parcels
type Parcel struct {
	gorm.Model              // Menambahkan kolom ID, CreatedAt, UpdatedAt, DeletedAt otomatis
	NomorResi      string   `gorm:"unique;not null" json:"nomor_resi"`
	NamaPengirim   string   `json:"nama_pengirim"`
	NamaPenerima   string   `json:"nama_penerima"`
	AlamatPenerima string   `json:"alamat_penerima"`
	StatusTerakhir string   `json:"status_terakhir"`
	TrackingStatus []TrackingStatus `gorm:"foreignKey:ParcelID" json:"tracking_status"` // Relasi ke tracking_status
}

// TrackingStatus - Struktur data untuk tabel tracking_status
type TrackingStatus struct {
	gorm.Model
	ParcelID    uint   `json:"parcel_id"`
	Status      string `json:"status"`
}

// CreateParcel - Fungsi untuk menambahkan parcel baru
func CreateParcel(db *gorm.DB, parcel *Parcel) error {
	if err := db.Create(parcel).Error; err != nil {
		return err
	}
	return nil
}

// GetAllParcels - Fungsi untuk mengambil semua data parcel
func GetAllParcels(db *gorm.DB) ([]Parcel, error) {
	var parcels []Parcel
	if err := db.Preload("TrackingStatus").Find(&parcels).Error; err != nil {
		return nil, err
	}
	return parcels, nil
}

// GetParcelByID - Fungsi untuk mengambil parcel berdasarkan ID
func GetParcelByID(db *gorm.DB, id uint) (Parcel, error) {
	var parcel Parcel
	if err := db.Preload("TrackingStatus").First(&parcel, id).Error; err != nil {
		return parcel, err
	}
	return parcel, nil
}

// UpdateParcel - Fungsi untuk memperbarui data parcel
func UpdateParcel(db *gorm.DB, id uint, updatedData Parcel) error {
	var parcel Parcel
	if err := db.First(&parcel, id).Error; err != nil {
		return err
	}
	if err := db.Model(&parcel).Updates(updatedData).Error; err != nil {
		return err
	}
	return nil
}

// DeleteParcel - Fungsi untuk menghapus parcel berdasarkan ID
func DeleteParcel(db *gorm.DB, id uint) error {
	if err := db.Where("id = ?", id).Delete(&Parcel{}).Error; err != nil {
		return err
	}
	return nil
}

// AddTrackingStatus - Fungsi untuk menambahkan tracking status ke parcel
func AddTrackingStatus(db *gorm.DB, parcelID uint, status string) error {
	var parcel Parcel
	if err := db.First(&parcel, parcelID).Error; err != nil {
		return errors.New("parcel tidak ditemukan")
	}

	trackingStatus := TrackingStatus{
		ParcelID: parcelID,
		Status:   status,
	}

	if err := db.Create(&trackingStatus).Error; err != nil {
		return err
	}

	// Perbarui status terakhir parcel
	parcel.StatusTerakhir = status
	if err := db.Save(&parcel).Error; err != nil {
		return err
	}

	return nil
}

// GetTrackingStatus - Fungsi untuk mengambil riwayat tracking status
func GetTrackingStatus(db *gorm.DB, parcelID uint) ([]TrackingStatus, error) {
	var statuses []TrackingStatus
	if err := db.Where("parcel_id = ?", parcelID).Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

// DeleteTrackingStatusByParcelID - Fungsi untuk menghapus semua tracking status berdasarkan Parcel ID
func DeleteTrackingStatusByParcelID(db *gorm.DB, parcelID uint) error {
	if err := db.Where("parcel_id = ?", parcelID).Delete(&TrackingStatus{}).Error; err != nil {
		return err
	}
	return nil
}
