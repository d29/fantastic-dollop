package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fake := flag.Bool("fake", false, "use stdout as fake device")
	flag.Parse()

	exit := make(chan struct{})
	output := make(chan string, 100)

	if *fake {
		go fakeDeviceLoop(output, exit)
		// go debounce(time.Second, output, fake_device_draw)
	} else {
		go deviceLoop(output, exit)
	}

	c := NewClient(output)
	go c.Run()

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt)

	<-interrupt
	log.Println("interrupt")
	close(exit)
	c.Close()
}
