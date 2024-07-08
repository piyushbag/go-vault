// This is awesome docs
package example

import "fmt"

type Demo struct{}

func (d *Demo) Hello() string {
	return "Hello, World!"
}

func Hello(name string) string {
	return "Hello, World!"
}

func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.\n", name)
		}
	}
}
