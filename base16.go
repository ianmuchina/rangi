package main

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Base16 struct {
	filename string
	Scheme   string `yaml:"scheme" json:"scheme"`
	Type     string `yaml:"type"   json:"type"`
	Author   string `yaml:"author" json:"author"`
	Base00   string `yaml:"base00" json:"base00"`
	Base01   string `yaml:"base01" json:"base01"`
	Base02   string `yaml:"base02" json:"base02"`
	Base03   string `yaml:"base03" json:"base03"`
	Base04   string `yaml:"base04" json:"base04"`
	Base05   string `yaml:"base05" json:"base05"`
	Base06   string `yaml:"base06" json:"base06"`
	Base07   string `yaml:"base07" json:"base07"`
	Base08   string `yaml:"base08" json:"base08"`
	Base09   string `yaml:"base09" json:"base09"`
	Base0A   string `yaml:"base0A" json:"base0A"`
	Base0B   string `yaml:"base0B" json:"base0B"`
	Base0C   string `yaml:"base0C" json:"base0C"`
	Base0D   string `yaml:"base0D" json:"base0D"`
	Base0E   string `yaml:"base0E" json:"base0E"`
	Base0F   string `yaml:"base0F" json:"base0F"`
}

func (b Base16) ToYaml() []byte {
	d, err := yaml.Marshal(&b)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return d
}

func (b Base16) ToYamlString() string {

	return string(b.ToYaml())
}
