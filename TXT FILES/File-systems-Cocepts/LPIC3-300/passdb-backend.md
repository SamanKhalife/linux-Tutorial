# passdb-backend

`passdb-backend` is a Samba configuration parameter that specifies the storage mechanism used for Samba's password database. This setting determines where and how Samba stores user account information and password hashes. The choice of backend affects performance, scalability, and integration with other directory services.

## Common Backends

- **tdbsam**:  
  - **Description**: A file-based backend that stores user information in a TDB (Trivial Database) file.
  - **Use Case**: It is the default and most common backend for small to medium-sized deployments.
  - **Example Configuration**:
    ```ini
    [global]
       passdb backend = tdbsam
    ```

- **smbpasswd**:  
  - **Description**: A traditional backend that uses the `/etc/samba/smbpasswd` file to store password hashes.
  - **Use Case**: Typically used in legacy setups; less common in modern environments due to security and scalability limitations.
  - **Example Configuration**:
    ```ini
    [global]
       passdb backend = smbpasswd
    ```

- **ldapsam**:  
  - **Description**: Uses an LDAP directory to store Samba account information. This backend is often used in environments that already use LDAP or Active Directory for centralized user management.
  - **Use Case**: Ideal for larger organizations or when integrating with a centralized directory service.
  - **Example Configuration**:
    ```ini
    [global]
       passdb backend = ldapsam:ldap://ldap.example.com
    ```

## Configuration Details

The `passdb-backend` setting is defined in the `[global]` section of the Samba configuration file (`smb.conf`). When selecting a backend, you should consider factors such as:

- **Integration Needs**:  
  For environments already using LDAP or Active Directory, `ldapsam` might be the best choice. For standalone servers, `tdbsam` is generally preferred.

- **Performance and Scalability**:  
  File-based backends like `tdbsam` work well for small-to-medium environments, while LDAP-based solutions are more scalable for large deployments.

- **Security**:  
  Ensure that the chosen backend and its associated storage (files or LDAP directory) are secured appropriately to protect sensitive user credentials.

## Example smb.conf Snippet

Below is an example configuration using the default `tdbsam` backend for local authentication:

```ini
[global]
   workgroup = MYGROUP
   server string = Samba Server
   security = user

   # Use tdbsam for storing user credentials and password hashes
   passdb backend = tdbsam
```

And an example configuration using `ldapsam`:

```ini
[global]
   workgroup = MYGROUP
   server string = Samba Server
   security = ADS
   realm = MYGROUP.COM
   encrypt passwords = yes

   # Use ldapsam for centralized user management via LDAP
   passdb backend = ldapsam:ldap://ldap.myorg.com
```

## Conclusion

The `passdb-backend` parameter is critical in determining how Samba handles user account data and password authentication. By choosing the appropriate backend (such as `tdbsam`, `smbpasswd`, or `ldapsam`), administrators can tailor Samba to meet the security, performance, and integration needs of their environment. Proper configuration and security practices ensure that user credentials are stored reliably and protected against unauthorized access.
