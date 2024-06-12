package main

import (
	"encoding/json"
	"github.com/k0kubun/pp/v3"
	"log"
	"os"
	"sort"
	"time"
)

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Requests int       `json:"requests"`
	Domain string      `json:"domain"`
}

type Logs []Log

func (l *Logs) Len() int {
	return len(*l)
}
func (l *Logs) Less(i, j int) bool {
	return (*l)[i].Requests > (*l)[j].Requests
}
func (l *Logs) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]

}

func main() {

	file ,err := os.Open("./log.json")
	if err != nil {
		log.Println(err)
	}
	logs := Logs{}

	json.NewDecoder(file).Decode(&logs)

	sort.Sort(&logs)

	pp.Println(logs)


}