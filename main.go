package main

import (
	"bufio"
	"html/template"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Config []string
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
			Config: []string{},
			Status: "Unknown",
			Alert:  "alert-warning",
			Msg:    "This means the config file was not found and will use the default config",
		}
	} else {
		// read the file, display any errors
		content, err := readFile(AppConfigFile)
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

// readFile returns the contents of a file as a slice of string
func readFile(file string) ([]string, error) {
	// set up the return slice of string
	var s = []string{}

	// open the file, report any errors
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	// close the file when the function exists
	defer f.Close()

	// read the file line by line and convert it into a slice of string
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		//if we get to the end of the file, break out of this loop
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		// append contents into a slice of string
		s = append(s, string(line))

	}

	return s, nil
}
