package data

import (
	"gophers_2/model"
)

var inventory []model.Item
var id int

func GetInventory() ([]model.Item){
	return inventory
}

func AddInventory(item model.Item) {
	item.ID = AutoIncrement()
	inventory = append(inventory, item)
}

func AutoIncrement() (int){
	id++
	return id
}

func BuyItemByID(id int, qty int) (bool, int, int, string){
	for i, item := range inventory {
		if item.ID == id {
			if item.Stock < qty {
				return false, 0, 0, ""
			}

			inventory[i].Stock -= qty
			return true, item.Price, item.Stock, item.Name
		}
	}
	return false, 0, 0, ""
}