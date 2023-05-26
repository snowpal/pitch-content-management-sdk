package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/snowpal/pitch-content-management-sdk/lib/config"
	"github.com/snowpal/pitch-content-management-sdk/lib/recipes"
	"os"
)

func main() {
	var err error
	if config.Files, err = config.InitConfigFiles(); err != nil {
		log.Fatal(err.Error())
		return
	}

	var recipeID int
	recipeIDInEnv := os.Getenv("RECIPE_ID")
	if len(recipeIDInEnv) == 0 {
		recipeID = 1
	} else {
		recipeID, err = strconv.Atoi(recipeIDInEnv)
		if err != nil {
			recipeID = 1
		}
	}

	switch recipeID {
	case 1:
		log.Info("Run Recipe1")
		recipes.RegisterFewUsers()
		break
	case 2:
		log.Info("Run Recipe2")
		recipes.GetResourceAttributes()
		break
	case 3:
		log.Info("Run Recipe3")
		recipes.CreatePrivateConversation()
		break
	case 4:
		log.Info("Run Recipe4")
		recipes.AddAndLinkResources()
		break
	case 5:
		log.Info("Run Recipe5")
		recipes.ShareBlock()
		break
	case 6:
		log.Info("Run Recipe6")
		recipes.GetAllKeys()
		break
	case 7:
		log.Info("Run Recipe7")
		recipes.AddFavorite()
		break
	case 8:
		log.Info("Run Recipe8")
		recipes.FetchScheduler()
		break
	case 9:
		log.Info("Run Recipe9")
		recipes.AddRelation()
		break
	case 10:
		log.Info("Run Recipe12")
		recipes.GrantAclOnCustomBlock()
		break
	case 11:
		log.Info("Run Recipe13")
		recipes.UpdateAttributes()
		break
	default:
		log.Info("pick a specific recipe to run")
	}
}
