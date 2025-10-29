// simulation of a concurrent cake shop with numerous parameters
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Shop struct {
	Verbose        bool
	Cakes          int           // number of cakes to bake
	BakeTime       time.Duration // time to bake one cake
	BakeStdDev     time.Duration // standard deviation of baking time
	BakeBuf        int           // buffer slots between baking and icing
	NumIcers       int           // number of cooks doing icing
	IceTime        time.Duration // time to ice one cake
	IceStdDev      time.Duration // standard deviation of icing time
	IceBuf         int           // buffer slots between icing and inscribing
	InscribeTime   time.Duration // time to inscribe one cake
	InscribeStdDev time.Duration // standard deviation of inscribing time
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
	for i := range s.Cakes {
		c := cake(i)
		if s.Verbose {
			fmt.Println("baking", c)
		}
		work(s.BakeTime, s.BakeStdDev)
		baked <- c
	}
	close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
	for c := range baked {
		if s.Verbose {
			fmt.Println("icing", c)
		}
		work(s.IceTime, s.IceStdDev)
		iced <- c
	}
}

func (s *Shop) inscriber(iced <-chan cake) {
	for range s.Cakes {
		c := <-iced
		if s.Verbose {
			fmt.Println("inscribing", c)
		}
		work(s.InscribeTime, s.InscribeStdDev)
		if s.Verbose {
			fmt.Println("==== finished cake", c)
		}
	}
}

// work blocks the calling goroutine for a period of time that is
// normally distributed around d with a standard deviation of stddev.
func work(d, stddev time.Duration) {
	delay := d + time.Duration(rand.NormFloat64()*float64(stddev))
	time.Sleep(delay)
}

func main() {
	shop := Shop{
		Verbose:        true,
		Cakes:          5,
		BakeTime:       2 * time.Second,
		BakeStdDev:     500 * time.Millisecond,
		BakeBuf:        2,
		NumIcers:       2,
		IceTime:        1 * time.Second,
		IceStdDev:      300 * time.Millisecond,
		IceBuf:         2,
		InscribeTime:   500 * time.Millisecond,
		InscribeStdDev: 100 * time.Millisecond,
	}

	baked := make(chan cake, shop.BakeBuf)
	iced := make(chan cake, shop.IceBuf)
	go shop.baker(baked)
	for range shop.NumIcers {
		go shop.icer(iced, baked)
	}
	shop.inscriber(iced)
}
