package input

import "fmt"

func SelectMenu() (int){
	var option int

	fmt.Println("\n=== TOKO KELONTONG INVENTORY MANAGER ===")
	fmt.Println("")
	fmt.Println("1. Tambah Barang ke Gudang")
	fmt.Println("2. Lihat Semua Stok Barang")
	fmt.Println("3. Beli Barang")
	fmt.Println("4. Keluar")
	fmt.Println("")
	fmt.Print("Pilih Menu (1-4): ")
	fmt.Scan(&option)
	fmt.Println()


	return option
}