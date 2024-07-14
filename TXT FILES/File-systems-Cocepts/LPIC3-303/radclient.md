# radclient

`radclient` is a versatile command-line utility used for sending RADIUS packets to a RADIUS server and receiving responses. It is part of the FreeRADIUS package and serves as a diagnostic tool for testing various aspects of RADIUS server functionality, such as authentication, authorization, and accounting (AAA).

### Key Features of `radclient`

- **Flexible Packet Generation**: Allows customization of RADIUS packets for different types of testing.
- **Diagnostic Tool**: Useful for troubleshooting and verifying RADIUS server configurations.
- **Command-Line Based**: Can be scripted for automated testing and integration into larger systems.

### Installation

`radclient` is included with the FreeRADIUS client utilities package. To install it, ensure you have the FreeRADIUS client utilities installed.

On Debian-based systems:
```bash
sudo apt-get install freeradius-utils
```

On Red Hat-based systems:
```bash
sudo yum install freeradius-utils
```

### Basic Usage

The basic syntax for `radclient` is:
```bash
radclient [options] <radius-server> <port> <shared-secret> <packet-type> [packet-attributes...]
```

- `<radius-server>`: The IP address or hostname of the RADIUS server.
- `<port>`: The port number on which the RADIUS server is listening (default is 1812).
- `<shared-secret>`: The shared secret between the client and the RADIUS server.
- `<packet-type>`: The type of RADIUS packet (`auth`, `acct`, `status`, `coa`, etc.).
- `[packet-attributes...]`: Additional attributes specific to the RADIUS packet type.

### Example Commands

#### Authentication Request

To send an authentication request (`Access-Request`) for user `alice` with password `password123`:
```bash
radclient -x <radius-server> auth <shared-secret> <<< "User-Name=alice,User-Password=password123"
```

#### Accounting Request

To send an accounting request (`Accounting-Request`) for user `alice`:
```bash
radclient -x <radius-server> acct <shared-secret> <<< "User-Name=alice,Acct-Status-Type=Start"
```

#### Status Server Request

To send a status server request (`Status-Server`) to check server status:
```bash
radclient -x <radius-server> status <shared-secret>
```

### Detailed Example

Consider you have a RADIUS server running on `192.168.1.100` with the shared secret `radiusSecret`. You want to test the authentication for user `bob` with the password `bobPassword`.

1. Open your terminal.
2. Execute the `radclient` command as follows for authentication:
   ```bash
   radclient -x 192.168.1.100 auth radiusSecret <<< "User-Name=bob,User-Password=bobPassword"
   ```

### Interpreting the Output

The `-x` option in `radclient` displays detailed output, showing the RADIUS packets exchanged between the client and server. Here’s an example output for an authentication request:

```plaintext
Sending Access-Request of id 45 to 192.168.1.100 port 1812
        User-Name = "bob"
        User-Password = "bobPassword"
        NAS-IP-Address = 127.0.1.1
        NAS-Port = 0
        Message-Authenticator = 0x00
Received Access-Accept Id 45 from 192.168.1.100:1812 to 0.0.0.0:0 length 20
```

- **Access-Request**: Indicates that an authentication request was sent to the RADIUS server.
- **Access-Accept**: Indicates that the RADIUS server accepted the authentication request, meaning the credentials were correct.

If the authentication fails, you might see `Access-Reject` instead of `Access-Accept`.

### Troubleshooting Tips

- **Shared Secret Mismatch**: Ensure that the shared secret specified in the `radclient` command matches the one configured on the RADIUS server.
- **User Credentials**: Verify that the username and password are correct and exist in the RADIUS server's user database.
- **Network Issues**: Ensure that the RADIUS server is reachable from the client machine and there are no firewall rules blocking the communication on port 1812.
- **RADIUS Server Logs**: Check the logs on the RADIUS server for any error messages or additional details about the failed authentication attempt.

### Conclusion

`radclient` is a powerful tool for testing and troubleshooting RADIUS server configurations. It allows administrators to simulate RADIUS client behavior and verify server responses in various scenarios. By mastering `radclient` and understanding its capabilities, administrators can effectively manage and troubleshoot their RADIUS environments.
