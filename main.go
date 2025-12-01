package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type configFile struct {
	Routes []route `yaml:"routes"`
}

type route struct {
	Route       string `yaml:"route"`
	ContentType string `yaml:"contentType"`
	Contents    string `yaml:"contents"`
}

func loadConfigFile(filename string) (configFile, error) {
	f, err := os.Open(filename)
	if err != nil {
		return configFile{}, fmt.Errorf("unable to open file: %w", err)
	}
	defer f.Close()

	var cf configFile
	if err := yaml.NewDecoder(f).Decode(&cf); err != nil {
		return configFile{}, fmt.Errorf("unable to decode yaml: %w", err)
	}

	return cf, nil
}

func main() {
	var configFile string
	var address string
	flag.StringVar(&configFile, "config-file", "/etc/flagpole.yml", "Config file location")
	flag.StringVar(&address, "listen-address", ":8080", "Address to listen on")
	flag.Parse()

	config, err := loadConfigFile(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to load config: ", err)
	}

	for _, v := range config.Routes {
		http.HandleFunc(v.Route, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.RemoteAddr, v.Route)

			contentType := v.ContentType
			if contentType == "" {
				contentType = http.DetectContentType([]byte(v.Contents))
			}

			w.Header().Set("Content-type", contentType)
			if _, err := w.Write([]byte(v.Contents)); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		})
	}

	log.Fatal(http.ListenAndServe(address, nil))
}
