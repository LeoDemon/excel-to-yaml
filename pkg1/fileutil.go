package pkg1

import (
	"errors"
	"io/ioutil"
	"log"
)

func ReadFileData(fileName string) (string, error) {
	if len(fileName) <= 0 {
		return "", errors.New("invalid fileName")
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("read file [%s] error: [%v]\n", fileName, err)
		return "", err
	}
	return string(content), nil
}

func SaveDataToFile(fileName string, content string) error {
	if len(fileName) <= 0 {
		return errors.New("invalid fileName")
	}
	if len(content) <= 0 {
		return errors.New("content is empty")
	}

	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
