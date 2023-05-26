package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-content-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"
)

type BlockByTemplateParam struct {
	KeyId        string
	TemplateId   string
	ExcludePods  bool
	ExcludeTasks bool
}

func AddBlockBasedOnTemplate(
	jwtToken string,
	reqBody request.AddBlockReqBody,
	blockParam BlockByTemplateParam,
) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	var route string
	route, err = helpers2.GetRoute(
		lib.RouteBlocksAddBlockBasedOnTemplate,
		blockParam.KeyId,
		blockParam.TemplateId,
		strconv.FormatBool(blockParam.ExcludePods),
		strconv.FormatBool(blockParam.ExcludeTasks),
	)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	return resBlock, nil
}
