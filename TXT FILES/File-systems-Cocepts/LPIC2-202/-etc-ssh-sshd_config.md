# /etc/ssh/sshd_config
The **`sshd_config`** file is the main configuration file for the OpenSSH SSH daemon (`sshd`). It controls various aspects of how `sshd` operates, including authentication methods, port settings, and user permissions. This file is typically located at `/etc/ssh/sshd_config` on most Linux systems.

Here is a detailed overview of common settings in **`/etc/ssh/sshd_config`**:

### Basic Configuration Options

1. **Port**:
   - Specifies the port number `sshd` listens on. The default is 22.
   - Example:
     ```sh
     Port 22
     ```

2. **ListenAddress**:
   - Specifies the IP address `sshd` should listen on. You can specify multiple addresses.
   - Example:
     ```sh
     ListenAddress 0.0.0.0
     ListenAddress :: 
     ```

3. **PermitRootLogin**:
   - Controls whether the root user can log in via SSH. Options include `yes`, `no`, `without-password`, `prohibit-password`.
   - Example:
     ```sh
     PermitRootLogin no
     ```

4. **PasswordAuthentication**:
   - Enables or disables password authentication. For improved security, consider using key-based authentication.
   - Example:
     ```sh
     PasswordAuthentication yes
     ```

5. **PubkeyAuthentication**:
   - Enables or disables public key authentication.
   - Example:
     ```sh
     PubkeyAuthentication yes
     ```

6. **PermitEmptyPasswords**:
   - Allows or disallows login with empty passwords.
   - Example:
     ```sh
     PermitEmptyPasswords no
     ```

7. **ChallengeResponseAuthentication**:
   - Enables or disables challenge-response authentication (such as OTP).
   - Example:
     ```sh
     ChallengeResponseAuthentication no
     ```

8. **UsePAM**:
   - Specifies whether to use Pluggable Authentication Modules (PAM) for authentication.
   - Example:
     ```sh
     UsePAM yes
     ```

9. **AllowUsers**:
   - Specifies which users are allowed to log in via SSH.
   - Example:
     ```sh
     AllowUsers user1 user2
     ```

10. **DenyUsers**:
    - Specifies which users are denied SSH access.
    - Example:
      ```sh
      DenyUsers user3 user4
      ```

11. **AllowGroups**:
    - Specifies which groups are allowed to log in via SSH.
    - Example:
      ```sh
      AllowGroups sshusers
      ```

12. **DenyGroups**:
    - Specifies which groups are denied SSH access.
    - Example:
      ```sh
      DenyGroups badusers
      ```

13. **X11Forwarding**:
    - Enables or disables X11 forwarding.
    - Example:
      ```sh
      X11Forwarding yes
      ```

14. **AllowTcpForwarding**:
    - Controls whether TCP forwarding is allowed.
    - Example:
      ```sh
      AllowTcpForwarding yes
      ```

15. **PrintMotd**:
    - Controls whether the message of the day (MOTD) is printed when a user logs in.
    - Example:
      ```sh
      PrintMotd yes
      ```

16. **Subsystem**:
    - Configures subsystems like `sftp`. By default, OpenSSH provides an SFTP subsystem.
    - Example:
      ```sh
      Subsystem sftp /usr/lib/openssh/sftp-server
      ```

17. **PermitTunnel**:
    - Allows or disallows tunnelled sessions.
    - Example:
      ```sh
      PermitTunnel no
      ```

18. **LogLevel**:
    - Sets the verbosity of logging. Levels include `QUIET`, `FATAL`, `ERROR`, `INFO`, `VERBOSE`, `DEBUG`, and `DEBUG1`, `DEBUG2`, `DEBUG3` for more detailed debugging.
    - Example:
      ```sh
      LogLevel INFO
      ```

### Example Configuration

Here is an example of a more secure **`sshd_config`**:

```sh
# Port to listen on
Port 22

# Allow only specific users
AllowUsers user1 user2

# Disable root login
PermitRootLogin no

# Enable public key authentication
PubkeyAuthentication yes

# Disable password authentication
PasswordAuthentication no

# Disable empty passwords
PermitEmptyPasswords no

# Enable X11 forwarding
X11Forwarding yes

# Log level
LogLevel VERBOSE

# Subsystem for SFTP
Subsystem sftp /usr/lib/openssh/sftp-server
```

### Reloading Configuration

After making changes to **`sshd_config`**, you need to reload or restart the `sshd` service to apply the changes:

- **Reload** (to apply configuration changes without disconnecting active sessions):
  ```sh
  sudo systemctl reload sshd
  ```

- **Restart** (to restart the `sshd` service):
  ```sh
  sudo systemctl restart sshd
  ```

### Summary

The **`sshd_config`** file is crucial for configuring how the SSH daemon operates, including security settings, authentication methods, and user permissions. Properly configuring this file helps secure SSH access and manage how users connect to the system.
