package main

import (
	"fmt"
	"math"
)

type Payment struct {
	Month            int
	MonthlyPayment   float64
	Interest         float64
	PrincipalPayment float64
	RemainingBalance float64
}

func loanSchedule(amount, annualRate float64, years int, extraPayment, deposit float64) []Payment {
	principal := amount - deposit
	monthlyRate := annualRate / 12 / 100
	months := years * 12

	monthlyPayment := principal * (monthlyRate * math.Pow(1+monthlyRate, float64(months))) / (math.Pow(1+monthlyRate, float64(months)) - 1)
	balance := principal
	schedule := []Payment{}

	for m := 1; m <= months; m++ {
		interest := balance * monthlyRate
		principalPayment := monthlyPayment - interest

		if m == 1 && extraPayment > 0 {
			principalPayment += extraPayment
		}

		balance -= principalPayment
		if balance < 0 {
			balance = 0
		}

		schedule = append(schedule, Payment{
			Month:            m,
			MonthlyPayment:   math.Round(monthlyPayment),
			Interest:         math.Round(interest),
			PrincipalPayment: math.Round(principalPayment),
			RemainingBalance: math.Round(balance),
		})
	
	if balance<=0{
		break
	}
}
   return schedule
}

func main(){
	amount:=100_000_000.0
	annualRate := 12.0
	extraPayment :=10_000_000.0

	for _, years := range []int{1,2,3,4,5} {
		fmt.Printf("\nLoan term: %d years(s)\n", years)
		schedule := loanSchedule(amount, annualRate,years,extraPayment,0)
		limit:=12
		if len(schedule) < limit{
			limit = len(schedule)
		}
		for _, row := range schedule[:limit]{
			fmt.Printf("Month: %d, Monthly Payment: %.0f, Interest: %.0f, Principal Payment: %.0f, Remaining Balance: %.0f\n",
				row.Month, row.MonthlyPayment, row.Interest, row.PrincipalPayment, row.RemainingBalance)
		}
	}
}

