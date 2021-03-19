package main

import (
	"encoding/json"
)

func (s Session) GetQueue() (CommandQueue, error) {
	results, err := s.dial("queue", "GET", nil)
	if err != nil {
		return CommandQueue{}, err
	}
	cq := &CommandQueue{}
	err = json.Unmarshal(results, cq)

	if err != nil {
		return CommandQueue{}, err
	}

	return *cq, nil
}

func (s Session) GetOfflineQueue() (OfflineQueue, error) {
	results, err := s.dial("queue/offline-commands", "GET", nil)
	if err != nil {
		return OfflineQueue{}, err
	}
	oq := &OfflineQueue{}
	err = json.Unmarshal(results, oq)

	if err != nil {
		return OfflineQueue{}, err
	}
	return *oq, nil
}

func (s Session) GetOnlineQueue(username string) (OnlineQueue, error) {
	results, err := s.dial("queue/online-commands/"+username, "GET", nil)
	if err != nil {
		return OnlineQueue{}, err
	}
	oq := &OnlineQueue{}
	err = json.Unmarshal(results, oq)

	if err != nil {
		return OnlineQueue{}, err
	}

	return *oq, nil
}

// TODO DeleteCommandQ
// func (s Session) DeleteCommand(commands []int) (bool, error) {
//
// 	lol, err := s.dial("queue", "DELETE", data)
// 	if err != nil {
// 		return false, err
// 	}
// 	fmt.Println(string(lol))
// 	return true, nil
// }
