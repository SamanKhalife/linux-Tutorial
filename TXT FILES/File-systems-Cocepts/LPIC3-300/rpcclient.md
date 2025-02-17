# rpcclient
`rpcclient` is a command-line tool that allows you to interact with the Remote Procedure Call (RPC) service on a Windows or Samba server. It is a part of the **Samba** suite of tools and is primarily used to communicate with a Windows server or Samba share via the SMB (Server Message Block) protocol.

The `rpcclient` tool provides a way to query, configure, and interact with various aspects of the Windows network environment, such as user and group information, services, shares, and more.

### Overview of `rpcclient`

`rpcclient` allows you to interact with and manage the Windows RPC services on remote systems. It can perform tasks such as listing users, querying shares, and managing accounts on a Windows server or Samba domain controller. 

It is commonly used for:
- Enumerating shares and users on a server.
- Querying user groups and managing user accounts.
- Performing administrative tasks on remote machines.
- Checking for service status and other system-level information.

### General Syntax
```bash
rpcclient [OPTIONS] <hostname> -U <username>
```

Where:
- **`hostname`** is the name or IP address of the remote Windows or Samba machine.
- **`-U <username>`** specifies the username to authenticate with (can include domain, e.g., `domain\user`).

### Common `rpcclient` Commands

1. **Connecting to a Remote Host**:
   To connect to a remote Windows or Samba machine, you need to specify the machine's hostname or IP and provide a username (and optionally a password).
   ```bash
   rpcclient <hostname> -U <username>
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator
   ```

2. **Listing Shares on a Remote Server**:
   You can list all shared resources (shares) available on the server.
   ```bash
   rpcclient <hostname> -U <username> -c "netshareenum"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "netshareenum"
   ```

3. **Enumerating Users**:
   To list all users on the remote machine:
   ```bash
   rpcclient <hostname> -U <username> -c "enumdomusers"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "enumdomusers"
   ```

4. **Querying Groups**:
   To list the groups on the remote server:
   ```bash
   rpcclient <hostname> -U <username> -c "enumdomgroups"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "enumdomgroups"
   ```

5. **Getting Information About a Specific User**:
   You can retrieve information about a specific user, such as their SID (Security Identifier).
   ```bash
   rpcclient <hostname> -U <username> -c "queryuser <username>"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "queryuser john"
   ```

6. **Changing User Password**:
   To change the password for a user:
   ```bash
   rpcclient <hostname> -U <username> -c "changepassword <target_user> <new_password>"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "changepassword john newpassword123"
   ```

7. **Checking Shares and Permissions**:
   You can query the shares and permissions for a given share:
   ```bash
   rpcclient <hostname> -U <username> -c "netshareenumall"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "netshareenumall"
   ```

8. **Listing User Groups for a Specific User**:
   To get the groups a user belongs to:
   ```bash
   rpcclient <hostname> -U <username> -c "getdomgroups <username>"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "getdomgroups john"
   ```

9. **Listing Available Services on a Remote Machine**:
   To list all active services on the remote system:
   ```bash
   rpcclient <hostname> -U <username> -c "srvsvc"
   ```
   Example:
   ```bash
   rpcclient 192.168.1.10 -U administrator -c "srvsvc"
   ```

10. **Executing Commands on Remote Machines**:
    Some versions of `rpcclient` allow for running commands remotely. You can use this feature to execute processes or scripts on remote systems.

### Advanced Commands

- **`net share`**:
  This command allows you to manage network shares, including creating or deleting shares.
  ```bash
  rpcclient <hostname> -U <username> -c "net share <share_name> /delete"
  ```

- **`adduser`**:
  Adds a user to the remote system. Typically used for administrative tasks like adding new users remotely.
  ```bash
  rpcclient <hostname> -U <username> -c "adduser <username> <password>"
  ```

- **`addgroup`**:
  Adds a group to the remote machine.
  ```bash
  rpcclient <hostname> -U <username> -c "addgroup <groupname>"
  ```

- **`deluser`**:
  Removes a user from the remote machine.
  ```bash
  rpcclient <hostname> -U <username> -c "deluser <username>"
  ```

### Example Use Cases

#### 1. List available shares on a server:
```bash
rpcclient 192.168.1.10 -U administrator -c "netshareenum"
```

#### 2. List all users in the domain:
```bash
rpcclient 192.168.1.10 -U administrator -c "enumdomusers"
```

#### 3. Query information about a specific user:
```bash
rpcclient 192.168.1.10 -U administrator -c "queryuser john"
```

#### 4. Change the password of a user:
```bash
rpcclient 192.168.1.10 -U administrator -c "changepassword john newpassword123"
```

### Security Considerations

- **Authentication**: `rpcclient` requires authentication (username and password). Ensure that credentials are securely managed.
- **Permissions**: Use of `rpcclient` requires appropriate permissions on the remote machine. Administrative privileges are often necessary to perform tasks like modifying user passwords, creating/deleting shares, etc.
- **Network Access**: Since `rpcclient` communicates over SMB (Port 445), ensure that the necessary ports are open and accessible between the client and the remote machine.

### Conclusion

`rpcclient` is an essential tool for interacting with Windows servers or Samba shares over the network. It can perform a wide variety of administrative tasks remotely, including querying system information, managing users and shares, and configuring services.
