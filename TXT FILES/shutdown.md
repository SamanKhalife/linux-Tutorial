



# help 
 
```
shutdown [OPTIONS...] [TIME] [WALL...]

Shut down the system.

Options:
     --help      Show this help
  -H --halt      Halt the machine
  -P --poweroff  Power-off the machine
  -r --reboot    Reboot the machine
  -h             Equivalent to --poweroff, overridden by --halt
  -k             Don't halt/power-off/reboot, just send warnings
     --no-wall   Don't send wall message before halt/power-off/reboot
  -c             Cancel a pending shutdown
     --show      Show pending shutdown

See the shutdown(8) man page for details.
```