# rndc
`rndc` (Remote Name Daemon Control) is a command-line tool used to manage and control the BIND (Berkeley Internet Name Domain) DNS server. It allows administrators to communicate with and configure the named daemon dynamically. Here’s a detailed guide on `rndc`, its subcommands, usage scenarios, and configuration:
`rndc` is used to send commands to the `named` daemon to perform various administrative tasks. These tasks include reloading configurations, flushing caches, viewing server status, and managing zones. 

### Key Subcommands

1. **Reload Configuration**
   - `rndc reload`: Reloads the configuration file and any new or changed zones.
   - Example:
     ```
     rndc reload
     ```

2. **Flush Cache**
   - `rndc flush`: Flushes the DNS cache.
   - Example:
     ```
     rndc flush
     ```

3. **Query Statistics**
   - `rndc stats`: Dumps statistics to the file specified in the BIND configuration.
   - Example:
     ```
     rndc stats
     ```

4. **Stop and Restart the Server**
   - `rndc stop`: Gracefully stops the `named` daemon.
   - Example:
     ```
     rndc stop
     ```
   - `rndc restart`: Restarts the `named` daemon.
   - Example:
     ```
     rndc restart
     ```

5. **View Server Status**
   - `rndc status`: Displays the current status of the `named` daemon.
   - Example:
     ```
     rndc status
     ```

6. **Reconfiguring Zones**
   - `rndc reconfig`: Reloads the configuration file and new or removed zones, but does not reload unchanged zones.
   - Example:
     ```
     rndc reconfig
     ```

7. **Managing Zones**
   - `rndc freeze <zone>`: Freezes dynamic updates to a zone.
   - Example:
     ```
     rndc freeze example.com
     ```
   - `rndc thaw <zone>`: Thaws a previously frozen zone, allowing updates.
   - Example:
     ```
     rndc thaw example.com
     ```
   - `rndc sync <zone>`: Writes the current zone file to disk.
   - Example:
     ```
     rndc sync example.com
     ```

8. **Trace and Debug**
   - `rndc trace`: Increases the debug level by one.
   - Example:
     ```
     rndc trace
     ```
   - `rndc trace level <level>`: Sets the debug level to a specific value.
   - Example:
     ```
     rndc trace level 3
     ```
   - `rndc notrace`: Sets the debug level to zero.
   - Example:
     ```
     rndc notrace
     ```

### Configuration

To use `rndc`, a key needs to be configured in both the `rndc` and `named` configuration files. The configuration typically involves creating a shared secret key that both `rndc` and `named` use to authenticate each other.

#### Configuring `rndc.conf`

Example `/etc/rndc.conf`:

```conf
options {
    default-server 127.0.0.1;
    default-key "rndc-key";
};

key "rndc-key" {
    algorithm hmac-sha256;
    secret "randomlygeneratedbase64key==";
};
```

#### Configuring `named.conf`

Example `/etc/named.conf`:

```conf
controls {
    inet 127.0.0.1 allow { localhost; } keys { "rndc-key"; };
};

key "rndc-key" {
    algorithm hmac-sha256;
    secret "randomlygeneratedbase64key==";
};
```

#### Generating the Key

Use the `rndc-confgen` tool to generate the key:

```bash
rndc-confgen -a -k rndc-key -b 256
```

This command generates the key and configures the default location of the `rndc` key file.

### Use Cases

1. **Dynamic DNS Management**:
   - Administrators can add, modify, or remove DNS records without restarting the DNS service, improving uptime and flexibility.

2. **Cache Management**:
   - Flushing or dumping the DNS cache helps in troubleshooting and ensuring that stale records do not cause resolution issues.

3. **Monitoring and Maintenance**:
   - Regularly checking the status and statistics of the DNS server helps in maintaining performance and identifying potential issues.

4. **Security and Troubleshooting**:
   - Freezing and thawing zones can be useful during maintenance or investigation of suspicious activities.

### Conclusion

