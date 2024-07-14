# /etc/raddb/
This directory houses various configuration files crucial for setting up and managing the FreeRADIUS server. Below are some of the key files and directories typically found under `/etc/raddb/`:

### Key Files and Directories in `/etc/raddb/`

1. **`clients.conf`**: Defines the clients (NAS - Network Access Server) that are allowed to authenticate against the RADIUS server. Each client entry includes IP addresses, shared secrets, and other parameters.

2. **`users`**: This file contains user-specific configuration details for authentication, authorization, and accounting. It defines user credentials, access policies, and group memberships.

3. **`radiusd.conf`**: As discussed earlier, this is the main configuration file for FreeRADIUS, containing global server settings, module configurations, authentication methods, and more.

4. **`sites-available/` and `sites-enabled/`**: These directories contain configuration snippets or "sites" that define different service profiles or policies (e.g., `default`, `inner-tunnel`). The `sites-enabled/` directory includes symbolic links to active sites.

5. **`mods-available/` and `mods-enabled/`**: These directories house configuration modules (`mods`) that extend FreeRADIUS functionality. `mods-enabled/` contains symbolic links to enabled modules.

6. **`certs/`**: This directory stores TLS/SSL certificates and keys used for secure RADIUS communication (if configured).

7. **`dictionary`**: Defines attribute-value pairs (AVPs) used in RADIUS communication. It maps human-readable names to RADIUS attribute codes and is essential for interpreting and generating RADIUS packets.

8. **`policy.d/`**: Contains policy configuration snippets that define specific rules or conditions for handling authentication, authorization, and accounting requests.

9. **`sql/`**: If using SQL-based backends for storing user credentials, accounting data, or session information, this directory contains SQL schema files and configurations.

10. **`log/`**: Typically not present by default, but may be used for storing FreeRADIUS log files if configured to a custom location.

### Usage and Configuration

- **Editing Configuration Files**: Use a text editor with appropriate permissions (usually requiring `sudo` or root access) to modify configuration files.
  
- **Restarting FreeRADIUS**: After making changes to configuration files under `/etc/raddb/`, restart the FreeRADIUS service to apply changes:
  
  ```bash
  sudo systemctl restart freeradius
  ```
  
- **Testing**: Use tools like `radtest` or `radclient` (as discussed earlier) to test authentication, authorization, and accounting functionalities after making configuration changes.

### Security Considerations

- **Secure Shared Secrets**: Ensure that shared secrets in `clients.conf` are strong and securely managed to prevent unauthorized access.
  
- **TLS/SSL Configuration**: If using TLS/SSL for secure RADIUS communication, regularly update certificates and keys, and adhere to best practices for certificate management.

### Conclusion

Understanding the structure and contents of `/etc/raddb/` is crucial for effectively configuring and managing FreeRADIUS deployments. It forms the backbone of RADIUS server operations, facilitating secure and efficient network access control through authentication, authorization, and accounting mechanisms.
