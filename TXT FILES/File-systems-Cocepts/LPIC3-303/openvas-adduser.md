### `openvas-adduser` Command

The `openvas-adduser` command is used to create new user accounts for the OpenVAS (Open Vulnerability Assessment System) scanner. OpenVAS is an open-source framework for vulnerability scanning and vulnerability management. Adding users to OpenVAS is an essential step for managing access control and user permissions.

#### Basic Usage

To add a new user to OpenVAS, open a terminal and use the `openvas-adduser` command:

```bash
sudo openvas-adduser
```

This command initiates an interactive process where you will be prompted to enter details for the new user.

#### Steps in Interactive Mode

1. **Enter the Username**:
    ```text
    Enter the username: <username>
    ```

2. **Select Roles**:
    You will be prompted to select roles for the user. Roles define the permissions and access levels within OpenVAS. Common roles include:
    - `Admin`: Full access to all features and configurations.
    - `User`: Limited access, typically to perform scans and view results.

    Example:
    ```text
    Roles (comma-separated, one of Admin,User): Admin
    ```

3. **Password Setup**:
    Enter a password for the new user. You will be prompted to confirm the password.
    ```text
    New password:
    Re-enter new password:
    ```

4. **Confirmation**:
    After entering the necessary details, you will receive a confirmation message indicating that the user has been successfully added.

#### Example

Here is an example session of adding a user named `john` with the `Admin` role:

```text
$ sudo openvas-adduser
Enter the username: john
Roles (comma-separated, one of Admin,User): Admin
New password:
Re-enter new password:
User john has been successfully added.
```

#### Advanced Options

In addition to the interactive mode, `openvas-adduser` can also be used with command-line options to specify user details non-interactively. This is useful for scripting and automation purposes.

Example:
```bash
sudo openvas-adduser --username=john --role=Admin --password='SecurePassword'
```

#### Security Considerations

1. **Strong Passwords**: Ensure that passwords are strong and meet the security policies of your organization.
2. **Role Assignment**: Assign roles carefully to limit access based on the principle of least privilege.
3. **Regular Audits**: Regularly review and audit user accounts and their roles to ensure proper access control.

#### Conclusion

The `openvas-adduser` command is a straightforward and essential tool for managing user accounts in OpenVAS. By following the interactive prompts or using command-line options, administrators can efficiently add and manage users, ensuring secure and controlled access to the OpenVAS vulnerability scanning system.
