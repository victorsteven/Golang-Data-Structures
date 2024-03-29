package main

import "fmt"

func main() {

	currency := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"JPY": "Japan Yen",
		"CHF": "Switzerland Franc",
	}

	//a. Adding to the map:
	currency["USD"] = "USA Dollar"

	fmt.Println("Currency with USD added: ", currency)

	//b. Remove from the map:
	delete(currency, "GBP")
	fmt.Println("Currency with GBP deleted: ", currency)

	//c. Replacing one entry with another:
	currency["AUD"] = "New Zealand Dollar"
	fmt.Println("Currency with AUD value replaced with NZD: ", currency)

	//Ranging through the map:
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}

	if value, ok := currency["USD"]; ok { //comma ok idiom
		fmt.Printf("The value %s is present\n", value)
		fmt.Println(ok) //This will print "true"
	} else {
		fmt.Println("We could not find that key")
	}
}
