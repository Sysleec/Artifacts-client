package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func LoadToken() (string, error) {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		return "", fmt.Errorf("failed to load config file: %w", err)
	}

	localConfig, err := ini.Load("config.local.ini")
	if err == nil {
		return loadTokenFromIni(localConfig)
	}

	return loadTokenFromIni(iniData)
}

func loadTokenFromIni(iniData *ini.File) (string, error) {
	generalSection := iniData.Section("General")

	tokenValue := generalSection.Key("TOKEN").String()

	if tokenValue == "" || tokenValue == "ACCOUNT_TOKEN" {
		return "", fmt.Errorf("please provide a valid token in the config.ini file")
	}
	return tokenValue, nil
}
