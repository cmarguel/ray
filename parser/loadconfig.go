package parser

import (
	"container/list"
	"ray/config"
	"ray/geom"
	"ray/light"
	"ray/mmath"
	"ray/shape"
	"strconv"
	"strings"
)

type stack struct {
	transforms *list.List
}

func (s stack) top() mmath.Transform {
	return s.transforms.Front().Value.(mmath.Transform)
}

func (s stack) push(tr mmath.Transform) {
	s.transforms.PushFront(tr)
}

func (s stack) pop() mmath.Transform {
	return s.transforms.Remove(s.transforms.Front()).(mmath.Transform)
}

func (s stack) mod(tr mmath.Transform) {
	s.pop()
	s.push(tr)
}

func (s stack) repl() {
	s.push(s.top())
}

func NewStack() stack {
	return stack{list.New()}
}

func LoadConfig(path string) config.Config {
	directives := GetDirectives(path)
	transforms := NewStack()
	transforms.push(mmath.NewTransform())

	conf := config.NewConfig()
	inAttribute := false
	var attribute *config.Attribute = nil
	for _, dir := range directives {
		if dir.Name == "AttributeBegin" || dir.Name == "TransformBegin" {
			inAttribute = true
			attribute = new(config.Attribute)
			attribute.Transform = nil
			transforms.repl()
			continue
		} else if dir.Name == "AttributeEnd" || dir.Name == "TransformEnd" {
			inAttribute = false
			conf.AddAttribute(*attribute)
			transforms.pop()
			continue
		}
		if inAttribute {
			handleAttribute(dir, attribute, transforms)
		} else {
			handleGlobal(dir, &conf, transforms)
		}
	}

	return conf
}

func handleAttribute(dir Directive, att *config.Attribute, transforms stack) {
	switch dir.Name {
	case "Shape":
		if contains(dir.Args, "trianglemesh") {
			inds := getIntSlice(dir.Args, "integer indices")
			p := getFloatSlice(dir.Args, "point P")
			mesh := shape.NewMesh(inds, makeVectors(p))
			if att.Transform != nil {
				mesh = mesh.Transform(*att.Transform)
			}
			att.Shapes = append(att.Shapes, mesh.Transform(transforms.top()))
		}
	case "Transform":
		t := floatSlice(dir.Args[0])
		m := mmath.NewMatrix4x4(
			t[0], t[1], t[2], t[3],
			t[4], t[5], t[6], t[7],
			t[8], t[9], t[10], t[11],
			t[12], t[13], t[14], t[15])
		tr := mmath.TransformFrom(m)
		att.Transform = &tr
		transforms.mod(tr)
	case "Translate":
		dx, _ := strconv.ParseFloat(dir.Args[0], 64)
		dy, _ := strconv.ParseFloat(dir.Args[1], 64)
		dz, _ := strconv.ParseFloat(dir.Args[2], 64)
		tr := transforms.top().Translate(dx, dy, dz)
		transforms.mod(tr)
	case "LightSource":
		p := getFloatSlice(dir.Args, "point from")
		v := transforms.top().Apply(geom.NewVector3(p[0], p[1], p[2]))
		c := getFloatSlice(dir.Args, "color I")
		l := light.NewPointLight(v.Vals()[0], v.Vals()[1], v.Vals()[2], c[0], c[1], c[2])
		att.Lights = append(att.Lights, l)
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

func handleGlobal(dir Directive, conf *config.Config, transforms stack) {
	switch dir.Name {
	case "Translate":
		conf.Translate.X, _ = strconv.ParseFloat(dir.Args[0], 64)
		conf.Translate.Y, _ = strconv.ParseFloat(dir.Args[1], 64)
		conf.Translate.Z, _ = strconv.ParseFloat(dir.Args[2], 64)
		tr := transforms.top().Translate(conf.Translate.X, conf.Translate.Y, conf.Translate.Z)
		transforms.mod(tr)
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
