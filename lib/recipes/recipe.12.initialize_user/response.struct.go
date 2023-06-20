package recipes

import "github.com/snowpal/pitch-content-management-sdk/lib/structs/response"

type KeyWithResources struct {
	Key    response.Key
	Blocks []BlockWithPods
	Pods   []response.Pod
}

type BlockWithPods struct {
	Block     response.Block
	BlockPods []response.Pod
}
