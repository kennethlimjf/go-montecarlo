package main

import (
	"fmt"
	"math/rand"
)

func main() {
	begin_amt := 200000.0
	annual_inv := 50000.0
	years := 10

	simulations := 1000000

	var base_ip [2]float64
	base_ip[0] = 0.05 // Average annual returns
	base_ip[1] = 0.12 // Standard deviation

	chns := make([]chan float64, simulations)
	for i := 0; i < simulations; i++ {
		chns[i] = make(chan float64)
		go calc_total_years_amt(begin_amt, base_ip, annual_inv, years, chns[i])
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
	fmt.Printf("This Monte Carlo Simulation says you can save $%10.2f in %d years if you start with $%.2f and invest $%.2f annually.\n", sum/float64(simulations), years, begin_amt, annual_inv)
}

func calc_total_years_amt(begin_amt float64, base_ip [2]float64, annual_inv float64, years int, c chan float64) float64 {
	sum := begin_amt
	for i := 0; i < years; i++ {
		ret := generate_rand_ret(base_ip[0], base_ip[1])
		sum += calc_annual_amt(begin_amt, ret, annual_inv)
	}
	c <- sum
	close(c)
	return sum
}

func generate_rand_ret(mean, std_dev float64) float64 {
	return rand.NormFloat64()*std_dev + mean
}

func calc_annual_amt(amt, ret, invest float64) float64 {
	return (amt*(1+ret) + invest)
}
