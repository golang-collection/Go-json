package tools

import (
	"Go-json/model"
	"bytes"
	"encoding/json"
)

/**
* @Author: super
* @Date: 2020-08-25 08:13
* @Description:
**/

func BookToJson(book *model.Book) (string, error) {
	data, err := json.Marshal(book)
	if err != nil {
		return "", err
	}
	// 格式化json输出
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "\t")
	if err != nil{
		return string(data), err
	}
	return out.String(), err
}

func JsonToBook(str string) (*model.Book, error) {
	book := &model.Book{}
	err := json.Unmarshal([]byte(str), book)
	if err != nil {
		return nil, err
	}

	return book, nil
}