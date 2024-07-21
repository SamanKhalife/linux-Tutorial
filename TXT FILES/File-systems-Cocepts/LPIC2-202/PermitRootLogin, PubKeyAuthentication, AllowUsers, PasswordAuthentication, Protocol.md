# PermitRootLogin, PubKeyAuthentication, AllowUsers, PasswordAuthentication, Protocol

The SSH configuration file (`sshd_config`) allows administrators to control various aspects of the SSH server's behavior. Here’s an overview of the specific directives you mentioned:

### 1. **PermitRootLogin**

**Purpose**: Controls whether the root user is allowed to log in via SSH.

**Syntax**:
```sh
PermitRootLogin yes | no | prohibit-password | forced-commands-only
```

- **`yes`**: Allows root login via SSH with any authentication method.
- **`no`**: Disables root login completely.
- **`prohibit-password`**: Disables password-based authentication for root, but allows other authentication methods such as public key.
- **`forced-commands-only`**: Allows root login only if the command specified in the `authorized_keys` file is executed, primarily used for restricted commands.

**Example**:
```sh
PermitRootLogin no
```
This configuration disables root logins via SSH.

### 2. **PubkeyAuthentication**

**Purpose**: Enables or disables public key authentication for SSH connections.

**Syntax**:
```sh
PubkeyAuthentication yes | no
```

- **`yes`**: Enables public key authentication.
- **`no`**: Disables public key authentication.

**Example**:
```sh
PubkeyAuthentication yes
```
This configuration enables public key authentication, which is often preferred for secure logins.

### 3. **AllowUsers**

**Purpose**: Specifies which users are allowed to log in via SSH. 

**Syntax**:
```sh
AllowUsers user1 user2 ...
```

- **`user1 user2 ...`**: A list of usernames who are allowed to log in. Wildcards can be used (e.g., `user*`).

**Example**:
```sh
AllowUsers alice bob
```
This configuration allows only the users `alice` and `bob` to log in via SSH.

### 4. **PasswordAuthentication**

**Purpose**: Controls whether password authentication is allowed for SSH connections.

**Syntax**:
```sh
PasswordAuthentication yes | no
```

- **`yes`**: Allows password-based authentication.
- **`no`**: Disables password-based authentication, requiring other authentication methods like public key.

**Example**:
```sh
PasswordAuthentication no
```
This configuration disables password authentication, which can enhance security by requiring keys or other methods.

### 5. **Protocol**

**Purpose**: Specifies the SSH protocol versions supported by the server.

**Syntax**:
```sh
Protocol 1 | 2
```

- **`1`**: Enables SSH protocol version 1 (generally considered outdated and less secure).
- **`2`**: Enables SSH protocol version 2 (recommended and more secure).
- **`1,2`**: Enables both versions (rarely used as version 1 is deprecated).

**Example**:
```sh
Protocol 2
```
This configuration enables only SSH protocol version 2, which is the recommended version for security reasons.

### Summary

- **`PermitRootLogin`**: Controls root access; set to `no` for better security.
- **`PubkeyAuthentication`**: Enables or disables public key authentication; recommended to enable.
- **`AllowUsers`**: Restricts SSH access to specific users.
- **`PasswordAuthentication`**: Controls if password authentication is allowed; setting to `no` enforces key-based authentication.
- **`Protocol`**: Specifies which SSH protocol versions are supported; `2` is preferred.

These directives can be used to enhance the security and control access to SSH services on a Unix-like system.
