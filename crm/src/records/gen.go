package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ChimeraCoder/gojson"
)

func main() {
	f := getFiles("./")
	for _, f := range f {
		b, err := ioutil.ReadFile(f.Name())
		if err != nil {
			log.Fatal("IOUTIL:", err)
		}

		slug := strings.TrimSuffix(f.Name(), ".json")
		capSlug := strings.ToUpper(slug[:1])

		output, err := gojson.Generate(strings.NewReader(string(b)), gojson.ParseJson, capSlug, "crm", []string{"json"}, false, false)
		if err != nil {
			log.Fatal("GOJSON: ", capSlug, err)
		}

		outputStr := strings.ReplaceAll(string(output), "\"`", ",omitempty\"`")
		err = ioutil.WriteFile(fmt.Sprintf("%s_type.go", capSlug), []byte(outputStr), 0700)
		if err != nil {
			log.Fatal("WRITE FILE: ", capSlug, err)
		}
	}
}

func getFiles(dir string) (f []os.FileInfo) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			f = append(f, file)
		}
	}
	return f
}
