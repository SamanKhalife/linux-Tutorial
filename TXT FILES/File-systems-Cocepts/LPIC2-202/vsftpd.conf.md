# vsftpd.conf

The `vsftpd.conf` file is the configuration file for the **vsftpd** (Very Secure FTP Daemon) server on Unix-like systems. `vsftpd` is known for its security and performance, and the `vsftpd.conf` file allows administrators to fine-tune various settings to control the behavior of the FTP server.

### Structure and Configuration Options

The `vsftpd.conf` file is a plain text file where each line represents a configuration option. Options are specified in the format:

```
option=value
```

Lines starting with `#` are comments and are ignored by the server. Configuration options can be grouped into sections based on functionality.

### Key Configuration Options

#### General Settings

- **`listen`**: Specifies whether vsftpd should run in standalone mode (i.e., listen on a specific port).
  ```sh
  listen=YES
  ```

- **`listen_ipv6`**: Use this option if you want vsftpd to listen for IPv6 connections. This option is mutually exclusive with `listen`.
  ```sh
  listen_ipv6=YES
  ```

- **`anonymous_enable`**: Enables or disables anonymous FTP access.
  ```sh
  anonymous_enable=NO
  ```

- **`local_enable`**: Allows local users to log in.
  ```sh
  local_enable=YES
  ```

- **`write_enable`**: Allows writing operations (such as uploads) if set to `YES`. It must be enabled if `local_enable` is enabled and you want to allow file modifications.
  ```sh
  write_enable=YES
  ```

- **`chroot_local_user`**: Restricts local users to their home directories. This improves security by preventing users from navigating to directories outside their home.
  ```sh
  chroot_local_user=YES
  ```

#### FTP Access Control

- **`deny_email_enable`**: Denies access to users with specific email addresses. Useful for blocking certain users from logging in.
  ```sh
  deny_email_enable=YES
  ```

- **`banned_email_file`**: Specifies a file containing email addresses to be banned if `deny_email_enable` is enabled.
  ```sh
  banned_email_file=/etc/vsftpd.banned_emails
  ```

- **`userlist_enable`**: Enables the user list functionality, which can restrict access to specified users.
  ```sh
  userlist_enable=YES
  ```

- **`userlist_file`**: Specifies a file containing a list of users who are allowed or denied access based on the `userlist_deny` option.
  ```sh
  userlist_file=/etc/vsftpd.user_list
  ```

#### Security Settings

- **`ssl_enable`**: Enables SSL/TLS for secure FTP connections.
  ```sh
  ssl_enable=YES
  ```

- **`rsa_cert_file`**: Specifies the path to the RSA certificate file.
  ```sh
  rsa_cert_file=/etc/ssl/certs/vsftpd.pem
  ```

- **`rsa_private_key_file`**: Specifies the path to the RSA private key file.
  ```sh
  rsa_private_key_file=/etc/ssl/private/vsftpd.key
  ```

- **`require_ssl_reuse`**: Enforces SSL/TLS session reuse. This can improve performance and security.
  ```sh
  require_ssl_reuse=YES
  ```

#### Performance Tuning

- **`max_clients`**: Specifies the maximum number of clients that can connect simultaneously.
  ```sh
  max_clients=100
  ```

- **`max_per_ip`**: Limits the number of simultaneous connections from a single IP address.
  ```sh
  max_per_ip=10
  ```

- **`local_umask`**: Sets the default umask for local users. This controls the default permissions for newly created files and directories.
  ```sh
  local_umask=022
  ```

### Example Configuration

Here’s a basic example of a `vsftpd.conf` file with common settings:

```sh
# vsftpd configuration file

# Run vsftpd in standalone mode
listen=YES

# Enable local user logins
local_enable=YES

# Allow writing operations
write_enable=YES

# Restrict local users to their home directories
chroot_local_user=YES

# Disable anonymous logins
anonymous_enable=NO

# Enable SSL/TLS for secure connections
ssl_enable=YES
rsa_cert_file=/etc/ssl/certs/vsftpd.pem
rsa_private_key_file=/etc/ssl/private/vsftpd.key

# Limit the number of simultaneous connections
max_clients=100
max_per_ip=10
```

### File Locations

- **Configuration File**: `/etc/vsftpd.conf`
- **SSL Certificates**: Common locations for SSL/TLS certificates are `/etc/ssl/certs/` and `/etc/ssl/private/`.

### Summary

The `vsftpd.conf` file is essential for configuring the vsftpd FTP server. It allows detailed control over how the server operates, including user access, security, and performance settings. Proper configuration of `vsftpd.conf` helps ensure that the FTP server meets security and operational requirements.
