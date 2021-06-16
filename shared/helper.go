package shared

import (
	"encoding/json"
	"fmt"
)

//Shared : ""
type Shared struct {
	commandArgs map[string]string
}

//NewShared : Shared Factory
func NewShared(commandArgs map[string]string) *Shared {
	return &Shared{commandArgs}
}

//GetCmdArg : ""
func (sh *Shared) GetCmdArg(key string) string {
	return sh.commandArgs[key]
}

//BsonToJSONPrint : ""
func BsonToJSONPrint(d interface{}) {
	b, err1 := json.Marshal(d)
	fmt.Println("err1", err1, string(b))
}
