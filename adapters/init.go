package adapters

import (
	"log"

	"github.com/MAAF72/go-boilerplate/adapters/database"
	"github.com/MAAF72/go-boilerplate/adapters/grule"
)

// Adapters adapters struct
type Adapters struct {
	Database *database.Database
	Grule    *grule.Grule
}

// Init init all adapters
func Init() Adapters {
	db, err := database.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	grule, err := grule.Init()
	if err != nil {
		log.Fatal(err)
	}

	return Adapters{
		Database: db,
		Grule:    grule,
	}
}
