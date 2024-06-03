# Systemd

Systemd is a modern init system and system manager that has become the default for many Linux distributions, including Fedora, CentOS, Debian, Ubuntu, and Arch Linux. It replaces the traditional SysVinit system, offering parallel startup of services, improved dependency management, and a unified way to manage system processes.

### Overview of Systemd

#### Key Concepts
- **Unit Files**: Configuration files that define how services, mount points, devices, sockets, and other system resources are managed.
- **Targets**: Group units into a logical unit for booting (similar to runlevels in SysVinit).
- **Journal**: A logging system that collects and manages logs from various sources.
- **Socket Activation**: Starts services on-demand based on incoming network requests.

### Basic Commands

#### Managing Services
- **Start a Service**:
  ```sh
  sudo systemctl start service_name
  ```
- **Stop a Service**:
  ```sh
  sudo systemctl stop service_name
  ```
- **Restart a Service**:
  ```sh
  sudo systemctl restart service_name
  ```
- **Reload Service Configuration**:
  ```sh
  sudo systemctl reload service_name
  ```
- **Check Status of a Service**:
  ```sh
  sudo systemctl status service_name
  ```
- **Enable a Service** (start at boot):
  ```sh
  sudo systemctl enable service_name
  ```
- **Disable a Service**:
  ```sh
  sudo systemctl disable service_name
  ```
- **Check if a Service is Enabled**:
  ```sh
  sudo systemctl is-enabled service_name
  ```

#### System Commands
- **Reboot the System**:
  ```sh
  sudo systemctl reboot
  ```
- **Shut Down the System**:
  ```sh
  sudo systemctl poweroff
  ```
- **Halt the System**:
  ```sh
  sudo systemctl halt
  ```
- **Enter Rescue Mode**:
  ```sh
  sudo systemctl rescue
  ```
- **Enter Emergency Mode**:
  ```sh
  sudo systemctl emergency
  ```

### Unit Files

Unit files describe various system resources and are located in directories like `/etc/systemd/system/` and `/lib/systemd/system/`. They have different types, such as:

- **Service Units** (`.service`): Describe services.
- **Target Units** (`.target`): Group other units for booting.
- **Mount Units** (`.mount`): Define mount points.
- **Socket Units** (`.socket`): Describe network sockets.
- **Timer Units** (`.timer`): Schedule tasks.

#### Example Service Unit File

Here is an example of a simple service unit file for a hypothetical web server (`/etc/systemd/system/example.service`):

```ini
[Unit]
Description=Example Web Service
After=network.target

[Service]
ExecStart=/usr/bin/example-web-service
Restart=always

[Install]
WantedBy=multi-user.target
```

- **[Unit]**: General information about the service, including dependencies.
- **[Service]**: Service-specific details, including the command to start the service.
- **[Install]**: Installation information, defining the target to which this service should be linked.

### Managing Logs with Journal

Systemd uses the `journal` to manage logs. The `journalctl` command allows you to view and query these logs.

#### Basic `journalctl` Commands
- **View All Logs**:
  ```sh
  sudo journalctl
  ```
- **View Logs for a Specific Unit**:
  ```sh
  sudo journalctl -u service_name
  ```
- **View Logs Since Boot**:
  ```sh
  sudo journalctl -b
  ```
- **Follow Real-Time Logs**:
  ```sh
  sudo journalctl -f
  ```
- **View Logs by Time**:
  ```sh
  sudo journalctl --since "2023-06-01" --until "2023-06-02"
  ```

### Targets

Targets in systemd are similar to runlevels in SysVinit. They group units into a logical unit for booting. Common targets include:

- **default.target**: The default target that systemd loads at boot.
- **multi-user.target**: Similar to runlevel 3 in SysVinit (multi-user mode).
- **graphical.target**: Similar to runlevel 5 in SysVinit (multi-user mode with graphical interface).
- **rescue.target**: Similar to single-user mode in SysVinit (maintenance mode).
- **emergency.target**: A minimal environment for emergency repairs.

### Example: Creating a Custom Service

Suppose you have a custom script that you want to run as a service. Here’s how to create a systemd service for it:

1. **Create Your Script** (e.g., `/usr/local/bin/my_script.sh`):
   ```sh
   #!/bin/bash
   echo "My script is running"
   # Your custom commands here
   ```
   Make it executable:
   ```sh
   chmod +x /usr/local/bin/my_script.sh
   ```

2. **Create a Service Unit File** (`/etc/systemd/system/my_script.service`):
   ```ini
   [Unit]
   Description=My Custom Script

   [Service]
   ExecStart=/usr/local/bin/my_script.sh

   [Install]
   WantedBy=multi-user.target
   ```

3. **Reload Systemd Configuration**:
   ```sh
   sudo systemctl daemon-reload
   ```

4. **Start and Enable the Service**:
   ```sh
   sudo systemctl start my_script.service
   sudo systemctl enable my_script.service
   ```

5. **Check the Status**:
   ```sh
   sudo systemctl status my_script.service
   ```

### Conclusion

Systemd is a powerful and flexible system and service manager that provides many advantages over traditional init systems like SysVinit. It supports parallel service startup, on-demand service activation, and centralized management of logs and services. Understanding how to manage services, configure unit files, and use `journalctl` for logging is essential for modern Linux system administration.
