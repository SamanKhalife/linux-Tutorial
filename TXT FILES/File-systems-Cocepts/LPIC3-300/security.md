# security

The `security` parameter in Samba's configuration (`smb.conf`) determines how Samba handles user authentication. This setting is critical because it defines the authentication model used for accessing Samba shares and interacting with Windows domains.

## Common Values

- **`security = user`**  
  - **Description**: Users are authenticated against a local Samba user database (e.g., stored in `smbpasswd` or TDB files).
  - **Use Case**: Ideal for standalone servers or small networks where centralized authentication is not required.
  - **Example**:
    ```ini
    [global]
    security = user
    ```

- **`security = share`**  
  - **Description**: Share-level authentication is used, meaning that access is controlled by a password on the share rather than individual user accounts.
  - **Use Case**: This is a legacy mode that is largely deprecated due to security concerns.
  - **Example**:
    ```ini
    [global]
    security = share
    ```

- **`security = ADS`**  
  - **Description**: Samba uses Active Directory for authentication, leveraging Kerberos for secure, ticket-based login.
  - **Use Case**: Best suited for enterprise environments where Samba is integrated with a Windows Active Directory.
  - **Example**:
    ```ini
    [global]
    workgroup = EXAMPLE
    realm = EXAMPLE.COM
    security = ADS
    encrypt passwords = yes
    ```

- **`security = server`**  
  - **Description**: Samba delegates authentication to another server.
  - **Use Case**: Used in scenarios where authentication is managed externally.
  - **Example**:
    ```ini
    [global]
    security = server
    ```

## How It Works

- **`security = user`**:  
  Each login attempt is checked against Samba's local user database. If the user exists and the password matches, access is granted.

- **`security = share`**:  
  Access is granted based on a shared password for the resource rather than per-user credentials. This model is less secure because it does not differentiate between individual users.

- **`security = ADS`**:  
  Samba, as a domain member, communicates with Active Directory using Kerberos. It authenticates users by obtaining Kerberos tickets and uses them to access shared resources. This mode requires a properly configured Kerberos environment (`/etc/krb5.conf`) and correct domain settings.

- **`security = server`**:  
  Samba forwards authentication requests to another server which is responsible for validating user credentials.

## Example smb.conf Snippet

```ini
[global]
   # For standalone authentication using a local database
   ; security = user

   # For Active Directory integration using Kerberos
   workgroup = EXAMPLE
   realm = EXAMPLE.COM
   security = ADS
   encrypt passwords = yes

   # For delegated authentication (less common)
   ; security = server

   # Other useful settings
   log file = /var/log/samba/%m.log
   max log size = 50
```

## Conclusion

The `security` parameter is a foundational setting in Samba that governs how users are authenticated. Selecting the appropriate mode—whether it's `user` for local authentication, `ADS` for Active Directory integration, `share` for legacy setups, or `server` for delegated authentication—is essential to both the security and the functionality of your Samba deployment. Make sure your configuration matches your environment's needs, and verify that supporting services (like Kerberos for `ADS`) are correctly configured.
