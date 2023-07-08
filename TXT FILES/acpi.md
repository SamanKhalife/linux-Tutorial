help

```
Usage: acpi [OPTION]...
Shows information from the /proc filesystem, such as battery status or
thermal information.

  -b, --battery            battery information
  -i, --details            show additional details if available:
                             - battery capacity information
                             - temperature trip points
  -a, --ac-adapter         ac adapter information
  -t, --thermal            thermal information
  -c, --cooling            cooling information
  -V, --everything         show every device, overrides above options
  -s, --show-empty         show non-operational devices
  -f, --fahrenheit         use fahrenheit as the temperature unit
  -k, --kelvin             use kelvin as the temperature unit
  -d, --directory <dir>    path to ACPI info (/sys/class resp. /proc/acpi)
  -p, --proc               use old proc interface instead of new sys interface
  -h, --help               display this help and exit
  -v, --version            output version information and exit

By default, acpi displays information on installed system batteries.
Non-operational devices, for example empty battery slots are hidden.
The default unit of temperature is degrees celsius.

Report bugs to Michael Meskes <meskes@debian.org>.
```