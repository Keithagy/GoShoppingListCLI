package main

import "fmt"

// [CAT] Prints existing categories and indices
func showCatTracker(catTracker []string) {
	fmt.Println("Existing Categories: ")
	for index, value := range catTracker {
		fmt.Println(index, "--", value)
	}

	if len(catTracker) == 0 {
		fmt.Println("None currently, please add a category first")
	}
	fmt.Println(" ")
}

// [CAT] Takes user-input category name and returns corresponding index to be saved in struct
func getCatIndex(catTracker []string, category string) int {
	for index, value := range catTracker {
		if category == value {
			return index
		}
	}
	fmt.Println("User-input category does not match any existing category.")
	return -1
}

// [CAT] Add Categories

func addCat(catTracker []string) []string {
	var category string

	for category == "" {
		fmt.Println("What is the category to be added?")
		fmt.Scanln(&category)

		if category == "" {
			fmt.Println("No input detected, please try again.")
		}
	}

	existingCheck := getCatIndex(catTracker, category)
	if existingCheck == -1 {
		catTracker = append(catTracker, category)
		fmt.Println("Category", category, "successfully added.")
	} else {
		fmt.Println("Category already exists as Index ", existingCheck)
	}

	showCatTracker(catTracker)

	return catTracker
}

// [CAT] Modify Categories

func modCat(catTracker []string) []string {
	showCatTracker(catTracker)
	inputindex := -100
	for inputindex < 0 || inputindex > (len(catTracker)-1) {
		fmt.Println("Input the INDEX of the category to be edited: ")
		fmt.Scanln(&inputindex)

		if inputindex < 0 || inputindex > (len(catTracker)-1) {
			fmt.Println("Invalid input, please try again.")
		}
	}

	fmt.Println("Category Selected:", inputindex, "--", catTracker[inputindex])
	oldname := catTracker[inputindex]

	var catname string
	for {
		fmt.Println("What will be the new name of the category?")
		fmt.Scanln(&catname)

		if catname == "" {
			fmt.Println("No input detected, please try again.")
		} else if getCatIndex(catTracker, catname) != -1 {
			fmt.Println("Category name already exists, please try again.")
		} else {
			break
		}
	}

	catTracker[inputindex] = catname
	fmt.Println("Category at index", inputindex, "renamed from", oldname, "to", catname)

	return catTracker
}

// [CAT] Delete Categories

func delCat(sl map[string]details, catTracker []string) (map[string]details, []string) {
	inputindex := -100
	for inputindex < 0 || inputindex > (len(catTracker)-1) {
		fmt.Println("Input the INDEX of the category to be deleted: ")
		fmt.Scanln(&inputindex)

		if inputindex < 0 || inputindex > (len(catTracker)-1) {
			fmt.Println("Invalid input, please try again.")
		}
	}

	fmt.Println("Category Selected:", inputindex, "--", catTracker[inputindex])

	// Delete items of that from sl

	fmt.Println("The following items belong to the category", catTracker[inputindex], "and will also be deleted:")
	for name, detail := range sl {
		if detail.category == inputindex {
			delete(sl, name)
			fmt.Printf(" %s |", name)
		} else if detail.category > inputindex {
			tmp := sl[name]
			sl[name] = details{category: tmp.category - 1, quantity: tmp.quantity, unitcost: tmp.unitcost}
		}
	}

	// Delete category from catTracker
	copy(catTracker[inputindex:], catTracker[inputindex+1:])
	catTracker[len(catTracker)-1] = ""
	catTracker = catTracker[:len(catTracker)-1]

	fmt.Println("Category has been deleted.")

	return sl, catTracker
}
