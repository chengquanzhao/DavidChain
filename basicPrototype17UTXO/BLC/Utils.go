package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/json"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func JsonToArry(jsonString string) []string {
	// json 2 []string
	var wo []string
	err := json.Unmarshal([]byte(jsonString), &wo)
	if err != nil {
		log.Panic(err)
	}
	return wo
}
