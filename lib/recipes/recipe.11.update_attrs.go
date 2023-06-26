package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/attributes"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-content-management-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-content-management-sdk/lib/helpers/recipes"
)

const (
	AttrsKeyName   = "Birds"
	AttrsBlockName = "Parrot"
)

// UpdateAttributes sign in, update key attributes, update block attributes, update pod attributes, update block pod attributes,
// get resource attributes
func UpdateAttributes() {
	log.Info("Objective: Update show/hide of key, block, pod & block pod attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Get displayable attributes")
	recipes.SleepBefore()
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}
	log.Info(resourceAttrs)

	log.Info("Update key attributes")
	recipes.SleepBefore()
	var key response.Key
	key, err = keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: AttrsKeyName,
			Type: lib.CustomKeyType,
		})
	if err != nil {
		return
	}
	err = attributes.UpdateKeyAttrs(
		user.JwtToken,
		key.ID,
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Attributes for Key %s updated successfully", key.Name))
	recipes.SleepAfter()

	log.Info("Update block attributes")
	recipes.SleepBefore()
	var block response.Block
	block, err = blocks.AddBlock(
		user.JwtToken,
		request.AddBlockReqBody{Name: AttrsBlockName},
		key.ID)
	if err != nil {
		return
	}
	err = attributes.UpdateBlockAttrs(
		user.JwtToken,
		common.ResourceIdParam{
			BlockId: block.ID,
			KeyId:   block.Key.ID,
		},
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf(".Attributes for block %s updated successfully", key.Name))
	recipes.SleepAfter()
}
