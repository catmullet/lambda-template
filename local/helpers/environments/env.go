package environments

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Environment string

const Development Environment = "development"
const Staging Environment = "staging"
const Production Environment = "production"

var ymlFileLocation = "env.yml"
var fileSearchCount = 0

type Yml struct {
	Development map[string]string `yaml:"development,omitempty"`
	Staging     map[string]string `yaml:"staging,omitempty"`
	Production  map[string]string `yaml:"production,omitempty"`
}

func LoadEnvironmentVariablesFromYml(environment Environment) {
	yml := Yml{}

	data, err := ioutil.ReadFile(ymlFileLocation)

	if err != nil {
		ymlFileLocation = "../" + ymlFileLocation
		LoadEnvironmentVariablesFromYml(environment)
		fileSearchCount++
		if fileSearchCount == 11 {
			panic(err)
		}
	}

	err = yaml.Unmarshal(data, &yml)

	if err != nil {
		panic(err.Error())
	}

	switch environment {
	case Development:
		setEnvironmentVariables(yml.Development)
		break
	case Staging:
		setEnvironmentVariables(yml.Staging)
		break
	case Production:
		setEnvironmentVariables(yml.Production)
		break
	default:

	}
}

func setEnvironmentVariables(variables map[string]string) {
	for key, value := range variables {
		formattedKey := FormatEnvironmentKey(key)
		os.Setenv(formattedKey, value)
		fmt.Println(formattedKey, os.Getenv(formattedKey))
	}
}

func FormatEnvironmentKey(key string) (value string) {
	values := split(key)
	for i, val := range values {
		value += strings.ToUpper(val)
		if i < len(values)-1 {
			value += "_"
		}
	}
	return
}

func split(src string) (entries []string) {

	if !utf8.ValidString(src) {
		return []string{src}
	}
	entries = []string{}
	var runes [][]rune
	lastClass := 0
	class := 0

	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}
	return
}
