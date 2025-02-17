# smbcontrol
The `smbcontrol` command is a tool used in **Samba** for sending control messages to the Samba server processes. It allows administrators to communicate with running Samba services (such as `smbd` and `nmbd`) without having to restart them. This can be very useful for performing various administrative tasks, such as reloading configurations, changing settings on the fly, or even sending debugging signals.

### Overview of `smbcontrol`

The `smbcontrol` command allows you to send specific control messages to Samba daemon processes. These messages can instruct the daemons to reload configurations, perform cleanup tasks, or even provide debugging output.

### General Syntax:
```bash
smbcontrol <process> <message> [<parameters>]
```

Where:
- `<process>` is the name of the Samba process to send the message to (e.g., `smbd`, `nmbd`, `winbindd`).
- `<message>` is the type of control message you want to send.
- `[<parameters>]` are any additional parameters required by the specific message type.

### Common Process Names
- `smbd`: The Samba server daemon responsible for file sharing and printer services.
- `nmbd`: The Samba NetBIOS name server daemon responsible for name resolution.
- `winbindd`: The Samba daemon that allows systems to authenticate users from Windows domains.
  
### Relevant Message Types and Their Use Cases

1. **`reload-config`**:
   - **Purpose**: This message causes the Samba daemon to reload its configuration file (`smb.conf`), applying any changes made without needing to restart the service.
   - **Example**:
     ```bash
     smbcontrol smbd reload-config
     ```
   - **Use Case**: You modify the Samba configuration file (`/etc/samba/smb.conf`) and want to apply the changes immediately without restarting the Samba service. This message triggers `smbd` to re-read its configuration file.

2. **`debug level`**:
   - **Purpose**: This message changes the debug level of the Samba daemon. Debug levels control the verbosity of logging, which can be useful for troubleshooting.
   - **Example**:
     ```bash
     smbcontrol smbd debug level 3
     ```
   - **Use Case**: You can dynamically change the debug level of the `smbd` process to obtain more detailed logs for debugging purposes. Level 3 will generate moderate verbosity, which might include additional details useful for diagnosing issues.

3. **`exit`**:
   - **Purpose**: This message causes the specified Samba daemon (e.g., `smbd`) to gracefully exit.
   - **Example**:
     ```bash
     smbcontrol smbd exit
     ```
   - **Use Case**: You can ask a running `smbd` process to shut down gracefully without having to restart the service manually. This is often used when performing system maintenance or updating Samba without needing a full system restart.

4. **`shutdown`**:
   - **Purpose**: This message instructs the Samba daemon to shut down immediately, causing it to stop serving shares and terminate the process.
   - **Example**:
     ```bash
     smbcontrol smbd shutdown
     ```
   - **Use Case**: Similar to `exit`, but it forces an immediate shutdown of the daemon without waiting for ongoing processes to finish.

5. **`reload`**:
   - **Purpose**: This is another command that forces the Samba server to reload configuration files, much like `reload-config`. However, it applies to both the main `smbd` process and any other related processes.
   - **Example**:
     ```bash
     smbcontrol smbd reload
     ```
   - **Use Case**: To apply configuration changes across all processes without restarting the whole service.

6. **`setparms`**:
   - **Purpose**: This command sends a parameter change to a running Samba daemon. It can be used to adjust parameters such as global Samba settings on-the-fly.
   - **Example**:
     ```bash
     smbcontrol smbd setparms "log level = 3"
     ```
   - **Use Case**: Change Samba’s log level without having to modify `smb.conf` directly and restart the service.

7. **`ping`**:
   - **Purpose**: Sends a simple ping to the Samba daemon to check if it is responsive.
   - **Example**:
     ```bash
     smbcontrol smbd ping
     ```
   - **Use Case**: Check if the Samba server is alive and responsive. This can be useful for monitoring or diagnostic scripts.

8. **`testparm`** (via `smbcontrol`):
   - **Purpose**: This message is used to request the Samba server to perform configuration validation, such as checking the current configuration for errors.
   - **Example**:
     ```bash
     smbcontrol smbd testparm
     ```
   - **Use Case**: Before applying a new configuration, you might want to validate the configuration file for syntax errors or incorrect settings.

9. **`reload` (nmbd)**:
   - **Purpose**: Reloads the **NetBIOS name resolution** settings, useful when there have been changes to name resolution parameters.
   - **Example**:
     ```bash
     smbcontrol nmbd reload
     ```
   - **Use Case**: Use this if you've changed any settings related to NetBIOS names, such as in `nmbd.conf` or DNS settings.

10. **`winbindd`** messages:
    - **Purpose**: Various messages for managing Winbind services such as restarting or reloading.
    - **Example**:
      ```bash
      smbcontrol winbindd reload
      ```
    - **Use Case**: This reloads the Winbind configuration, especially helpful when changes have been made to domain settings or other authentication parameters.

### Example Usage

#### Reload Samba Configuration
If you make changes to `/etc/samba/smb.conf` and want to apply them without restarting the service:

```bash
smbcontrol smbd reload-config
```

#### Increase Debug Level for Samba Daemon
If you need to troubleshoot Samba issues and need more detailed logs:

```bash
smbcontrol smbd debug level 3
```

#### Gracefully Shut Down Samba Daemon
If you need to shut down the `smbd` daemon without restarting the entire system:

```bash
smbcontrol smbd exit
```

#### Check if `nmbd` is Responsive
To check if the `nmbd` process (NetBIOS name server) is running and responsive:

```bash
smbcontrol nmbd ping
```

#### Monitor or Adjust Samba Performance
If you're running a production server and need to adjust performance parameters, you can adjust settings dynamically without restarting:

```bash
smbcontrol smbd setparms "log level = 3"
```

### Conclusion

The `smbcontrol` command provides a powerful mechanism for managing and controlling Samba services on the fly. It can be used for tasks ranging from reloading configurations, changing logging levels, monitoring status, and shutting down processes, to troubleshooting and debugging. By using `smbcontrol`, Samba administrators can manage their servers dynamically without needing to stop or restart the services, ensuring minimal disruption.
