# /var/log/samba/
The **`/var/log/samba/`** directory is the default location where Samba logs its activities, including error messages, system information, and various Samba service logs. This directory contains log files that help system administrators monitor and troubleshoot Samba servers, share access, authentication issues, and other Samba-related services.

### Key Log Files in `/var/log/samba/`:

1. **`log.smbd`**:
   - This is the main log file for the **Samba daemon (smbd)**, which handles file and printer sharing.
   - **Contents**: It includes logs related to Samba file sharing operations, such as access to shared files, authentication events, and errors.
   - **Usage**: Admins can inspect this file to identify issues like failed login attempts, permission errors, or file access problems.

   Example:
   ```bash
   tail -f /var/log/samba/log.smbd
   ```

2. **`log.nmbd`**:
   - This file logs events related to the **NetBIOS Name Service (nmbd)**, which is responsible for name resolution and browsing services in Samba.
   - **Contents**: Includes information on name registration, resolution requests, and responses, and general network browsing events.
   - **Usage**: Useful for troubleshooting name resolution issues or understanding how Samba is interacting with networked Windows machines.

   Example:
   ```bash
   tail -f /var/log/samba/log.nmbd
   ```

3. **`log.winbindd`**:
   - This file logs events for the **Winbind** daemon, which is used to integrate Unix-like systems with Windows-based networks, especially for user and group mapping.
   - **Contents**: Contains logs for user authentication, domain membership status, and information about the communication between Samba and Active Directory.
   - **Usage**: Critical for troubleshooting issues related to domain integration, user authentication, and group membership.

   Example:
   ```bash
   tail -f /var/log/samba/log.winbindd
   ```

4. **`log.<hostname>`**:
   - Samba also creates individual log files for each client that connects to the server. These files are named `log.<hostname>`, where `<hostname>` corresponds to the client's machine name or IP address.
   - **Contents**: These logs contain information about specific connections from clients, including authentication attempts, file access, and any errors related to the individual client.
   - **Usage**: Useful for tracking activity or troubleshooting issues with specific clients.

   Example:
   ```bash
   tail -f /var/log/samba/log.<hostname>
   ```

5. **`log.samba`**:
   - Some Samba installations might have this file, which logs general information about the server, such as startup messages, process tracking, and performance statistics.
   - **Usage**: Typically used for general logging and information, especially during troubleshooting.

   Example:
   ```bash
   tail -f /var/log/samba/log.samba
   ```

6. **`log.level`**:
   - The level of detail in Samba logs is determined by the `log level` setting in `smb.conf`. You can configure the level of logging to capture different amounts of detail, ranging from errors only to verbose debugging information.
   - **Usage**: If you're encountering issues, you might want to increase the `log level` temporarily to get more detailed logs. This can be done in the `smb.conf` file under the `[global]` section:

     ```ini
     log level = 3
     ```

   After modifying the log level, check the logs for more detailed information about the problem.

### Managing Samba Logs:

1. **Setting Log Level**:
   - To control the verbosity of logs, modify the `log level` in the `smb.conf` file. For example:
   
     ```ini
     [global]
     log level = 2
     ```

   This setting will log more detailed information for Samba operations. Higher values (up to 10) will produce more detailed logs, but may impact performance.

2. **Rotating Logs**:
   - Samba logs can become large over time, especially in a production environment. To manage this, you can configure **log rotation** to archive and compress older logs periodically. 
   - **Logrotate** is typically used for this purpose. The configuration for Samba log rotation can be found in `/etc/logrotate.d/samba`.

   Example `/etc/logrotate.d/samba` configuration:

   ```bash
   /var/log/samba/log.* {
       weekly
       rotate 4
       compress
       missingok
       notifempty
   }
   ```

3. **Clearing Logs**:
   - If logs grow too large or you need to reset them, you can clear the logs manually:
   
     ```bash
     > /var/log/samba/log.smbd
     > /var/log/samba/log.nmbd
     ```

### Example of Troubleshooting with Samba Logs:

1. **Failed Authentication**:
   If users are unable to authenticate to the Samba server, you can examine the `log.smbd` and `log.<hostname>` files to identify the cause. Look for messages that indicate incorrect passwords or authentication errors.

   Example output from `log.smbd`:
   ```
   [2021/05/01 14:35:30.123456,  0] smbd/negotiation.c:477(reject_dos_request)
   reject_dos_request: failed to authenticate user: user1
   ```

2. **Access Denied on File Share**:
   If a user is denied access to a share, the issue might be with permissions or share configuration. Check the logs for any permission errors in `log.smbd` and `log.<hostname>`.

   Example:
   ```
   [2021/05/01 14:35:32.789012,  1] smbd/service.c:1326(make_connection_snum)
   user1 denied access to share Documents (user1).
   ```

3. **Name Resolution Problems**:
   If there are issues with name resolution (such as clients not being able to find the Samba server), check the `log.nmbd` file for error messages related to name resolution or browsing.

   Example:
   ```
   [2021/05/01 14:35:35.987654,  0] nmbd/nmbd_namequery.c:193(query_name_response)
   query_name_response: No name server found for domain 'MYDOMAIN'
   ```

---

### Conclusion:
The **`/var/log/samba/`** directory is the key location for Samba’s logs. It contains various files that help administrators troubleshoot and monitor Samba's operations, such as file sharing, user authentication, and NetBIOS name resolution. By reviewing these logs, administrators can identify and resolve issues with Samba services, improve performance, and ensure that the Samba server is functioning as expected.
