package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/request"

	recipes "github.com/snowpal/pitch-content-management-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	CopyKeyName   = "Insurance"
	CopyBlockName = "Car Insurance"
)

func GrantAclOnCustomBlock() {
	log.Info("Objective: Add Custom Block, Share Block, Grant Read Access, Copy Block, Grant Admin Access")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var key response.Key
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: CopyKeyName,
			Type: lib.CustomKeyType,
		})
	if err != nil {
		return
	}

	log.Info("Add custom block")
	recipes.SleepBefore()
	var block response.Block
	block, err = blocks.AddBlock(
		user.JwtToken,
		request.AddBlockReqBody{Name: CopyBlockName},
		key.ID)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Block %s added successfully", block.Name))
	recipes.SleepAfter()

	log.Info("Share block with read access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Block %s shared with %s with read access level", block.Name, lib.ReadUser))
	recipes.SleepAfter()

	log.Info("Copy block and see acl is not copied")
	recipes.SleepBefore()
	var anotherBlock response.Block
	anotherBlock, err = copyBlock(user, block)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Block %s copied but %s don't have access on copied block", block.Name, lib.ReadUser))
	recipes.SleepAfter()

	log.Info("Share block with admin access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareBlock(user, anotherBlock, lib.AdminUser, lib.AdminAcl)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Block %s shared with %s with admin access", block.Name, lib.ReadUser))
	recipes.SleepAfter()
}

func copyBlock(user response.User, block response.Block) (response.Block, error) {
	resBlock, err := blocks.CopyBlock(
		user.JwtToken,
		request.CopyMoveBlockParam{
			BlockId:       block.ID,
			KeyId:         block.Key.ID,
			TargetKeyId:   block.Key.ID,
			AllPods:       true,
			AllTasks:      true,
			AllChecklists: true,
		},
	)
	if err != nil {
		return resBlock, err
	}
	return resBlock, nil
}
