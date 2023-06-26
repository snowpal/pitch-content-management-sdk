package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	"github.com/snowpal/pitch-content-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareData(user response.User, anotherUserEmail string, keyWithResources KeyWithResources) error {
	log.Info(fmt.Sprintf("Share key data with %s", anotherUserEmail))

	block := keyWithResources.Blocks[0].Block
	log.Info(fmt.Sprintf("Share block %s for custom key, %s", block.Name, keyWithResources.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	blockPod := keyWithResources.Blocks[0].BlockPods[0]
	log.Info(fmt.Sprintf("Share block pod %s for custom key, %s and for %s", blockPod.Name, keyWithResources.Key.Name, block.Name))
	err = recipes.SearchUserAndShareBlockPod(user, blockPod, anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	pod := keyWithResources.Pods[0]
	log.Info(fmt.Sprintf("Share key pod %s for custom key, %s", pod.Name, keyWithResources.Key.Name))
	err = recipes.SearchUserAndSharePod(user, pod, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	return nil
}
