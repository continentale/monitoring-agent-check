## monitoring-agent-check mem

checks mem values of the target

### Synopsis

Checks the return values of the agent from the mem endpoint

```
monitoring-agent-check mem [flags]
```

### Options

```
      --convert string   The unit in which the data is converted. Supported is KB, MB, GB, TB (default "GB")
  -h, --help             help for mem
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
* [monitoring-agent-check mem used](monitoring-agent-check_mem_used.md)	 - A brief description of your command

###### Auto generated by spf13/cobra on 17-Nov-2022
