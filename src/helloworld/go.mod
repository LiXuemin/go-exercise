module helloworld

go 1.16

require (
	github.com/myuser/calculator v0.0.0
	github.com/lixuemin/fizzbuzz v0.0.0
)
replace (
	github.com/myuser/calculator => ../calculator
	github.com/lixuemin/fizzbuzz => ../challenge
)