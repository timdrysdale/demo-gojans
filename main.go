package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	jans "github.com/timdrysdale/gojans/pkg/jansAuth/client"
)

func main() {

	var reg strfmt.Registry
	ja := jans.NewHTTPClient(reg)
	fmt.Println(prettyStringStruct((*ja).ClientInfo))

}

// prettyStringStruct returns struct formatted into pretty string
func prettyStringStruct(t interface{}) string {

	json, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return string(json)
}
