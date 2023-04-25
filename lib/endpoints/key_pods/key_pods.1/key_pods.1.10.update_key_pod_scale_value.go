package keyPods

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-content-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/request"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"
)

func UpdateKeyPodScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	podParam common.ResourceIdParam,
) (response.UpdatePodScaleValue, error) {
	resPodScaleValue := response.UpdatePodScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteKeyPodsUpdateKeyPodScaleValue,
		podParam.PodId,
		podParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	err = json.Unmarshal(body, &resPodScaleValue)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}
	return resPodScaleValue, nil
}
