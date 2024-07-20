# httpd, apache2

`httpd` and `apache2` are both command-line utilities used for managing the Apache HTTP Server, but they are typically associated with different Linux distributions and versions of Apache. Here's a breakdown of each:

### httpd

- **Usage**: `httpd` is the command used to start and manage the Apache HTTP Server on distributions like CentOS, Red Hat Enterprise Linux (RHEL), and Fedora.
  
- **Commonly Used Options**:
  - **start**: Starts the Apache server.
  - **stop**: Stops the Apache server.
  - **restart**: Restarts the Apache server.
  - **graceful**: Gracefully restarts Apache without dropping connections.
  - **configtest**: Checks the Apache configuration for syntax errors.
  - **status**: Displays the current status of Apache (running or stopped).

#### Example Usage:
```bash
# Start Apache
sudo service httpd start

# Stop Apache
sudo service httpd stop

# Restart Apache
sudo service httpd restart

# Gracefully restart Apache
sudo service httpd graceful

# Check Apache configuration syntax
sudo service httpd configtest

# Get current status of Apache
sudo service httpd status
```

### apache2

- **Usage**: `apache2` is commonly used on Debian-based distributions such as Ubuntu to manage the Apache HTTP Server.

- **Commonly Used Options**:
  - **start**: Starts the Apache server.
  - **stop**: Stops the Apache server.
  - **restart**: Restarts the Apache server.
  - **reload**: Reloads the Apache configuration without stopping the server.
  - **force-reload**: Forces a reload of Apache's configuration.
  - **status**: Displays the current status of Apache.

#### Example Usage:
```bash
# Start Apache
sudo service apache2 start

# Stop Apache
sudo service apache2 stop

# Restart Apache
sudo service apache2 restart

# Reload Apache configuration
sudo service apache2 reload

# Force reload of Apache configuration
sudo service apache2 force-reload

# Get current status of Apache
sudo service apache2 status
```

### Differences

- **Command Name**: The main difference is the command name (`httpd` vs `apache2`), which varies based on the Linux distribution.

- **Options**: While they share similar options (`start`, `stop`, `restart`, `status`), there are some differences (`graceful`, `configtest`, `reload`, `force-reload`) that are tailored to the specific distribution's way of managing services.

- **Compatibility**: Although `httpd` is more commonly associated with RHEL-based systems and `apache2` with Debian-based systems, both can often be found on different distributions due to aliases or symbolic links.

### Recommendations

- **Distribution Choice**: Use `httpd` on CentOS, RHEL, Fedora, and other RHEL-based distributions. Use `apache2` on Debian, Ubuntu, and derivatives.
  
- **Systemd Integration**: For systems using systemd, consider using `systemctl` commands (`systemctl start apache2`, `systemctl stop apache2`) for better integration and management.

Both `httpd` and `apache2` provide essential functionality for managing Apache HTTP Server instances on Linux, making them indispensable tools for web server administration tasks.
