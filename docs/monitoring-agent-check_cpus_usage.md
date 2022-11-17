## monitoring-agent-check cpus usage

A brief description of your command

### Synopsis

A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

```
monitoring-agent-check cpus usage [flags]
```

### Options

```
  -h, --help   help for usage
```

### Options inherited from parent commands

```
      --critical string   The value on which field the value is checked (default "95")
      --filter string     Defines the filter for the request. A dot (.) Means no filter at all
      --host string       Defines the filter for the request. A dot (.) Means no filter at all (default "localhost")
      --mode string       Defines the filter for the request. A dot (.) Means no filter at all
      --on string         The value on which field the value is checked (default "available")
      --perCPU            
      --perf              Defines if perfData is added to the command (default true)
      --port int          Defines the filter for the request. A dot (.) Means no filter at all (default 20480)
      --secure            Defines the filter for the request. A dot (.) Means no filter at all
      --token string      Defines the filter for the request. A dot (.) Means no filter at all (default ".")
      --verbose           Prints out debug messages for developing
      --warning string    The value on which field the value is checked (default "90")
```

### SEE ALSO

* [monitoring-agent-check cpus](monitoring-agent-check_cpus.md)	 - checks cpu values of the target

###### Auto generated by spf13/cobra on 17-Nov-2022