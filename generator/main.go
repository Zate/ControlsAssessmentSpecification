package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// Control Struct
type Control struct {
	Name        string       `json:"name" yaml:"name"`
	Number      string       `json:"number" yaml:"number"`
	Description string       `json:"description" yaml:"description"`
	Why         string       `json:"why" yaml:"why"`
	Subcontrols []Subcontrol `json:"subcontrols" yaml:"subcontrols"`
}

// Subcontrol struct
type Subcontrol struct {
	Name         string      `json:"name" yaml:"name"`
	Number       string      `json:"number" yaml:"number"`
	Description  string      `json:"description" yaml:"description"`
	Header       Header      `json:"header" yaml:"header"`
	Dependencies []string    `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Inputs       []Input     `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Assumptions  []string    `json:"assumptions,omitempty" yaml:"assumptions,omitempty"`
	Operations   []Operation `json:"operations,omitempty" yaml:"operations,omitempty"`
	Measures     []string    `json:"measures" yaml:"measures"`
	Metrics      []Metric    `json:"metrics" yaml:"metrics"`
}

// Input struct
type Input struct {
	Name   string `json:"name" yaml:"name"`
	Source string `json:"source,omitempty" yaml:"source,omitempty"`
}

// Operation struct
type Operation struct {
	Name  string   `json:"name" yaml:"name"`
	Steps []string `steps:"steps,omitempty" yaml:"steps,omitempty"`
}

// Metric struct
type Metric struct {
	Name        string `json:"name" yaml:"name"`
	Metric      string `json:"metric,omitempty" yaml:"metric,omitempty"`
	Calculation string `json:"calculation,omitempty" yaml:"calculation,omitempty"`
}

// Header struct
type Header struct {
	AssetType            string   `json:"assetType,omitempty" yaml:"assetType,omitempty"`
	SecurityFunction     string   `json:"securityFunction,omitempty" yaml:"securityFunction,omitempty"`
	ImplementationGroups []string `json:"implementationGroups,omitempty" yaml:"implementationGroups,omitempty"`
}

func readIndex(i int) (c Control) {
	rst, err := ioutil.ReadFile("../control-" + strconv.Itoa(i) + "/index.rst")
	if err != nil {
		log.Fatal(err)
	}
	rsts := string(rst)
	c.Name = clean(parseTitle(rsts))
	c.Number = strconv.Itoa(i)
	c.Description = clean(parseControlDescription(rsts))
	c.Why = clean(parseControlWhy(rsts))
	return c
}

func clean(s string) string {
	strings.ReplaceAll(s, "\n", " ")
	return s
}

