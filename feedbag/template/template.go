package template

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Template struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Event     string `json:"event"`
	Condition string `json:"condition"`
	Size      string `json:"size"`
	Template  string `json:"template"`
}

func ParseDir(templatesDir string) ([]*Template, error) {
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, err
	}

	templates := []*Template{}

	for _, fileInfo := range files {
		if strings.HasSuffix(fileInfo.Name(), ".tmpl") {
			filePath := fmt.Sprintf("%s/%s", templatesDir, fileInfo.Name())
			log.Println("Parsing template:", filePath)

			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Println("Error opening template file:", filePath)
				continue
			}

			template, err := parseTemplate(string(content))
			if err != nil {
				log.Println(err.Error())
			} else {
				templates = append(templates, template)
			}
		}
	}

	return templates, nil
}

// func parseTemplate(dir string, file os.FileInfo) (*Template, error) {
func parseTemplate(tmplString string) (*Template, error) {
	var err error
	template := &Template{}

	configBlock := false
	configComplete := false

	lines := strings.Split(tmplString, "\n")
	for i, l := range lines {
		line := strings.Trim(l, " \t\n\r")

		if i == 0 {
			if !strings.HasPrefix(line, "###") {
				return nil, errors.New("First line of template must be opening `###` marker")
			}
			configBlock = true
			continue
		}

		if configBlock {
			if strings.HasPrefix(line, "###") {
				configBlock = false
				configComplete = true
				continue
			}

			// Read in the config line key: value pair and set it on template
			data := strings.SplitN(line, ":", 2)
			trimChars := " \"'"
			switch strings.Trim(data[0], trimChars) {
			case "id":
				template.Id = strings.Trim(data[1], trimChars)
			case "name":
				template.Name = strings.Trim(data[1], trimChars)
			case "event":
				template.Event = strings.Trim(data[1], trimChars)
			case "condition":
				template.Condition = strings.Trim(data[1], trimChars)
			case "size":
				template.Size = strings.Trim(data[1], trimChars)
			}
			continue
		}

		// Config block is done, so take the rest of the lines as the Template and exit
		template.Template = strings.Join(lines[i:], "\n")
		break
	}

	if !configComplete {
		return nil, errors.New(`Invalid config block. No closing '###' marker`)
	}

	if len(template.Id) == 0 {
		return nil, errors.New(`Template is missing "id" field`)
	}

	if len(template.Name) == 0 {
		return nil, errors.New(`Template is missing "name" field`)
	}

	if len(template.Event) == 0 {
		return nil, errors.New(`Template is missing "event" field`)
	}

	return template, err
}
