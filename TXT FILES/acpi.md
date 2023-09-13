# acpi

ACPI, or Advanced Configuration and Power Interface, is a set of specifications that define how the BIOS, operating system, and hardware interact to manage power consumption and thermal management on x86-based computers.

The acpi command has a number of options that can be used to view and control ACPI-related settings. Some of the most commonly used acpi commands are:

- info: This command displays information about the ACPI implementation on your system.
- power: This command displays the current power state of your system.
- sleep: This command puts your system to sleep.
- suspend: This command puts your system into suspend mode.
- hibernate: This command puts your system into hibernation mode.

# help

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
