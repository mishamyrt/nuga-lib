package main

import (
	"encoding/json"
	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/hid"
	"github.com/mishamyrt/nuga-lib/internal/slices"
	"github.com/mishamyrt/nuga-lib/light"
	"os"
)

func collectTemplate() (*nuga.SimulationTemplate, error) {
	err := nuga.Init()
	if err != nil {
		return nil, err
	}
	handle, err := hid.Open()
	if err != nil {
		return nil, err
	}
	f := light.New(handle)
	params, err := f.GetRawEffects()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	colors, err := f.GetRawColors()
	if err != nil {
		return nil, err
	}
	return &nuga.SimulationTemplate{
		Name:     handle.Info.Model,
		Firmware: handle.Info.Firmware,
		Lights: &light.SimulationTemplate{
			Colors: slices.Cast[byte, int](colors),
			Params: slices.Cast[byte, int](params),
		},
	}, nil
}

func dumpTemplate(filePath string) error {
	template, err := collectTemplate()
	if err != nil {
		return err
	}
	content, err := json.Marshal(template)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, []byte(content), 0644)
}
