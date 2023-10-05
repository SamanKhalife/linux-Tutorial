# acpid

The acpid daemon is a service that manages power management on Linux systems that support ACPI. It is responsible for handling ACPI events, such as lid close, power button press, and battery low, and taking appropriate action, such as putting the system to sleep or suspending it.


The acpid daemon is started automatically when the system boots. It can be stopped and started using the following commands:

```
systemctl stop acpid
systemctl start acpid
```

The acpid daemon can be configured using the /etc/acpi/events file. This file contains a list of ACPI events that the acpid daemon should handle, as well as the action that the daemon should take when each event occurs.


# help 

```
Usage: acpid [OPTIONS]
  -c, --confdir      Set the configuration directory.
  -C, --clientmax    Set the limit on non-root socket connections.
  -d, --debug        Increase debugging level.
  -e, --eventfile    Use the specified file for events.
  -f, --foreground   Run in the foreground.
  -l, --logevents    Log all event activity.
  -g, --socketgroup  Set the group on the socket file.
  -m, --socketmode   Set the permissions on the socket file.
  -s, --socketfile   Use the specified socket file.
  -S, --nosocket     Do not listen on a UNIX socket (overrides -s).
  -p, --pidfile      Use the specified PID file.
  -L, --lockfile     Use the specified lockfile to stop processing.
  -n, --netlink      Force netlink/input layer mode. (overrides -e)
  -r, --dropaction   Define the pseudo-action to drop an event.
  -t, --tpmutefix    Fixup for ThinkPad mute-repeat behaviour.
  -v, --version      Print version information.
  -h, --help         Print this message.
  ```

