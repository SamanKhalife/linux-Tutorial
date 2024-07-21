# Fail2ban

**Fail2ban** is a popular intrusion prevention software tool designed to protect Linux servers from various types of attacks, such as brute-force attacks. It works by monitoring log files for suspicious activity and then dynamically updating firewall rules to block offending IP addresses for a specified duration. This can help prevent unauthorized access and mitigate attacks.

### Key Components and Concepts

1. **Jails**
   - Jails are configurations that define what log files to monitor, what patterns to look for, and how to react to those patterns.
   - Each jail can have its own set of rules and can be customized for different services like SSH, HTTP, or FTP.

2. **Filters**
   - Filters define patterns to search for in log files. They use regular expressions to match entries that indicate suspicious behavior.
   - Each jail uses a filter to identify the activity that should trigger an action.

3. **Actions**
   - Actions specify what should be done when suspicious activity is detected. The most common action is to update firewall rules to block the offending IP address.
   - Actions can also include sending alerts or executing custom scripts.

### Configuration Files

- **Main Configuration File**: `/etc/fail2ban/fail2ban.conf`
  - Contains general settings for Fail2ban, such as logging configuration and the path to the main configuration directory.

- **Jail Configuration File**: `/etc/fail2ban/jail.conf`
  - This file defines the jails and their settings. It includes the default settings and can be overridden by custom configurations in `/etc/fail2ban/jail.d/`.

- **Jail Directory**: `/etc/fail2ban/jail.d/`
  - Contains custom jail configurations that override the settings in `jail.conf`. This is the recommended place for user-specific jail configurations.

- **Filter Directory**: `/etc/fail2ban/filter.d/`
  - Contains filter definitions that are referenced by the jails. Each filter corresponds to a service or log file and defines patterns to match.

- **Action Directory**: `/etc/fail2ban/action.d/`
  - Contains action definitions, which specify what Fail2ban should do when a filter matches. Actions might include updating firewall rules or sending notifications.

### Common Commands

- **Start Fail2ban**:
  ```sh
  sudo systemctl start fail2ban
  ```
  
- **Stop Fail2ban**:
  ```sh
  sudo systemctl stop fail2ban
  ```

- **Restart Fail2ban**:
  ```sh
  sudo systemctl restart fail2ban
  ```

- **Check Fail2ban Status**:
  ```sh
  sudo fail2ban-client status
  ```

- **Check Status of a Specific Jail**:
  ```sh
  sudo fail2ban-client status <jail_name>
  ```

- **Unban an IP Address**:
  ```sh
  sudo fail2ban-client set <jail_name> unbanip <IP_address>
  ```

### Example Configuration

**Jail Configuration (`/etc/fail2ban/jail.local`)**:
```ini
[ssh]
enabled  = true
port     = ssh
logpath  = /var/log/auth.log
backend  = auto
bantime  = 3600
findtime = 600
maxretry = 5
```

- **`enabled`**: Whether the jail is active.
- **`port`**: The port on which the service is running (e.g., `ssh` for SSH).
- **`logpath`**: The path to the log file that Fail2ban will monitor.
- **`backend`**: The log reading method.
- **`bantime`**: Duration (in seconds) that an IP will be banned.
- **`findtime`**: The time window (in seconds) during which the failed attempts are counted.
- **`maxretry`**: The maximum number of failed login attempts before an IP is banned.

**Filter Example (`/etc/fail2ban/filter.d/sshd.conf`)**:
```ini
[Definition]
failregex = ^%(__prefix_line)sFailed password for .+ from <HOST> port \d+ ssh2$
ignoreregex =
```

- **`failregex`**: The regular expression used to identify failed login attempts in the log file.
- **`ignoreregex`**: Regular expressions for lines to ignore.

### Summary

Fail2ban is a versatile and effective tool for protecting Linux servers from unauthorized access and brute-force attacks. By configuring jails, filters, and actions, administrators can customize Fail2ban to secure various services, monitor log files, and take automatic actions to mitigate threats. Proper configuration and regular monitoring are key to maintaining effective protection.
