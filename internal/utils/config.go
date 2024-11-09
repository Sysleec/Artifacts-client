package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

func LoadTokens() (map[string]string, error) {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	localConfig, err := ini.Load("config.local.ini")
	if err == nil {
		return loadTokensFromIni(localConfig)
	}

	return loadTokensFromIni(iniData)
}

func loadTokensFromIni(iniData *ini.File) (map[string]string, error) {
	generalSection := iniData.Section("ACCOUNTS")

	result := make(map[string]string)

	for _, section := range generalSection.Keys() {
		result[strings.ToLower(section.Name())] = section.Value()
	}

	return result, nil
}
