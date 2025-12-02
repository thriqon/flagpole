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
	Route    string            `yaml:"route"`
	Headers  map[string]string `yaml:"headers"`
	Status   int               `yaml:"status"`
	Contents string            `yaml:"body"`
}

func (r route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.RemoteAddr, r.Route)

	if r.Status == 0 {
		r.Status = 200
	}

	for k, v := range r.Headers {
		w.Header().Set(k, v)
	}
	if w.Header().Get("Content-type") == "" {
		w.Header().Set("Content-type", http.DetectContentType([]byte(r.Contents)))
	}

	w.WriteHeader(r.Status)

	if _, err := w.Write([]byte(r.Contents)); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
		log.Fatalf("unable to read configfile %s: %v", configFile, err)
	}

	for _, v := range config.Routes {
		http.Handle(v.Route, v)
	}

	log.Println("starting server on ", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
