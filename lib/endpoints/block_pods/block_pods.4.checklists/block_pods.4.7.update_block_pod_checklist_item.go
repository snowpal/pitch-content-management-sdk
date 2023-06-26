package blockPods

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-content-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"
)

func UpdateBlockPodChecklistItem(
	jwtToken string,
	reqBody request.ChecklistItemReqBody,
	checklistParam request.ChecklistIdParam,
) (response.ChecklistItem, error) {
	resChecklistItem := response.ChecklistItem{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resChecklistItem, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlockPodsUpdateBlockPodChecklistItem,
		*checklistParam.PodId,
		*checklistParam.ChecklistId,
		*checklistParam.ChecklistItemId,
		checklistParam.KeyId,
		*checklistParam.BlockId,
	)
	if err != nil {
		return resChecklistItem, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resChecklistItem, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resChecklistItem, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resChecklistItem, err
	}

	err = json.Unmarshal(body, &resChecklistItem)
	if err != nil {
		return resChecklistItem, err
	}
	return resChecklistItem, nil
}
