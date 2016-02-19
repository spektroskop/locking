package main

type Dispatcher struct {
	Request chan map[string]bool
	Return  chan map[string]bool
	Save    chan bool
	data    map[string]bool
}

func (d *Dispatcher) Run() {
	d.Request <- make(map[string]bool)

	go func() {
		for {
			select {
			case <-d.Save:
			case data := <-d.Return:
				d.Request <- data
			}
		}
	}()
}

func main() {

}
