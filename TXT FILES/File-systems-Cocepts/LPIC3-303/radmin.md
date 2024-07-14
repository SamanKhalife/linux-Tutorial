# radmin

`radmin` is a control and monitoring utility for the FreeRADIUS server, providing administrators with a command-line interface to interact with the FreeRADIUS server in real-time. It allows administrators to perform various tasks, such as managing users, checking server status, and manipulating the server's configuration without needing to restart the server.

### Key Features of `radmin`

- **Real-time Control**: Manage and configure the FreeRADIUS server dynamically.
- **Monitoring**: Check the status of the server and view statistics.
- **User Management**: Add, remove, or modify user configurations.
- **Debugging**: Assist in debugging and troubleshooting the server.
- **Virtual Server Management**: Enable or disable virtual servers.

### Installation

`radmin` is included with the FreeRADIUS server package. If you have FreeRADIUS installed, you should already have `radmin` available. If not, install FreeRADIUS first.

For Debian-based systems:
```bash
sudo apt-get install freeradius
```

For Red Hat-based systems:
```bash
sudo yum install freeradius
```

### Usage

To use `radmin`, you typically need administrative privileges. Here are some common tasks and commands you can perform with `radmin`:

#### Starting `radmin`

```bash
sudo radmin
```

This command starts the `radmin` interface, where you can enter various commands to control and monitor the FreeRADIUS server.

### Common `radmin` Commands

1. **Check Server Status**
   ```bash
   radmin> status
   ```
   Displays the status of the FreeRADIUS server, including uptime, requests processed, and other statistics.

2. **Show Server Statistics**
   ```bash
   radmin> stats client
   radmin> stats detail
   ```
   Displays detailed statistics about the server's performance and client interactions.

3. **Manage Virtual Servers**
   - **Enable a Virtual Server**
     ```bash
     radmin> server start <server_name>
     ```
   - **Disable a Virtual Server**
     ```bash
     radmin> server stop <server_name>
     ```

4. **Manage Clients**
   - **Add a Client**
     ```bash
     radmin> add client <client_name> <ip_address> <secret>
     ```
   - **Remove a Client**
     ```bash
     radmin> del client <client_name>
     ```

5. **Check Authenticated Users**
   ```bash
   radmin> show users
   ```
   Lists the currently authenticated users.

6. **Reload Configuration**
   ```bash
   radmin> reload
   ```
   Reloads the FreeRADIUS server configuration without restarting the server.

7. **Debugging and Logging**
   - **View Debug Output**
     ```bash
     radmin> debug level <level>
     ```
   - **Set Logging Level**
     ```bash
     radmin> log level <level>
     ```

### Example Usage Scenarios

#### Adding a New Client

To add a new client named `example-client` with an IP address `192.168.1.100` and a secret `example-secret`:
```bash
sudo radmin
radmin> add client example-client 192.168.1.100 example-secret
```

#### Enabling a Virtual Server

To enable a virtual server named `example-server`:
```bash
sudo radmin
radmin> server start example-server
```

#### Checking Server Status and Statistics

To check the overall status and detailed statistics of the server:
```bash
sudo radmin
radmin> status
radmin> stats detail
```

### Security Considerations

- **Restricted Access**: Ensure `radmin` is only accessible by authorized administrators.
- **Secure Communications**: Use secure channels when remotely accessing the FreeRADIUS server.
- **Strong Authentication**: Implement strong authentication mechanisms for users and clients.

### Conclusion

`radmin` is a powerful tool for managing and monitoring the FreeRADIUS server in real-time. It provides administrators with the ability to dynamically control server operations, manage users and clients, and troubleshoot issues efficiently. By leveraging `radmin`, administrators can ensure the FreeRADIUS server is running optimally and securely.