`rndc` is an essential tool for managing BIND DNS servers, offering a wide range of commands for dynamic configuration, monitoring, and troubleshooting. Proper configuration and understanding of its subcommands enable administrators to effectively control and maintain their DNS infrastructure, ensuring efficient and secure DNS operations.




# help 

```
Query or send control commands to the system manager.

Unit Commands:
  list-units [PATTERN...]             List units currently in memory
  list-sockets [PATTERN...]           List socket units currently in memory,
                                      ordered by address
  list-timers [PATTERN...]            List timer units currently in memory,
                                      ordered by next elapse
  is-active PATTERN...                Check whether units are active
  is-failed PATTERN...                Check whether units are failed
  status [PATTERN...|PID...]          Show runtime status of one or more units
  show [PATTERN...|JOB...]            Show properties of one or more
                                      units/jobs or the manager
  cat PATTERN...                      Show files and drop-ins of specified units
  help PATTERN...|PID...              Show manual for one or more units
  list-dependencies [UNIT...]         Recursively show units which are required
                                      or wanted by the units or by which those
                                      units are required or wanted
  start UNIT...                       Start (activate) one or more units
  stop UNIT...                        Stop (deactivate) one or more units
  reload UNIT...                      Reload one or more units
  restart UNIT...                     Start or restart one or more units
  try-restart UNIT...                 Restart one or more units if active
  reload-or-restart UNIT...           Reload one or more units if possible,
                                      otherwise start or restart
  try-reload-or-restart UNIT...       If active, reload one or more units,
                                      if supported, otherwise restart
  isolate UNIT                        Start one unit and stop all others
  kill UNIT...                        Send signal to processes of a unit
  clean UNIT...                       Clean runtime, cache, state, logs or
                                      configuration of unit
  freeze PATTERN...                   Freeze execution of unit processes
  thaw PATTERN...                     Resume execution of a frozen unit
  set-property UNIT PROPERTY=VALUE... Sets one or more properties of a unit
  bind UNIT PATH [PATH]               Bind-mount a path from the host into a
                                      unit's namespace
  mount-image UNIT PATH [PATH [OPTS]] Mount an image from the host into a
                                      unit's namespace
  service-log-level SERVICE [LEVEL]   Get/set logging threshold for service
  service-log-target SERVICE [TARGET] Get/set logging target for service
  reset-failed [PATTERN...]           Reset failed state for all, one, or more
                                      units
Unit File Commands:
  list-unit-files [PATTERN...]        List installed unit files
  enable [UNIT...|PATH...]            Enable one or more unit files
  disable UNIT...                     Disable one or more unit files
  reenable UNIT...                    Reenable one or more unit files
  preset UNIT...                      Enable/disable one or more unit files
                                      based on preset configuration
  preset-all                          Enable/disable all unit files based on
                                      preset configuration
  is-enabled UNIT...                  Check whether unit files are enabled
  mask UNIT...                        Mask one or more units
  unmask UNIT...                      Unmask one or more units
  link PATH...                        Link one or more units files into
                                      the search path
  revert UNIT...                      Revert one or more unit files to vendor
                                      version
  add-wants TARGET UNIT...            Add 'Wants' dependency for the target
                                      on specified one or more units
  add-requires TARGET UNIT...         Add 'Requires' dependency for the target
                                      on specified one or more units
  edit UNIT...                        Edit one or more unit files
  get-default                         Get the name of the default target
  set-default TARGET                  Set the default target

Machine Commands:
  list-machines [PATTERN...]          List local containers and host

Job Commands:
  list-jobs [PATTERN...]              List jobs
  cancel [JOB...]                     Cancel all, one, or more jobs

Environment Commands:
  show-environment                    Dump environment
  set-environment VARIABLE=VALUE...   Set one or more environment variables
  unset-environment VARIABLE...       Unset one or more environment variables
  import-environment VARIABLE...      Import all or some environment variables

Manager State Commands:
  daemon-reload                       Reload systemd manager configuration
  daemon-reexec                       Reexecute systemd manager
  log-level [LEVEL]                   Get/set logging threshold for manager
  log-target [TARGET]                 Get/set logging target for manager
  service-watchdogs [BOOL]            Get/set service watchdog state

System Commands:
  is-system-running                   Check whether system is fully running
  default                             Enter system default mode
  rescue                              Enter system rescue mode
  emergency                           Enter system emergency mode
  halt                                Shut down and halt the system
  poweroff                            Shut down and power-off the system
  reboot                              Shut down and reboot the system
  kexec                               Shut down and reboot the system with kexec
  exit [EXIT_CODE]                    Request user instance or container exit
  switch-root ROOT [INIT]             Change to a different root file system
  suspend                             Suspend the system
  hibernate                           Hibernate the system
  hybrid-sleep                        Hibernate and suspend the system
  suspend-then-hibernate              Suspend the system, wake after a period of
                                      time, and hibernate
Options:
  -h --help              Show this help
     --version           Show package version
     --system            Connect to system manager
     --user              Connect to user service manager
  -H --host=[USER@]HOST  Operate on remote host
  -M --machine=CONTAINER Operate on a local container
  -t --type=TYPE         List units of a particular type
     --state=STATE       List units with particular LOAD or SUB or ACTIVE state
     --failed            Shortcut for --state=failed
  -p --property=NAME     Show only properties by this name
  -P NAME                Equivalent to --value --property=NAME
  -a --all               Show all properties/all units currently in memory,
                         including dead/empty ones. To list all units installed
                         on the system, use 'list-unit-files' instead.
  -l --full              Don't ellipsize unit names on output
  -r --recursive         Show unit list of host and local containers
     --reverse           Show reverse dependencies with 'list-dependencies'
     --with-dependencies Show unit dependencies with 'status', 'cat',
                         'list-units', and 'list-unit-files'.
     --job-mode=MODE     Specify how to deal with already queued jobs, when
                         queueing a new job
  -T --show-transaction  When enqueuing a unit job, show full transaction
     --show-types        When showing sockets, explicitly show their type
     --value             When showing properties, only print the value
     --check-inhibitors=MODE
                         Specify if checking inhibitors before shutting down,
                         sleeping or hibernating
  -i                     Shortcut for --check-inhibitors=no
     --kill-who=WHO      Whom to send signal to
  -s --signal=SIGNAL     Which signal to send
     --what=RESOURCES    Which types of resources to remove
     --now               Start or stop unit after enabling or disabling it
     --dry-run           Only print what would be done
                         Currently supported by verbs: halt, poweroff, reboot,
                             kexec, suspend, hibernate, suspend-then-hibernate,
                             hybrid-sleep, default, rescue, emergency, and exit.
  -q --quiet             Suppress output
     --wait              For (re)start, wait until service stopped again
                         For is-system-running, wait until startup is completed
     --no-block          Do not wait until operation finished
     --no-wall           Don't send wall message before halt/power-off/reboot
     --no-reload         Don't reload daemon after en-/dis-abling unit files
     --legend=BOOL       Enable/disable the legend (column headers and hints)
     --no-pager          Do not pipe output into a pager
systemctl [OPTIONS...] COMMAND ...

Query or send control commands to the system manager.

Unit Commands:
  list-units [PATTERN...]             List units currently in memory
  list-sockets [PATTERN...]           List socket units currently in memory,
                                      ordered by address
  list-timers [PATTERN...]            List timer units currently in memory,
                                      ordered by next elapse
  is-active PATTERN...                Check whether units are active
  is-failed PATTERN...                Check whether units are failed
  status [PATTERN...|PID...]          Show runtime status of one or more units
  show [PATTERN...|JOB...]            Show properties of one or more
                                      units/jobs or the manager
  cat PATTERN...                      Show files and drop-ins of specified units
  help PATTERN...|PID...              Show manual for one or more units
  list-dependencies [UNIT...]         Recursively show units which are required
                                      or wanted by the units or by which those
                                      units are required or wanted
  start UNIT...                       Start (activate) one or more units
  stop UNIT...                        Stop (deactivate) one or more units
  reload UNIT...                      Reload one or more units
  restart UNIT...                     Start or restart one or more units
  try-restart UNIT...                 Restart one or more units if active
  reload-or-restart UNIT...           Reload one or more units if possible,
                                      otherwise start or restart
  try-reload-or-restart UNIT...       If active, reload one or more units,
                                      if supported, otherwise restart
  isolate UNIT                        Start one unit and stop all others
  kill UNIT...                        Send signal to processes of a unit
  clean UNIT...                       Clean runtime, cache, state, logs or
                                      configuration of unit
  freeze PATTERN...                   Freeze execution of unit processes
  thaw PATTERN...                     Resume execution of a frozen unit
  set-property UNIT PROPERTY=VALUE... Sets one or more properties of a unit
  bind UNIT PATH [PATH]               Bind-mount a path from the host into a
                                      unit's namespace
  mount-image UNIT PATH [PATH [OPTS]] Mount an image from the host into a
                                      unit's namespace
  service-log-level SERVICE [LEVEL]   Get/set logging threshold for service
  service-log-target SERVICE [TARGET] Get/set logging target for service
  reset-failed [PATTERN...]           Reset failed state for all, one, or more
                                      units
Unit File Commands:
  list-unit-files [PATTERN...]        List installed unit files
  enable [UNIT...|PATH...]            Enable one or more unit files
  disable UNIT...                     Disable one or more unit files
  reenable UNIT...                    Reenable one or more unit files
  preset UNIT...                      Enable/disable one or more unit files
                                      based on preset configuration
  preset-all                          Enable/disable all unit files based on
                                      preset configuration
  is-enabled UNIT...                  Check whether unit files are enabled
  mask UNIT...                        Mask one or more units
  unmask UNIT...                      Unmask one or more units
  link PATH...                        Link one or more units files into
                                      the search path
  revert UNIT...                      Revert one or more unit files to vendor
                                      version
  add-wants TARGET UNIT...            Add 'Wants' dependency for the target
                                      on specified one or more units
  add-requires TARGET UNIT...         Add 'Requires' dependency for the target
                                      on specified one or more units
  edit UNIT...                        Edit one or more unit files
  get-default                         Get the name of the default target
  set-default TARGET                  Set the default target

Machine Commands:
  list-machines [PATTERN...]          List local containers and host

Job Commands:
  list-jobs [PATTERN...]              List jobs
  cancel [JOB...]                     Cancel all, one, or more jobs

Environment Commands:
  show-environment                    Dump environment
  set-environment VARIABLE=VALUE...   Set one or more environment variables
  unset-environment VARIABLE...       Unset one or more environment variables
  import-environment VARIABLE...      Import all or some environment variables

Manager State Commands:
  daemon-reload                       Reload systemd manager configuration
  daemon-reexec                       Reexecute systemd manager
  log-level [LEVEL]                   Get/set logging threshold for manager
  log-target [TARGET]                 Get/set logging target for manager
  service-watchdogs [BOOL]            Get/set service watchdog state

System Commands:
  is-system-running                   Check whether system is fully running
  default                             Enter system default mode
  rescue                              Enter system rescue mode
  emergency                           Enter system emergency mode
  halt                                Shut down and halt the system
  poweroff                            Shut down and power-off the system
  reboot                              Shut down and reboot the system
  kexec                               Shut down and reboot the system with kexec
  exit [EXIT_CODE]                    Request user instance or container exit
  switch-root ROOT [INIT]             Change to a different root file system
  suspend                             Suspend the system
  hibernate                           Hibernate the system
  hybrid-sleep                        Hibernate and suspend the system
  suspend-then-hibernate              Suspend the system, wake after a period of
                                      time, and hibernate
Options:
  -h --help              Show this help
     --version           Show package version
     --system            Connect to system manager
     --user              Connect to user service manager
  -H --host=[USER@]HOST  Operate on remote host
  -M --machine=CONTAINER Operate on a local container
  -t --type=TYPE         List units of a particular type
     --state=STATE       List units with particular LOAD or SUB or ACTIVE state
     --failed            Shortcut for --state=failed
```



## breakdown

```

```
