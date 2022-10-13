package main

import "fmt"
import "time"

func main() {
	exec(5, "")
	time.Sleep(3 * time.Second)
}

func exec(times int, id string) {

	if len(id) > 0 {
		fmt.Println(id)
		id += "."
	}

	for i := 0; i < times; i++ {
		nextId := fmt.Sprintf("%s%d", id, i)
		go exec(times-1, nextId)
	}
}
