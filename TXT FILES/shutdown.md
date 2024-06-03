# shutdown

The `shutdown` command in Linux is used to bring the system down in a safe and controlled manner. It ensures that all processes are terminated gracefully, filesystems are synchronized, and any remaining data in memory is written to disk before the system is powered off or rebooted.

### Basic Syntax

The basic syntax of the `shutdown` command is as follows:

```sh
shutdown [OPTIONS] [TIME] [MESSAGE]
```

- **OPTIONS**: Options to control the behavior of the shutdown (e.g., `-h`, `-r`).
- **TIME**: When to perform the shutdown (e.g., `now`, `+5`).
- **MESSAGE**: An optional message that is broadcast to all logged-in users.

### Common Options

- **`-h`**: Halt the system after shutdown.
- **`-r`**: Reboot the system after shutdown.
- **`-P`**: Power off the system (similar to `-h` but ensures power off).
- **`-c`**: Cancel a scheduled shutdown.
- **`-k`**: Send a warning message only; do not actually shut down.

### Examples

#### Immediate Shutdown

To shut down the system immediately:

```sh
sudo shutdown -h now
```

#### Scheduled Shutdown

To schedule a shutdown in 10 minutes:

```sh
sudo shutdown -h +10
```

To schedule a shutdown at a specific time (e.g., 22:00 or 10:00 PM):

```sh
sudo shutdown -h 22:00
```

#### Immediate Reboot

To reboot the system immediately:

```sh
sudo shutdown -r now
```

#### Scheduled Reboot

To schedule a reboot in 5 minutes:

```sh
sudo shutdown -r +5
```

#### Cancel a Scheduled Shutdown

To cancel a scheduled shutdown or reboot:

```sh
sudo shutdown -c
```

#### Send a Warning Message

To warn users about a scheduled shutdown without actually shutting down:

```sh
sudo shutdown -k +15 "System will go down for maintenance in 15 minutes."
```

### Using `systemctl` for Shutdown and Reboot

With systems that use `systemd`, you can also use the `systemctl` command to perform shutdown and reboot operations. This method is often preferred on modern systems.

#### Shutdown

To shut down the system immediately:

```sh
sudo systemctl poweroff
```

#### Reboot

To reboot the system immediately:

```sh
sudo systemctl reboot
```

#### Halt

To halt the system (stop all processes but do not power off):

```sh
sudo systemctl halt
```

#### Suspend

To suspend the system:

```sh
sudo systemctl suspend
```

#### Hybrid Sleep

To put the system into hybrid sleep (combination of suspend and hibernate):

```sh
sudo systemctl hybrid-sleep
```

### Handling User Sessions

The `shutdown` command and `systemctl` can also handle active user sessions. For instance, if users are logged into the system, a warning message can be broadcast to all logged-in users, giving them time to save their work before the system shuts down.

### Custom Broadcast Messages

You can include a custom message with the shutdown command to inform users about the reason for the shutdown or any other pertinent details.

#### Example with Custom Message

```sh
sudo shutdown -h +30 "The system will shut down in 30 minutes for scheduled maintenance. Please save your work."
```

### Conclusion

The `shutdown` command is a vital tool for system administrators to safely and properly bring a Linux system down. Understanding the various options and how to use them effectively ensures that shutdowns and reboots are handled gracefully, minimizing potential data loss and disruption.



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
