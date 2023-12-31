package main

import (
	"encoding/json"
	"nuga/internal/slices"
	"nuga/pkg/device"
	"nuga/pkg/features"
	"nuga/pkg/features/light"
	"nuga/pkg/hid"
	"os"
)

func collectTemplate() (*features.SimulationTemplate, error) {
	err := hid.Init()
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
	info, err := handle.GetInfo()
	if err != nil {
		return nil, err
	}
	colors, err := f.GetRawColors()
	if err != nil {
		return nil, err
	}
	return &features.SimulationTemplate{
		Name:     info.Name,
		Firmware: device.FormatVersion(info.Firmware),
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
