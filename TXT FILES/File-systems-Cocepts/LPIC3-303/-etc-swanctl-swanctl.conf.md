# /etc/swanctl/swanctl.conf

The `swanctl.conf` file located in `/etc/swanctl/` is the primary configuration file for the `swanctl` command, which is part of the strongSwan VPN suite. This file is used to define the connection profiles, secrets, and other configurations necessary for managing IKEv2 and IPsec connections.

#### Purpose

The `swanctl.conf` file is used to configure the `swanctl` command-line tool, which provides an interface for managing and monitoring IKEv2 connections established by the Charon daemon in strongSwan. This configuration file includes definitions for connections, pools, and secrets.

### Structure of `swanctl.conf`

The configuration file is written in a hierarchical, section-based format. The main sections typically include `connections`, `secrets`, and `pools`.

### Example `swanctl.conf`

Here is a basic example of a `swanctl.conf` file:

```ini
connections {
    myvpn {
        version = 2
        local_addrs = 192.0.2.1
        remote_addrs = 192.0.2.2

        local {
            auth = pubkey
            certs = serverCert.pem
            id = "CN=myvpnserver"
        }
        remote {
            auth = pubkey
            certs = clientCert.pem
            id = "CN=myvpnclient"
        }

        children {
            net {
                local_ts = 10.0.0.0/24
                remote_ts = 0.0.0.0/0
                dpd_action = restart
                esp_proposals = aes256-sha256-modp2048
            }
        }
    }
}

secrets {
    private {
        file = serverKey.pem
    }
    eap {
        id = "username"
        secret = "password"
    }
}

pools {
    mypool {
        addrs = 10.3.0.0/24
        dns = 10.3.0.1
    }
}
```

### Sections and Options

1. **connections**:
    - **myvpn**: Name of the connection profile.
      - **version**: Specifies the IKE version (1 or 2).
      - **local_addrs**: Local IP address.
      - **remote_addrs**: Remote IP address.
      - **local**: Local authentication settings.
        - **auth**: Authentication method (e.g., `pubkey`, `psk`).
        - **certs**: Path to the certificate file.
        - **id**: Identity of the local peer.
      - **remote**: Remote authentication settings.
        - **auth**: Authentication method for the remote peer.
        - **certs**: Path to the certificate file of the remote peer.
        - **id**: Identity of the remote peer.
      - **children**: Defines child SAs.
        - **net**: Name of the child SA.
          - **local_ts**: Traffic selectors for local traffic.
          - **remote_ts**: Traffic selectors for remote traffic.
          - **dpd_action**: Dead Peer Detection (DPD) action.
          - **esp_proposals**: ESP encryption and integrity algorithms.

2. **secrets**:
    - **private**: Specifies the private key.
      - **file**: Path to the private key file.
    - **eap**: Defines EAP (Extensible Authentication Protocol) secrets.
      - **id**: EAP identity (username).
      - **secret**: EAP secret (password).

3. **pools**:
    - **mypool**: Name of the address pool.
      - **addrs**: Range of IP addresses.
      - **dns**: DNS server address.

### Managing the Configuration

#### Loading Configuration

To apply the configuration in `swanctl.conf`, use the following command:

```bash
sudo swanctl --load-all
```

This command loads the connections, secrets, and pools defined in the configuration file.

#### Monitoring and Managing Connections

- **List all connections**:
  ```bash
  sudo swanctl --list-conns
  ```

- **Initiate a connection**:
  ```bash
  sudo swanctl --initiate --child myvpn
  ```

- **Terminate a connection**:
  ```bash
  sudo swanctl --terminate --child myvpn
  ```

- **Check the status of connections**:
  ```bash
  sudo swanctl --list-sas
  ```

### Conclusion

The `/etc/swanctl/swanctl.conf` file is an essential component for configuring and managing IKEv2 and IPsec connections with the strongSwan VPN suite. By properly defining connections, authentication methods, and address pools, administrators can create secure VPN configurations tailored to their network requirements. The modular approach of this file allows for clear and organized management of VPN settings, ensuring a robust and secure VPN infrastructure.
