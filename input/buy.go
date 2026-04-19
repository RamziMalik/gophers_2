package input

import (
	"fmt"
)

func InputSelectItem() (int, int){
	var id, qty int

	fmt.Print("\nPilih ID Barang yang mau dibeli: ")
	fmt.Scan(&id)
	fmt.Print("Jumlah yang mau dibeli: ")
	fmt.Scan(&qty)

	return id, qty
}

func InputPayment() (int){
	var paid int

	fmt.Print("Masukkan Uang Anda: ")
	fmt.Scan(&paid)

	return paid
}