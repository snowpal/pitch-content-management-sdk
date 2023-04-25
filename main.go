package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-content-management-sdk/lib/recipes"
)

func runRecipe(recipeID int) {
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
		recipes.AddAndLinkResources()
		break
	case 4:
		log.Info("Run Recipe5")
		recipes.ShareBlock()
		break
	case 5:
		log.Info("Run Recipe5")
		recipes.GetAllKeys()
		break
	case 6:
		log.Info("Run Recipe6")
		recipes.AddFavorite()
		break
	case 7:
		log.Info("Run Recipe7")
		recipes.FetchScheduler()
		break
	case 8:
		log.Info("Run Recipe8")
		recipes.AddRelation()
		break
	case 9:
		log.Info("Run Recipe9")
		recipes.GrantAclOnCustomBlock()
		break
	case 10:
		log.Info("Run Recipe10")
		recipes.UpdateAttributes()
		break
	default:
		log.Info("pick a specific recipe to run")
	}
}

func main() {
	for i := 1; i <= 10; i++ {
		runRecipe(i)
	}
}
