[![Go](https://github.com/Henelik/tricaster/actions/workflows/go.yml/badge.svg)](https://github.com/Henelik/tricaster/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/Henelik/tricaster/branch/master/graph/badge.svg)](https://codecov.io/gh/Henelik/tricaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/Henelik/qtrade-api-go)](https://goreportcard.com/report/github.com/Henelik/qtrade-api-go)
[![License: GPL-2.0](https://img.shields.io/badge/License-GPL2-yellow.svg)](https://opensource.org/licenses/GPL-2.0)

# tricaster

A multi-threaded raytracer written in Go.

<img src="/renders/refraction_test.png" alt="reflections" width="1024"/>

## Features

* Uses backward raytracing
* Phong shading
* Toggleable shadows
* Reflection
* Refraction
* Procedural texture pipeline
* Anti-aliasing
* Can be configured to run on any number of threads
* Scenes can be loaded from YAML

## Planned features

* Allow modifying background color
* Allow use of multiple lights
* Allow shadow support for transparent objects
* Add configurable shadow strength
* Add support for multiple types of lights
* Allow scene lighting with sphere-projected images
* Add Constructive Solid Geometry
