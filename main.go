package main

import (
	"fmt"
)

// Struct for shopping list items
type details struct {
	category int
	quantity int
	unitcost float64
}

// Global Variables
var (
	sl         map[string]details // Shopping List
	catTracker []string           // Categories tracker

	slsaver         []map[string]details // Saves multiple shopping lists
	catTrackersaver [][]string           // Saves catTrackers corresponding to respective shopping lists
	accessedsave    int                  // Tracks the save index being accessed by the app

	killswitch bool // Makes sure both loader and app terminate togther
)

func init() {
	slsaver = make([]map[string]details, 0)
	catTrackersaver = make([][]string, 0)
	killswitch = false
}

func main() {
	initloader()
}

// [MAIN] Shopping List Application
func slapp() {
application:
	for {
		if killswitch == true {
			break
		}

		//  Main Menu
		fmt.Println(" ")
		fmt.Println("-------------------------")
		fmt.Println("Shopping List Application")
		fmt.Println("Save Index:", accessedsave)
		fmt.Println("-------------------------")
		fmt.Println("1. View entire shopping list")
		fmt.Println("2. Generate shopping list report")
		fmt.Println("3. Add items")
		fmt.Println("4. Modify items")
		fmt.Println("5. Delete items")
		fmt.Println("6. View or add/modify/delete categories")
		fmt.Println("7. Print current data fields")
		fmt.Println("8. Save and exit to Shopping List Loader")
		fmt.Println("9. Terminate program (DOES NOT SAVE CHANGES)")
		fmt.Println(" ")

		// Input user choice
		var menuchoice int
		fmt.Printf("Select your choice (INPUT NUMBERS ONLY): ")
		fmt.Scanln(&menuchoice)

		// Exiting the application
		if menuchoice == 9 {
			fmt.Println("Terminating Program...")
			killswitch = true
			break
		}

		// Switch statement for user control
		switch menuchoice {
		case 1:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 1. View entire shopping list")
			viewSL(sl)
		case 2:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 2. Generate shopping list report")
			genSLrep(sl)
		case 3:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 3. Add items")
			sl = addItems(sl)
		case 4:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 4. Modify existing items")
			sl = modifyItems(sl)
		case 5:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 5. Delete items")
			sl = delItems(sl)
		case 6:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 6. View or add/modify/delete categories")
			sl, catTracker = catEdit(sl, catTracker)
		case 7:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 7. Print current data fields")
			printdata(sl)
		case 8:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 8. Save and exit to Shopping List Loader")
			saveExit(sl, catTracker)
			break application
		default:
			fmt.Println("Invalid option")
		}
	}
}

// [MAIN] 1. View entire shopping list
func viewSL(sl map[string]details) {
	fmt.Println("Shopping List Contents:")
	fmt.Println(" ")

	for itemname, itemdet := range sl {
		fmt.Println("Category: ", itemdet.getCatName(), "|| Item: ", itemname, "|| Quantity: ", itemdet.quantity, "|| Unit Cost: ", itemdet.unitcost)
	}

	if len(sl) == 0 {
		fmt.Println("Shopping List is currently empty.")
	} else {
		fmt.Println("Shopping List ends here.")
	}
}

// [MAIN] 2. Generate shopping list report

func genSLrep(sl map[string]details) {
	for {
		fmt.Println(" ")
		fmt.Println("-------------------------")
		fmt.Println("Shopping list report")
		fmt.Println("-------------------------")
		fmt.Println("1. Total cost of each category")
		fmt.Println("2. List of item by category")
		fmt.Println("3. Main Menu")
		fmt.Println(" ")

		var reportchoice int
		fmt.Println("Select your choice (INPUT NUMBERS ONLY): ")
		fmt.Scanln(&reportchoice)

		if reportchoice == 3 {
			fmt.Println("Exiting...")
			break
		}

		switch reportchoice {
		case 1:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 1. Total cost of each category")
			costbyCat(sl)
		case 2:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 2. List of item by category")
			listbyCat(sl)
		default:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Invalid option")
		}
	}
}

// [MAIN] 3. Add items

func addItems(sl map[string]details) map[string]details {
	var userinputItemname string
	var userinputcategory int
	var userinputItemdet details

	for {
		fmt.Println("What is the name of your item? (CASE SENSITIVE)")
		fmt.Scanln(&userinputItemname)

		if userinputItemname == "" || userinputItemname == " " {
			fmt.Scanln("Item name cannot be blank, please try again.")
		} else {
			break
		}
	}

	_, presentcheck := sl[userinputItemname]
	if presentcheck == true {
		fmt.Println("This item already exists, please edit the existing entry using Option 4 in the main menu.")
		fmt.Println("Exiting... ")
		return sl
	}

	for {
		fmt.Println("What category does it belong to? Enter the corresponding INDEX NUMBER.")
		fmt.Println("Enter -9 to add a new category.")
		showCatTracker(catTracker)
		fmt.Scanln(&userinputcategory)

		if userinputcategory < len(catTracker) && userinputcategory >= 0 {
			fmt.Println("Valid category input, Category", catTracker[userinputcategory])
			userinputItemdet.category = userinputcategory
			break
		} else if userinputcategory == -9 {
			catTracker = addCat(catTracker)
			userinputItemdet.category = len(catTracker) - 1
			fmt.Println(userinputItemname, "will be added under new category,", catTracker[userinputItemdet.category], "(Index", userinputItemdet.category, ")")
			break
		} else {
			fmt.Println("Invalid entry, please check existing categories and retry.")
		}
	}

	for {
		fmt.Println("How many units are there? (Input a POSITIVE INTEGER)")
		fmt.Scanln(&userinputItemdet.quantity)

		if userinputItemdet.quantity <= 0 {
			fmt.Println("Item cannot have zero / negative quantity. Please retry.")
		} else {
			break
		}

	}

	for {
		fmt.Println("How much does it cost per unit? (Input a POSITIVE NUMBER UP TO 2 DECIMAL PLACES)")
		fmt.Scanln(&userinputItemdet.unitcost)

		if userinputItemdet.unitcost <= 0 {
			fmt.Println("Item cannot have zero / negative unit cost. Please retry.")
		} else {
			break
		}

	}

	sl[userinputItemname] = userinputItemdet
	return sl
}

