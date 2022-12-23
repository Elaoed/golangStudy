package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

/**
结构体标签说明, 这个字段在不同的包里面不同的作用说明
*/
type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func findTag(str interface{}) {
	// 有Elem和没有有什么区别
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, ", doc: ", tagdoc)
	}

}

func main() {
	var re resume
	findTag(&re)

	movie := Movie{"喜剧之王", 2000, 10, []string{"星爷", "张柏芝"}}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {

	}
	fmt.Println(myMovie)

}
