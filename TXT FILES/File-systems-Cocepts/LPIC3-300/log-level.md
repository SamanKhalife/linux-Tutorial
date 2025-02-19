# log level (or debuglevel) in Samba

The **`log level`** (also referred to as **`debuglevel`**) setting in Samba is used to control the verbosity of the log messages generated by Samba services, including the `smbd`, `nmbd`, and `winbindd` daemons. By adjusting the **log level**, administrators can gather more or less detailed information about Samba's operation, which is useful for troubleshooting, debugging, and performance monitoring.

The `log level` configuration is set in the `smb.conf` file, under the `[global]` section, and can be adjusted to generate logs of varying detail, from basic operational information to in-depth debugging logs.

### Key Features of `log level`:
- **Control verbosity**: You can control how much information is logged. This is particularly useful when troubleshooting or during debugging.
- **Granular control**: Samba allows you to set the log level for different components of the system, enabling fine-grained control over what gets logged.
- **Helps with diagnostics**: When issues occur, increasing the log level can help administrators gather enough information to diagnose problems with Samba.

---

### General Syntax:

In the `smb.conf` file, the `log level` is set under the `[global]` section. The syntax is:

```ini
[global]
log level = <level>
```

### `log level` Values:

The `log level` can be set to different values, with each level increasing the amount of logging information:

1. **`log level = 0`** (No debugging information):
   - This is the default level for minimal logging, where only critical error messages are logged.
   - **Use Case**: Production environments where you do not want excessive log output.

2. **`log level = 1`** (Errors only):
   - Logs error messages only. Useful for when you want to track only serious problems without cluttering the log with non-essential information.
   - **Use Case**: Basic logging for servers where you only need to know when something goes wrong.

3. **`log level = 2`** (Warnings and errors):
   - Logs both warnings and errors, offering more insight into issues that might not be critical but could still require attention.
   - **Use Case**: General use for production environments with some tolerance for log noise.

4. **`log level = 3`** (Informational messages):
   - Logs informational messages, including standard operations and routine activities.
   - **Use Case**: Monitoring Samba's operation and capturing more context about the system's behavior.

5. **`log level = 4`** (Debugging level):
   - Logs debugging messages that help track Samba’s internal processes, like connections and resource handling.
   - **Use Case**: Useful during the initial setup or when troubleshooting specific Samba functions.
  
6. **`log level = 5` to `log level = 10`** (Higher verbosity levels):
   - **Level 5-10**: These levels produce increasingly verbose logging, with level 10 being the most detailed. It includes everything from high-level operations to low-level debug information about the internal workings of Samba, including timing, connection details, and even packet-level details.
   - **Use Case**: When performing deep debugging or diagnosing complex issues. These levels may significantly impact performance due to the large amount of log data generated.

### Example Configuration in `smb.conf`:

To adjust the log level in your Samba configuration, modify the `log level` directive in the `smb.conf` file:

```ini
[global]
log level = 3  # Logs informational messages, warnings, and errors
```

This will capture logs at the "Informational" level, logging basic operational details along with any warnings or errors.

---

### Log Level for Specific Components:

You can fine-tune logging for specific Samba subsystems by setting the log level for individual components. This allows you to gather more detailed logs for a specific service without overloading the log files with unnecessary information.

For example:

```ini
[global]
log level = 1    # Global log level
log level = 3 smb:client   # Set log level for the SMB client to level 3
log level = 2 smb:server   # Set log level for the SMB server to level 2
```

In this case:
- **`smb:client`** will log at level 3 (informational messages).
- **`smb:server`** will log at level 2 (warnings and errors).

This allows you to focus on specific parts of the Samba service, such as debugging client-side or server-side issues.

---

### Use Cases:

1. **Troubleshooting Connection Issues**:
   If you're experiencing issues with connecting to a Samba share, you may want to increase the log level to level 3 or 4 to gather detailed logs on what happens during the connection process. This could help identify problems with authentication, network issues, or share access.

   Example:
   ```ini
   log level = 3  # Logs information about the connection attempts
   ```

2. **Monitoring Samba Performance**:
   If you're performing a performance audit of your Samba server, you might want to use a higher log level (5 or higher) to see detailed logs about file access times, resource consumption, and other internal processes. However, keep in mind that higher log levels can impact system performance.

3. **Security Auditing**:
   Increasing the log level can help during security audits by providing insights into failed login attempts, invalid access, or suspicious activity. 

   Example:
   ```ini
   log level = 5  # Detailed log to capture every operation, useful for security audits
   ```

4. **Debugging Samba Failures**:
   When something goes wrong, increasing the log level to 10 will allow you to capture all the internal workings of Samba, making it easier to pinpoint the exact cause of an issue.

   Example:
   ```ini
   log level = 10  # Maximum verbosity for deep debugging
   ```

---

### Viewing Samba Logs:
After adjusting the log level, you can view the logs by examining the files in the `/var/log/samba/` directory. The most common files are `log.smbd`, `log.nmbd`, `log.winbindd`, and `log.<hostname>` for individual client logs.

To view logs in real-time, you can use:

```bash
tail -f /var/log/samba/log.smbd
```

This will allow you to see the log messages as they are written, which is especially useful when debugging a live system.

---

### Performance Considerations:

- **Log Verbosity and System Load**: Increasing the `log level` to a high value (e.g., 10) can significantly impact the performance of your Samba server, especially in a high-traffic environment. High verbosity can slow down the server and fill up disk space quickly.
- **Use in Production**: It’s recommended to use higher log levels (5 or higher) only temporarily for troubleshooting and to revert to lower levels (e.g., 1 or 2) in production environments to minimize overhead.

---

### Conclusion:

The **`log level`** directive in Samba provides administrators with control over how much information is logged by Samba services. By adjusting the log level, you can gather more detailed information for debugging, troubleshooting, and monitoring Samba’s performance. Care should be taken when using high log levels in production environments, as they can negatively affect performance and disk space.
