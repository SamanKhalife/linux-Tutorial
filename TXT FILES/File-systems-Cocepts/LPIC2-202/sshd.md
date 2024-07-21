# sshd
**`sshd`** (OpenSSH Daemon) is the server-side component of the OpenSSH suite, responsible for handling incoming SSH connections and managing secure remote login sessions. It allows users to securely connect to a remote machine over a network.

### Important `sshd` Command-Line Options

When starting `sshd`, you can use various command-line options to control its behavior. Here are some of the most important options:

#### Starting `sshd`:

- **`-D`**: Run `sshd` in the foreground. Useful for debugging or running `sshd` in environments where a daemon is not desirable.
  ```sh
  /usr/sbin/sshd -D
  ```

- **`-f config_file`**: Specify a custom configuration file instead of the default `/etc/ssh/sshd_config`.
  ```sh
  /usr/sbin/sshd -f /path/to/custom_sshd_config
  ```

- **`-p port`**: Specify a port number for `sshd` to listen on. By default, `sshd` listens on port 22.
  ```sh
  /usr/sbin/sshd -p 2222
  ```

- **`-h`**: Print a help message showing all available options.
  ```sh
  /usr/sbin/sshd -h
  ```

- **`-t`**: Test the configuration file for errors without starting `sshd`.
  ```sh
  /usr/sbin/sshd -t
  ```

- **`-u`**: Enable user-based options. This is useful for overriding default options with user-specific settings.
  ```sh
  /usr/sbin/sshd -u
  ```

- **`-a`**: Enable authentication methods specified in the configuration file. This option is not commonly used directly but is more relevant in internal or script contexts.

#### Configuration File Options

The main configuration file for `sshd` is `/etc/ssh/sshd_config`. Some key settings you can configure in this file include:

- **`Port`**: Specifies the port number for `sshd` to listen on.
  ```sh
  Port 22
  ```

- **`PermitRootLogin`**: Controls whether root can log in via SSH. Options include `yes`, `no`, `without-password`, or `prohibit-password`.
  ```sh
  PermitRootLogin no
  ```

- **`PasswordAuthentication`**: Enables or disables password authentication. For better security, consider using key-based authentication.
  ```sh
  PasswordAuthentication yes
  ```

- **`PubkeyAuthentication`**: Enables or disables public key authentication.
  ```sh
  PubkeyAuthentication yes
  ```

- **`AllowUsers`**: Specifies which users are allowed to log in via SSH.
  ```sh
  AllowUsers user1 user2
  ```

- **`DenyUsers`**: Specifies which users are denied SSH access.
  ```sh
  DenyUsers user3 user4
  ```

- **`PermitEmptyPasswords`**: Allows or denies login with empty passwords. It is generally advised to disable this for security reasons.
  ```sh
  PermitEmptyPasswords no
  ```

- **`AllowTcpForwarding`**: Controls whether TCP forwarding is permitted.
  ```sh
  AllowTcpForwarding yes
  ```

- **`X11Forwarding`**: Enables or disables X11 forwarding.
  ```sh
  X11Forwarding yes
  ```

- **`ChallengeResponseAuthentication`**: Controls whether challenge-response authentication is allowed (e.g., for 2FA).
  ```sh
  ChallengeResponseAuthentication no
  ```

- **`UsePAM`**: Specifies whether to use Pluggable Authentication Modules (PAM) for authentication.
  ```sh
  UsePAM yes
  ```

### Example Usage

- **Start `sshd` with a custom configuration file and in the foreground**:
  ```sh
  /usr/sbin/sshd -f /etc/ssh/sshd_config.custom -D
  ```

- **Test the `sshd` configuration for errors**:
  ```sh
  /usr/sbin/sshd -t
  ```

- **Start `sshd` on a non-default port**:
  ```sh
  /usr/sbin/sshd -p 2222
  ```

### Summary

The `sshd` command is central to managing SSH connections on a server. Understanding its command-line options and configuration file settings allows administrators to securely and efficiently manage remote access. For robust security, ensure that `sshd` is configured to use key-based authentication, restrict root access, and implement other best practices.
