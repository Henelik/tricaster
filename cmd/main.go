//go:build !test
// +build !test

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/Henelik/tricaster/pkg/renderer"
	"gopkg.in/yaml.v3"
)

var filename string

func init() {
	flag.StringVar(&filename, "f", "scene.yml", "file path for the scene to render")
}

func main() {
	start := time.Now()

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

	scene := renderer.NewScene(config)

	fmt.Printf("rendering scene %s\n", filename)

	scene.Render()

	fmt.Printf("render took %s\n", time.Since(start))
}
