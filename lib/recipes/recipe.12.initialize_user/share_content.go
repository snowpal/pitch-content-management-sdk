package recipes

import (
	"github.com/snowpal/pitch-content-management-sdk/lib"
	"github.com/snowpal/pitch-content-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareContent(user response.User, anotherUserEmail string, keyWithResources KeyWithResources) error {
	log.Info("Share key content with ", anotherUserEmail)

	log.Info("Share block ", keyWithResources.Blocks[0].Block.Name, " for custom key, ", keyWithResources.Key.Name)
	err := recipes.SearchUserAndShareBlock(user, keyWithResources.Blocks[0].Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	log.Info("Share block pod ", keyWithResources.Blocks[0].BlockPods[0].Name, " for custom key, ", keyWithResources.Key.Name, " and for ", keyWithResources.Blocks[0].Block.Name)
	err = recipes.SearchUserAndShareBlockPod(user, keyWithResources.Blocks[0].BlockPods[0], anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	log.Info("Share key pod ", keyWithResources.Pods[0].Name, " for custom key, ", keyWithResources.Key.Name)
	err = recipes.SearchUserAndSharePod(user, keyWithResources.Pods[0], anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	return nil
}
