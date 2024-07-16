# swanctl

The `swanctl` command is a powerful tool in the strongSwan suite used for managing and configuring IPsec and IKEv2 connections. It allows you to load, list, initiate, terminate, and monitor VPN connections based on the configuration files in the `/etc/swanctl/` directory.

#### Purpose

`swanctl` provides a command-line interface for managing the strongSwan IPsec daemon (Charon). It interacts with the configurations defined in `swanctl.conf` and other related files in `/etc/swanctl/`.

### Common `swanctl` Commands

1. **Loading Configuration**
   - **Load all configurations**: This command reads all configurations defined in `/etc/swanctl/` and loads them into the strongSwan daemon.
     ```bash
     sudo swanctl --load-all
     ```

2. **Listing Configurations and Status**
   - **List all connection configurations**: Displays the connection profiles defined in the configuration files.
     ```bash
     sudo swanctl --list-conns
     ```
   - **List active security associations (SAs)**: Shows currently active IPsec tunnels and their status.
     ```bash
     sudo swanctl --list-sas
     ```
   - **List loaded pools**: Lists IP address pools loaded into strongSwan.
     ```bash
     sudo swanctl --list-pools
     ```

3. **Managing Connections**
   - **Initiate a connection**: Manually start a VPN connection based on a specific child SA name.
     ```bash
     sudo swanctl --initiate --child <child-sa-name>
     ```
   - **Terminate a connection**: Manually stop a VPN connection based on a specific child SA name.
     ```bash
     sudo swanctl --terminate --child <child-sa-name>
     ```

4. **Managing Certificates and Keys**
   - **List certificates**: Display all loaded certificates.
     ```bash
     sudo swanctl --list-certs
     ```
   - **Load a certificate**: Load a specific certificate file.
     ```bash
     sudo swanctl --load-creds
     ```

5. **Debugging and Logging**
   - **Show active log levels**: Display the current logging levels and modules.
     ```bash
     sudo swanctl --loglevel
     ```

### Example Usage of `swanctl`

#### Loading Configuration

To load all configurations, use:
```bash
sudo swanctl --load-all
```

#### Listing Connections

To list all configured connections:
```bash
sudo swanctl --list-conns
```

Example output:
```plaintext
myvpn: IKEv2, reauthentication every 14400s, no rekeying, dpd delay 30s
  local:  %any
  remote: %any
  local public key authentication:
    id: "CN=myvpnserver"
    cert: "CN=myvpnserver"
  remote public key authentication:
    id: "CN=myvpnclient"
    cert: "CN=myvpnclient"
  child:  net, local_ts 10.0.0.0/24, remote_ts 0.0.0.0/0, dpd action restart
```

#### Initiating and Terminating Connections

To initiate a connection named `net`:
```bash
sudo swanctl --initiate --child net
```

To terminate the same connection:
```bash
sudo swanctl --terminate --child net
```

### Configuration in `/etc/swanctl/`

#### `swanctl.conf`

Here is a detailed example of a `swanctl.conf` configuration file:

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
            certs = certs/clientcert.pem
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

### Conclusion

The `swanctl` command is a crucial tool for managing strongSwan's IPsec and IKEv2 VPN connections. By leveraging the configurations defined in `/etc/swanctl/`, administrators can easily load, initiate, and monitor secure VPN connections. Proper understanding and usage of `swanctl` commands ensure efficient and secure management of VPN infrastructure.
