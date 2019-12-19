package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"

	"github.com/go-tee/di"
	"github.com/go-tee/di/config"
	"github.com/go-tee/di/ext"
	"github.com/go-tee/di/utils"
)

var files utils.StringSlice
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
		conf, err := config.ParseConfig(files...)
		if err != nil {
			log.Fatalf("Error loading config: %s", err)
		}
		merged, err := yaml.Marshal(conf)
		if err != nil {
			log.Fatalf("Error building merged YAML: %s", err)
		}
		if showConfig {
			fmt.Println(string(merged))
			return
		}

		c := di.NewCompiler(conf)
		// c.AddExtension("parameters", ext.NewParametersExtension())
		c.MustAddExtension("services", ext.NewServicesExtension())
		c.MustAddExtension("navigation", ext.NewNavigationExtension())

		fmt.Printf("Result: %s\n", conf)

		if err = c.Compile("main", output); err != nil {
			log.Fatalln("Compiler error:", err)
		}

		log.Println("Success")
	}
}
