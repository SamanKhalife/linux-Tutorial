# radiusd

`radiusd` refers to the RADIUS (Remote Authentication Dial-In User Service) daemon, which is a server-side application responsible for authenticating, authorizing, and accounting (AAA) for remote users who connect to a network service. The RADIUS protocol is widely used for network access, including Wi-Fi, VPN, and dial-up connections.

### Key Features of RADIUS

- **Authentication**: Verifies the identity of the user or device attempting to connect.
- **Authorization**: Determines the level of access granted to the authenticated user.
- **Accounting**: Tracks the usage of network resources by users.

### FreeRADIUS: A Popular Implementation

FreeRADIUS is the most widely deployed RADIUS server and is known for its high performance, flexibility, and extensive features.

### Installation

On a Debian-based system, you can install FreeRADIUS using:
```bash
sudo apt-get install freeradius
```

On a Red Hat-based system, use:
```bash
sudo yum install freeradius
```

### Configuration Files

The main configuration files for FreeRADIUS are typically located in `/etc/freeradius/` or `/etc/raddb/`. Key configuration files include:

1. **radiusd.conf**: The main configuration file for the RADIUS daemon.
2. **clients.conf**: Specifies clients (network devices) allowed to connect to the RADIUS server.
3. **users**: Defines user authentication details.
4. **eap.conf**: Configures EAP (Extensible Authentication Protocol) methods for authentication.
5. **sites-available/**: Contains configurations for different virtual servers.
6. **mods-available/**: Contains modules that can be enabled or disabled.

### Key Configuration Files and Directories

1. **radiusd.conf**
   - **Description**: Main configuration file for the RADIUS daemon.
   - **Location**: `/etc/freeradius/radiusd.conf` or `/etc/raddb/radiusd.conf`
   - **Example Content**:
     ```plaintext
     prefix = /usr
     exec_prefix = /usr
     sysconfdir = ${prefix}/etc
     localstatedir = ${prefix}/var
     sbindir = ${exec_prefix}/sbin
     logdir = ${localstatedir}/log/radius
     raddbdir = ${sysconfdir}/raddb
     radacctdir = ${logdir}/radacct
     ```
   
2. **clients.conf**
   - **Description**: Defines clients allowed to connect to the RADIUS server.
   - **Location**: `/etc/freeradius/clients.conf` or `/etc/raddb/clients.conf`
   - **Example Content**:
     ```plaintext
     client localhost {
         ipaddr = 127.0.0.1
         secret = testing123
         require_message_authenticator = no
         nastype = other
     }
     ```

3. **users**
   - **Description**: Defines user authentication details.
   - **Location**: `/etc/freeradius/users` or `/etc/raddb/users`
   - **Example Content**:
     ```plaintext
     alice Cleartext-Password := "password123"
     bob Cleartext-Password := "password456"
     ```

4. **eap.conf**
   - **Description**: Configures EAP methods for authentication.
   - **Location**: `/etc/freeradius/eap.conf` or `/etc/raddb/eap.conf`
   - **Example Content**:
     ```plaintext
     eap {
         default_eap_type = md5
         timer_expire = 60
         ignore_unknown_eap_types = no
         cisco_accounting_username_bug = no
         md5 {
         }
     }
     ```

5. **sites-available/**
   - **Description**: Contains configurations for different virtual servers.
   - **Location**: `/etc/freeradius/sites-available/` or `/etc/raddb/sites-available/`
   - **Example Content**:
     ```plaintext
     default {
         authorize {
             preprocess
             chap
             mschap
             digest
             suffix
             eap {
                 ok = return
             }
         }
         authenticate {
             Auth-Type PAP {
                 pap
             }
         }
     }
     ```

6. **mods-available/**
   - **Description**: Contains modules that can be enabled or disabled.
   - **Location**: `/etc/freeradius/mods-available/` or `/etc/raddb/mods-available/`
   - **Example Modules**:
     - **sql**: Integrates SQL databases.
     - **ldap**: Integrates LDAP for user authentication.
     - **eap**: Configures EAP for wireless authentication.

### Starting and Managing the RADIUS Daemon

To start the RADIUS daemon:
```bash
sudo systemctl start freeradius
```

To enable the RADIUS daemon to start on boot:
```bash
sudo systemctl enable freeradius
```

To check the status of the RADIUS daemon:
```bash
sudo systemctl status freeradius
```

### Testing RADIUS Configuration

Use the `radtest` utility to test the RADIUS server configuration:
```bash
radtest alice password123 localhost 0 testing123
```
This command tests authentication for the user `alice` with the password `password123` against the RADIUS server running on `localhost`.

### Troubleshooting

1. **Logs**: Check the logs for errors and issues.
   - Location: `/var/log/radius/radius.log`

2. **Debug Mode**: Run FreeRADIUS in debug mode for detailed output.
   ```bash
   sudo freeradius -X
   ```

3. **Common Issues**:
   - Incorrect client configuration in `clients.conf`.
   - Mismatched shared secrets between the RADIUS server and clients.
   - Incorrect user credentials or attributes in the `users` file.

### Security Considerations

- **Strong Shared Secrets**: Use strong, unique shared secrets in `clients.conf`.
- **Secure Protocols**: Configure and use secure authentication protocols (e.g., EAP-TLS).
- **Access Control**: Limit access to the RADIUS server by properly configuring `clients.conf` and firewall rules.

### Conclusion

The `radiusd` daemon is essential for implementing robust AAA services in network environments. Proper configuration and management of FreeRADIUS, along with regular monitoring and security practices, ensure a secure and efficient RADIUS deployment. By understanding and utilizing the configuration files and tools provided by FreeRADIUS, administrators can effectively manage network access and authentication.
