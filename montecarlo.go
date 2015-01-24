package main

import (
	"fmt"
	"math/rand"
)

func main() {
	begin_amt := 200000.0
	annual_inv := 50000.0
	years := 10

	portfolio_returns := 0.08
	portfolio_sd := 0.04

	simulations := 1000000

	chns := make([]chan float64, simulations)
	for i := 0; i < simulations; i++ {
		chns[i] = make(chan float64)
		go calc_total_years_amt(begin_amt, portfolio_returns, portfolio_sd, annual_inv, years, chns[i])
	}

	projections := make([]float64, simulations)
	for i, c := range chns {
		final_amt := <-c
		projections = append(projections, final_amt)
		fmt.Printf("%-18s[%d]: $%10.2f\n", "Projected Savings", i, final_amt)
	}

	sum := 0.0
	for _, v := range projections {
		sum += v
	}
	fmt.Printf("You can save $%10.2f in %d years if you start with $%.2f and invest $%.2f annually.\n", sum/float64(simulations), years, begin_amt, annual_inv)
}

func calc_total_years_amt(begin_amt float64, p_ret float64, p_sd float64, annual_inv float64, years int, c chan float64) float64 {
	sum := begin_amt
	for i := 0; i < years; i++ {
		ret := generate_rand_ret(p_ret, p_sd)
		sum += calc_annual_returns(sum, ret) + annual_inv
	}
	c <- sum
	close(c)
	return sum
}

func generate_rand_ret(mean, std_dev float64) float64 {
	return rand.NormFloat64()*std_dev + mean
}

func calc_annual_returns(amt, ret float64) float64 {
	return amt * ret
}
