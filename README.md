# World of Warcraft Community API Client for Golang

Provides Golang structs and an API client for consuming WoW Community
API data

See [the official
documentation](http://blizzard.github.io/api-wow-docs/) for an
overview of what data is avaialable, then check out [api_client.go's
docs](http://godoc.org/github.com/capoferro/wow#ApiClient) for the
functions used to retrieve said data.

## All APIs Supported:

* Achievement
* Auction
* BattlePet
* Challenge Mode
* Character Profile
* Item
* Guild Profile
* PvP
* Quest
* Realm Status
* Recipe
* Spell
* Data Resources

Todo:

* Verify signature with real account

## Usage

```go
package main

import (
  "fmt"
  "github.com/capoferro/wow"
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
```
