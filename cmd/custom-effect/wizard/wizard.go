package wizard

import (
	"fmt"

	"github.com/mishamyrt/nuga-lib"
	"github.com/mishamyrt/nuga-lib/features/light"
	"github.com/mishamyrt/nuga-lib/layout"
)

// Wizard is a light customization wizard
type Wizard struct {
	initialState *light.Effects
	dev          *nuga.Device
	activeIndex  uint32
	Keys         layout.BacklightTemplate
}

// New creates Wizard instance
func New(dev *nuga.Device) (*Wizard, error) {
	state, err := dev.Features.Light.GetEffects()
	if err != nil {
		return nil, err
	}
	w := &Wizard{
		dev:          dev,
		initialState: state,
		Keys:         make(layout.BacklightTemplate),
		activeIndex:  0,
	}
	return w, nil
}

// Restore restores initial state
func (w *Wizard) Restore() error {
	return w.dev.Features.Light.SetEffects(w.initialState)
}

// Start starts wizard
func (w *Wizard) Start() error {
	customEffect := &light.Effects{
		Backlight: w.initialState.Backlight,
		Sidelight: w.initialState.Sidelight,
		Halo:      w.initialState.Halo,
		Debounce:  w.initialState.Debounce,
	}
	customEffect.Backlight.Mode = &light.BacklightCustom
	customEffect.Halo.Mode = &light.HalolightOff
	customEffect.Sidelight.Mode = &light.SidelightOff
	req := highlightKey(0)
	err := w.dev.Handle.Send(req)
	if err != nil {
		return err
	}
	return w.dev.Features.Light.SetEffects(customEffect)
}

// HandleKey handles key
func (w *Wizard) HandleKey(keyName string) error {
	if keyName == "" {
		return nil
	}
	name := layout.KeyName(keyName)
	code := layout.FindKeyCode(name)
	if code == 0 {
		return fmt.Errorf("unknown key: %s", keyName)
	}
	w.Keys[name] = w.activeIndex
	return nil
}

// Next sets next state
func (w *Wizard) Next() error {
	w.activeIndex++
	req := highlightKey(w.activeIndex)
	return w.dev.Handle.Send(req)
}
