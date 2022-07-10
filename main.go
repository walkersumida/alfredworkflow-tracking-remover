package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

// Item is Alfred's item struct.
type Item struct {
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Menu is Alfred's menu struct.
type Menu struct {
	Items []Item `json:"items"`
}

type BlockList struct {
	BlockList []Block `json:"block_list"`
}

type Block struct {
	Param string `json:"param"`
}

func outputFormat(item Item) {
	var menu Menu
	item.Icon = "./icon.png"
	menu.Items = append(menu.Items, item)

	menuJSON, _ := json.Marshal(menu)
	fmt.Println(string(menuJSON))
}

func extractParam(u url.Values, p string) (string) {
	u.Del(p)
	return u.Encode()
}

func main() {
	flag.Parse()
	arg := flag.Arg(0)

	u, err := url.Parse(arg)
	if err != nil {
		fmt.Println("invalid url")
		return
	}

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var params BlockList
	err = json.Unmarshal(byteValue, &params)
	if err != nil {
		fmt.Println(err)
	}

	q, _ := url.ParseQuery(u.RawQuery)

	for _, b := range params.BlockList {
		u.RawQuery = extractParam(q, b.Param)
	}

	var item Item
	item.Title = fmt.Sprint(u)
	item.Arg = fmt.Sprint(u)

	outputFormat(item)
}
