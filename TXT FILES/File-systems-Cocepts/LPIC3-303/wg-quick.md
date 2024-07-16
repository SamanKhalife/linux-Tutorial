# wg-quick
The `wg-quick` command is a utility that simplifies the management of WireGuard interfaces. It uses configuration files to bring up and down WireGuard interfaces with a single command, handling the necessary IP and routing configurations automatically.


#### Purpose

`wg-quick` automates the setup of WireGuard interfaces, allowing for quick and easy management of VPN connections. It reads configuration files from `/etc/wireguard/` and applies the settings to create or remove interfaces.

### Common Commands and Options

#### Bringing Up an Interface

To bring up a WireGuard interface using `wg-quick`, use the following command:

```bash
sudo wg-quick up <interface-name>
```

Example:

```bash
sudo wg-quick up wg0
```

#### Bringing Down an Interface

To bring down a WireGuard interface, use:

```bash
sudo wg-quick down <interface-name>
```

Example:

```bash
sudo wg-quick down wg0
```

### Configuration File Structure

The configuration files for `wg-quick` are typically located in `/etc/wireguard/` and have the `.conf` extension. Each file corresponds to a WireGuard interface.

#### Example Configuration File: `/etc/wireguard/wg0.conf`

```ini
[Interface]
PrivateKey = <base64-encoded-private-key>
Address = 10.0.0.1/24
ListenPort = 51820
DNS = 1.1.1.1

[Peer]
PublicKey = <base64-encoded-peer-public-key>
AllowedIPs = 10.0.0.2/32
Endpoint = peer.example.com:51820
PersistentKeepalive = 25
```

- **[Interface] Section**:
  - **PrivateKey**: The private key for this interface.
  - **Address**: The IP address(es) assigned to this interface.
  - **ListenPort**: The port on which WireGuard listens for incoming connections.
  - **DNS**: Optional, sets the DNS servers to use while the interface is up.

- **[Peer] Section**:
  - **PublicKey**: The public key of the peer.
  - **AllowedIPs**: IP addresses that are allowed to be routed to this peer.
  - **Endpoint**: The endpoint of the peer (hostname or IP address and port).
  - **PersistentKeepalive**: Optional, keeps the connection alive by sending keepalive packets.

### Example Commands

#### Starting and Stopping an Interface

Start the WireGuard interface `wg0`:

```bash
sudo wg-quick up wg0
```

Stop the WireGuard interface `wg0`:

```bash
sudo wg-quick down wg0
```

### Environment Variables

`wg-quick` supports several environment variables that can modify its behavior. For instance:

- **WG_QUICK_USERSPACE_IMPLEMENTATION**: Specifies a userspace implementation (e.g., `boringtun`).
- **WG_QUICK_PERSISTENT_KEEPALIVE**: Sets the default PersistentKeepalive interval for peers.

Example:

```bash
WG_QUICK_USERSPACE_IMPLEMENTATION=boringtun sudo wg-quick up wg0
```

### Using `wg-quick` with Systemd

`wg-quick` can be integrated with systemd for automatic management of WireGuard interfaces at boot or shutdown.

#### Example Systemd Service Unit

Create a service file `/etc/systemd/system/wg-quick@wg0.service`:

```ini
[Unit]
Description=WireGuard via wg-quick(8) for %I
After=network.target
Wants=network-online.target
Before=network-online.target
Requires=systemd-networkd.service

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/bin/wg-quick up %I
ExecStop=/usr/bin/wg-quick down %I

[Install]
WantedBy=multi-user.target
```

Enable the service:

```bash
sudo systemctl enable wg-quick@wg0
```

Start the service:

```bash
sudo systemctl start wg-quick@wg0
```

### Conclusion

`wg-quick` is a powerful utility for managing WireGuard interfaces, streamlining the process of setting up and tearing down VPN connections. By leveraging configuration files and simple commands, administrators can quickly deploy secure VPNs. Integration with systemd further enhances its usability, allowing for automatic management of VPN interfaces.
