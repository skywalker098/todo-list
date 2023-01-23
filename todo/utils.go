package todo

import (
	"encoding/json"
	"io"
	"os"
)

// #1.Save the todoStore to the Json file___________________________________
func (t *todoList) SaveToJson() {
	//create a json file if not exist
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		_, err := os.Create("db.json")
		if err != nil {
			panic(err)
		}
	}

	//open the json file
	file, err := os.OpenFile("db.json", os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}

	//marshal the data from todoStore
	data, err := json.Marshal(t.todoStore)
	if err != nil {
		panic(err)
	}

	//write the data to the file
	_, err = file.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	//close the json file
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

// #2.Load the data from the json file_____________________________________
func (t todoList) LoadFromJson() {
	//create a json file if not exist
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		// log.Fatalf("db.json does not exist")
		_, err := os.Create("db.json")
		if err != nil {
			panic(err)
		}

	}

	//open the json file
	file, err := os.OpenFile("db.json", os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}

	//convert the file to byte array
	fileByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//check if the file is non empty then unmarshal it

	if len(fileByte) > 0 {
		//log.Fatalf("db.json is empty")
		//unmarshal the data from json file to t.todoStore
		err = json.Unmarshal([]byte(fileByte), &t.todoStore)
		if err != nil {
			panic(err)
		}
	}

	// //unmarshal the data from json file to t.todoStore
	// err = json.Unmarshal([]byte(fileByte), &t.todoStore)
	// if err != nil {
	// 	panic(err)
	// }

	//close the json file
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
