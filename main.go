package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/aleksrutins/terminated/utils"
	"gopkg.in/yaml.v3"
)

func main() {
	content, err := ioutil.ReadFile("credits.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	credits := make(map[string]string)

	err = yaml.Unmarshal(content, &credits)
	if err != nil {
		fmt.Println(err)
		return
	}

	for role, name := range credits {
		text := fmt.Sprintf("\x1b[32m%s\x1b[0m\n%s", role, name)
		utils.TypeText(text, 50*time.Millisecond)
	}
}
