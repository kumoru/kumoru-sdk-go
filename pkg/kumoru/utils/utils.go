package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Pprint(data string) {
	buf := new(bytes.Buffer)
	json.Indent(buf, []byte(data), "", "  ")
	fmt.Println(buf)
}
