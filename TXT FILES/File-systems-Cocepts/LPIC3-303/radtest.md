# radtest
`radtest` is a command-line utility that comes with FreeRADIUS. It is used to test RADIUS authentication by sending Access-Request packets to a RADIUS server and displaying the server's response. This tool is invaluable for administrators when they need to verify the functionality and configuration of their RADIUS server setup.

### Key Features of `radtest`

- **Simplicity**: Provides an easy way to test RADIUS authentication.
- **Diagnostic Tool**: Useful for troubleshooting and verifying RADIUS server configurations.
- **Command-Line Based**: Can be easily scripted for automated testing.

### Installation

`radtest` is included with the FreeRADIUS client utilities package. To install it, you need to install the FreeRADIUS client utilities.

On Debian-based systems:
```bash
sudo apt-get install freeradius-utils
```

On Red Hat-based systems:
```bash
sudo yum install freeradius-utils
```

### Basic Usage

The basic syntax for `radtest` is as follows:
```bash
radtest [options] <username> <password> <radius-server> <nas-port> <secret> [nas-ip-address]
```

- `<username>`: The username you want to authenticate.
- `<password>`: The password for the user.
- `<radius-server>`: The IP address or hostname of the RADIUS server.
- `<nas-port>`: The NAS port number. Typically set to 0 for testing.
- `<secret>`: The shared secret between the NAS and the RADIUS server.
- `[nas-ip-address]`: (Optional) The IP address of the NAS.

### Example Commands

#### Basic Authentication Test

To test authentication for user `alice` with password `password123` against a RADIUS server at `localhost` with shared secret `testing123`:
```bash
radtest alice password123 localhost 0 testing123
```

#### Specifying NAS IP Address

To specify the NAS IP address as `192.168.1.1`:
```bash
radtest alice password123 localhost 0 testing123 192.168.1.1
```

### Detailed Example

Consider you have a RADIUS server running on `192.168.1.100` with the shared secret `radiusSecret`. You want to test the authentication for user `bob` with the password `bobPassword`.

1. Open your terminal.
2. Execute the `radtest` command as follows:
   ```bash
   radtest bob bobPassword 192.168.1.100 0 radiusSecret
   ```

### Interpreting the Output

The `radtest` command will provide output that helps you understand the RADIUS server's response. Here’s an example output:

```plaintext
Sending Access-Request of id 45 to 192.168.1.100 port 1812
        User-Name = "bob"
        User-Password = "bobPassword"
        NAS-IP-Address = 127.0.1.1
        NAS-Port = 0
        Message-Authenticator = 0x00
        Cleartext-Password = "bobPassword"
Received Access-Accept Id 45 from 192.168.1.100:1812 to 0.0.0.0:0 length 20
```

- **Access-Request**: Indicates that an authentication request was sent to the RADIUS server.
- **Access-Accept**: Indicates that the RADIUS server accepted the authentication request, meaning the credentials were correct.

If the authentication fails, you might see `Access-Reject` instead of `Access-Accept`.

### Troubleshooting Tips

- **Shared Secret Mismatch**: Ensure that the shared secret specified in the `radtest` command matches the one configured on the RADIUS server.
- **User Credentials**: Verify that the username and password are correct and exist in the RADIUS server's user database.
- **Network Issues**: Ensure that the RADIUS server is reachable from the client machine and there are no firewall rules blocking the communication on port 1812.
- **RADIUS Server Logs**: Check the logs on the RADIUS server for any error messages or additional details about the failed authentication attempt.

### Conclusion

`radtest` is a powerful and straightforward tool for testing and verifying RADIUS authentication setups. It helps administrators diagnose and resolve configuration issues, ensuring that the RADIUS server is functioning correctly. By understanding how to use `radtest` and interpret its output, administrators can efficiently manage and troubleshoot their RADIUS environments.
