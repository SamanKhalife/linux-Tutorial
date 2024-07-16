# wg

The `wg` command is a tool provided by WireGuard for managing WireGuard interfaces and peers. It allows you to configure, view, and manage WireGuard settings directly from the command line.

#### Purpose

The `wg` command provides various subcommands and options to create, configure, and monitor WireGuard VPN interfaces and their peers. It allows administrators to perform tasks such as generating keys, setting configuration parameters, and displaying interface status.

### Common `wg` Commands and Options

#### 1. Key Management

- **Generate a private key**:
  ```bash
  wg genkey
  ```
  Example:
  ```bash
  sudo wg genkey | tee privatekey
  ```

- **Generate a public key from a private key**:
  ```bash
  wg pubkey < privatekey
  ```
  Example:
  ```bash
  sudo wg pubkey < privatekey > publickey
  ```

#### 2. Interface Management

- **Bring up a WireGuard interface**:
  ```bash
  sudo ip link add dev wg0 type wireguard
  ```

- **Assign an IP address to the interface**:
  ```bash
  sudo ip address add dev wg0 10.0.0.1/24
  ```

- **Set the private key for the interface**:
  ```bash
  sudo wg set wg0 private-key /path/to/privatekey
  ```

- **Set the listen port for the interface**:
  ```bash
  sudo wg set wg0 listen-port 51820
  ```

- **Bring up the interface**:
  ```bash
  sudo ip link set up dev wg0
  ```

#### 3. Peer Management

- **Add a peer to the interface**:
  ```bash
  sudo wg set wg0 peer <base64-encoded-peer-public-key> allowed-ips 10.0.0.2/32 endpoint peer.example.com:51820
  ```

- **Set persistent keepalive interval**:
  ```bash
  sudo wg set wg0 peer <base64-encoded-peer-public-key> persistent-keepalive 25
  ```

#### 4. Viewing Configuration and Status

- **View the current configuration of an interface**:
  ```bash
  sudo wg show wg0
  ```

- **View the full configuration in configuration file format**:
  ```bash
  sudo wg showconf wg0
  ```

- **View all WireGuard interfaces and their status**:
  ```bash
  sudo wg
  ```

### Example Commands

#### Generating Key Pairs

Generate a private key and derive the public key:
```bash
sudo wg genkey | tee /etc/wireguard/privatekey | wg pubkey > /etc/wireguard/publickey
```

#### Setting Up an Interface and Adding a Peer

1. Create the interface:
   ```bash
   sudo ip link add dev wg0 type wireguard
   ```

2. Assign an IP address:
   ```bash
   sudo ip address add dev wg0 10.0.0.1/24
   ```

3. Set the private key:
   ```bash
   sudo wg set wg0 private-key /etc/wireguard/privatekey
   ```

4. Set the listen port:
   ```bash
   sudo wg set wg0 listen-port 51820
   ```

5. Add a peer:
   ```bash
   sudo wg set wg0 peer <peer-public-key> allowed-ips 10.0.0.2/32 endpoint peer.example.com:51820 persistent-keepalive 25
   ```

6. Bring up the interface:
   ```bash
   sudo ip link set up dev wg0
   ```

#### Viewing the Status of the Interface

To see the current status and configuration of the interface:
```bash
sudo wg show wg0
```

### Configuration File Example

Here is an example configuration file `/etc/wireguard/wg0.conf`:

```ini
[Interface]
PrivateKey = <your-private-key>
Address = 10.0.0.1/24
ListenPort = 51820

[Peer]
PublicKey = <peer-public-key>
AllowedIPs = 10.0.0.2/32
Endpoint = peer.example.com:51820
PersistentKeepalive = 25
```

To bring up the interface using the configuration file:
```bash
sudo wg-quick up wg0
```

To bring it down:
```bash
sudo wg-quick down wg0
```

### Conclusion

The `wg` command is a central tool for managing WireGuard VPN interfaces and peers. By leveraging its capabilities, administrators can efficiently set up, configure, and monitor secure VPN connections. The simplicity and performance of WireGuard, combined with the flexibility of the `wg` command, make it an excellent choice for modern VPN setups.
