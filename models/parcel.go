package models

import (
	"database/sql"
	"log"
	"sistem-tracking/config"
)

type Parcel struct {
	ID            int    `json:"id"`
	NomorResi     string `json:"nomor_resi"`
	NamaPengirim  string `json:"nama_pengirim"`
	NamaPenerima  string `json:"nama_penerima"`
	AlamatPenerima string `json:"alamat_penerima"`
	StatusTerakhir string `json:"status_terakhir"`
	CreatedAt     string `json:"created_at"`
}

// GetAllParcels - Mengambil semua data pengiriman barang
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

// GetParcelByID - Mengambil data pengiriman berdasarkan ID
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


// CreateParcel - Menambahkan data pengiriman baru
func CreateParcel(parcel Parcel) error {
	_, err := config.DB.Exec("INSERT INTO parcels (nomor_resi, nama_pengirim, nama_penerima, alamat_penerima, status_terakhir) VALUES (?, ?, ?, ?, ?)",
		parcel.NomorResi, parcel.NamaPengirim, parcel.NamaPenerima, parcel.AlamatPenerima, parcel.StatusTerakhir)
	if err != nil {
		log.Println("Gagal menambahkan data pengiriman:", err)
		return err
	}
	return nil
}

// UpdateParcel - Mengupdate data pengiriman
func UpdateParcel(id int, parcel Parcel) error {
	_, err := config.DB.Exec("UPDATE parcels SET nomor_resi = ?, nama_pengirim = ?, nama_penerima = ?, alamat_penerima = ?, status_terakhir = ? WHERE id = ?",
		parcel.NomorResi, parcel.NamaPengirim, parcel.NamaPenerima, parcel.AlamatPenerima, parcel.StatusTerakhir, id)
	if err != nil {
		log.Println("Gagal mengupdate data pengiriman:", err)
		return err
	}
	return nil
}

// DeleteParcel - Menghapus data pengiriman
func DeleteParcel(id int) error {
	_, err := config.DB.Exec("DELETE FROM parcels WHERE id = ?", id)
	if err != nil {
		log.Println("Gagal menghapus data pengiriman:", err)
		return err
	}
	return nil
}
