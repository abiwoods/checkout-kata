package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var db dataSource

func main() {
	exitApp := false
	reader := bufio.NewReader(os.Stdin)
	db = mockDB{}
	checkout := GetCheckout()

	for !exitApp {
		fmt.Println("Input action - scan, checkout, exit:")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}

		standardisedInput := standardiseString(input)
		fmt.Println(standardisedInput)

		switch standardisedInput {
		case "SCAN":
			handleScan(reader, checkout)
			break
		case "CHECKOUT":
			handleCheckout(checkout)
			exitApp = true
			break
		case "EXIT":
			fmt.Println("Exiting checkout app")
			exitApp = true
			break
		default:
			fmt.Println("Input not recognised, please input valid option")
			break
		}
	}
}

func handleScan(reader *bufio.Reader, checkout Checkout) {
	fmt.Println("Please select product to scan:")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	standardisedInput := standardiseString(input)

	err = checkout.scan(standardisedInput)
	if err != nil {
		fmt.Println(err)
		return
	} 

	fmt.Println(fmt.Sprintf("Succesfully scanned product %s", standardisedInput))
}

func handleCheckout(checkout Checkout) {
	price := checkout.getTotalPrice()

	fmt.Println(fmt.Sprintf("Your total price is %d", price))
}

func standardiseString(oldString string) string {
	upperCase := strings.ToUpper(oldString)
	spacesTrimmed := strings.ReplaceAll(upperCase, " ", "")
	return strings.TrimSuffix(spacesTrimmed, "\n")
}
