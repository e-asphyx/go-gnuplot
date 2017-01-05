package gnuplot

import (
	"os"
	"os/signal"
)

func WaitSigInt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
