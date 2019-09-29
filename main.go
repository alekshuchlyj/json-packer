package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func pack(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	for key, value := range input {
		if reflect.ValueOf(value).Kind() == reflect.Map {
			temp := pack(value.(map[string]interface{}))
			for k, v := range temp {
				output[key + "." + k] = v
			}
		} else {
			output[key] = value
		}
	}

	return output
}

func unpack(keys []string, value interface{}, output map[string]interface{}) {
	currentKey, keys := removeIndex(keys, 0)

	if len(keys) != 0 {
		if _, ok := output[currentKey]; !ok {
			output[currentKey] = make(map[string]interface{})
		}
		unpack(keys, value, output[currentKey].(map[string]interface{}))
	} else {
		output[currentKey] = value
	}
}

func removeIndex(arr []string, index int) (string, []string) {
	remVal := arr[0]

	if index < len(arr) {
		arr = append(arr[:index], arr[index+1:]...)
	} else {
		fmt.Println("Invalid position")
	}

	return remVal, arr
}

func mapCleaning(input map[string]interface{}) {
	for key := range input {
		delete(input, key )
	}
}

func viewIn() {
	fmt.Printf("%.*s\n", 45, "--------------------------------------------------")
	fmt.Printf("|%21s %-21s|\n", "Json", "Input")
	fmt.Printf("%.*s\n", 45, "--------------------------------------------------")

}

func viewOut() {
	fmt.Printf("%.*s\n", 45, "--------------------------------------------------")
	fmt.Printf("|%20s %-22s|\n", "Json", "Output")
	fmt.Printf("%.*s\n", 45, "--------------------------------------------------")
}

func viewDashLine()  {
	fmt.Printf("%.*s\n", 45, "--------------------------------------------------")
}

func main() {
	fmt.Printf("\n\n%27s \n", "~Packing~")
	viewIn()

	jsonStr := []byte( `{"a":"1", "b":{"c":"2", "d":{"e":"3"}, "g":{"h":"5"}}, "f":"4"}` )
	input := make(map[string]interface{})

	if err := json.Unmarshal(jsonStr, &input); err != nil{
		panic("Deserializing error")
	} else {
		fmt.Println(string(jsonStr))
	}

	viewOut()

	output := pack(input)
	for key, val := range output {
		fmt.Printf("|%8s: %-33v|\n", key, val)
	}

	viewDashLine()

	mapCleaning(input)
	mapCleaning(output)

	fmt.Printf("\n\n%28s \n", "~Unpacking~")
	viewIn()

	jsonStr = []byte( `{"a":"1", "b.c":"2", "b.d.e":"3", "b.g.h": "5", "f":"4"	}` )

	if err := json.Unmarshal(jsonStr, &input); err!= nil {
		panic("Deserializing error")
	} else {
		for key, val := range input {
			fmt.Printf("|%8s: %-33v|\n", key, val)

		}
	}

	temp := pack(input)

	for key, value := range temp {
		keys := strings.Split(key, ".")
		unpack(keys, value, output)
	}

	viewOut()

	if jsonData, err := json.Marshal(output); err!= nil {
		panic("Deserializing error")
	} else {
		fmt.Println(string(jsonData))
	}

	viewDashLine()

	mapCleaning(input)
	mapCleaning(output)
}