// [MAIN] 4. Modify Existing Items

func modifyItems(sl map[string]details) map[string]details {

	var usermodname string

	for {
		viewSL(sl)
		fmt.Println(" ")
		fmt.Println("Enter the NAME of the item (e.g. Bread, Coke, Water, etc) you wish to modify. (CASE SENSITIVE)")
		fmt.Scanln(&usermodname)

		_, presentcheck := sl[usermodname]

		if presentcheck == false {
			fmt.Println("Item not found, please check and retry.")
		} else {
			break
		}
	}

	usermoddet := sl[usermodname]
	originalname := usermodname
	fmt.Println("Enter new name. Enter for no change.")
	fmt.Scanln(&usermodname)

	for {
		showCatTracker(catTracker)
		fmt.Println("Enter INDEX of new category. Enter for no change")
		fmt.Println("Enter -9 to save to new category.")
		fmt.Scanln(&usermoddet.category)

		if usermoddet.category < len(catTracker) && usermoddet.category >= 0 {
			// Report modified category of the item
			fmt.Println("Category selected:", catTracker[usermoddet.category])
			break
		} else if usermoddet.category == -9 {
			// Create new category
			catTracker = addCat(catTracker)
			usermoddet.category = len(catTracker) - 1
			fmt.Println("Item will now be stored under newly created category,", catTracker[usermoddet.category], "(Index", usermoddet.category, ")")
			break
		} else {
			// Invalid input, exit without changes
			fmt.Println("Invalid entry, please check existing categories and retry.")
		}
	}

	for {
		fmt.Println("Enter new quantity. Enter for no change.")
		fmt.Scanln(&usermoddet.quantity)

		if usermoddet.quantity <= 0 {
			fmt.Println("Item cannot have zero / negative quantity. Please retry.")
		} else {
			break
		}

	}

	for {
		fmt.Println("Enter new unit cost. Enter for no change.")
		fmt.Scanln(&usermoddet.unitcost)

		if usermoddet.unitcost <= 0 {
			fmt.Println("Item cannot have zero / negative unit cost. Please retry.")
		} else {
			break
		}

	}

	fmt.Println(" ")
	fmt.Println("The item:")
	fmt.Println("Category: ", sl[originalname].getCatName(), "|| Item: ", originalname, "|| Quantity: ", sl[originalname].quantity, "|| Unit Cost: ", sl[originalname].unitcost)

	delete(sl, originalname)
	sl[usermodname] = usermoddet

	fmt.Println(" ")
	fmt.Println("Has been modified to:")
	fmt.Println("Category: ", sl[usermodname].getCatName(), "|| Item: ", usermodname, "|| Quantity: ", sl[usermodname].quantity, "|| Unit Cost: ", sl[usermodname].unitcost)

	return sl

}

// [Main] 5. Delete items

func delItems(sl map[string]details) map[string]details {

	if len(sl) == 0 {
		fmt.Println("Shopping list is currently empty, nothing to delete.")
		fmt.Println("Exiting...")
		return sl
	}

	var usermodname string

	for {
		viewSL(sl)

		fmt.Println(" ")
		fmt.Println("Enter the NAME of the item (e.g. Bread, Coke, Water, etc) you wish to delete. (CASE SENSITIVE)")
		fmt.Scanln(&usermodname)

		_, presentcheck := sl[usermodname]

		if presentcheck == false {
			fmt.Println("Item not found, please check and retry.")
		} else {
			break
		}
	}

	delete(sl, usermodname)
	fmt.Println(usermodname, "has been deleted. Exiting...")

	return sl
}

// [MAIN] 6. View or add/modify/delete categories

func catEdit(sl map[string]details, catTracker []string) (map[string]details, []string) {
	for {
		fmt.Println(" ")
		fmt.Println("-------------------------")
		fmt.Println("View or add/modify/delete categories")
		fmt.Println("-------------------------")

		fmt.Println(" ")

		showCatTracker(catTracker)

		fmt.Println("1. Add categories")
		fmt.Println("2. Modify categories")
		fmt.Println("3. Delete categories")
		fmt.Println("4. Main Menu")
		fmt.Println(" ")

		var choice int
		fmt.Println("Select your choice: ")
		fmt.Scanln(&choice)

		if choice == 4 {
			fmt.Println("Exiting...")
			break
		}

		switch choice {
		case 1:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 1. Add categories")
			catTracker = addCat(catTracker)
		case 2:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 2. Modify categories")
			catTracker = modCat(catTracker)
		case 3:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Selected: 3. Delete categories")
			sl, catTracker = delCat(sl, catTracker)
		default:
			fmt.Println(" ")
			fmt.Println(" ")
			fmt.Println("Invalid option")
		}
	}
	return sl, catTracker
}

// [MAIN] 7. Print current data fields

func printdata(sl map[string]details) {
	if len(sl) == 0 {
		fmt.Println("No data found.")
	} else {
		for name, details := range sl {
			fmt.Println(name, "-", details)
		}
	}
}

// [MAIN] 8. Save and exit to Shopping List Loader

func saveExit(sl map[string]details, catTracker []string) {
	slsaver[accessedsave] = sl
	catTrackersaver[accessedsave] = catTracker
}
