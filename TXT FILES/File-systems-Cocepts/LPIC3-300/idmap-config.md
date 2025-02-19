# idmap-config

The `idmap config` directives in Samba are used to configure how Windows Security Identifiers (SIDs) are mapped to Unix user IDs (UIDs) and group IDs (GIDs). This mapping is essential for integrating Samba with Windows Active Directory or when operating in mixed environments, ensuring that file permissions and identity management work seamlessly between Windows and Unix systems.

## Overview

- **Purpose**:  
  `idmap config` specifies which backend to use for the SID-to-UID/GID mapping and defines the numerical range for these mappings. This configuration allows Samba to translate Windows account identities into Unix IDs, ensuring that domain users have proper file access permissions on Unix systems.

- **Backends**:  
  Samba supports several ID mapping backends, including:
  - **tdb / tdb2**: File-based mapping using a local database.
  - **rid**: Algorithmically derives UID/GID values from the Relative Identifier (RID) portion of the SID.
  - **rfc2307**: Uses LDAP attributes (e.g., uidNumber and gidNumber) based on the RFC 2307 schema.
  - **ldap**: Stores and retrieves mappings from an LDAP directory.
  - **ad**: Retrieves Unix attributes directly from Active Directory (AD) when AD is extended with Unix attributes.

## Basic Configuration

The `idmap config` settings are placed in the `[global]` section of your Samba configuration file (`smb.conf`). Here’s a basic example:

```ini
[global]
   # Default mapping for any domain not explicitly configured
   idmap config * : backend = tdb2
   idmap config * : range = 10000-20000

   # Specific mapping for your Active Directory domain (e.g., EXAMPLE)
   idmap config EXAMPLE : backend = rid
   idmap config EXAMPLE : range = 20000-30000
```

- **`backend`**:  
  Specifies the mapping backend to use (e.g., `tdb2`, `rid`, `rfc2307`, `ldap`, or `ad`).

- **`range`**:  
  Defines the range of UIDs and GIDs that Samba will assign to Windows accounts for that particular domain or mapping. This range should not overlap with locally managed Unix accounts.

## Detailed Options

Each mapping backend might support additional options. Here are a few examples:

- **For `idmap_rid`**:
  ```ini
  idmap config EXAMPLE : backend = rid
  idmap config EXAMPLE : range = 20000-30000
  ```
  The `rid` backend uses the Relative Identifier from the Windows SID to calculate Unix IDs.

- **For `idmap_rfc2307`**:
  ```ini
  idmap config EXAMPLE : backend = rfc2307
  idmap config EXAMPLE : range = 30000-40000
  ```
  This backend reads Unix attributes directly from the LDAP directory (or AD with Unix extensions).

- **For `idmap_ldap`**:
  ```ini
  idmap config EXAMPLE : backend = ldap
  idmap config EXAMPLE : ldap_url = "ldap://ldap.example.com"
  idmap config EXAMPLE : ldap_base_dn = "ou=IDMapping,dc=example,dc=com"
  idmap config EXAMPLE : ldap_bind_dn = "cn=admin,dc=example,dc=com"
  idmap config EXAMPLE : ldap_password = your_password
  ```
  The `ldap` backend stores mapping information in an LDAP directory.

## Best Practices

- **Non-overlapping Ranges**:  
  Ensure that the UID/GID ranges defined for different backends or domains do not overlap with local system accounts.

- **Consistency**:  
  Use a consistent mapping strategy across your environment to simplify administration and troubleshooting.

- **Backup and Documentation**:  
  Document your `idmap config` settings and back up your Samba configuration file, so you can easily restore or replicate the setup if needed.

- **Testing**:  
  After making changes to `idmap config`, use commands like `wbinfo -u` and `getent passwd` to verify that Windows domain users are correctly mapped to Unix IDs.

## Conclusion

The `idmap config` directives are central to Samba's ability to integrate with Windows domains by mapping Windows SIDs to Unix UIDs/GIDs. By carefully selecting the appropriate backend and defining a proper range, administrators can ensure seamless interoperability and consistent file permission management across mixed-OS environments. Proper configuration and testing of these settings are crucial for a secure and well-functioning Samba deployment.
