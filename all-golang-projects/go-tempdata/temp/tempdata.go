package temp

import "time"

func Tempdata(vars []*int, wait int) {
	done := make(chan bool) // true false

	go func() {
		defer close(done)
		time.Sleep(time.Duration(wait) * time.Second)
		for _, v := range vars {
			*v = 0 // sets vars to 0
		}
	}()

	<-done // wait for the goroutine to finish
}
