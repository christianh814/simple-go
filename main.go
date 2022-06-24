package main

import (
	"html/template"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Config string
	Status string
	Msg    string
	Alert  string
}

// static port for now
const HttpPort string = "8080"

// Locate where the index template is
const IndexHtml string = "html/index.tmpl"

// App configuration file location
const AppConfigFile string = "/etc/myapp/test.conf"

func main() {
	//Set up handler for /
	http.HandleFunc("/", appRoot)

	// try to start the app and log output
	log.Info("Starting server on port " + HttpPort)
	log.Fatal(http.ListenAndServe(":"+HttpPort, nil))
}

// appRoot is the HTTP root of the application
func appRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(IndexHtml))
	config := AppConfig{}

	// If file exists, get the contents, if not set defaults
	if _, err := os.Stat(AppConfigFile); os.IsNotExist(err) {
		config = AppConfig{
			Config: "",
			Status: "Unknown",
			Alert:  "alert-warning",
			Msg:    "This means the config file was not found and will use the default config",
		}
	} else {
		// read the file, display any errors
		content, err := os.ReadFile(AppConfigFile)
		if err != nil {
			log.Fatal(err)
		}

		config = AppConfig{
			Config: string(content),
			Status: "Success",
			Alert:  "alert-success",
			Msg:    "This means that the app was able to successfully read the config file.",
		}

	}

	// Display index page from template
	tmpl.Execute(w, config)

}
