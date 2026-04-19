package handler

import (
	"fmt"
	"gophers_2/model"
	"gophers_2/input"
	"gophers_2/data"
)

func AddItem() {
	name, price, stock := input.InputItem()

	item := model.Item{
		Name: name,
		Price: price,
		Stock: stock,
	}

	data.AddInventory(item)

	fmt.Println("[SISTEM]: Barang berhasil ditambahkan ke gudang!")
	fmt.Println()
	fmt.Println("-------------------------------------------")

	return
}