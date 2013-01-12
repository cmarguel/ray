package config

import (
	"ray/geom"
	"ray/light"
	"ray/mmath"
	"ray/shape"
)

type Attribute struct {
	Shapes    []shape.Shape
	Transform *mmath.Transform
	Lights    []light.PointLight
}

type Config struct {
	Translate  geom.Vector3
	Fov        float64
	FilmX      int
	FilmY      int
	Attributes []Attribute
}

func NewConfig() Config {
	config := *new(Config)

	config.Fov = 90.
	config.Translate = geom.NewVector3(0., 0., 0.)
	config.FilmX = 0
	config.FilmY = 0

	config.Attributes = make([]Attribute, 0)

	return config
}

func (c *Config) AddAttribute(a Attribute) {
	c.Attributes = append(c.Attributes, a)
}
