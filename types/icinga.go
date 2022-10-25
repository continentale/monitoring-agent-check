package types

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type Threshold struct {
	Raw        string
	Up         float64
	UpString   string
	Down       float64
	DownString string
	Negate     bool
}
type Icinga struct {
	ExitCode         int
	ExitWording      string
	PluginOutput     string
	LongPluginOutput string
	PerfData         string

	Warning  Threshold
	Critical Threshold
}

func NewIcinga(output, warning, critical string) *Icinga {
	icinga := &Icinga{
		ExitCode:     0,
		ExitWording:  "OK",
		PluginOutput: output,
	}
	icinga.Warning = parseThreshold(warning)
	icinga.Critical = parseThreshold(critical)

	return icinga
}

func (i *Icinga) Evaluate(value float64, label, longOutputOK, longOutputWarning, longOutputCritical string) {
	if i.Critical.DownString == "" && i.Critical.UpString == "" {
		// Check for value, because strings are empty
		// Is in the range
		if (value > i.Critical.Down && value < i.Critical.Up && i.Critical.Negate) || // is in the range with negate
			((value <= i.Critical.Down || value >= i.Critical.Up) && !i.Critical.Negate) { // is out the range
			// Is in the range
			i.setStatus(2, label, longOutputCritical)
			return
		}
	}

	if i.Warning.DownString == "" && i.Warning.UpString == "" {
		// Check for value, because strings are empty
		if (value > i.Warning.Down && value < i.Warning.Up && i.Warning.Negate) ||
			(value <= i.Warning.Down && value >= i.Warning.Up && !i.Warning.Negate) {
			// Is in the range
			i.setStatus(1, label, longOutputWarning)
			return
		}
	}

	i.setStatus(0, label, longOutputOK)
}

func (i *Icinga) GenerateOutput() {
	output := i.PluginOutput

	if i.PerfData != "" {
		output = output + " | " + i.PerfData
	}

	fmt.Println(i.ExitWording + " - " + output)
	if i.LongPluginOutput != "" {
		fmt.Println(i.LongPluginOutput)
	}
}

func (i *Icinga) ParseToPerfData(data gjson.Result) {

	data.ForEach(func(key, value gjson.Result) bool {
		i.PerfData += fmt.Sprintf("'%s'=%f;;;;", key.String(), value.Float())
		return true
	})
}

func (i *Icinga) setStatus(code int, output, longOutput string) {
	if code == 3 {
		i.ExitCode = code
		i.PluginOutput = output
		i.ExitWording = i.parseCodeToWord(code)
	}

	if code > i.ExitCode && code != 3 {
		i.ExitCode = code
		i.PluginOutput = output
		i.ExitWording = i.parseCodeToWord(code)
	}

	if longOutput != "" {
		i.LongPluginOutput += longOutput + "\n"
	}

}

func (i *Icinga) parseCodeToWord(code int) string {
	switch code {
	case 0:
		return "OK"
	case 1:
		return "WARNING"
	case 2:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

func parseThreshold(threshold string) Threshold {
	thresholdCopy := threshold
	number, err := strconv.ParseFloat(thresholdCopy, 64)

	if err != nil {
		if strings.Contains(thresholdCopy, ":") {
			// threshold is more complicated
			result := Threshold{
				Raw: threshold,
			}
			if strings.HasPrefix(thresholdCopy, "@") {
				result.Negate = true

				thresholdCopy = threshold[1:]
			}

			splittedString := strings.Split(thresholdCopy, ":")
			down, err := strconv.ParseFloat(splittedString[0], 64)
			if err != nil {
				if splittedString[0] == "~" {
					result.Down = -math.MaxFloat64
				} else {
					result.DownString = splittedString[0]
				}
			} else {
				result.Down = down
			}

			if len(splittedString) > 1 {
				up, err := strconv.ParseFloat(splittedString[1], 64)

				if err != nil {
					result.UpString = splittedString[1]
				} else {
					result.Up = up
				}
			}

			return result
		} else {
			log.Println("Could not parse threshold")
		}
	}

	return Threshold{
		Raw:  threshold,
		Down: 0,
		Up:   number,
	}

}

func (i *Icinga) AddPerfData(value float64, data string) {
	i.PerfData += fmt.Sprintf("'%s'=%f;%f;%f;;", data, value, i.Warning.Up, i.Critical.Up)
}
