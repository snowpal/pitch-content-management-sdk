package recipes

import (
	"github.com/snowpal/pitch-content-management-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

var KeyName = "Insurances"
var BlockName = "Life Insurance"
var KeyPodName = "Health Insurance"
var BlockPodName = "Term Insurance"

func CreateContent(user response.User) (KeyWithResources, error) {
	var err error
	var keyWithResources KeyWithResources

	log.Info("Creating custom key")
	key, err := recipes.AddCustomKey(user, KeyName)
	if err != nil {
		return keyWithResources, err
	}

	var blockWithPods BlockWithPods

	log.Info("Creating block inside ", KeyName, " key.")
	var block response.Block
	block, err = recipes.AddBlock(user, BlockName, key)
	if err != nil {
		return keyWithResources, err
	}
	blockWithPods.Block = block

	log.Info("Creating block pod inside ", BlockName, " block.")
	var blockPod response.Pod
	blockPod, err = recipes.AddPodToBlock(user, BlockPodName, block)
	if err != nil {
		return keyWithResources, err
	}
	blockWithPods.BlockPods = append(blockWithPods.BlockPods, blockPod)

	keyWithResources.Blocks = append(keyWithResources.Blocks, blockWithPods)

	log.Info("Creating key pod inside ", KeyName, " key.")
	var keyPod response.Pod
	keyPod, err = recipes.AddPod(user, KeyPodName, key)
	if err != nil {
		return keyWithResources, err
	}
	keyWithResources.Pods = append(keyWithResources.Pods, keyPod)

	return keyWithResources, err
}
