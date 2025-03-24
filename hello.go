package main

import "fmt"

func main() {
	var smsSendingLimit int = 0
	var costPerSMS float64 = 0.0
	var hasPermission bool = false
	var username string = "raj"

	fmt.Printf("%d %f %t %s\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
