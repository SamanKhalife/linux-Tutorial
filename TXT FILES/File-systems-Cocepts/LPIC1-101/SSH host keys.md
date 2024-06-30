# SSH host keys
SSH host keys are cryptographic keys used by the SSH (Secure Shell) protocol to verify the identity of a server to a client. These keys play a crucial role in securing SSH connections by preventing man-in-the-middle attacks and ensuring that the client is connecting to the correct server.

### Key Concepts of SSH Host Keys

#### Purpose of SSH Host Keys

1. **Server Authentication**: SSH host keys allow clients to verify that they are connecting to the correct server. When a client connects to a server for the first time, the server presents its host key, which the client can then store for future reference.
2. **Security**: By ensuring that the server’s identity can be verified, SSH host keys help protect against man-in-the-middle attacks.
3. **Encryption**: Host keys are part of the initial key exchange process that sets up encrypted communication between the client and server.

#### Types of SSH Host Keys

SSH supports several types of host keys, each using a different cryptographic algorithm:

1. **RSA**: One of the oldest and most widely used algorithms. Suitable for most use cases.
2. **ECDSA**: Based on elliptic curve cryptography, offering strong security with smaller key sizes.
3. **ED25519**: A newer elliptic curve algorithm that provides high security and performance.
4. **DSA**: An older algorithm that is less commonly used today due to security concerns.

### Location of SSH Host Keys

SSH host keys are typically stored in the `/etc/ssh/` directory on the server:

- RSA: `/etc/ssh/ssh_host_rsa_key` and `/etc/ssh/ssh_host_rsa_key.pub`
- ECDSA: `/etc/ssh/ssh_host_ecdsa_key` and `/etc/ssh/ssh_host_ecdsa_key.pub`
- ED25519: `/etc/ssh/ssh_host_ed25519_key` and `/etc/ssh/ssh_host_ed25519_key.pub`
- DSA: `/etc/ssh/ssh_host_dsa_key` and `/etc/ssh/ssh_host_dsa_key.pub`

### Managing SSH Host Keys

#### Viewing Existing Host Keys

You can view the public part of the SSH host keys using the following command:

```sh
sudo cat /etc/ssh/ssh_host_rsa_key.pub
sudo cat /etc/ssh/ssh_host_ecdsa_key.pub
sudo cat /etc/ssh/ssh_host_ed25519_key.pub
sudo cat /etc/ssh/ssh_host_dsa_key.pub
```

#### Generating New Host Keys

To generate new SSH host keys, you can use the `ssh-keygen` command. This might be necessary if the keys have been compromised or if you want to upgrade to a more secure key type.

Example for generating an RSA key:

```sh
sudo ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key -N ''
```

Replace `rsa` with `ecdsa`, `ed25519`, or `dsa` to generate keys of other types. The `-N ''` part ensures that the key has no passphrase.

#### Configuring SSH to Use Host Keys

Ensure that your `sshd_config` file (typically located at `/etc/ssh/sshd_config`) includes the correct paths to the host keys:

```sh
HostKey /etc/ssh/ssh_host_rsa_key
HostKey /etc/ssh/ssh_host_ecdsa_key
HostKey /etc/ssh/ssh_host_ed25519_key
HostKey /etc/ssh/ssh_host_dsa_key
```

After making changes to the `sshd_config` file, restart the SSH service to apply the changes:

```sh
sudo systemctl restart sshd
```

#### Verifying SSH Host Keys on the Client Side

When a client connects to an SSH server for the first time, it stores the server’s host key in the `~/.ssh/known_hosts` file. The next time the client connects to the server, it verifies the server’s identity using this stored key.

You can manually add a server’s host key to the `known_hosts` file to avoid the prompt during the first connection:

```sh
ssh-keyscan -H <server-ip-or-hostname> >> ~/.ssh/known_hosts
```

### Example: SSH Host Key Management

#### Viewing the Current Host Key Fingerprints

```sh
sudo ssh-keygen -l -f /etc/ssh/ssh_host_rsa_key.pub
sudo ssh-keygen -l -f /etc/ssh/ssh_host_ecdsa_key.pub
sudo ssh-keygen -l -f /etc/ssh/ssh_host_ed25519_key.pub
```

#### Generating a New ED25519 Host Key

```sh
sudo ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key -N ''
```

#### Ensuring SSHD Uses the Correct Host Keys

Edit `/etc/ssh/sshd_config`:

```sh
sudo nano /etc/ssh/sshd_config
```

Add or verify the following lines:

```sh
HostKey /etc/ssh/ssh_host_rsa_key
HostKey /etc/ssh/ssh_host_ecdsa_key
HostKey /etc/ssh/ssh_host_ed25519_key
```

Restart the SSH service:

```sh
sudo systemctl restart sshd
```

### Conclusion

SSH host keys are fundamental to the security and integrity of SSH connections. Proper management of these keys ensures that clients can reliably verify the identity of servers and establish secure communication channels. Understanding how to view, generate, and configure SSH host keys is crucial for maintaining a secure SSH environment.
