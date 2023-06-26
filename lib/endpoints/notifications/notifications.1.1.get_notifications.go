package notifications

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-content-management-sdk/lib"
	helpers2 "github.com/snowpal/pitch-content-management-sdk/lib/helpers"
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/response"
)

func GetNotifications(jwtToken string) ([]response.Notification, error) {
	resNotifications := response.Notifications{}
	route, err := helpers2.GetRoute(lib.RouteNotificationsGetNotifications)
	if err != nil {
		return resNotifications.Notifications, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resNotifications.Notifications, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resNotifications.Notifications, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resNotifications.Notifications, err
	}

	err = json.Unmarshal(body, &resNotifications)
	if err != nil {
		return resNotifications.Notifications, err
	}
	return resNotifications.Notifications, nil
}
