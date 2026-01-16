package docs

import (
	"encoding/json"
	"strings"
)

type TagsSummary struct {
	Tags    []string `json:"tags"`
	Summary string   `json:"summary"`
}

type tree struct {
	Paths map[string]map[string]TagsSummary `json:"paths"`
}

var apiTree *tree

func init() {
	err := json.Unmarshal([]byte(SwaggerInfo.ReadDoc()), &apiTree)
	if err != nil {
		panic(err)
	}
}

func GetApiTagsAndSummary(url, method string) *TagsSummary {
	url = strings.ReplaceAll(url, SwaggerInfo.BasePath, "")
	methodTree, ok := apiTree.Paths[url]
	if !ok {
		return nil
	}
	ts, ok := methodTree[strings.ToLower(method)]
	if !ok {
		return nil
	}
	return &ts
}