func readfile(dir string, filename string, i int, z int) (s Subcontrol) {
	s.Number = strconv.Itoa(i) + "." + strconv.Itoa(z)
	rst, err := ioutil.ReadFile("../" + dir + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	rsts := string(rst)
	s.Name = clean(parseTitle(rsts))
	s.Description = clean(parseSlug(rsts))
	s.Header = parseHeader(rsts)
	s.Dependencies = parseDependencies(rsts)
	s.Inputs = parseInputs(rsts)
	s.Operations = parseOperations(rsts)
	s.Measures = parseMeasures(rsts)
	s.Metrics = parseMetrics(rsts, i, z)
	return s
}

func parseHeader(rst string) (h Header) {
	mr := regexp.MustCompile(`\.\. list-table::\s+:header-rows: 1\s+\* - Asset Type\s+ - Security Function\s+- Implementation Groups\s+\* - (.*)\s+- (.*)\s+- (.*)\s+`)
	mm := mr.FindAllStringSubmatch(rst, -1)
	if len(mm) > 0 {
		h.AssetType = strings.Trim(mm[0][1], "\t- *")
		h.SecurityFunction = strings.Trim(mm[0][2], "\t- *")
		h.ImplementationGroups = strings.Split(strings.Trim(mm[0][3], "\t- *"), ", ")
	}
	return h
}

func parseMetrics(rst string, i int, z int) (m []Metric) {
	mn := regexp.MustCompile(`(?m)^(?P<name>[ \S]+)\s+\^{3,}\s+\.\. list-table::\s+\* - \*\*Metric\*\*\s+- \|(?P<metric>[ \S]+)\s+\* - \*\*Calculation\*\*\s+- (?:|\|)(?P<calculation>[ \S]+|)`)
	mmn := mn.FindAllStringSubmatch(rst, -1)
	q := len(mmn)
	m = make([]Metric, q)
	if q > 0 {
		for k := 0; k < len(mmn); k++ {
			if mmn[k][1] != "" {
				m[k].Name = mmn[k][1]
			} else {
				m[k].Name = ""
			}
			if mmn[k][2] != "" {
				m[k].Metric = mmn[k][2]
			} else {
				m[k].Metric = ""
			}
			if mmn[k][3] != "" {
				m[k].Calculation = mmn[k][3]
			} else {
				m[k].Calculation = ""
			}

		}
	}
	return m
}

func parseDependencies(rst string) []string {
	start := "Dependencies\n------------\n"
	end := "\n\nI"
	p := findInString(rst, start, end)
	if checkAssumptions(p) {
		p = strings.Split(p, "\n\nAssumption")[0]
	}
	z := strings.Split(p, "\n")
	for j := range z {
		z[j] = strings.Replace(z[j], "* Sub-control ", "", 1)
	}
	return z
}

func parseInputs(rst string) []Input {
	start := "Inputs\n"
	end := "\n\nO"
	p := findInString(rst, start, end)
	p = strings.Split(p, "-----\n")[1]
	if checkAssumptions(p) {
		p = strings.Split(p, "\n\nAssumption")[0]
	}
	z := strings.Split(p, "\n")
	m := make([]Input, len(z))
	for j := range z {
		m[j].Name = strings.Replace(z[j], "#. ", "", 1)
	}
	return m
}

func parseOperations(rst string) []Operation {
	start := "Operations\n"
	end := "\n\nM"
	p := findInString(rst, start, end)
	p = strings.Split(p, "-----\n")[1]
	if checkAssumptions(p) {
		p = strings.Split(p, "\n\nAssumption")[0]
	}
	z := strings.Split(p, "\n")
	m := make([]Operation, len(z))
	for j := range z {
		m[j].Name = strings.Replace(z[j], "#. ", "", 1)
	}
	return m
}

func parseMeasures(rst string) []string {
	start := "\nMeasures\n"
	end := "\n\nMetrics"
	p := findInString(rst, start, end)
	p = strings.Split(p, "---\n")[1]
	if checkAssumptions(p) {
		p = strings.Split(p, "\n\nAssumption")[0]
	}
	z := strings.Split(p, "\n")
	for j := range z {
		z[j] = strings.Trim(z[j], "\t- *")
	}
	return z
}

func checkAssumptions(rst string) bool {
	if strings.Contains(rst, "\n\nAssumption") {
		return true
	}
	return false
}

func parseControlDescription(rst string) string {
	start := "=====\n\n"
	end := "\n\n**"
	return findInString(rst, start, end)
}

func parseControlWhy(rst string) string {
	start := "**\n\n"
	end := "\n\n.. toctree::"
	return findInString(rst, start, end)
}

func parseTitle(rst string) string {
	start := ""
	end := "\n====="
	return findInString(rst, start, end)
}

func parseSlug(rst string) string {
	sr := regexp.MustCompile(`(?mi)^[ \S]+\n+={3,}\s+([\S ]+)\s+`)
	sm := sr.FindAllStringSubmatch(rst, 1)
	return sm[0][1]
}

func findInString(value, a, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func genFiles(z int, i int) (fx []string) {
	fx = append(fx, "index.rst")
	for q := 1; q < z; q++ {
		sub := "control-" + strconv.Itoa(i) + "." + strconv.Itoa(q) + ".rst"
		fx = append(fx, sub)
	}
	return fx
}

func main() {
	var C = []Control{}
	for i := 1; i < 21; i++ {
		dir := "control-" + strconv.Itoa(i)
		files, err := ioutil.ReadDir("../" + dir)
		if err != nil {
			log.Fatal(err)
		}
		c := readIndex(i)
		f := genFiles(len(files), i)
		for s, file := range f {
			if strings.HasPrefix(file, dir) {
				x := readfile(dir, file, i, s)
				c.Subcontrols = append(c.Subcontrols, x)
			}
		}
		C = append(C, c)
	}
	// e, err := json.Marshal(C)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(e))
	y, err := yaml.Marshal(C)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(y))
}
