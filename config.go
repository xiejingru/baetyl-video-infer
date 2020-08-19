package main

import (
	"github.com/baetyl/baetyl-go/v2/mqtt"
	"github.com/baetyl/baetyl-go/v2/http"
)

// Config the configuration of this video inference module
type Config struct {
	Function http.ClientConfig `yaml:"function,omitempty" json:"function,omitempty"`
	Broker   mqtt.ClientConfig `yaml:"broker" json:"broker"`
	Video    VideoInfo         `yaml:"video" json:"video"`
	Infer    InferInfo         `yaml:"infer" json:"infer"`
	Process  ProcessInfo       `yaml:"process" json:"process"`
}

// VideoInfo the video configuration
type VideoInfo struct {
	URL   string `yaml:"uri" json:"uri" default:"0" validate:"nonzero"`
	Limit struct {
		FPS float64 `yaml:"fps" json:"fps"`
	} `yaml:"limit" json:"limit"`
}

// InferInfo the inference configuration
type InferInfo struct {
	Model   string `yaml:"model" json:"model" validate:"nonzero"`
	Config  string `yaml:"config" json:"config" validate:"nonzero"`
	Backend string `yaml:"backend" json:"backend"`
	Device  string `yaml:"device" json:"device" default:"cpu"`
}

// ProcessInfo the image process configuration
type ProcessInfo struct {
	Before struct {
		Scale float64 `yaml:"scale" json:"scale" default:"1.0"`
		Width int     `yaml:"width" json:"width"`
		Hight int     `yaml:"hight" json:"hight"`
		Mean  struct {
			V1 float64 `yaml:"v1" json:"v1"`
			V2 float64 `yaml:"v2" json:"v2"`
			V3 float64 `yaml:"v3" json:"v3"`
			V4 float64 `yaml:"v4" json:"v4"`
		} `yaml:"mean" json:"mean"`
		SwapRB bool `yaml:"swaprb" json:"swaprb"`
		Crop   bool `yaml:"crop" json:"crop"`
	} `yaml:"before" json:"before"`
	After struct {
		Function struct {
			Name string `yaml:"name" json:"name"`
		} `yaml:"function" json:"function"`
	} `yaml:"after" json:"after"`
}