# reboot

There are many tools and commands to shut down and reboot the system, one of these tools is reboot, which can be used to reboot or shut down the system without using the graphical environment.


	To turn off the system using reboot, you can use -p or --poweroff:

	root@Saman:~# reboot -p
	or
	root@Saman:~# reboot --poweroff

	Stopping or halting the system is a state in which all CPU activities are stopped, but the system is still on. This mode can be used for low-level repairs (hardware).

	root@Saman:~# reboot --halt 

	Shutdown or quick system reboot:

	root@Saman:~# reboot -f
	or
	root@Saman:~# reboot --force

	System reboot without sending message to users:

	root@Saman:~# reboot --no-wall
	or
	root@Saman:~# reboot -d


help

```
reboot [OPTIONS...] [ARG]

Reboot the system.

Options:
     --help      Show this help
     --halt      Halt the machine
  -p --poweroff  Switch off the machine
     --reboot    Reboot the machine
  -f --force     Force immediate halt/power-off/reboot
  -w --wtmp-only Don't halt/power-off/reboot, just write wtmp record
  -d --no-wtmp   Don't write wtmp record
     --no-wall   Don't send wall message before halt/power-off/reboot
```
man

```
NAME
       halt, poweroff, reboot - Halt, power-off or reboot the machine

SYNOPSIS
       halt [OPTIONS...]

       poweroff [OPTIONS...]

       reboot [OPTIONS...]

DESCRIPTION
       halt, poweroff, reboot may be used to halt, power-off, or reboot the machine. All three commands take the same options.

OPTIONS
       The following options are understood:

       --help
           Print a short help text and exit.

       --halt
           Halt the machine, regardless of which one of the three commands is invoked.

       -p, --poweroff
           Power-off the machine, regardless of which one of the three commands is invoked.

       --reboot
           Reboot the machine, regardless of which one of the three commands is invoked.

       -f, --force
           Force immediate halt, power-off, reboot. Do not contact the init system.

       -w, --wtmp-only
           Only write wtmp shutdown entry, do not actually halt, power-off, reboot.

       -d, --no-wtmp
           Do not write wtmp shutdown entry.

       -n, --no-sync
           Don't sync hard disks/storage media before halt, power-off, reboot.

       --no-wall
           Do not send wall message before halt, power-off, reboot.

EXIT STATUS
       On success, 0 is returned, a non-zero failure code otherwise.

NOTES
       These commands are implemented in a way that preserves basic compatibility with the original SysV commands.  systemctl(1) verbs
       halt, poweroff, reboot provide the same functionality with some additional features.

       Note that on many SysV systems halt used to be synonymous to poweroff, i.e. both commands would equally result in powering the
       machine off. systemd is more accurate here, and halt results in halting the machine only (leaving power on), and poweroff is
       required to actually power it off.

SEE ALSO
       systemd(1), systemctl(1), shutdown(8), wall(1)

```

