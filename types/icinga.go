package types

type Icinga struct {
	ExitCode         int
	ExitWording      string
	PluginOutput     string
	LongPluginOutput string
	PerfData         string
}

func NewIcinga() {
	return Icinga{
		ExitCode: 0,
		ExitWording: "OK",
		PluginOutput: "Everything works fine",
	}
}

func (i *Icinga) 
