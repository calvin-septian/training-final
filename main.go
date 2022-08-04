package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"training-final/database"
	"training-final/helper"
	"training-final/services"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	flag "github.com/spf13/pflag"
)

type configuration struct {
	ConnectionString string `yaml:"ConnectionString"`
}

type options struct {
	configFilename string
}

var (
	config = configuration{}
)

func main() {
	args := loadOptions()
	if args.configFilename == "" {
		args.configFilename = "config.yaml"
	}

	err := parseConfig(args.configFilename)
	if err != nil {
		fmt.Println(err)
	}

	sql := database.NewSQLConnection(config.ConnectionString)
	database.DbConn = *sql
	defer database.DbConn.LocalDB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", services.UsersHandler)
	r.HandleFunc("/photo/{id}", services.PhotosHandler)
	r.HandleFunc("/comment/{id}", services.CommentsHandler)
	r.HandleFunc("/socialmedia/{id}", services.SocialMediasHandler)

	r.Use(helper.IsAuthorized)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
	}
	srv.ListenAndServe()
}

func loadOptions() options {
	o := options{}

	flag.StringVar(&o.configFilename, "config", "", "Path to the config files")

	flag.CommandLine.SortFlags = false

	flag.Parse()

	return o
}

func parseConfig(configFile string) error {

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
		return err
	}

	return nil
}
