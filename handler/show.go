package handler

import (
	"fmt"
	"gophers_2/data"
)

func ShowInventory() {
	fmt.Println("=== DAFTAR STOK GUDANG ===")

	for _, item := range data.GetInventory() {
		fmt.Printf("ID: %-v | Nama: %-15v | Stok: %v | Harga: Rp %v\n", 
		item.ID, item.Name, item.Stock, item.Price)
	}

	fmt.Println("===========================")
	fmt.Println("Total Jenis Barang: ", len(data.GetInventory()))
	fmt.Println("\n--------------------------------------")
}