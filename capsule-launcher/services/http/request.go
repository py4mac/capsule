package capsulehttp

import (
	"encoding/json"
	"github.com/bots-garden/capsule/commons"
	"github.com/gin-gonic/gin"
)

// GetJsonStringFromPayloadRequest :
func GetJsonStringFromPayloadRequest(c *gin.Context) (string, error) {
	jsonMap := make(map[string]interface{})
	if err := c.Bind(&jsonMap); err != nil {
		return "", err
	}
	// Convert map to json string
	jsonBytes, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// GetHeadersStringFromHeadersRequest :
func GetHeadersStringFromHeadersRequest(c *gin.Context) string {
	var headersMap = make(map[string]string)
	for key, values := range c.Request.Header {
		headersMap[key] = values[0]
	}
	headersSlice := commons.CreateSliceFromMap(headersMap)
	headersParameter := commons.CreateStringFromSlice(headersSlice, commons.StrSeparator)

	return headersParameter
}
