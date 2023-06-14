package app

import (
	"html/template"
	"net/http"
	"os"

	"github.com/christianh814/simple-go/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// AppConfig type
type AppConfig struct {
	Config []string
	Status string
	Msg    string
	Alert  string
}

// Info type
type Info struct {
	Alert  string
	Status string
	Info   map[string]string
	Msg    string
}

// Locate where the templates are
const IndexHtml string = "html/index.tmpl"
const InfoHtml string = "html/info.tmpl"

// App configuration file location
const AppConfigFile string = "/etc/myapp/test.conf"

// appRoot is the HTTP root of the application
func appRoot(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(IndexHtml))
	config := AppConfig{}

	// If file exists, get the contents, if not set defaults
	if _, err := os.Stat(AppConfigFile); os.IsNotExist(err) {
		config = AppConfig{
			Config: []string{},
			Status: "Unknown",
			Alert:  "alert-warning",
			Msg:    "This means the config file was not found and will use the default config",
		}
	} else {
		// read the file, display any errors
		content, err := utils.ReadFile(AppConfigFile)
		if err != nil {
			log.Fatal(err)
		}

		config = AppConfig{
			Config: content,
			Status: "Success",
			Alert:  "alert-success",
			Msg:    "This means that the app was able to successfully read the config file.",
		}

	}

	// Display index page from template
	tmpl.Execute(w, config)
}

// appInfo is the route that returns the pod information
func appInfo(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(InfoHtml))
	info := map[string]string{
		"Hostname": os.Getenv("HOSTNAME"),
		"App IP":   os.Getenv("SIMPLE_GO_SERVICE_HOST"),
		"App Port": os.Getenv("SIMPLE_GO_PORT_8080_TCP_PORT"),
	}

	config := Info{
		Info:   info,
		Status: "App Info:",
		Alert:  "alert-primary",
		Msg:    "This information is for debugging only.",
	}

	// Display index page from template
	tmpl.Execute(w, config)
}
