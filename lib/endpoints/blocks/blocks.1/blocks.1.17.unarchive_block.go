package blocks

import (
	"net/http"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	"github.com/snowpal/pitch-content-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
)

func UnarchiveBlock(jwtToken string, blockParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlocksUnarchiveBlock,
		blockParam.BlockId,
		blockParam.KeyId,
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
