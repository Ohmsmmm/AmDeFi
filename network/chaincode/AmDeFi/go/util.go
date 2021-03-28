package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

//main * &
//var reportImageAsStruct []ReportImages // ถังเปล่า//A01
//err := parsingStringToStruct(args[0] ของที่จะใส่ถัง , &reportImageAsStruct ถังเปล่า//A01)

func parsingStringToStruct(get string, dataType *[]string) error { //dataType//A51 :: A01
	functionName := "[getReportImages]" //*dataType // ถังเปล่า
	println(functionName)
	//var reportImageAsStruct []ReportImages
	var jsonData = []byte(get) //ของที่จะใส่ถัง => ของเหลว
	// ByteArray to json
	err := json.Unmarshal(jsonData, dataType) //ถังที่มีน้ำ
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		return err
	}
	println(functionName + " successfully")
	return nil
}

func hashString(password string) string {
	s := password
	h := sha256.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return "0x" + sha1_hash
}
