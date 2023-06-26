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

func GetBlocksLinkedToPods(jwtToken string, blockParam common.ResourceIdParam) ([]response.Block, error) {
	resBlocks := response.Blocks{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetBlocksLinkedToPod,
		blockParam.PodId,
		blockParam.KeyId,
	)
	if err != nil {
		return resBlocks.Blocks, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resBlocks.Blocks, err
	}

	helpers2.AddUserHeaders(jwtToken, req)
	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resBlocks.Blocks, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resBlocks.Blocks, err
	}

	err = json.Unmarshal(body, &resBlocks)
	if err != nil {
		return resBlocks.Blocks, err
	}
	return resBlocks.Blocks, nil
}
