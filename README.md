# Simple Monte Carlo Simulation on Stock Portfolio

A simple Monte Carlo Simulation (implemented in Go langauge) that projects the final investment amount of a Stock Portfolio. In this model, an investor starts with an initial investment amount, and invests annually a fixed amount of money.

By default, the simulation runs a million times. Within each simulation, the returns (in percent) generated for a stock portfolio follows a normal distribution with the mean and standard deviation specified by the variables `portfolio_returns` and `portfolio_sd` in the `montecarlo.go` file.

The final projected amount is the average of the final investment amount returned from the simulations.

## Usage

Set the following variables to that you wish to simulate:

```golang
# Initial amount the investment starts with
begin_amt = 200000

# Amount of money that is invested at the end of each year
annual_inv = 50000

# Number of years to invest
years = 10

# Average return for a stock portfolio, in percentage
portfolio_returns = 0.12

# Standard deviation of the average return for a stock portfolio
portfolio_sd = 0.25
```

In the terminal, execute:

    $ go run montecarlo.go


