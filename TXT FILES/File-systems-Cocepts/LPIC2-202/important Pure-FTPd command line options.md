# Important Pure-FTPd command line options

**Pure-FTPd** is a free, secure, production-quality, and standard-compliant FTP server. It supports a wide array of features and configurations. Below are some of the most important command-line options for running and managing Pure-FTPd:

### Starting and Stopping the Server

- **`-l`**: Specify the list of authentication methods (e.g., `pam`, `puredb`, `mysql`).
  ```sh
  pure-ftpd -l pam
  ```

- **`-A`**: Enables automatic creation of user directories.
  ```sh
  pure-ftpd -A
  ```

- **`-d`**: Enable debugging mode for verbose output, useful for troubleshooting.
  ```sh
  pure-ftpd -d
  ```

- **`-E`**: Disable anonymous logins.
  ```sh
  pure-ftpd -E
  ```

- **`-e`**: Redirect logs to the syslog.
  ```sh
  pure-ftpd -e
  ```

- **`-F`**: Run in the foreground (do not daemonize).
  ```sh
  pure-ftpd -F
  ```

- **`-h`**: Show the help message with available options.
  ```sh
  pure-ftpd -h
  ```

- **`-k`**: Enable FTP over SSL/TLS.
  ```sh
  pure-ftpd -k
  ```

- **`-P`**: Specify the public IP address for the server. Useful for NAT configurations.
  ```sh
  pure-ftpd -P 203.0.113.1
  ```

- **`-p`**: Define the port or port range for the FTP service.
  ```sh
  pure-ftpd -p 21
  ```

- **`-r`**: Enable support for IPv6.
  ```sh
  pure-ftpd -r
  ```

- **`-s`**: Enable the use of a chroot environment for users. This restricts them to their home directories.
  ```sh
  pure-ftpd -s
  ```

- **`-T`**: Enable a custom banner for the FTP server.
  ```sh
  pure-ftpd -T "Welcome to Pure-FTPd!"
  ```

- **`-x`**: Disable IP spoofing protection.
  ```sh
  pure-ftpd -x
  ```

- **`-Y`**: Enable support for FTP over TLS/SSL.
  ```sh
  pure-ftpd -Y
  ```

- **`-y`**: Disable support for FTP over TLS/SSL.
  ```sh
  pure-ftpd -y
  ```

### Configuration Files

- **`/etc/pure-ftpd/pure-ftpd.conf`**: The primary configuration file for Pure-FTPd where most server settings are defined.
  
### Examples

- **Run Pure-FTPd with SSL/TLS Support and in Debug Mode**:
  ```sh
  pure-ftpd -k -d
  ```

- **Run Pure-FTPd in Foreground with Specified Port and Public IP**:
  ```sh
  pure-ftpd -F -p 21 -P 203.0.113.1
  ```

- **Start Pure-FTPd with Chroot Environment and Disable Anonymous Logins**:
  ```sh
  pure-ftpd -s -E
  ```

### Summary

Pure-FTPd offers a variety of command-line options to configure its behavior and features. From enabling SSL/TLS encryption to setting up IPv6 support and controlling the FTP service's behavior, these options allow for fine-tuned management of the FTP server. For production use, it's crucial to configure Pure-FTPd according to your specific needs and security requirements.
