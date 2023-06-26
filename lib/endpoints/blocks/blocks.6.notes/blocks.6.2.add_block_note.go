package blocks

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

func AddBlockNote(
	jwtToken string,
	reqBody request.NoteReqBody,
	commentParam request.NoteIdParam,
) (response.Note, error) {
	resNote := response.Note{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resNote, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlocksAddBlockNote,
		*commentParam.BlockId,
		commentParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resNote, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resNote, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resNote, err
	}

	err = json.Unmarshal(body, &resNote)
	if err != nil {
		return resNote, err
	}
	return resNote, nil
}
