# auditd.conf
The `auditd.conf` file is the main configuration file for the Linux Audit Framework (`auditd`). It controls various global settings for the audit daemon (`auditd`) and influences how audit events are collected, stored, and managed. Here’s an overview of what you can typically find in `auditd.conf` and its significance:

### Purpose of `auditd.conf`

The `auditd.conf` file is designed to:
- Define global configuration settings for `auditd`, including log storage, rotation policies, and administrative notifications.
- Specify the behavior of the audit daemon (`auditd`) regarding event collection, buffering, and storage.
- Configure system-wide audit policies and defaults for audit rules.

### Common Configuration Options

1. **Log File Configuration**:
   - **log_file**: Specifies the path where audit logs are stored.
   - **max_log_file**: Sets the maximum size of each audit log file before rotation.
   - **num_logs**: Defines the number of audit log files to retain (rotation count).

2. **Disk Space Management**:
   - **space_left**: Specifies the minimum free space required on the audit log partition before discarding old logs.
   - **space_left_action**: Defines the action to take (e.g., email notification) when disk space runs low.

3. **Buffering and Flush Behavior**:
   - **max_log_file_action**: Action taken when a log file reaches its maximum size (`ROTATE`, `IGNORE`, `SYSLOG`, `SUSPEND`, etc.).
   - **flush**: Controls how often buffered audit records are written to disk (`INCREMENTAL`, `INCREMENTAL_ASYNC`, `SYNC`, `ASYNC`).

4. **System-Wide Settings**:
   - **priority_boost**: Sets the audit daemon priority boost.
   - **freq**: Frequency of log file flushing.
   - **num_logs**: Number of log files to keep.

5. **Admin Notifications**:
   - **admin_space_left**: Notifies admins when the disk space threshold (`space_left`) is reached.
   - **admin_space_left_action**: Action taken when the disk space threshold is reached (`SUSPEND`, `ROTATE`, `SYSLOG`, etc.).

### Example `auditd.conf` Configuration

Here’s a simplified example of what `auditd.conf` might look like:

```plaintext
log_file = /var/log/audit/audit.log
max_log_file = 50
num_logs = 10
space_left = 75
space_left_action = SUSPEND
admin_space_left = 5
admin_space_left_action = SUSPEND
flush = INCREMENTAL
```

### Advanced Settings

- **Configuration Directives**: Explore additional directives like `dispatcher`, `write_logs`, and `name_format` to tailor audit logging and behavior according to specific requirements.

- **Integration with Log Management**: Integrate `auditd` with centralized logging solutions or SIEM (Security Information and Event Management) systems for comprehensive security monitoring.

### Security Considerations

- **File Permissions**: Ensure proper permissions (`chmod`, `chown`) are set on audit log files and directories to restrict access.

- **Monitoring and Review**: Regularly monitor audit logs (`audit.log`) for suspicious activities and security incidents.

### Conclusion

The `auditd.conf` file plays a crucial role in configuring the behavior and settings of the Linux Audit Framework (`auditd`). By understanding and correctly configuring `auditd.conf`, administrators can enhance system security, comply with regulatory requirements, and effectively monitor and manage audit logs.
