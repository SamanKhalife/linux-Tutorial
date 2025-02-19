# [homes]

The `[homes]` share in Samba is a special, dynamic share that automatically provides each user with access to their own home directory. This feature simplifies user management by eliminating the need to create and maintain individual share definitions for each user.

## Purpose

- **Automatic Home Directory Access**:  
  When a user connects to the Samba server, the `[homes]` share dynamically maps to their personal home directory based on their username.

- **Centralized Configuration**:  
  Administrators can define a single `[homes]` share in the `smb.conf` file, and Samba will automatically resolve and provide each user's home directory without additional configuration.

- **Enhanced Security**:  
  Typically configured with settings that restrict access so that only the owner can access their home directory.

## Typical Configuration

A common configuration in the Samba configuration file (`smb.conf`) for the `[homes]` share looks like this:

```ini
[homes]
   comment = Home Directories
   browseable = no
   writable = yes
   valid users = %S
   create mask = 0700
   directory mask = 0700
   template homedir = /home/%U
   template shell = /bin/bash
```

### Explanation of Options

- **`comment`**:  
  A descriptive label for the share (used by network browsers).

- **`browseable = no`**:  
  Hides the `[homes]` share from being listed in network browsing, enhancing security by preventing users from seeing other users' home directories.

- **`writable = yes`**:  
  Allows users to write to their home directories.

- **`valid users = %S`**:  
  Restricts access so that only the user accessing the share can see their home directory. `%S` is dynamically replaced with the session's username.

- **`create mask` and `directory mask`**:  
  Set the default permissions for newly created files and directories to ensure privacy (typically granting full permissions to the owner only).

- **`template homedir = /home/%U`**:  
  Defines the path template for user home directories, where `%U` is replaced by the username.

- **`template shell = /bin/bash`**:  
  Sets the default login shell for users when accessed via Samba (if applicable).

## How It Works

- When a user logs in, Samba dynamically constructs the path to their home directory using the `template homedir` setting.
- The `[homes]` share is automatically mapped to that directory, and access is restricted by the `valid users` directive.
- This process ensures that each user only accesses their own home directory and that file permissions remain secure.

## Use Cases

- **Enterprise and Domain Environments**:  
  Provides centralized management of user home directories in environments where users log in to a domain, simplifying the provisioning of personal storage.

- **Small Office/Home Networks**:  
  Eliminates the need to manually configure shares for each user, allowing for an easy-to-manage file-sharing solution.

## Security Considerations

- **Access Control**:  
  The `valid users` directive helps ensure that only the owner can access their home directory.
- **Permissions**:  
  The `create mask` and `directory mask` settings enforce strict file and directory permissions, protecting user data.
- **Visibility**:  
  Setting `browseable = no` keeps the share hidden from general network listings, reducing the risk of unauthorized access.

## Conclusion

The `[homes]` share in Samba is a powerful feature for dynamically providing each user with their personal home directory. Its automatic mapping, combined with robust security settings, simplifies administration and enhances security in both small and large environments. By properly configuring `[homes]` in your `smb.conf`, you can ensure a consistent, secure, and user-friendly file sharing experience.
