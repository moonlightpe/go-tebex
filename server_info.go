package main

import (
	"encoding/json"
)

func (s Session) GetInformation() (ServerInfo, error) {
	results, err := s.dial("information", "GET", nil)
	if err != nil {
		return ServerInfo{}, err
	}

	info := &ServerInfo{}
	err = json.Unmarshal(results, info)
	if err != nil {
		return ServerInfo{}, err
	}

	return *info, nil
}
