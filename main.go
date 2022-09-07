package gradecryptdb

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func readKey(key string) string {
	if hasKey(key) == true {
		file, err := os.Open(key)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		ret := (string(data))
		return ret
	}
	return ""
}
func addKey(key string, value string) {
	if hasKey(key) == false {
		newFile, err = os.Create(key)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(newFile)
		newFile.Close()
		err2 := ioutil.WriteFile(key, []byte(value), 0666)
		if err2 != nil {
			log.Fatal(err2)
		}
	}

}
func removeKey(key string) {
	if hasKey(key) == true {
		err := os.Remove(key)
		if err != nil {
			log.Fatal(err)
		}
	}

}
func hasKey(key string) bool {

	_, err := os.Stat(key)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
