package icinga

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type Threshold struct {
	Raw    string
	Up     float64
	Down   float64
	Negate bool
}
type Icinga struct {
	ExitCode            int
	ExitWording         string
	PluginOutput        string
	LongPluginOutput    string
	PerfData            string
	hasLongPluginOutput bool
	Warning             Threshold
	Critical            Threshold
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

func (i *Icinga) GetStatus() int {
	return i.ExitCode
}

func (i *Icinga) Evaluate(value float64, problemMessagem, outputOK, outputWarning, outputCritical string) {
	if ((value <= i.Critical.Down || value >= i.Critical.Up) && !i.Critical.Negate) ||
		((value >= i.Critical.Down && value <= i.Critical.Up) && i.Critical.Negate) {
		i.setStatus(2, outputCritical, "")
		i.PluginOutput = outputCritical
	} else if ((value <= i.Warning.Down || value >= i.Warning.Up) && !i.Warning.Negate) ||
		((value >= i.Warning.Down && value <= i.Warning.Up) && i.Warning.Negate) {
		i.setStatus(1, outputWarning, "")
		i.PluginOutput = outputWarning
	} else {
		i.PluginOutput = outputOK
	}

}

func (i *Icinga) MultiEvaluate(value float64, problemMessage, outputOK, outputWarning, outputCritical string) {
	if ((value < i.Critical.Down || value > i.Critical.Up) && !i.Critical.Negate) ||
		((value >= i.Critical.Down && value <= i.Critical.Up) && i.Critical.Negate) {
		i.setStatus(2, outputCritical, "")
		i.addLongPluginOutput(outputCritical)
		i.PluginOutput = problemMessage
	} else if ((value < i.Warning.Down || value > i.Warning.Up) && !i.Warning.Negate) ||
		((value >= i.Warning.Down && value <= i.Warning.Up) && i.Warning.Negate) {
		i.setStatus(1, outputWarning, "")
		i.addLongPluginOutput(outputWarning)
		i.PluginOutput = problemMessage
	} else {
		i.addLongPluginOutput(outputOK)
	}
}

func (i *Icinga) GenerateOutput(perfData bool) {
	output := i.PluginOutput

	if i.PerfData != "" && perfData {
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

func (i *Icinga) addLongPluginOutput(message string) {
	if i.hasLongPluginOutput {
		i.LongPluginOutput += "\n" + message
	} else {
		i.LongPluginOutput = message
		i.hasLongPluginOutput = true
	}

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
		negate := false
		// threshold is more complicated and should be more parsed
		if strings.HasPrefix(thresholdCopy, "@") {
			// is a negate of the threshold. Add negation and delete the token
			thresholdCopy = thresholdCopy[1:]
			negate = true
		}

		// split into parts to parse the value before and after the colon
		parts := strings.Split(thresholdCopy, ":")
		up := 0.
		down := 0.

		if len(parts) == 2 {
			if parts[0] == "~" {
				down = -math.MaxFloat64
			} else {
				firstPart, err := strconv.ParseFloat(parts[0], 64)

				if err != nil {
					log.Fatal("Cannot parse first part of threshold", err)
				}
				down = firstPart
			}

			if parts[1] == "" {
				up = math.MaxFloat64
			} else {
				secondPart, err := strconv.ParseFloat(parts[1], 64)

				if err != nil {
					log.Fatal("Cannot parse second part of threshold", err)
				}
				up = secondPart
			}

		}

		return Threshold{
			Raw:    threshold,
			Up:     up,
			Down:   down,
			Negate: negate,
		}

	} else {
		return Threshold{
			Raw:    threshold,
			Up:     number,
			Down:   0,
			Negate: false,
		}
	}
}

func (i *Icinga) AddPerfData(value float64, data string) {
	i.PerfData += fmt.Sprintf("'%s'=%f;%f;%f;;", data, value, i.Warning.Up, i.Critical.Up)
}
