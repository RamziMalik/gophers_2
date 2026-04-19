package calculator

func CalculateTotalPrice(qty int, price int) (int){
	var totalPrice = qty * price
	return totalPrice
}

func CalculateChange(price int, paid int) (int){
	var change int = paid - price
	return change
}