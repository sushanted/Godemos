package main

import (
	"fmt"
	"strconv"
)

func main() {
	for _, value := range []interface{}{31, "412", nil, true, false} {
		fmt.Printf("Value of %T(%v) is int(%v)\n", value, value, getValue(value))
	}
}

func getValue(any interface{}) int {

	switch any := any.(type) { // NOTE : the target type here is "type" and it is collected into any variable
	case nil:
		return 0
	case bool:
		if any {
			return 1
		}
		return 0
	case int: // for each case any is of that case type, no need to cast it for every case
		return any
	case string:
		if value, err := strconv.Atoi(any); err == nil {
			return value
		}
		return -1
	default:
		return -1
	}
}
