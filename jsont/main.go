// 当前程序的包名
package main

// 导入其它的包
import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	map2json2map()
}

func map2json2map() {

	map1 := make(map[string]interface{})
	map1["1"] = "hello"
	map1["2"] = []byte("hello")
	map1["age"] = 25
	map1["books"] = [...]float32{2.1, 3.3, 4.4}

	//return []byte
	str, err := json.MarshalIndent([]interface{}{2, true, "hello"}, "", "\t")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("map to json", string(str))

	//json([]byte) to map
	map2 := make(map[string]interface{})
	err = json.Unmarshal([]byte("{\"name\" : [2, true, 3.5, \"h\"]}"), &map2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("json to map ", map2)
	fmt.Println("The value of key1 is", map2["name"])
	fmt.Println(reflect.TypeOf(map2["name"]))
}
