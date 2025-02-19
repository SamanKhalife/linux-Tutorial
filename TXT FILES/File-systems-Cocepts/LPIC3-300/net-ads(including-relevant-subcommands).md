# net ads
The **`net ads`** command is a Samba utility used to perform a variety of Active Directory (AD) operations when Samba is running in AD domain controller (or member server) mode. It forms a subcommand group under the `net` command and is central to managing interactions with a Windows AD environment. Below is an overview of its primary functions and the most commonly used subcommands.

## Overview

**`net ads`** is designed to help administrators perform tasks such as joining or leaving an AD domain, retrieving domain information, managing trust relationships, and performing searches against the directory. Because Samba implements many Windows Active Directory features, the `net ads` command serves as the primary interface for these operations from the Linux/Unix command line.

## Common Subcommands and Their Functions

### 1. **`join`**
- **Purpose**: Adds the Samba server as a member of an Active Directory domain.
- **Usage**:  
  ```bash
  net ads join -U <domain_admin>
  ```
- **Example**:  
  ```bash
  net ads join -U Administrator
  ```
  This command will prompt for the Administrator’s password and attempt to join the domain. Successful execution registers the machine account in AD.

### 2. **`leave`**
- **Purpose**: Removes the Samba server from the Active Directory domain.
- **Usage**:  
  ```bash
  net ads leave -U <domain_admin>
  ```
- **Example**:  
  ```bash
  net ads leave -U Administrator
  ```
  Use this when decommissioning a server or when it should no longer be part of the domain.

### 3. **`testjoin`**
- **Purpose**: Checks whether the Samba server is properly joined to the AD domain.
- **Usage**:  
  ```bash
  net ads testjoin
  ```
- **Example**:  
  Running this command will return a simple status indicating whether the server’s domain join is valid or if there are any issues (for example, expired machine passwords).

### 4. **`info`**
- **Purpose**: Retrieves information about the current Active Directory domain configuration.
- **Usage**:  
  ```bash
  net ads info
  ```
- **Example**:  
  This command outputs domain details like domain name, domain SID, DC information, and other configuration parameters, which are useful for auditing and troubleshooting.

### 5. **`changetrustpw`**
- **Purpose**: Changes the trust password for the machine account used by Samba in the AD domain.
- **Usage**:  
  ```bash
  net ads changetrustpw -U <domain_admin>
  ```
- **Example**:  
  ```bash
  net ads changetrustpw -U Administrator
  ```
  This ensures that the machine account’s password remains secure and synchronized with the AD domain.

### 6. **`search`**
- **Purpose**: Performs LDAP searches against the Active Directory to locate objects such as users, groups, or computers.
- **Usage**:  
  ```bash
  net ads search "<LDAP_filter>" [attributes]
  ```
- **Example**:  
  ```bash
  net ads search "(sAMAccountName=john.doe)" cn mail
  ```
  This will search for a user with the `sAMAccountName` "john.doe" and return the common name (cn) and mail attributes.

### 7. **`spn`**
- **Purpose**: Manages Service Principal Names (SPNs) for the Samba server’s machine account.
- **Usage**:  
  ```bash
  net ads spn add <spn> -U <domain_admin>
  net ads spn delete <spn> -U <domain_admin>
  ```
- **Example**:  
  ```bash
  net ads spn add HTTP/myserver.example.com -U Administrator
  ```
  This adds an SPN for HTTP services on the machine, which is important for Kerberos authentication.

### 8. **`verify`**
- **Purpose**: Checks the health of the domain trust relationship and ensures that the AD integration is functioning properly.
- **Usage**:  
  ```bash
  net ads verify
  ```
- **Example**:  
  This command helps verify that the server’s trust relationship with the domain is intact and functioning as expected.

## Additional Subcommands

Depending on the Samba version and specific environment, you might also find subcommands for:
- **`lookup`**: To resolve names to SIDs or vice versa.
- **`dns update`**: To force a DNS update for the domain controller records (in environments where Samba handles DNS updates).


## Practical Considerations

- **Security**:  
  Use appropriate domain administrator credentials when performing operations that modify the domain state (e.g., join, leave, changetrustpw). Ensure that these commands are run in secure environments to prevent credential exposure.

- **Troubleshooting**:  
  If you encounter issues with domain join or trust relationships, start by running `net ads testjoin` and `net ads info` to gather diagnostic information.

- **Scripting**:  
  These subcommands are scriptable, allowing for automation of tasks such as periodic domain health checks or automated re-joining of servers that may have lost trust relationships.

- **Documentation & Community Resources**:  
  Samba’s official documentation and community forums (such as Samba mailing lists and the Samba Wiki) provide detailed explanations and troubleshooting tips for each subcommand.

## Conclusion

The **`net ads`** command group is essential for managing Samba’s integration with Active Directory. Its subcommands—such as `join`, `leave`, `testjoin`, `info`, `changetrustpw`, `search`, `spn`, and `verify`—allow administrators to perform critical domain operations, maintain trust relationships, and troubleshoot issues in mixed Windows/Linux environments. Mastering these commands is key to ensuring that your Samba AD deployment remains secure, stable, and fully interoperable with Windows domain services.
