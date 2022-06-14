package main

import "fmt"

func main() {
	var (
		applePrice          = 5.99
		pearPrice   float64 = 7
		totalBudget float64 = 23
	)

	// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?", applePrice*9+pearPrice*8, "грн.")

	// 2. Скільки груш ми можемо купити?
	fmt.Println("2. Скільки груш ми можемо купити?", int(totalBudget/pearPrice))

	// 3. Скільки яблук ми можемо купити?
	fmt.Println("3. Скільки яблук ми можемо купити?", int(totalBudget/applePrice))

	// 4. Чи ми можемо купити 2 груші та 2 яблука?
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", applePrice*2+pearPrice*2 < totalBudget)
}
