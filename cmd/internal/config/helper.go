package config

import (
	"errors"
	"log"
	"strconv"
)

func GetIntParam(value string, err error) (int64, error) {
	loginId := value
	if loginId == "" || err != nil {
		log.Println("Need to login to account api")
		return -1, errors.New("need to login to account api")
	}

	loginIdInt, err := strconv.ParseInt(loginId, 10, 64)
	if err != nil || loginIdInt <= 0 {
		log.Println(err)
		return -1, errors.New("invalid id value")
	}
	return loginIdInt, nil
}
