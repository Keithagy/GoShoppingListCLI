package main

import "fmt"

// [LOADER] Initializes the loader.

func initloader() {
	for {
		if killswitch == true {
			break
		}

		if len(slsaver) == 0 {
			fmt.Println("No shopping lists saved currently. Initializing...")
			loadSL(newsave())
			slapp()
		} else {

			fmt.Println(" ")
			fmt.Println("-------------------------")
			fmt.Println("Shopping List Application")
			fmt.Println("Save Selector")
			fmt.Println("-------------------------")
			fmt.Println(" ")
			fmt.Println("Existing Saves:")
			for index := range slsaver {
				fmt.Println("Save Index:", index)
				fmt.Println("Shopping list:", slsaver[index])
				fmt.Println("Categories:", catTrackersaver[index])
				fmt.Println("------------------------------------")
			}
			fmt.Println(" ")
			fmt.Println("1. Create new save")
			fmt.Println("2. Access existing save")
			fmt.Println("3. Delete existing save")
			fmt.Println("4. Terminate Program")

			// Input user choice
			var menuchoice int
			fmt.Printf("Select your choice (INPUT NUMBERS ONLY): ")
			fmt.Scanln(&menuchoice)

			// Exiting the application
			if menuchoice == 4 {
				fmt.Println("Terminating Program...")
				killswitch = true
				break
			}

			// Switch statement for user control
			switch menuchoice {
			case 1:
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Println("Selected: 1. Create new save")
				newsave()
			case 2:
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Println("Selected: 2. Access existing save")
				fmt.Println(" ")
				accessindex := -100
				for {
					fmt.Println("Enter the save index to be accessed:")
					fmt.Scanln(&accessindex)

					if accessindex >= 0 && accessindex < len(slsaver) {
						break
					} else {
						fmt.Println("Invalid entry, please recheck save indices and try again")
					}
				}
				loadSL(accessindex)
				slapp()
			case 3:
				fmt.Println(" ")
				fmt.Println(" ")
				fmt.Println("Selected: 3. Delete existing save")
				deleteSL()
			}

		}

	}
}

// [LOADER] Creates a new save slot. Returns the slice index of the new save slot. Returns -1 in case of error.

func newsave() int {
	tmpsl := make(map[string]details)
	tmpcatTracker := make([]string, 0)

	slsaver = append(slsaver, tmpsl)
	catTrackersaver = append(catTrackersaver, tmpcatTracker)

	if len(slsaver) == len(catTrackersaver) {
		fmt.Println("New save successfully created at index", len(slsaver)-1)
		return len(slsaver) - 1
	}

	fmt.Print("ERROR! Lengths of slsaver and catTracker saver not equal")
	return -1

}

//  [LOADER] Deletes an existing save slot. Cannot delete SL if it is the only SL saved.

func deleteSL() {
	// Print existing saves
	for index := range slsaver {
		fmt.Println("Save Index:", index)
		fmt.Println("Shopping list:", slsaver[index])
		fmt.Println("Categories:", catTrackersaver[index])
		fmt.Println("------------------------------------")
	}

	if len(slsaver) <= 1 {
		fmt.Printf("Only one shopping list exists at the moment. Deletion not allowed.")
	} else {
		var userinput int
		for {
			fmt.Println("Input the Save index to be deleted")
			fmt.Scanln(&userinput)

			if userinput < len(slsaver) && userinput >= 0 {
				fmt.Println("Valid input, Save index", userinput, "will be deleted.")
				break
			} else {
				fmt.Println("Invalid entry, please check existing saves and retry.")
			}
		}
		copy(slsaver[userinput:], slsaver[userinput+1:])
		slsaver[len(slsaver)-1] = nil
		slsaver = slsaver[:len(slsaver)-1]

		copy(catTrackersaver[userinput:], catTrackersaver[userinput+1:])
		catTrackersaver[len(catTrackersaver)-1] = nil
		catTrackersaver = catTrackersaver[:len(catTrackersaver)-1]
	}
}

// [LOADER] Accesses an existing save slot and makes it modifiable within the app.

func loadSL(saveindex int) {
	sl = slsaver[saveindex]
	catTracker = catTrackersaver[saveindex]
	accessedsave = saveindex
}
