package tools

import (
	"Go-json/model"
	"encoding/json"
)

/**
* @Author: super
* @Date: 2020-08-25 08:13
* @Description:
**/

func BookToJson(book *model.Book) (string, error) {
	str, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func JsonToBook(str string) (*model.Book, error) {
	book := &model.Book{}
	err := json.Unmarshal([]byte(str), book)
	if err != nil {
		return nil, err
	}
	return book, nil
}