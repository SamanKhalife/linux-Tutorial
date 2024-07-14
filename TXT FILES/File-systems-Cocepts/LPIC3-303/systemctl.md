# systemctl
`systemctl` is a command-line utility used to interact with `systemd`, the init system and service manager used by many Linux distributions. `systemctl` provides a comprehensive interface for managing services, inspecting the state of the system, and controlling system behavior.

### Overview of `systemctl`

#### Purpose

`systemctl` is used to:

- Manage system services (start, stop, restart, enable, disable).
- Query and change system states.
- Control system boot process.
- Inspect logs and system status.

### Basic Usage

The general syntax for `systemctl` commands is:

```bash
systemctl [options] <command> [name]
```

- `[options]`: Various options to control the behavior of `systemctl`.
- `<command>`: The specific action to perform (e.g., start, stop, status).
- `[name]`: The name of the service, unit, or target to act upon.

### Common `systemctl` Commands

1. **Starting and Stopping Services**

   - **Start a service**:
     ```bash
     systemctl start <service_name>
     ```
     Example:
     ```bash
     systemctl start apache2
     ```

   - **Stop a service**:
     ```bash
     systemctl stop <service_name>
     ```
     Example:
     ```bash
     systemctl stop apache2
     ```

   - **Restart a service**:
     ```bash
     systemctl restart <service_name>
     ```
     Example:
     ```bash
     systemctl restart apache2
     ```

   - **Reload a service** (if the service supports it):
     ```bash
     systemctl reload <service_name>
     ```
     Example:
     ```bash
     systemctl reload apache2
     ```

2. **Enabling and Disabling Services**

   - **Enable a service** (start at boot):
     ```bash
     systemctl enable <service_name>
     ```
     Example:
     ```bash
     systemctl enable apache2
     ```

   - **Disable a service** (do not start at boot):
     ```bash
     systemctl disable <service_name>
     ```
     Example:
     ```bash
     systemctl disable apache2
     ```

   - **Check if a service is enabled**:
     ```bash
     systemctl is-enabled <service_name>
     ```
     Example:
     ```bash
     systemctl is-enabled apache2
     ```

3. **Checking Service Status**

   - **Get the status of a service**:
     ```bash
     systemctl status <service_name>
     ```
     Example:
     ```bash
     systemctl status apache2
     ```

4. **Listing Services and Units**

   - **List all services**:
     ```bash
     systemctl list-units --type=service
     ```

   - **List all units (including services, sockets, targets, etc.)**:
     ```bash
     systemctl list-units
     ```

   - **List failed units**:
     ```bash
     systemctl --failed
     ```

5. **Managing System States**

   - **Reboot the system**:
     ```bash
     systemctl reboot
     ```

   - **Power off the system**:
     ```bash
     systemctl poweroff
     ```

   - **Suspend the system**:
     ```bash
     systemctl suspend
     ```

   - **Hibernate the system**:
     ```bash
     systemctl hibernate
     ```

6. **Masking and Unmasking Services**

   - **Mask a service** (prevent it from starting):
     ```bash
     systemctl mask <service_name>
     ```
     Example:
     ```bash
     systemctl mask apache2
     ```

   - **Unmask a service** (allow it to start):
     ```bash
     systemctl unmask <service_name>
     ```
     Example:
     ```bash
     systemctl unmask apache2
     ```

### Advanced Usage

1. **Viewing Logs with `journalctl`**

   - **View logs for a specific service**:
     ```bash
     journalctl -u <service_name>
     ```
     Example:
     ```bash
     journalctl -u apache2
     ```

2. **Editing Unit Files**

   - **Edit a service unit file**:
     ```bash
     systemctl edit <service_name>
     ```

   - **Reload systemd manager configuration after editing unit files**:
     ```bash
     systemctl daemon-reload
     ```

### Conclusion

`systemctl` is an essential tool for managing and controlling services and system states on a Linux system running `systemd`. It provides powerful capabilities for administrators to maintain and troubleshoot system services effectively.
