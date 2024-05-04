<p align="center">
    <img src="./docs/logo@2x.png" alt="Nuga library Logo" width="150" />
    <h3 align="center">Nuga library</h3>
    <p align="center">Go interface for BYK916 NuPhy® keyboards</p>
    <p align="center">
        <a href="https://github.com/mishamyrt/nuga-lib/actions/workflows/quality-assurance.yaml">
            <img src="https://github.com/mishamyrt/nuga-lib/actions/workflows/quality-assurance.yaml/badge.svg" />
        </a>
        <a href="https://github.com/mishamyrt/nuga-lib/tags">
            <img src="https://img.shields.io/github/v/tag/mishamyrt/nuga-lib?sort=semver" />
        </a>
    </p>
</p>

---

This library provides the ability to control keyboard parameters through the HID USB interface. It is used in [Nuga.app](https://github.com/mishamyrt/nuga-app).

## Usage

First of all, add a library module to your project.

```sh
go get -u github.com/mishamyrt/nuga-lib@latest
```

Use the `nuga.Open()` method to get the keyboard controller. It allows you to control the keyboard. For example, the brightness can be changed.

```go
package main

import (
	"github.com/mishamyrt/nuga-lib"
)

func main() {
	// Setup HID
	nuga.Init()
	defer nuga.Exit()
	// Open connection with keyboard
	device, _ := nuga.Open()
	// Read current effects
	effects, _ := device.Features.Light.GetEffects()
	// Set brightness to 50%
	effects.Backlight.SetBrightness(2)
	// Write effects
	_ = device.Features.Light.SetEffects(effects)
}
```

## Terminology

- Mode — keyboard light mode;
- Effect — combination of color, speed, brightness and light mode.

## Protocol

The library is based on reverse-engineering of the keyboard protocol. The knowledge that was obtained is recorded in the [`docs`](./docs/) folder.

## Trademarks

NuPhy® is a registered trademark of NuPhy Studio. Nuga is an unofficial product and is not affiliated with NuPhy Studio.
