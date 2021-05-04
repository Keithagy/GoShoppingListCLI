package main

import "fmt"

// [SL] Method for type details

func (itemdet details) getCatName() string {
	return catTracker[itemdet.category]
}

// [SL] List of items by category

func listbyCat(sl map[string]details) {
	if len(sl) == 0 {
		fmt.Println("Shopping List is currently empty. Returning to report menu...")
	} else {
		// Initialize slice of maps by category
		listbyCat := make([]map[string]details, len(catTracker))
		for index := range listbyCat {
			listbyCat[index] = make(map[string]details)
		}

		for itemname, itemdet := range sl {
			listbyCat[itemdet.category][itemname] = itemdet
		}

		for index, catlist := range listbyCat {
			fmt.Println("|| Category:", catTracker[index], "||")
			for itemname, itemdet := range catlist {
				fmt.Println("Item: ", itemname, "|| Quantity: ", itemdet.quantity, "|| Unit Cost: ", itemdet.unitcost)
			}

			if len(catlist) == 0 {
				fmt.Println("Category is currently empty.")
				fmt.Println(" ")

			} else {
				fmt.Println("Category ends here.")
				fmt.Println(" ")

			}

		}
	}
}

// [SL] Cost of items by category

func costbyCat(sl map[string]details) {
	if len(sl) == 0 {
		fmt.Println("Shopping List is currently empty. Returning to report menu...")
	} else {
		// Initialize slice of maps by category
		listbyCat := make([]map[string]details, len(catTracker))
		for index := range listbyCat {
			listbyCat[index] = make(map[string]details)
		}

		for itemname, itemdet := range sl {
			listbyCat[itemdet.category][itemname] = itemdet
		}

		for index, catlist := range listbyCat {
			if len(listbyCat) == 0 {
				fmt.Println(" ")
				fmt.Println("Category is currently empty. Next category...")
			}

			catsum := 0.00

			for _, itemdet := range catlist {
				catsum += (float64(itemdet.quantity) * itemdet.unitcost)
			}

			fmt.Printf("|| Index %d: %s -- Cost: $%.2f ||\n", index, catTracker[index], catsum)
		}
	}
}
