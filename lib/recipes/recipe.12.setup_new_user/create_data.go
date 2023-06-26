package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	blockPods "github.com/snowpal/pitch-content-management-sdk/lib/endpoints/block_pods/block_pods.1"
	keyPods "github.com/snowpal/pitch-content-management-sdk/lib/endpoints/key_pods/key_pods.1"
)

var KeyName = "Insurances"
var BlockName = "Life Insurance"
var KeyPodName = "Health Insurance"
var BlockPodName = "Term Insurance"

func CreateData(user response.User) (KeyWithResources, error) {
	var err error
	var keyWithResources KeyWithResources

	log.Info("Creating custom key")
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: KeyName,
			Type: lib.CustomKeyType,
		})
	if err != nil {
		return keyWithResources, err
	}

	var blockWithPods BlockWithPods

	log.Info(fmt.Sprintf("Creating block inside %s key.", KeyName))
	var block response.Block
	block, err = blocks.AddBlock(
		user.JwtToken,
		request.AddBlockReqBody{Name: BlockName},
		key.ID)
	if err != nil {
		return keyWithResources, err
	}
	blockWithPods.Block = block

	log.Info(fmt.Sprintf("Creating block pod inside %s block.", BlockName))
	var blockPod response.Pod
	blockPod, err = blockPods.AddBlockPod(
		user.JwtToken,
		request.AddPodReqBody{Name: BlockPodName},
		common.ResourceIdParam{BlockId: block.ID, KeyId: key.ID},
	)
	if err != nil {
		return keyWithResources, err
	}
	blockWithPods.BlockPods = append(blockWithPods.BlockPods, blockPod)

	keyWithResources.Blocks = append(keyWithResources.Blocks, blockWithPods)

	log.Info(fmt.Sprintf("Creating key pod inside %s key.", KeyName))
	var keyPod response.Pod
	keyPod, err = keyPods.AddKeyPod(
		user.JwtToken,
		request.AddPodReqBody{Name: KeyPodName},
		key.ID)
	if err != nil {
		return keyWithResources, err
	}
	keyWithResources.Pods = append(keyWithResources.Pods, keyPod)

	return keyWithResources, err
}
