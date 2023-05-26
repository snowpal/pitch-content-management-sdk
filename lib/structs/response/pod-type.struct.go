package response

import (
	common2 "github.com/snowpal/pitch-content-management-sdk/lib/structs/common"
)

type PodTypes struct {
	PodTypes []PodType `json:"podTypes"`
}

type PodType struct {
	ID   string             `json:"id"`
	Name string             `json:"podTypeName"`
	Pods *[]common2.SlimPod `json:"pods"`

	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
