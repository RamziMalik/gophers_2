package handler

import (
	"fmt"
	"gophers_2/data"
	"gophers_2/calculator"
	"gophers_2/input"
)

func BuyItem() {
	fmt.Println("=== MENU PEMBELIAN ===")

	for _, item := range data.GetInventory() {
		fmt.Printf("ID: %v | Nama: %-15v | Stok: %v | Harga: Rp %v\n", 
		item.ID, item.Name, item.Stock, item.Price)
	}

	id, qty := input.InputSelectItem()
	success, price, stock, name := data.BuyItemByID(id, qty)
	totalPrice := calculator.CalculateTotalPrice(qty, price)

	fmt.Println("\n[SISTEM]: Total Harga: Rp", totalPrice)
	paid := input.InputPayment()
	change := calculator.CalculateChange(totalPrice, paid)

	if success && change >= 0 {
		fmt.Println("\n[SISTEM]: Transaksi Berhasil!")
		fmt.Println("[SISTEM]: Kembalian Anda: Rp", change)
		fmt.Printf("[SISTEM]: Stok %v sekarang: %v\n", name, stock - qty)

	} else {
		fmt.Println("\n[SISTEM]: Transaksi dibatalkan")
	}

	fmt.Println("\n---------------------------------------")
}