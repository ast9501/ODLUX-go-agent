package internal

import (
	. "agent/model"
	"os"
	"io/ioutil"
	"path/filepath"
	"log"
	"encoding/json"
)

func LoadOauthProvider() []byte {
	absPath, _ := filepath.Abs("config/oauth-provider.config.json")
	jsonFile, err := os.Open(absPath)
	if err != nil {
		log.Println(err)
		return nil
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var provider Provider

	json.Unmarshal(byteValue, &provider)
	log.Println(provider.Id)
	log.Println(provider.Title)
	log.Println(provider.Url)
	
	return byteValue
}