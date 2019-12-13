package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/gooff/di/ds"
	"github.com/gooff/di/ext"
)

var files ds.StringSlice
var showConfig bool
var output string

func main() {
	flag.Var(&files, "f", "List of YAML config files")
	flag.BoolVar(&showConfig, "show-config", false, "Shown merged configuration")
	flag.StringVar(&output, "o", "di.go", "Output file")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
	} else {
		config := ds.ParseConfig(files...)
		merged, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("Error building merged YAML: %s", err)
		}
		if showConfig {
			fmt.Println(string(merged))
			return
		}

		c := ext.NewCompiler(config)
		// c.AddExtension("parameters", ext.NewParametersExtension())
		c.AddExtension("services", ext.NewServicesExtension())
		c.AddExtension("navigation", ext.NewNavigationExtension())

		fmt.Printf("Result: %s\n", config)

		if err = c.Compile("main", output); err != nil {
			log.Fatalln("Compiler error:", err)
		}

		log.Println("Success")
	}
}
