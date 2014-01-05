package main

import (
  "fmt"
  "github.com/bluepojo/wow"
)

func main() {
	// Create a new client. The second parameter can be left empty if you
	// wish to use the region's default locale.
	client, _ := wow.NewApiClient("US", "")

	capo, _ := client.GetCharacterWithFields("Runetotem", "Capoferro", []string{"items"})
	className, _ := capo.Class()

	fmt.Printf("%s is the greatest %s ever.\nHe has an ilvl of %d and %d achievement points.\n", 
		capo.Name,
		className,
		capo.Items.AverageItemLevel,
		capo.AchievementPoints)
}
