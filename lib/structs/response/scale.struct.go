package response

import (
	"github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
)

type Scales struct {
	Scales []Scale `json:"scales"`
}

type Scale struct {
	ID          string   `json:"id"`
	Name        string   `json:"scaleName"`
	Type        *string  `json:"scaleType"`
	ScaleValues []string `json:"scaleValues"`

	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
