package templates

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

type Postblog struct {
	BuildTime time.Time
	Asset     struct {
		Cover   string
		Favicon string
	}
	Site struct {
		Name        string
		Author      string
		Description string
		Link        string
		JSScript    string `yaml:"script_js"`
	}
	Navigation []struct {
		Name string
		Link string
	}
	Theme struct {
		Primary string
		Info    string
		Link    string
		Success string
		Warning string
		Danger  string
	}
}

type Page struct {
	Name       string
	Title      string
	SubTitle   string
	Link       string
	Background string
	Cards      []struct {
		Size  int16
		Title string
		Text  string
	}
}

func LoadPostblog(path string) Postblog {
	config := Postblog{}
	config.BuildTime = time.Now()
	template, _ := ioutil.ReadFile(path)
	_ = yaml.Unmarshal(template, &config)
	return config
}

func LoadPage(path string) Page {
	page := Page{}
	template, _ := ioutil.ReadFile(path)
	_ = yaml.Unmarshal(template, &page)
	return page
}

func (p Postblog) Save(path string) {
	postblogYaml, _ := yaml.Marshal(p)
	_ = ioutil.WriteFile(path, postblogYaml, 0777)
}

func (p Page) Save(path string) {
	pageYaml, _ := yaml.Marshal(p)
	_ = ioutil.WriteFile(path, pageYaml, 0777)
}
