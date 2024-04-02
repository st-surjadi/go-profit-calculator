package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	fmt.Printf("Earning Before Tax: %.2f\n", ebt)
	fmt.Printf("Earning After Tax (Profit): %.2f\n", profit)
	fmt.Printf("Ratio (EBT/Profit): %.2f", ratio)
}

func outputText() {

}
