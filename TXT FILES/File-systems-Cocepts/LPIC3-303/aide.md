# aide

**AIDE (Advanced Intrusion Detection Environment)** is an open-source host-based intrusion detection system (HIDS) that monitors and detects unauthorized changes to file systems on Unix-like operating systems. It compares current file and directory states against a previously generated database of checksums and attributes to detect any modifications, additions, or deletions. Here’s an overview of AIDE and its key functionalities:

### Purpose of AIDE

The main purpose of AIDE is to:
- Detect and alert administrators to unauthorized changes to critical system files, directories, and configurations.
- Provide integrity checking and monitoring capabilities to detect potential security breaches and compromises.
- Help maintain the integrity and security of system configurations and files over time.

### Key Features and Functionality

1. **File Integrity Checking**: AIDE performs integrity checks by:
   - Generating and maintaining a database (or baseline) of checksums and attributes for files and directories on the system.
   - Periodically comparing current file states with the baseline to detect any changes in file content, permissions, timestamps, and other attributes.

2. **Configuration Flexibility**: AIDE offers flexibility in configuration:
   - Administrators can define which files, directories, or attributes to monitor and which to exclude from monitoring.
   - Customizable policies allow tuning of alerts and notifications based on detected changes.

3. **Reporting and Alerts**: AIDE provides various reporting and alerting mechanisms:
   - Generates reports listing detected changes, including details such as file path, modification type, and timestamp.
   - Supports email notifications and integration with logging systems to alert administrators about suspicious activities.

4. **Integration and Automation**: AIDE can be integrated into system monitoring and management workflows:
   - Supports automated scheduled scans and checks through cron jobs or other scheduling mechanisms.
   - Integrates with system monitoring tools and security incident response processes.

### Basic Usage

Here are some common commands and their usage with AIDE:

- **Initialize AIDE**: Initialize AIDE for the first time to create the initial database (configuration file `/etc/aide.conf` is required):
  ```bash
  sudo aide --init
  ```

- **Update AIDE Database**: Update the AIDE database after system changes or before periodic checks:
  ```bash
  sudo aide --update
  ```

- **Run AIDE Check**: Perform a file system integrity check using AIDE:
  ```bash
  sudo aide --check
  ```

- **Generate AIDE Report**: Generate a report of changes detected by AIDE:
  ```bash
  sudo aide --check > /var/log/aide/aide_report.txt
  ```

### Example Output

Upon running AIDE, you might see output similar to the following in the report:

```
AIDE found differences between database and filesystem!!
Start timestamp: 2023-07-14 15:30:00
Summary:
  Total number of entries: 35000
  Added entries: 1
  Removed entries: 0
  Changed entries: 2
---------------------------------------------------
Added entries:
---------------------------------------------------

changed: /etc/passwd
changed: /etc/shadow
```

### Security Considerations

- **Secure Baseline**: Create the initial AIDE database on a secure and trusted system to establish a reliable baseline for comparison.
  
- **Regular Updates**: Periodically update the AIDE database to reflect legitimate changes and maintain accurate integrity checks.

- **Monitoring**: Monitor AIDE reports and alerts actively to respond promptly to detected changes and potential security incidents.

### Conclusion

AIDE is a powerful tool for detecting unauthorized changes to critical files and directories on Unix-like systems, providing administrators with essential capabilities to monitor and maintain system integrity. By leveraging AIDE effectively, administrators can enhance the security posture of their systems and mitigate potential risks associated with unauthorized modifications.
