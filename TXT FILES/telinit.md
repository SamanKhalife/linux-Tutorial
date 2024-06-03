# telinit

The `telinit` command is used to change the runlevel of a system in Unix-like operating systems that use the SysVinit init system. Runlevels are preset operating states that define what system services are available. Each runlevel has its own configuration and purpose.

### Overview of Runlevels

- **0**: Halt the system
- **1**: Single-user mode (maintenance or emergency mode)
- **2**: Multi-user mode without networking (varies by distribution)
- **3**: Multi-user mode with networking (text mode)
- **4**: Undefined (can be user-defined)
- **5**: Multi-user mode with networking and graphical interface (default for many distributions)
- **6**: Reboot the system

### Usage of `telinit`

The `telinit` command is a symbolic link to the `init` command. It sends signals to the `init` process to change the system runlevel. When you use `telinit`, you are essentially instructing the `init` process to transition the system to a specified runlevel.

#### Basic Syntax

```sh
telinit [runlevel]
```

### Example Commands

#### Changing to Runlevel 1 (Single-User Mode)

Single-user mode is typically used for maintenance purposes, where you need minimal services running:

```sh
sudo telinit 1
```

#### Changing to Runlevel 3 (Multi-User Mode without GUI)

Runlevel 3 is used for full multi-user mode with networking but without a graphical interface:

```sh
sudo telinit 3
```

#### Changing to Runlevel 5 (Multi-User Mode with GUI)

Runlevel 5 is used for full multi-user mode with networking and a graphical user interface:

```sh
sudo telinit 5
```

#### Rebooting the System (Runlevel 6)

To reboot the system:

```sh
sudo telinit 6
```

#### Halting the System (Runlevel 0)

To halt the system:

```sh
sudo telinit 0
```

### Checking Current Runlevel

To check the current runlevel of the system, use the `runlevel` command:

```sh
runlevel
```

This command outputs two characters: the previous and the current runlevel. For example:

```plaintext
N 5
```

This means there was no previous runlevel (`N`) and the current runlevel is 5.

### Transitioning from SysVinit to Systemd

In systems using `systemd`, the concept of runlevels is replaced by targets, which offer more flexibility. However, `systemd` maintains compatibility with traditional runlevels.

#### Equivalent Commands in `systemd`

To change the runlevel (target) in `systemd`, use the `systemctl isolate` command:

- **Single-User Mode** (rescue.target):
  ```sh
  sudo systemctl isolate rescue.target
  ```

- **Multi-User Mode without GUI** (multi-user.target):
  ```sh
  sudo systemctl isolate multi-user.target
  ```

- **Multi-User Mode with GUI** (graphical.target):
  ```sh
  sudo systemctl isolate graphical.target
  ```

- **Reboot the System**:
  ```sh
  sudo systemctl reboot
  ```

- **Halt the System**:
  ```sh
  sudo systemctl poweroff
  ```

### Conclusion

The `telinit` command is an essential tool for managing the runlevel of a Unix-like system using SysVinit. Understanding how to change runlevels is crucial for system maintenance and administration. With the transition to `systemd`, similar functionalities are maintained through targets, providing a more modern and flexible approach to system state management.

