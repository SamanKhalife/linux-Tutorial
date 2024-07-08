# SysV Init Overview

SysV init (System V initialization) is one of the traditional init systems used to manage the startup and shutdown of a Unix-like operating system. It relies on configuration files, scripts, and utilities to control the boot process, system services, and runlevels. Three key components in SysV init systems are the `/etc/inittab` file, the `init` command, and the `telinit` command.

### /etc/inittab

The `/etc/inittab` file is the primary configuration file for the SysV init system. It defines the default runlevel, system initialization scripts, and how the system should handle certain key sequences.

#### Structure of `/etc/inittab`

The file consists of lines with the following format:
```
id:runlevels:action:process
```

- **id**: A unique identifier for the entry.
- **runlevels**: The runlevels for which the entry is valid (e.g., `12345`).
- **action**: Specifies what action to take (e.g., `respawn`, `wait`, `once`).
- **process**: The command or script to execute.

#### Example `/etc/inittab`

```bash
# Default runlevel
id:5:initdefault:

# System initialization
si::sysinit:/etc/init.d/rcS

# Runlevel 0 (halt)
l0:0:wait:/etc/init.d/rc 0

# Runlevel 1 (single-user mode)
l1:1:wait:/etc/init.d/rc 1

# Runlevel 2-5 (multi-user mode)
l2:2:wait:/etc/init.d/rc 2
l3:3:wait:/etc/init.d/rc 3
l4:4:wait:/etc/init.d/rc 4
l5:5:wait:/etc/init.d/rc 5

# Runlevel 6 (reboot)
l6:6:wait:/etc/init.d/rc 6

# respawn agetty on tty1-6 (login prompts)
1:2345:respawn:/sbin/agetty tty1 38400
2:2345:respawn:/sbin/agetty tty2 38400
3:2345:respawn:/sbin/agetty tty3 38400
4:2345:respawn:/sbin/agetty tty4 38400
5:2345:respawn:/sbin/agetty tty5 38400
6:2345:respawn:/sbin/agetty tty6 38400
```

### init Command

The `init` command is the parent of all processes on the system. It is the first process started by the kernel during the boot process and remains running until the system is shut down. The init process reads the `/etc/inittab` file to determine the default runlevel and other startup instructions.

#### Common `init` Usage

- **Change Runlevel**: The `init` command can be used to change the current runlevel.
  ```bash
  init 3  # Change to runlevel 3
  init 6  # Reboot the system (runlevel 6)
  init 0  # Halt the system (runlevel 0)
  ```

### telinit Command

The `telinit` command is a symbolic link to the `init` command and is used to communicate with the init process to change runlevels or instruct it to perform other actions.

#### Common `telinit` Usage

- **Change Runlevel**: The `telinit` command can be used similarly to `init` to change the runlevel.
  ```bash
  telinit 3  # Change to runlevel 3
  telinit 6  # Reboot the system (runlevel 6)
  telinit 0  # Halt the system (runlevel 0)
  ```

- **Re-examine /etc/inittab**: Instruct init to re-read the `/etc/inittab` file.
  ```bash
  telinit q
  ```

### Runlevels

Runlevels are predefined states that define what system services are running. Traditional SysV runlevels include:

- **0**: Halt the system.
- **1**: Single-user mode.
- **2-5**: Multi-user modes (distribution-specific).
- **6**: Reboot the system.

Different distributions may have specific uses for each runlevel, but typically:

- **Runlevel 3**: Multi-user mode with text login.
- **Runlevel 5**: Multi-user mode with graphical login (GUI).

### Conclusion

Understanding `/etc/inittab`, `init`, and `telinit` is essential for managing systems using SysV init. These components work together to define how a system starts, which services are running, and how the system responds to changes in runlevels. While many modern distributions have moved to `systemd`, knowing SysV init remains valuable for working with older or certain enterprise systems.
