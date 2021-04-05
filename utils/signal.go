package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// NewSignal for creating a new signal
func NewSignal() (stopCh chan struct{}) {
	sig := make(chan struct{})

	return sig
}

// osSignalHandler for storing a signal
type osSignalHandler struct {
	signal chan os.Signal
}

// NewOSSignal for creating a new signal
func NewOSSignal() osSignalHandler {
	osSig := osSignalHandler{}

	osSig.signal = make(chan os.Signal, 2)
	signal.Notify(
		osSig.signal,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	return osSig
}

// Wait for waiting for an OS signal
func (s *osSignalHandler) Wait() {
	<-s.signal
}

func SetupSignalHandler() (stopCh <-chan struct{}) {

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(
		c,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
	)

	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()

	return stop
}
