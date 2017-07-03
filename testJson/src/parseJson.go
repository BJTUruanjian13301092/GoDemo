package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Return []Desc `json:"return"`
}
type Desc struct {
	Field  []string `json:"field"`
	Start  int64    `json:"start"`
	Token  string   `json:"token"`
	Expire int64    `json:"expire"`
	User   string   `json:"user"`
	Eauth  string   `json:"eauth"`
}

func main() {

	str := `{"return": [
                        {  
                            "field": [".*"],  
                            "start": 1473841133,  
                            "token": "token1",  
                            "expire": 1473884333,  
                            "user": "xiaochuan",  
                            "eauth": "ss"  
                        },  
                        {  
                            "field": [".*"],  
                            "start": 1473841133,  
                            "token": "token2",  
                            "expire": 1473884333,  
                            "user": "xiaochuan",  
                            "eauth": "sr"  
                        }  
                            ]  
                }`

	t_struct := T{}
	err := json.Unmarshal([]byte(str), &t_struct)
	if err != nil {
		fmt.Println("error is %v\n", err)
	} else {
		fmt.Printf("%v\n", t_struct)
		fmt.Printf("%s\n", t_struct.Return[1].Token)
	}
}