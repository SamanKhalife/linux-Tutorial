# Additional Samba Commands: smbcontrol, smbstatus, testparm, smbpasswd, nmblookup

In addition to the main Samba daemons (`smbd`, `nmbd`, and `winbindd`), there are several useful command-line utilities that help manage and troubleshoot Samba. These include `smbcontrol`, `smbstatus`, `testparm`, `smbpasswd`, and `nmblookup`.

#### smbcontrol

`smbcontrol` is used to send messages to Samba daemons (`smbd`, `nmbd`, `winbindd`) to control their behavior or request information.

- **Usage**: 
  ```bash
  smbcontrol <daemon> <message-type> [parameters]
  ```

- **Common Commands**:
  - **smbcontrol all reload-config**: Reloads the configuration file for all running daemons.
  - **smbcontrol smbd debug 10**: Sets the debug level of `smbd` to 10.
  - **smbcontrol winbindd ping**: Sends a ping to the `winbindd` daemon.

- **Example**:
  ```bash
  sudo smbcontrol all reload-config
  ```

#### smbstatus

`smbstatus` is used to display current Samba connections, open files, and locked files.

- **Usage**:
  ```bash
  smbstatus [options]
  ```

- **Common Options**:
  - **-S**: Show share mode locks.
  - **-L**: Show byte range locks.
  - **-p**: Show process IDs of the connected clients.

- **Example**:
  ```bash
  smbstatus
  ```

  Output example:
  ```plaintext
  Samba version 4.13.3
  PID     Username     Group        Machine  Protocol Version  Encryption           Signing             
  ----------------------------------------------------------------------------------------------------------------
  1234    john         users        192.168.1.5 (ipv4:192.168.1.5:12345) NT1
  ```

#### testparm

`testparm` is a utility to check the Samba configuration file (`smb.conf`) for syntax errors.

- **Usage**:
  ```bash
  testparm [configuration file] [hostname host IP]
  ```

- **Common Usage**:
  - Validate the default configuration file:
    ```bash
    testparm
    ```

- **Example**:
  ```bash
  testparm /etc/samba/smb.conf
  ```

  This checks the specified configuration file for syntax errors and displays the loaded configuration.

#### smbpasswd

`smbpasswd` is used to manage Samba user accounts and passwords.

- **Usage**:
  ```bash
  smbpasswd [options] [username]
  ```

- **Common Commands**:
  - **smbpasswd -a username**: Adds a new Samba user.
  - **smbpasswd -x username**: Deletes a Samba user.
  - **smbpasswd username**: Changes the password for the specified user.

- **Example**:
  ```bash
  sudo smbpasswd -a john
  ```

  This adds a new Samba user named `john` and prompts for a password.

#### nmblookup

`nmblookup` is used for NetBIOS name resolution, similar to `nslookup` but for NetBIOS names.

- **Usage**:
  ```bash
  nmblookup [options] name
  ```

- **Common Options**:
  - **-A IP-address**: Queries the specified IP address for its NetBIOS name.
  - **-R**: Sends the query as a broadcast.

- **Example**:
  ```bash
  nmblookup -A 192.168.1.5
  ```

  This queries the NetBIOS names registered by the specified IP address.

### Summary of Commands and Usage

| Command     | Purpose                                      | Common Usage                                              | Example Command                                      |
|-------------|----------------------------------------------|-----------------------------------------------------------|------------------------------------------------------|
| `smbcontrol`| Sends control messages to Samba daemons      | `smbcontrol all reload-config`                            | `sudo smbcontrol smbd debug 10`                      |
| `smbstatus` | Displays current Samba connections and files | `smbstatus`                                               | `smbstatus`                                          |
| `testparm`  | Checks Samba configuration for syntax errors | `testparm`                                                | `testparm /etc/samba/smb.conf`                       |
| `smbpasswd` | Manages Samba user accounts and passwords    | `smbpasswd -a username`, `smbpasswd username`             | `sudo smbpasswd -a john`                             |
| `nmblookup` | Performs NetBIOS name queries                | `nmblookup -A IP-address`                                 | `nmblookup -A 192.168.1.5`                           |

### Conclusion

These Samba command-line utilities (`smbcontrol`, `smbstatus`, `testparm`, `smbpasswd`, and `nmblookup`) provide essential tools for managing, troubleshooting, and configuring Samba services. Understanding their usage and options enhances an administrator's ability to maintain and secure Samba servers in a network environment.
