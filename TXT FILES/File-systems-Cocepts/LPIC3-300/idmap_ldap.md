# idmap_ldap

`idmap_ldap` is a Samba backend that stores the mapping of Windows Security Identifiers (SIDs) to Unix user IDs (UIDs) and group IDs (GIDs) in an LDAP directory. This centralized approach is ideal for environments that already use LDAP for managing user and group information, ensuring consistent identity data across both Windows and Unix systems.

## Overview

- **Centralized Mapping**:  
  `idmap_ldap` integrates with your LDAP directory to store and retrieve UID/GID mappings for Windows accounts, eliminating the need for a local mapping file.

- **LDAP Integration**:  
  It leverages standard LDAP attributes (e.g., `uidNumber`, `gidNumber`, and `unixHomeDirectory`) to derive the mapping, ensuring that changes made in LDAP are automatically reflected in Samba.

- **Scalable and Consistent**:  
  Ideal for larger or enterprise environments, `idmap_ldap` ensures that Windows and Unix identities are managed consistently through a centralized directory.

## Key Features

- **Centralized Storage**: Stores SID-to-UID/GID mappings in a dedicated subtree of your LDAP directory.
- **Seamless Integration**: Works well in environments with existing LDAP (or AD with Unix extensions), reducing administrative overhead.
- **Consistent Identity Mapping**: Ensures that every Windows account is assigned a consistent UID/GID based on LDAP attributes.
- **Flexible Configuration**: Easily configurable on a per-domain basis in the Samba configuration file (`smb.conf`).

## Configuration

To configure `idmap_ldap` for a specific domain in your Samba configuration (`smb.conf`), add the following lines in the `[global]` section:

```ini
[global]
    # Default backend for non-domain-specific mappings
    idmap config * : backend = tdb2
    idmap config * : range = 10000-20000

    # Use idmap_ldap for your specific domain (replace YOURDOMAIN with your actual domain name)
    idmap config YOURDOMAIN : backend = ldap
    idmap config YOURDOMAIN : ldap_url = "ldap://ldap.example.com"
    idmap config YOURDOMAIN : ldap_base_dn = "ou=IDMapping,dc=example,dc=com"
    idmap config YOURDOMAIN : ldap_bind_dn = "cn=admin,dc=example,dc=com"
    idmap config YOURDOMAIN : ldap_password = your_password
```

- **`backend = ldap`**: Tells Samba to use the LDAP-based mapping for the specified domain.
- **`ldap_url`**: The URL of your LDAP server.
- **`ldap_base_dn`**: The base distinguished name in LDAP where mapping data is stored.
- **`ldap_bind_dn`** and **`ldap_password`**: Credentials for binding to LDAP to retrieve and update the mapping information.

## Use Cases

- **Enterprise Environments**:  
  Centralize identity mapping in organizations that use LDAP for user management, ensuring consistent UID and GID assignments for domain accounts.

- **Mixed OS Setups**:  
  Simplify administration in environments with both Windows and Unix systems by maintaining a single source of truth for user and group identifiers.

- **Centralized Management**:  
  Leverage existing LDAP infrastructure to manage Unix attributes, reducing the complexity of handling separate local mapping databases.

## Troubleshooting

- **LDAP Connectivity**:  
  Ensure that the Samba server can connect to your LDAP server. Tools like `ldapsearch` can help verify connectivity.

- **Credentials**:  
  Double-check the `ldap_bind_dn` and `ldap_password` for correctness and that they have sufficient permissions in LDAP.

- **UID/GID Range Conflicts**:  
  Verify that the UID/GID range specified for `idmap_ldap` does not overlap with local system accounts or other mapping backends.

## Conclusion

`idmap_ldap` is a robust backend for Samba that leverages an existing LDAP directory to manage the mapping of Windows SIDs to Unix UIDs and GIDs. By centralizing identity mapping, it simplifies administration in mixed environments and ensures consistency across systems. With proper configuration and troubleshooting, `idmap_ldap` can greatly enhance the integration of Unix systems into Windows domains.
