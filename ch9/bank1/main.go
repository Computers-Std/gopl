// Package bank provides a concurrency-safe bank with one account
package bank

var (
	deposits = make(chan int) // send amount to deposit
	balances = make(chan int) // recieve balance
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances // read request
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// If someone is ready to receive a value from the 'balances'
			// channel, the 'teller' hands over the current value of the
			// local 'balance' variable
		}
		// Because, the operations happens inside the 'select' block,
		// the 'teller' can only do one thing at a time. It is either
		// updating the balance (via deposits) OR reporting the
		// balance (via balances). It can never do both simultaneously
	}
}

func init() {
	go teller() // start the monitor goroutine
}
