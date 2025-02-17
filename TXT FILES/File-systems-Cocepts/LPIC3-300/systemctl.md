# systemctl
The `systemctl` command is a powerful utility used to interact with **systemd**, the init system and service manager on many Linux distributions. It allows you to manage services, system states, and various system resources.

### Overview of `systemctl`

`systemctl` provides control over **systemd** services, which includes starting, stopping, restarting, enabling, disabling, and checking the status of services. Additionally, it can be used to manage system states (e.g., reboot, shutdown) and inspect logs and configurations.

### General Syntax
```bash
systemctl [OPTIONS] COMMAND [NAME...]
```

Where:
- **COMMAND** is the action you want to perform (e.g., start, stop, status).
- **NAME** is the name of the service, target, or unit you want to manage (e.g., `nginx.service`, `docker`, or `multi-user.target`).

### Common `systemctl` Commands

1. **Starting a Service**:
   - Starts a systemd service.
   ```bash
   systemctl start <service-name>
   ```
   - Example: To start the `apache2` service:
     ```bash
     systemctl start apache2
     ```

2. **Stopping a Service**:
   - Stops a running systemd service.
   ```bash
   systemctl stop <service-name>
   ```
   - Example: To stop the `apache2` service:
     ```bash
     systemctl stop apache2
     ```

3. **Restarting a Service**:
   - Restarts a service by stopping and then starting it again.
   ```bash
   systemctl restart <service-name>
   ```
   - Example: To restart the `apache2` service:
     ```bash
     systemctl restart apache2
     ```

4. **Reloading a Service**:
   - Reloads a service’s configuration without stopping it. Some services support reloading (like web servers).
   ```bash
   systemctl reload <service-name>
   ```
   - Example: To reload the `apache2` service:
     ```bash
     systemctl reload apache2
     ```

5. **Checking the Status of a Service**:
   - Displays the current status of a service, including whether it's running, its PID, and any recent log entries.
   ```bash
   systemctl status <service-name>
   ```
   - Example: To check the status of `apache2`:
     ```bash
     systemctl status apache2
     ```

6. **Enabling a Service to Start at Boot**:
   - Enables a service to start automatically when the system boots.
   ```bash
   systemctl enable <service-name>
   ```
   - Example: To enable `apache2` to start at boot:
     ```bash
     systemctl enable apache2
     ```

7. **Disabling a Service from Starting at Boot**:
   - Disables a service from automatically starting when the system boots.
   ```bash
   systemctl disable <service-name>
   ```
   - Example: To disable `apache2` from starting at boot:
     ```bash
     systemctl disable apache2
     ```

8. **Masking a Service**:
   - Prevents a service from being started, either manually or automatically, by creating a symbolic link to `/dev/null`.
   ```bash
   systemctl mask <service-name>
   ```
   - Example: To mask the `apache2` service:
     ```bash
     systemctl mask apache2
     ```

9. **Unmasking a Service**:
   - Reverses the effect of `mask` by removing the symbolic link.
   ```bash
   systemctl unmask <service-name>
   ```
   - Example: To unmask `apache2`:
     ```bash
     systemctl unmask apache2
     ```

10. **Viewing Logs for a Service**:
    - Displays the logs for a specific service managed by **systemd**.
    ```bash
    journalctl -u <service-name>
    ```
    - Example: To view the logs for `apache2`:
      ```bash
      journalctl -u apache2
      ```

11. **Viewing the System Logs**:
    - Displays the entire system journal, including messages from all services and system events.
    ```bash
    journalctl
    ```

12. **Rebooting the System**:
    - Reboots the entire system.
    ```bash
    systemctl reboot
    ```

13. **Shutting Down the System**:
    - Shuts down the system gracefully.
    ```bash
    systemctl poweroff
    ```

14. **Suspending the System**:
    - Suspends the system (puts it to sleep).
    ```bash
    systemctl suspend
    ```

15. **Viewing All Active Units**:
    - Lists all active systemd units (services, sockets, etc.).
    ```bash
    systemctl list-units
    ```

16. **Viewing All Loaded Units**:
    - Lists all loaded systemd units, including inactive ones.
    ```bash
    systemctl list-units --all
    ```

17. **Checking Systemd Targets**:
    - Lists the systemd targets. Targets represent different states or "runlevels" in systemd. For example, `multi-user.target` is similar to runlevel 3 in traditional SysVinit systems.
    ```bash
    systemctl list-units --type=target
    ```

18. **Switching Runlevels (Changing Targets)**:
    - You can switch between different systemd targets (which are similar to runlevels).
    ```bash
    systemctl isolate <target-name>
    ```
    - Example: To switch to the multi-user target (which is equivalent to runlevel 3):
      ```bash
      systemctl isolate multi-user.target
      ```

19. **Checking System Boot Logs**:
    - You can view the logs related to the system’s boot process.
    ```bash
    journalctl -b
    ```

### Example Usage

#### 1. Start the Apache service:
```bash
systemctl start apache2
```

#### 2. Check the status of a service:
```bash
systemctl status apache2
```

#### 3. Enable a service to start at boot:
```bash
systemctl enable apache2
```

#### 4. Disable a service from starting at boot:
```bash
systemctl disable apache2
```

#### 5. Restart a service after making configuration changes:
```bash
systemctl restart apache2
```

#### 6. Check all active services:
```bash
systemctl list-units --type=service
```

#### 7. Reboot the system:
```bash
systemctl reboot
```

#### 8. View logs for the `nginx` service:
```bash
journalctl -u nginx
```

#### 9. Power off the system:
```bash
systemctl poweroff
```

### Conclusion

The `systemctl` command is an essential tool for managing a system running **systemd**. It provides control over services, targets (runlevels), and logs, enabling administrators to configure, troubleshoot, and maintain their Linux systems effectively. From simple service management to complex system configurations, `systemctl` is the go-to utility for modern Linux administration.
