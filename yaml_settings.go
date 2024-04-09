package mttools

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func LoadYamlSettingFromFile(path string, settings interface{}) error {
	//try to load only if it exists
	if !IsFileExists(path) {
		return errors.New("File is not accessible: " + path)
	}

	yamlFile, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, settings)
	if err != nil {
		return err
	}

	return nil
}

func SaveYamlSettingToFile(path string, comment string, settings interface{}) error {
	node_yaml := &yaml.Node{}

	if err := node_yaml.Encode(settings); err != nil {
		log.Fatalln(err)
	}

	//setting header comment
	node_yaml.HeadComment = strings.ReplaceAll(comment, "\n", "\n# ") +
		"\n# Created on: " + time.Now().Format(time.RFC3339) +
		"\n#\n\n"

	//adding comments
	r := reflect.TypeOf(settings).Elem()

	for _, option_yaml := range node_yaml.Content {
		option_yaml.HeadComment = settingsOptionYamlComment(r, option_yaml.Value)
	}

	// unmarshalling to raw yaml
	file_yaml, err := yaml.Marshal(node_yaml)

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Print(string(file_yaml)); os.Exit(0)

	if err := os.WriteFile(path, file_yaml, 0644); err != nil {
		return err
	}

	return nil
}

func settingsOptionYamlComment(r reflect.Type, yaml_field string) string {
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		if field.Type.Kind() == reflect.Struct {
			//embedded struct, need recursion
			comment := settingsOptionYamlComment(field.Type, yaml_field)

			if comment != "" {
				return comment
			}
		}

		tag := field.Tag.Get("yaml")

		if tag == "" {
			tag = strings.ToLower(r.Field(i).Name)
		} else {
			tag = strings.TrimSpace(strings.Split(tag, ",")[0])
		}

		//fmt.Println(r.Field(i).Name, tag)
		if tag == yaml_field {
			return r.Field(i).Tag.Get("yaml_comment")
		}
	}

	return ""
}

func PrintYamlSettings(settings interface{}) {
	yaml, err := yaml.Marshal(settings)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(yaml))
}
