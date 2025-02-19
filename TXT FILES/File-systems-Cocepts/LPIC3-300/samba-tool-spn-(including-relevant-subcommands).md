# samba-tool spn

`samba-tool spn` is a Samba utility used to manage Service Principal Names (SPNs) for machine accounts in an Active Directory environment. SPNs are essential for Kerberos authentication, as they uniquely identify service instances so that clients can request and obtain service tickets. This command group allows administrators to add, delete, list, and look up SPNs, ensuring that services are correctly registered and avoiding issues such as duplicate SPNs.

## Key Features and Purpose

- **SPN Management**:  
  Allows administrators to add, remove, and query SPNs associated with domain accounts.
  
- **Kerberos Authentication**:  
  SPNs are used by Kerberos to associate a service instance with a service logon account, which is critical for secure authentication.

- **Service Identification**:  
  Helps ensure that clients can correctly locate and authenticate to services by providing accurate SPN records.

## Common Subcommands

### 1. `add`
- **Purpose**:  
  Adds a new SPN to a specified account.
- **Usage**:
  ```bash
  samba-tool spn add <SPN> -U <admin_user>
  ```
- **Example**:
  ```bash
  samba-tool spn add HTTP/server.example.com -U Administrator
  ```
- **Notes**:  
  The `<SPN>` should follow the format `service/hostname[:port]`. You will be prompted for the administrator's password.

### 2. `delete`
- **Purpose**:  
  Removes an existing SPN from a specified account.
- **Usage**:
  ```bash
  samba-tool spn delete <SPN> -U <admin_user>
  ```
- **Example**:
  ```bash
  samba-tool spn delete HTTP/server.example.com -U Administrator
  ```
- **Notes**:  
  Use this command to resolve duplicate SPN conflicts or remove SPNs for decommissioned services.

### 3. `list`
- **Purpose**:  
  Lists all SPNs associated with a specified account, or all SPNs in the domain if no account is provided.
- **Usage**:
  ```bash
  samba-tool spn list [account]
  ```
- **Example**:
  ```bash
  samba-tool spn list server$
  ```
- **Notes**:  
  Typically used to verify the SPNs registered for a machine account (note the trailing `$` in machine accounts).

### 4. `lookup`
- **Purpose**:  
  Searches for a given SPN in the domain and returns the associated account.
- **Usage**:
  ```bash
  samba-tool spn lookup <SPN>
  ```
- **Example**:
  ```bash
  samba-tool spn lookup HTTP/server.example.com
  ```
- **Notes**:  
  Useful for verifying that an SPN is correctly associated with the intended account.

## Additional Considerations

- **Administrator Privileges**:  
  Most operations require administrative rights. Ensure that you run the commands with an account that has sufficient privileges in the domain.

- **DNS and Kerberos Integration**:  
  SPNs rely on correct DNS naming and Kerberos configurations. Verify that your DNS records and Kerberos settings are properly configured to avoid SPN resolution issues.

- **Handling Duplicate SPNs**:  
  If you encounter errors related to duplicate SPNs, use the `list` and `lookup` subcommands to diagnose conflicts and determine which account holds the duplicate entry.

## Conclusion

The `samba-tool spn` command group is an essential tool for managing SPNs in Samba-based Active Directory environments. By using subcommands such as `add`, `delete`, `list`, and `lookup`, administrators can ensure that SPNs are correctly registered, supporting robust Kerberos authentication and proper service discovery across the domain.
