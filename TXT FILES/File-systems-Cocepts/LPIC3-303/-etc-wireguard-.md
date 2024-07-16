# /etc/wireguard/

The `/etc/wireguard/` directory is used for storing configuration files for the WireGuard VPN. WireGuard is a modern, high-performance VPN that aims to be faster and simpler than traditional VPN solutions like IPsec and OpenVPN. Configuration files in this directory define the VPN interfaces, peers, and other settings necessary to establish secure connections.

#### Purpose

The `/etc/wireguard/` directory holds configuration files for WireGuard interfaces. Each configuration file typically corresponds to a network interface, such as `wg0.conf`.

### Common Files and Subdirectories

1. **Interface Configuration Files**: These files define the settings for each WireGuard interface, usually named `wg0.conf`, `wg1.conf`, etc.
2. **Keys**: Private and public keys used for authentication and encryption can be stored here or elsewhere, referenced in the configuration files.

### Example Directory Structure

```plaintext
/etc/wireguard/
├── wg0.conf
└── wg1.conf
```

### Example Configuration Files

#### `/etc/wireguard/wg0.conf`

This file defines the settings for the WireGuard interface `wg0`.

```ini
[Interface]
PrivateKey = <base64-encoded-private-key>
Address = 10.0.0.1/24
ListenPort = 51820

[Peer]
PublicKey = <base64-encoded-peer-public-key>
AllowedIPs = 10.0.0.2/32
Endpoint = peer.example.com:51820
PersistentKeepalive = 25
```

- **PrivateKey**: The private key for this interface.
- **Address**: The IP address assigned to this interface.
- **ListenPort**: The port on which WireGuard listens for incoming connections.
- **Peer**: Defines a peer that this interface can communicate with.
  - **PublicKey**: The public key of the peer.
  - **AllowedIPs**: IP addresses that are allowed to be routed to this peer.
  - **Endpoint**: The endpoint of the peer (hostname or IP address and port).
  - **PersistentKeepalive**: An optional setting to keep the connection alive.

#### `/etc/wireguard/wg1.conf`

Another example configuration for a different WireGuard interface.

```ini
[Interface]
PrivateKey = <base64-encoded-private-key>
Address = 10.0.1.1/24
ListenPort = 51821

[Peer]
PublicKey = <base64-encoded-peer-public-key>
AllowedIPs = 10.0.1.2/32
Endpoint = peer2.example.com:51821
PersistentKeepalive = 25
```

### Managing WireGuard

#### Starting and Stopping Interfaces

To bring up a WireGuard interface, use the `wg-quick` utility:

```bash
sudo wg-quick up wg0
```

To bring down the interface:

```bash
sudo wg-quick down wg0
```

#### Viewing Interface Status

To view the status and details of a WireGuard interface, use:

```bash
sudo wg show wg0
```

### Example Commands

- **Generate a Key Pair**:
  ```bash
  wg genkey | tee privatekey | wg pubkey > publickey
  ```
  - `privatekey`: The generated private key.
  - `publickey`: The corresponding public key.

- **Show WireGuard Status**:
  ```bash
  sudo wg show
  ```

- **Show Configuration of a Specific Interface**:
  ```bash
  sudo wg showconf wg0
  ```

### Security Considerations

1. **Protect Private Keys**: Ensure that private keys are stored securely and permissions are set appropriately (e.g., `chmod 600 privatekey`).
2. **Firewall Configuration**: Ensure that the WireGuard listening port is allowed through any firewalls.

### Conclusion

The `/etc/wireguard/` directory and its configuration files are essential for setting up and managing WireGuard VPN interfaces. By properly configuring these files and using the `wg` and `wg-quick` utilities, administrators can establish secure and efficient VPN connections. WireGuard’s simplicity and performance make it a powerful tool for modern VPN setups.
