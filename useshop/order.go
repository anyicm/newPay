package useshop

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func GetSha256(str string, key string) string {
	srcByte := []byte(str)
	srcKey := []byte(key)
	sha256Bytes := hmac.New(sha256.New, srcKey)
	sha256Bytes.Write(srcByte)
	return hex.EncodeToString(sha256Bytes.Sum(nil))
}

func GetSha512(str string) string {
	srcByte := []byte(str)
	sha512Bytes := sha512.New()
	sha512Bytes.Write(srcByte)
	h512 := sha512Bytes.Sum(nil)
	sha512Str := hex.EncodeToString(h512)
	h := md5.New()
	h.Write([]byte(sha512Str))
	return hex.EncodeToString(h.Sum(nil))
}

func StrDecrypt(str string, key string) (err error, res string) {

	if str == "" {
		return errors.New("str params error"), ""
	}
	decodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	var newStr strings.Builder
	strLen := len(decodeString)
	keyLen := len(key)
	for i := 0; i < strLen; i++ {
		k := i % keyLen
		tmpStr := rune(decodeString[i]) ^ rune(key[k])
		if tmpStr > rune(127) {
			newStr.WriteString(strings.Trim(strconv.QuoteToASCII(string(tmpStr)), "\""))
		} else {
			newStr.WriteString(string(tmpStr))
		}
	}
	return nil, newStr.String()
}

func resToApiParams(str string) (err error, outData *ApiParams) {
	outData = new(ApiParams)
	err = json.Unmarshal([]byte(str), outData)
	if err != nil {
		return
	}
	return
}

func StructToMapJson(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {

		jsonKey := t.Field(i).Tag.Get("json")

		if jsonKey != "-" {
			data[jsonKey] = v.Field(i).Interface()
		}

	}

	return data
}

func MapKeySort(m map[string]interface{}) (sortStr string, err error) {
	s := make([]string, 0, len(m))
	newMap := make(map[string]interface{})
	for key := range m {
		s = append(s, key)
	}
	sort.Strings(s)
	for _, k := range s {
		newMap[k] = m[k]
	}
	newStr, err := json.Marshal(newMap)
	if err != nil {
		return
	}
	return string(newStr), nil
}
