package example

import (
	"fmt"
)

func ExampleHello() {
	greeting := Hello("Hello, World!")
	fmt.Println(greeting)

	// Output:
	// Hello, World!
}

func ExamplePage() {
	checkIns := map[string]bool{
		"John": true,
		"Jane": true,
		"Joe":  false,
		"Jill": false,
		"Jack": true,
		"Jen":  false,
	}
	Page(checkIns)

	// Unordered output:
	// Paging Joe; please see the front desk to check in.
	// Paging Jill; please see the front desk to check in.
	// Paging Jen; please see the front desk to check in.
}
