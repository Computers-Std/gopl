// Exercise 9.1: Add a function Withdraw(amount int) bool to the
// gopl.io/ch9/bank1 program. The result should indicate whether the
// transaction succeeded or failed due to insufficient funds.
//
// The message sent to the monitor goroutine must contain both the
// amount to withdraw and a new channel over which the monitor
// goroutine can send the boolean result back to Withdraw.

package bank

type transaction struct {
	amount int       // amount to withdraw
	status chan bool // bi-directional chan to access boolean result
}

var (
	deposits  = make(chan int) // send amount to deposit
	balances  = make(chan int) // recieve balance
	withdraws = make(chan transaction)
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances // read request
}

func Withdraw(amount int) bool {
	statusChan := make(chan bool)
	withdraws <- transaction{
		amount: amount,
		status: statusChan,
	}
	return <-statusChan
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case ta := <-withdraws:
			if balance >= ta.amount {
				balance -= ta.amount // make transaction
				ta.status <- true
			} else {
				ta.status <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
