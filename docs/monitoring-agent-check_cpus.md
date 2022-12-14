## monitoring-agent-check cpus

checks cpu values of the target

### Synopsis

checks the return values of the agent from the cpu endpoint

		example output from the agent with perCPU = false: 
		[
			{
				"cpu": "cpu-total",
				"user": 94.32,
				"system": 76.49,
				"idle": 17205.19,
				"nice": 0.17,
				"iowait": 18.29,
				"irq": 0,
				"softirq": 33.28,
				"steal": 0,
				"guest": 0,
				"guestNice": 0
			}
		]
		
		Now you can check if you have enough cpu resources available with the cpuusage command or check the load on unix systems with the load command

		./monitoring-agent cpus usage --help
		./monitoring-agent cpus load  --help

```
monitoring-agent-check cpus [flags]
```

### Options

```
  -h, --help     help for cpus
      --perCPU   
```

### Options inherited from parent commands

```
      --critical string   The value on which field the value is checked (default "95")
      --filter string     Defines the filter for the request. A dot (.) Means no filter at all
      --host string       Defines the filter for the request. A dot (.) Means no filter at all (default "localhost")
      --mode string       Defines the filter for the request. A dot (.) Means no filter at all
      --on string         The value on which field the value is checked (default "available")
      --perf              Defines if perfData is added to the command (default true)
      --port int          Defines the filter for the request. A dot (.) Means no filter at all (default 20480)
      --secure            Defines the filter for the request. A dot (.) Means no filter at all
      --token string      Defines the filter for the request. A dot (.) Means no filter at all (default ".")
      --verbose           Prints out debug messages for developing
      --warning string    The value on which field the value is checked (default "90")
```

### SEE ALSO

* [monitoring-agent-check](monitoring-agent-check.md)	 - A brief description of your application
* [monitoring-agent-check cpus load](monitoring-agent-check_cpus_load.md)	 - A brief description of your command
* [monitoring-agent-check cpus usage](monitoring-agent-check_cpus_usage.md)	 - A brief description of your command

###### Auto generated by spf13/cobra on 17-Nov-2022
