# /etc/swanctl/

The `/etc/swanctl/` directory is used to store configuration files for the `swanctl` command-line utility, which is part of the strongSwan VPN suite. This directory typically contains various configuration files, including `swanctl.conf` for connection definitions and additional files for certificates, private keys, and other settings related to IPsec and IKEv2 configurations.

#### Purpose

The `/etc/swanctl/` directory centralizes the configuration and management of IPsec and IKEv2 connections, secrets, and related settings for strongSwan. This modular configuration approach allows for organized and maintainable VPN settings.

### Common Files and Subdirectories

1. **swanctl.conf**: The main configuration file for defining connection profiles, secrets, and address pools.
2. **certs/**: Directory for storing certificate files.
3. **private/**: Directory for storing private key files.
4. **conf.d/**: Directory for additional configuration snippets, if modular configurations are used.

### Example Directory Structure

```plaintext
/etc/swanctl/
├── certs/
│   └── mycert.pem
├── private/
│   └── mykey.pem
├── conf.d/
│   ├── connection1.conf
│   └── connection2.conf
└── swanctl.conf
```

### Example Configuration Files

#### `/etc/swanctl/swanctl.conf`

This file defines the main connection profiles, secrets, and pools.

```ini
connections {
    myvpn {
        version = 2
        local_addrs = 192.0.2.1
        remote_addrs = 192.0.2.2

        local {
            auth = pubkey
            certs = certs/mycert.pem
            id = "CN=myvpnserver"
        }
        remote {
            auth = pubkey
            certs = certs/remote_cert.pem
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
        file = private/mykey.pem
    }
}

pools {
    mypool {
        addrs = 10.3.0.0/24
        dns = 10.3.0.1
    }
}
```

#### `/etc/swanctl/certs/mycert.pem`

This directory holds the local certificate used for authentication.

```plaintext
-----BEGIN CERTIFICATE-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr6Vm
...
-----END CERTIFICATE-----
```

#### `/etc/swanctl/private/mykey.pem`

This directory contains the private key corresponding to the certificate.

```plaintext
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDX4
...
-----END PRIVATE KEY-----
```

### Managing the Configuration

#### Loading the Configuration

To apply the configuration defined in `swanctl.conf`, use the following command:

```bash
sudo swanctl --load-all
```

This command loads connections, secrets, and pools defined in the configuration files.

#### Common swanctl Commands

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

The `/etc/swanctl/` directory and its configuration files are crucial for managing strongSwan VPN connections. By organizing configuration into modular files, administrators can maintain a clear and manageable VPN setup. Properly configuring `swanctl.conf` and storing related certificates and keys in appropriate subdirectories ensures a secure and efficient VPN infrastructure. Always remember to load the configuration after making changes and use `swanctl` commands to manage the VPN connections effectively.
