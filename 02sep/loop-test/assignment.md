## Assignment: Compound Interest Accumulator
Calculate the total value of an investment over a set number of years with annually compounding interest, where a fixed annual contribution is added at the end of each year.

Problem Statement
Complete the calculateInvestment() function. It should return the total accumulated balance (as a float64) given an initial deposit, a fixed annual contribution, an annual interest rate, and the duration in years.

## Rules:

Year 1: Interest applies to the initial deposit, then the annual contribution is added at the end of the year.

Subsequent Years: Interest applies to the starting balance of that year (including previous contributions and accrued interest), then the fixed contribution is added at the end of the year.

Interest rate is passed as a percentage (e.g., 5.0 for 5%), so remember to convert it to a decimal multiplier inside the loop.

## Expected Breakdown (3-Year Example)
Start: $1000.00

End of Year 1: $1000.00 * 1.05 + $500.00 = $1550.00

End of Year 2: $1550.00 * 1.05 + $500.00 = $2127.50

End of Year 3: $2127.50 * 1.05 + $500.00 = $2733.88 + $97.75 = $2831.63