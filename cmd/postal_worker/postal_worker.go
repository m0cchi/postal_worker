package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/m0cchi/postal_worker/config"
	"github.com/m0cchi/postal_worker/model"
	"github.com/m0cchi/postal_worker/module"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
	"regexp"
)

var modules map[string]module.PostalModule

func init() {
	modules = make(map[string]module.PostalModule)

}

// HandleAPI is sugoi
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}

	// Unmarshal
	var postalMatter model.PostalMatter
	err = json.Unmarshal(b, &postalMatter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}

	to := postalMatter.To
	for _, t := range to {
		m, ok := modules[t.ModuleName]
		if ok {
			err := m.Exec(postalMatter, t)
			if err != nil {
				http.Error(w, err.Error(), 500)
				log.Println(err)
				return
			}
		}
	}

	err = json.NewEncoder(w).Encode(postalMatter)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println(err)
		return
	}
}

func loadModule(modulePath string, conf *config.Config) error {
	log.Println("load ", modulePath)
	p, err := plugin.Open(modulePath)
	if err != nil {
		log.Println(err)
		return err
	}
	m, err := p.Lookup("Module")
	if err != nil {
		log.Println(err)
		return err
	}
	castM := m.(module.PostalModule)
	err = castM.Init(conf)
	if err != nil {
		log.Println(err)
		return err
	}
	name := castM.GetModuleName()
	modules[name] = castM
	return nil
}

func initModules(modulesDirPath string, conf *config.Config) error {
	files, err := ioutil.ReadDir(modulesDirPath)
	if err != nil {
		return err
	}
	reSoFile := regexp.MustCompile(".+\\.so$")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if reSoFile.FindString(file.Name()) != "" {
			modulePath := filepath.Join(modulesDirPath, file.Name())
			loadModule(modulePath, conf)
		}
	}
	return nil
}

func main() {
	log.Println("init postal_worker")
	configPath := flag.String("c", "", "config path")
	flag.Parse()

	conf, err := config.NewConfig(*configPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = conf.Validate()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = initModules(conf.Module.Dir, conf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("config path: %v", *configPath)
	log.Println("start postal workerd")

	http.HandleFunc("/api/register", HandleAPI)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port), nil))
}
