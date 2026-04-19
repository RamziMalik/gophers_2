package input

import (
	"fmt"
	"bufio"
	"os"
)

func InputItem() (string, int, int){
	var price, stock int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Nama Barang: ")
	scanner.Scan()
	fmt.Print("Masukkan Harga: ")
	fmt.Scan(&price)
	fmt.Print("Masukan Stok Awal: ")
	fmt.Scan(&stock)

	name := scanner.Text()


	return name, price, stock
}