// use channel of capacity 1 to ensure that at most one goroutine
// access a shared variable at a time. A semaphore that counts only to
// 1 is called a binary semaphore.
package bank

var (
	sema    = make(chan struct{}, 1) // a binary semaphore gaurding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
