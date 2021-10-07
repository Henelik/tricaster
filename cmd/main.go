//+build !test

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Henelik/tricaster/pkg/renderer"
	"gopkg.in/yaml.v3"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "scene.yml", "file path for the scene to render")
}

func main() {
	flag.Parse()

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	config := new(renderer.Configuration)

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", config)

	scene := renderer.NewScene(config)

	scene.Render()
}
