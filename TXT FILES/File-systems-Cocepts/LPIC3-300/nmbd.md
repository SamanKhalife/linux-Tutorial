# nmbd (NetBIOS Name Service Daemon)

`nmbd` is another key component of the **Samba** suite, responsible for providing **NetBIOS** services. It handles the **NetBIOS over IP** (NBNS) protocol, which allows machines on the same network to locate and communicate with each other by their NetBIOS names. This is primarily for compatibility with Windows systems, which rely on NetBIOS for name resolution in older networking environments.

#### **Key Functions of `nmbd`**:
1. **NetBIOS Name Resolution**: Provides mapping between NetBIOS names (human-readable) and IP addresses. This allows clients to access servers and services using friendly names instead of IP addresses.
2. **Network Browsing**: Enables browsing and discovery of shared resources like files and printers on the local network (Windows workgroup or domain).
3. **WINS Server**: Can act as a WINS (Windows Internet Name Service) server, which centralizes the NetBIOS name-to-IP address mapping across a network, improving name resolution in larger networks.

#### **Configuration and Management**:
- The `nmbd` service is configured through the **`/etc/samba/smb.conf`** file alongside `smbd`.
- It is usually installed and run together with the `smbd` service.
- Like `smbd`, it is managed using **systemd** or **service** commands.

#### **Common Commands**:
- **Start/Stop/Restart `nmbd`**:
  ```bash
  sudo systemctl start nmbd
  sudo systemctl stop nmbd
  sudo systemctl restart nmbd
  ```

- **Check Status**:
  ```bash
  sudo systemctl status nmbd
  ```

#### **Example Configuration** (`/etc/samba/smb.conf`):
```bash
[global]
   workgroup = WORKGROUP
   server string = Samba Server %v
   security = user
   wins support = yes   # nmbd acting as a WINS server
```

- **`wins support = yes`**: Configures `nmbd` to act as a WINS server. This helps in name resolution, especially for larger networks.
- **`workgroup`**: Defines the Windows workgroup the Samba server will participate in.

#### **Key Files and Directories**:
- **`/etc/samba/smb.conf`**: Main Samba configuration file, also used by `nmbd`.
- **`/var/log/samba/`**: Log files for Samba, including logs for the `nmbd` service.

#### **NetBIOS Name Resolution**:
NetBIOS names are used by Windows systems to identify each other on local networks. When a machine attempts to resolve a name, `nmbd` either responds to the request (if it knows the name) or forwards it to a WINS server.

#### **Common NetBIOS-Related Files**:
- **`/var/cache/samba/wins.dat`**: Contains the database of NetBIOS names and IP addresses if `nmbd` is acting as a WINS server.
- **`/var/log/samba/log.nmbd`**: The log file for `nmbd` service, useful for troubleshooting NetBIOS name resolution issues.

#### **Security Considerations**:
- **WINS Security**: If the Samba server is acting as a WINS server, ensure only trusted systems can access it. Misconfigured or untrusted clients can introduce invalid name records, affecting name resolution.
- **Legacy Protocol**: NetBIOS is considered somewhat outdated. For modern environments, it's recommended to use DNS-based name resolution over NetBIOS wherever possible.

#### **Integration with `smbd`**:
- **`nmbd` and `smbd` Work Together**: While `smbd` handles the file-sharing and authentication, `nmbd` facilitates the browsing and name resolution across a Windows-style network. Both services are often used in tandem, but `nmbd` is not always necessary in environments that rely solely on DNS.

#### **Disabling `nmbd`**:
If you're using DNS and no longer need NetBIOS name resolution:
- **Disable the Service**:
  ```bash
  sudo systemctl disable nmbd --now
  ```

