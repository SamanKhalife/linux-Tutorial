# radwho
`radwho` is a command-line utility provided by FreeRADIUS for querying and displaying information about currently active RADIUS sessions. It is particularly useful for administrators who need to monitor and manage RADIUS sessions in real-time.

### Key Features of `radwho`

- **Real-time Session Monitoring**: Shows active RADIUS sessions.
- **Detailed Information**: Provides details such as username, NAS (Network Access Server) IP address, session start time, and session duration.
- **Command-Line Interface**: Suitable for scripting and automation.

### Installation

`radwho` is typically included with the FreeRADIUS server package. Ensure you have the FreeRADIUS server utilities installed to use `radwho`.

On Debian-based systems:
```bash
sudo apt-get install freeradius-utils
```

On Red Hat-based systems:
```bash
sudo yum install freeradius-utils
```

### Basic Usage

The basic syntax for `radwho` is:
```bash
radwho [options]
```

### Example Commands

#### Display Active Sessions

To display all active RADIUS sessions:
```bash
radwho
```

### Output Format

The output of `radwho` typically includes columns such as:
- **Username**: The username of the authenticated user.
- **NAS-IP-Address**: The IP address of the NAS (Network Access Server) where the user is connected.
- **Start Time**: The time when the session started.
- **Session Duration**: How long the session has been active.

Here’s an example of how the output might look:
```plaintext
Username       NAS-IP-Address         Start Time           Session Duration
------------------------------------------------------------------------------
alice          192.168.1.1            2024-07-15 10:30:00  1h 15m
bob            10.0.0.5               2024-07-15 09:45:00  2h 30m
```

### Additional Options

- `-v, --verbose`: Provides more detailed information about each session.
- `-h, --help`: Displays help information and usage instructions for `radwho`.

### Usage Scenarios

#### Monitoring Active Sessions

Administrators can use `radwho` to monitor and manage active RADIUS sessions. It helps in identifying currently connected users, their session details, and monitoring session durations.

#### Scripting and Automation

Since `radwho` provides a command-line interface, it can be incorporated into scripts and automated monitoring systems to periodically check and report on active RADIUS sessions.

### Conclusion

`radwho` is a valuable tool for administrators managing RADIUS servers, providing insights into active user sessions and facilitating real-time monitoring and management. By leveraging `radwho`, administrators can ensure smooth operation of authentication services and promptly address any session-related issues.
