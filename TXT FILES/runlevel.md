# runlevel

In Linux, a runlevel is a state that the system is in. It determines which services are running and what the system is capable of doing. There are seven standard runlevels, numbered 0 to 6.

* **Runlevel 0:** This is the halt runlevel. The system is completely shut down.
* **Runlevel 1:** This is the single-user runlevel. Only the most essential services are running. This runlevel is used for maintenance and troubleshooting.
* **Runlevel 2:** This is the multiuser, no networking runlevel. All essential services are running, but networking is not. This runlevel is used for systems that do not need to be connected to a network.
* **Runlevel 3:** This is the multiuser, text mode runlevel. All essential services are running, and networking is enabled. This is the default runlevel for most Linux systems.
* **Runlevel 4:** This is the multiuser, graphical mode runlevel. All essential services are running, networking is enabled, and the graphical user interface (GUI) is started. This runlevel is used for systems that need a GUI.
* **Runlevel 5:** This is the reboot runlevel. The system is rebooted.
* **Runlevel 6:** This is the poweroff runlevel. The system is powered off.

The runlevel of the system can be changed by editing the `/etc/inittab` file. The `init` daemon reads this file when the system boots up and starts the services that are listed for the current runlevel.

You can also change the runlevel of the system by using the `init` command. For example, to change the runlevel to 3, you would use the following command:

```
init 3
```

This command will stop all services that are not needed in runlevel 3 and start all services that are needed in runlevel 3.

The runlevel of the system can also be changed by using the `shutdown` command. For example, to change the runlevel to 0 and then halt the system, you would use the following command:

```
shutdown -r now
```

This command will stop all services, change the runlevel to 0, and then halt the system.

The runlevel of the system can be a useful tool for managing the services that are running on your system. By understanding the different runlevels and how to change them, you can control which services are running and what the system is capable of doing.



# help 

```
runlevel [OPTIONS...]

Prints the previous and current runlevel of the init system.

Options:
     --help      Show this help

See the runlevel(8) man page for details.
```

