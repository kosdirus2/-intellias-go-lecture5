package main

import "fmt"

func main() {
	var (
		applePrice          = 5.99
		pearPrice   float64 = 7
		totalBudget float64 = 23
	)

	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	priceFor9Apples := applePrice * 9
	priceFor8Pears := pearPrice * 8
	priceFor9ApplesAnd8Pears := priceFor9Apples + priceFor8Pears
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", priceFor9ApplesAnd8Pears, "грн.")

	// 2. Скільки груш ми можемо купити?
	maxAmountOfPearsToBuyWithCurrentBudget := int(totalBudget / pearPrice)
	fmt.Println("2. Скільки груш ми можемо купити?", maxAmountOfPearsToBuyWithCurrentBudget)

	// 3. Скільки яблук ми можемо купити?
	maxAmountOfApplesToBuyWithCurrentBudget := int(totalBudget / applePrice)
	fmt.Println("3. Скільки яблук ми можемо купити?", maxAmountOfApplesToBuyWithCurrentBudget)

	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	priceFor2Apples := applePrice * 2
	priceFor2Pears := pearPrice * 2
	priceFor2ApplesAnd2Pears := priceFor2Apples + priceFor2Pears
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", priceFor2ApplesAnd2Pears < totalBudget)
}
