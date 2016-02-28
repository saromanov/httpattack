package httpattack

import (
   "encoding/json"
   "io/ioutil"
)


type Actions struct {

}
func loadConfig(path string) error {
	file, error := ioutil.ReadFile(path)
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        return 
    }
    fmt.Printf("%s\n", string(file))
}