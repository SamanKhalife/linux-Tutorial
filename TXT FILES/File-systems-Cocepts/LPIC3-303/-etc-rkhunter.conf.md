# /etc/rkhunter.conf

The `/etc/rkhunter.conf` file is the main configuration file for Rootkit Hunter (rkhunter), a popular security tool used to detect rootkits and other security vulnerabilities on Unix-based systems. This configuration file allows administrators to customize various settings and behavior of rkhunter to suit their specific security requirements. Here’s an overview of what you might find in `/etc/rkhunter.conf` and its significance:

### Purpose of `/etc/rkhunter.conf`

The primary purpose of `/etc/rkhunter.conf` is to:
- Define and configure the behavior of rkhunter during system scans.
- Specify which tests (checks) should be performed and how they should be executed.
- Customize security options, paths, and command-line options used by rkhunter.

### Typical Configuration Options

1. **Basic Settings**:
   - **UPDATE_MIRRORS**: Specifies whether to update rkhunter’s mirrors for downloading updates.
   - **DB_UPDATE**: Controls automatic updating of the rkhunter database.
   - **AUTO_X_DETECT**: Automatically detects if X is being used (GUI environment).

2. **Logging and Output**:
   - **LOGFILE**: Specifies the file where rkhunter logs its output.
   - **REPORT_EMAIL**: Email address to which reports should be sent.
   - **MAIL-ON-WARNING**: Sends an email if warnings are found during scans.

3. **Scan Settings**:
   - **SCAN_MODE**: Defines the mode of operation (`CRON`, `ONETIME`, `INTERACTIVE`).
   - **SHOW_SKIP**: Displays skipped tests in the rkhunter output.
   - **SUSPSCAN**: Enables the scanning of suspended binaries (Linux).
   - **ALLOWHIDDENDIR**: Specifies hidden directories that should not be scanned.

4. **Paths and Commands**:
   - **TMPDIR**: Temporary directory used during scans.
   - **BINDIR**: Directory where rkhunter commands are located.
   - **SCRIPTDIR**: Directory where rkhunter scripts are located.

5. **Advanced Options**:
   - **AUTO_UPDATE**: Automatically updates rkhunter’s database and program files.
   - **ALLOW_SSH_ROOT_USER**: Allows root user login over SSH (not recommended for security reasons).
   - **DISABLE_TESTS**: Disables specific rkhunter tests from being executed.

### Example Configuration

Here’s a simplified example of what `/etc/rkhunter.conf` might look like:

```plaintext
UPDATE_MIRRORS=1
DB_UPDATE=1
AUTO_X_DETECT=1

LOGFILE=/var/log/rkhunter.log
REPORT_EMAIL=admin@example.com
MAIL-ON-WARNING="yes"

SCAN_MODE=CRON
SHOW_SKIP="yes"
SUSPSCAN=1
ALLOWHIDDENDIR="/dev/.udev"

TMPDIR=/var/lib/rkhunter/tmp
BINDIR=/usr/local/bin
SCRIPTDIR=/usr/local/lib/rkhunter/scripts

AUTO_UPDATE="yes"
ALLOW_SSH_ROOT_USER="no"
DISABLE_TESTS="suspscan hidden_procs"
```

### Security Considerations

- **Configuration Review**: Regularly review and update `/etc/rkhunter.conf` to ensure it reflects current security practices and addresses new threats.
  
- **Log Monitoring**: Monitor the rkhunter log file (`/var/log/rkhunter.log`) for warnings, errors, and suspicious activities.

- **Testing and Validation**: Test rkhunter configurations in a controlled environment to verify effectiveness without causing disruptions.

### Conclusion

`/etc/rkhunter.conf` is a critical component of configuring rkhunter to detect rootkits and other security vulnerabilities on Unix-based systems. By understanding and correctly configuring this file, administrators can enhance their system’s security posture, detect potential threats, and respond promptly to security incidents.
