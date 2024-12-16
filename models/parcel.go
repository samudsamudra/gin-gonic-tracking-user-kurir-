package models

import (
	"database/sql"
	"errors"
	"log"
	"sistem-tracking/config"
)

// Struktur data untuk tabel parcels
type Parcel struct {
	ID            int    `json:"id"`
	NomorResi     string `json:"nomor_resi"`
	NamaPengirim  string `json:"nama_pengirim"`
	NamaPenerima  string `json:"nama_penerima"`
	AlamatPenerima string `json:"alamat_penerima"`
	StatusTerakhir string `json:"status_terakhir"`
	CreatedAt     string `json:"created_at"`
}

// Struktur data untuk tabel tracking_status
type TrackingStatus struct {
	ID          int    `json:"id"`
	ParcelID    int    `json:"parcel_id"`
	Status      string `json:"status"`
	WaktuUpdate string `json:"waktu_update"`
}

// Fungsi untuk mengambil semua data pengiriman barang
func GetAllParcels() ([]Parcel, error) {
	rows, err := config.DB.Query("SELECT * FROM parcels")
	if err != nil {
		log.Println("Gagal mengambil data pengiriman:", err)
		return nil, err
	}
	defer rows.Close()

	var parcels []Parcel
	for rows.Next() {
		var parcel Parcel
		err := rows.Scan(&parcel.ID, &parcel.NomorResi, &parcel.NamaPengirim, &parcel.NamaPenerima, &parcel.AlamatPenerima, &parcel.StatusTerakhir, &parcel.CreatedAt)
		if err != nil {
			log.Println("Gagal membaca data pengiriman:", err)
			return nil, err
		}
		parcels = append(parcels, parcel)
	}
	return parcels, nil
}

// Fungsi untuk mengambil data pengiriman berdasarkan ID
func GetParcelByID(id int) (Parcel, error) {
	var parcel Parcel
	err := config.DB.QueryRow("SELECT * FROM parcels WHERE id = ?", id).Scan(
		&parcel.ID,
		&parcel.NomorResi,
		&parcel.NamaPengirim,
		&parcel.NamaPenerima,
		&parcel.AlamatPenerima,
		&parcel.StatusTerakhir,
		&parcel.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Data pengiriman tidak ditemukan:", err)
			return parcel, err
		}
		log.Println("Gagal mengambil data pengiriman:", err)
		return parcel, err
	}

	return parcel, nil
}

// Fungsi untuk menambahkan data pengiriman baru
func CreateParcel(parcel Parcel) error {
	_, err := config.DB.Exec("INSERT INTO parcels (nomor_resi, nama_pengirim, nama_penerima, alamat_penerima, status_terakhir) VALUES (?, ?, ?, ?, ?)",
		parcel.NomorResi, parcel.NamaPengirim, parcel.NamaPenerima, parcel.AlamatPenerima, parcel.StatusTerakhir)
	if err != nil {
		log.Println("Gagal menambahkan data pengiriman:", err)
		return err
	}
	return nil
}

// Fungsi untuk mengupdate data pengiriman
func UpdateParcel(id int, parcel Parcel) error {
	_, err := config.DB.Exec("UPDATE parcels SET nomor_resi = ?, nama_pengirim = ?, nama_penerima = ?, alamat_penerima = ?, status_terakhir = ? WHERE id = ?",
		parcel.NomorResi, parcel.NamaPengirim, parcel.NamaPenerima, parcel.AlamatPenerima, parcel.StatusTerakhir, id)
	if err != nil {
		log.Println("Gagal mengupdate data pengiriman:", err)
		return err
	}
	return nil
}

// Fungsi untuk menghapus data pengiriman
func DeleteParcel(id int) error {
	_, err := config.DB.Exec("DELETE FROM parcels WHERE id = ?", id)
	if err != nil {
		log.Println("Gagal menghapus data pengiriman:", err)
		return err
	}
	return nil
}

// Fungsi untuk menambahkan status baru untuk parcel
func AddTrackingStatus(parcelID int, status string) error {
	// Validasi keberadaan parcel
	exists, err := ParcelExists(parcelID)
	if err != nil {
		return err
	}
	if !exists {
		log.Println("Parcel tidak ditemukan")
		return errors.New("parcel tidak ditemukan")
	}

	// Tambahkan tracking status
	_, err = config.DB.Exec("INSERT INTO tracking_status (parcel_id, status) VALUES (?, ?)", parcelID, status)
	if err != nil {
		log.Println("Gagal menambahkan tracking status:", err)
		return err
	}

	// Perbarui status terakhir di tabel parcels
	_, err = config.DB.Exec("UPDATE parcels SET status_terakhir = ? WHERE id = ?", status, parcelID)
	if err != nil {
		log.Println("Gagal memperbarui status terakhir parcel:", err)
		return err
	}

	return nil
}

// Fungsi untuk mengambil semua riwayat status untuk parcel tertentu
func GetTrackingStatus(parcelID int) ([]TrackingStatus, error) {
	rows, err := config.DB.Query("SELECT * FROM tracking_status WHERE parcel_id = ? ORDER BY waktu_update ASC", parcelID)
	if err != nil {
		log.Println("Gagal mengambil tracking status:", err)
		return nil, err
	}
	defer rows.Close()

	var statuses []TrackingStatus
	for rows.Next() {
		var status TrackingStatus
		err := rows.Scan(&status.ID, &status.ParcelID, &status.Status, &status.WaktuUpdate)
		if err != nil {
			log.Println("Gagal membaca tracking status:", err)
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}

// Fungsi untuk menghapus semua tracking status berdasarkan Parcel ID
func DeleteTrackingStatusByParcelID(parcelID int) error {
	_, err := config.DB.Exec("DELETE FROM tracking_status WHERE parcel_id = ?", parcelID)
	if err != nil {
		log.Println("Gagal menghapus tracking status:", err)
		return err
	}
	return nil
}

// Fungsi untuk mengecek apakah parcel dengan ID tertentu ada
func ParcelExists(parcelID int) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM parcels WHERE id = ?)", parcelID).Scan(&exists)
	if err != nil {
		log.Println("Gagal mengecek keberadaan parcel:", err)
		return false, err
	}
	return exists, nil
}
