package main

import (
	"fmt"
	"gophers_2/input"
	"gophers_2/handler"
)

func main() {
	for {
		option := input.SelectMenu()

		switch option {
		case 1:
			handler.AddItem()
		case 2:
			handler.ShowInventory()
		case 3:
			handler.BuyItem()
		case 4:
			fmt.Println("[SISTEM]: Program selesai. Selamat beristirahat, Juragan!")
			return
		default:
			fmt.Println("[SISTEM]: 404 Not found!")
		}
	}
}