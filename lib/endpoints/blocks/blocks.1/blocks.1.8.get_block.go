package blocks

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-content-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"
)

func GetBlock(jwtToken string, blockParam common.ResourceIdParam) (response.Block, error) {
	resBlock := response.Block{}
	route, err := helpers2.GetRoute(lib.RouteBlocksGetBlock, blockParam.BlockId, blockParam.KeyId)
	if err != nil {
		return resBlock, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		return resBlock, err
	}
	return resBlock, nil
}
