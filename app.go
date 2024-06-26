package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	outputText("Revenue: ")
	fmt.Scan(&revenue)
	outputText("Expenses: ")
	fmt.Scan(&expenses)
	outputText("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculator(revenue, expenses, taxRate)

	fmt.Printf("Earning Before Tax: %.2f\n", ebt)
	fmt.Printf("Earning After Tax (Profit): %.2f\n", profit)
	fmt.Printf("Ratio (EBT/Profit): %.2f", ratio)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	return ebt, profit, ratio
}
