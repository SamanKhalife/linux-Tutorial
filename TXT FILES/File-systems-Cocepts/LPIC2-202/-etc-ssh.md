# /etc/ssh
The **`/etc/ssh`** directory contains configuration files and keys for the SSH (Secure Shell) service on a Unix-like system. This directory is critical for managing SSH server settings, including authentication methods, security policies, and key management. Here is a detailed overview of the main files and their purposes found in this directory:

### Key Files in `/etc/ssh`

1. **`sshd_config`**:
   - **Purpose**: The main configuration file for the SSH daemon (`sshd`), which defines the server's settings.
   - **Location**: `/etc/ssh/sshd_config`
   - **Common Settings**: Port number, authentication methods, root login permissions, and more.
   - **Example**:
     ```sh
     Port 22
     PermitRootLogin no
     PubkeyAuthentication yes
     ```

2. **`ssh_config`**:
   - **Purpose**: The configuration file for the SSH client (`ssh`), which specifies default options for SSH connections.
   - **Location**: `/etc/ssh/ssh_config`
   - **Common Settings**: Default user, port, and other options used by the `ssh` client.
   - **Example**:
     ```sh
     Host *
       ForwardAgent yes
       ForwardX11 yes
       ServerAliveInterval 60
     ```

3. **`ssh_host_rsa_key`**:
   - **Purpose**: The private RSA key used by the SSH server for RSA key-based authentication.
   - **Location**: `/etc/ssh/ssh_host_rsa_key`
   - **Permissions**: Should be read/write for root only (`600`).

4. **`ssh_host_rsa_key.pub`**:
   - **Purpose**: The public RSA key corresponding to the private RSA key.
   - **Location**: `/etc/ssh/ssh_host_rsa_key.pub`
   - **Permissions**: Should be readable by everyone (`644`).

5. **`ssh_host_ecdsa_key`**:
   - **Purpose**: The private ECDSA key used by the SSH server for ECDSA key-based authentication.
   - **Location**: `/etc/ssh/ssh_host_ecdsa_key`
   - **Permissions**: Should be read/write for root only (`600`).

6. **`ssh_host_ecdsa_key.pub`**:
   - **Purpose**: The public ECDSA key corresponding to the private ECDSA key.
   - **Location**: `/etc/ssh/ssh_host_ecdsa_key.pub`
   - **Permissions**: Should be readable by everyone (`644`).

7. **`ssh_host_ed25519_key`**:
   - **Purpose**: The private Ed25519 key used by the SSH server for Ed25519 key-based authentication.
   - **Location**: `/etc/ssh/ssh_host_ed25519_key`
   - **Permissions**: Should be read/write for root only (`600`).

8. **`ssh_host_ed25519_key.pub`**:
   - **Purpose**: The public Ed25519 key corresponding to the private Ed25519 key.
   - **Location**: `/etc/ssh/ssh_host_ed25519_key.pub`
   - **Permissions**: Should be readable by everyone (`644`).

9. **`ssh_known_hosts`**:
   - **Purpose**: Contains the host keys of known SSH servers. Used by the SSH client to verify the identity of the server.
   - **Location**: `/etc/ssh/ssh_known_hosts`
   - **Permissions**: Should be readable by everyone (`644`).

10. **`moduli`**:
    - **Purpose**: Contains a list of DH (Diffie-Hellman) groups used for key exchange. This file is used by the SSH server and client to select appropriate key exchange parameters.
    - **Location**: `/etc/ssh/moduli`
    - **Permissions**: Should be readable by everyone (`644`).

11. **`ssh_config.d/`**:
    - **Purpose**: Directory for additional SSH client configuration snippets. It allows for modular configuration management.
    - **Location**: `/etc/ssh/ssh_config.d/`
    - **Permissions**: Should be readable by everyone (`755`).

12. **`sshd_config.d/`**:
    - **Purpose**: Directory for additional SSH server configuration snippets. Similar to `ssh_config.d/`, but for the SSH server.
    - **Location**: `/etc/ssh/sshd_config.d/`
    - **Permissions**: Should be readable by everyone (`755`).

### Example Configuration Files

**`/etc/ssh/sshd_config`**:
```sh
# Listen on port 22
Port 22

# Allow only specific users
AllowUsers user1 user2

# Disable root login
PermitRootLogin no

# Enable public key authentication
PubkeyAuthentication yes

# Disable password authentication
PasswordAuthentication no

# Log level
LogLevel VERBOSE

# Subsystem for SFTP
Subsystem sftp /usr/lib/openssh/sftp-server
```

**`/etc/ssh/ssh_config`**:
```sh
# Global settings for all SSH clients
Host *
  ForwardAgent yes
  ForwardX11 yes
  ServerAliveInterval 60
```

### Permissions

- **Private Key Files**: Should be owned by root and have permissions set to `600` to prevent unauthorized access.
- **Public Key Files**: Should be world-readable (`644`), as they need to be accessed by clients.
- **Configuration Files**: Typically world-readable (`644`), but ownership should be set to root.

### Summary

The **`/etc/ssh`** directory is crucial for configuring SSH services on a Unix-like system. It contains key files for server authentication, configuration files for both client and server, and directories for modular configuration. Proper configuration and security of these files are essential for maintaining secure SSH communications and preventing unauthorized access.
