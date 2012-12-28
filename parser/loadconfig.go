package parser

import (
	"ray/config"
	"ray/geom"
	"ray/shape"
	"strconv"
	"strings"
)

func LoadConfig(path string) config.Config {
	directives := GetDirectives(path)

	conf := config.NewConfig()
	inAttribute := false
	var attribute *config.Attribute = nil
	for _, dir := range directives {
		if dir.Name == "AttributeBegin" {
			inAttribute = true
			attribute = new(config.Attribute)
			continue
		} else if dir.Name == "AttributeEnd" {
			inAttribute = false
			conf.AddAttribute(*attribute)
			continue
		}
		if inAttribute {
			handleAttribute(dir, attribute)
		} else {
			handleGlobal(dir, &conf)
		}
	}

	return conf
}

func handleAttribute(dir Directive, att *config.Attribute) {
	switch dir.Name {
	case "Shape":
		if contains(dir.Args, "trianglemesh") {
			inds := getIntSlice(dir.Args, "integer indices")
			p := getFloatSlice(dir.Args, "point P")
			att.Shapes = append(att.Shapes, shape.NewMesh(inds, makeVectors(p)))
		}
	}
}

func makeVectors(f []float64) []geom.Vector3 {
	v := make([]geom.Vector3, len(f)/3)
	for i := 0; i < len(f); i += 3 {
		v[i/3].X = f[i]
		v[i/3].Y = f[i+1]
		v[i/3].Z = f[i+2]
	}
	return v
}

func handleGlobal(dir Directive, conf *config.Config) {
	switch dir.Name {
	case "Translate":
		conf.Translate.X, _ = strconv.ParseFloat(dir.Args[0], 64)
		conf.Translate.Y, _ = strconv.ParseFloat(dir.Args[1], 64)
		conf.Translate.Z, _ = strconv.ParseFloat(dir.Args[2], 64)
	case "Camera":
		conf.Fov = getFloat(dir.Args, "float fov")
	}
}

func contains(str []string, s string) bool {
	s = "\"" + s + "\""
	for _, v := range str {
		if s == v {
			return true
		}
	}
	return false
}

func get(args []string, key string) (string, bool) {
	key = "\"" + key + "\""
	for i, v := range args {
		if v == key {
			if i+1 < len(args) {
				return args[i+1], true
			}
		}
	}
	return "", false
}

func getFloat(args []string, key string) float64 {
	arr, _ := get(args, key)
	return floatSlice(arr)[0]
}

func getFloatSlice(args []string, key string) []float64 {
	s, _ := get(args, key)
	return floatSlice(s)
}

func getIntSlice(args []string, key string) []int {
	s, _ := get(args, key)
	return intSlice(s)
}

func intSlice(s string) []int {
	elems := strings.Split(s, " ")
	ints := make([]int, len(elems))
	for i := range ints {
		ints[i], _ = strconv.Atoi(elems[i])
	}
	return ints
}

func floatSlice(s string) []float64 {
	elems := strings.Split(s, " ")
	floats := make([]float64, len(elems))
	for i := range floats {
		floats[i], _ = strconv.ParseFloat(elems[i], 64)
	}
	return floats
}