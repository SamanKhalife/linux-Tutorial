# idmap_rfc2307

**idmap_rfc2307** is an ID mapping backend for Samba that integrates with LDAP directories using the RFC 2307 schema. It is designed to map Windows Security Identifiers (SIDs) to Unix user IDs (UIDs) and group IDs (GIDs) by reading and, optionally, updating UNIX attributes (such as uidNumber, gidNumber, and homeDirectory) stored in the directory.

### Key Features and Use Cases

- **LDAP-Based Mapping**:  
  - Unlike file-based backends (e.g., idmap_tdb2) or algorithmic backends (e.g., idmap_rid), **idmap_rfc2307** derives UID/GID mappings from attributes in an LDAP directory.  
  - It uses RFC 2307-standard attributes (commonly found in Active Directory if extended for Unix interoperability, or in OpenLDAP) to determine the UID and GID for a given Windows account.

- **Centralized Identity Management**:  
  - In environments where Unix account information is stored in a central LDAP directory, idmap_rfc2307 allows Samba to use those same attributes.  
  - This supports seamless integration between Unix systems and Windows domains, enabling single sign-on (SSO) and consistent identity management.

- **Bidirectional Updates**:  
  - Depending on configuration, idmap_rfc2307 can not only read but also update UNIX attributes in the LDAP directory, ensuring that UID and GID information remains consistent across platforms.

### Typical Configuration

To use idmap_rfc2307, you'll need to configure it in your Samba configuration file (`smb.conf`). A typical configuration might look like:

```ini
[global]
   # Default backend for non-AD domains or fallback
   idmap config * : backend = tdb2
   idmap config * : range = 10000-20000

   # Use idmap_rfc2307 for a specific domain (e.g., EXAMPLE)
   idmap config EXAMPLE : backend = rfc2307
   idmap config EXAMPLE : range = 30000-40000
```

- **`backend = rfc2307`**:  
  This setting specifies that the mapping for the domain (here, "EXAMPLE") will be derived from LDAP RFC 2307 attributes.
- **`range`**:  
  Although idmap_rfc2307 pulls UID/GID values from the LDAP directory, the range parameter can be used to validate that the values fall within expected limits, ensuring they don’t conflict with locally managed accounts.

- **LDAP Integration**:  
  Ensure that your LDAP directory (e.g., Active Directory or OpenLDAP) is configured to include the RFC 2307 schema attributes. In Active Directory, this might involve extending the schema to include attributes like `uidNumber`, `gidNumber`, and `unixHomeDirectory`.

### Advantages

- **Consistency Across Platforms**:  
  Using idmap_rfc2307, Unix and Windows accounts can share the same identity information, reducing administrative overhead and simplifying permission management.
  
- **Centralized Administration**:  
  UID and GID values are managed centrally within your LDAP directory, meaning changes made in LDAP are automatically reflected in Samba’s ID mapping.
  
- **Enhanced Interoperability**:  
  This backend is particularly beneficial in environments with mixed operating systems, where maintaining consistency between Windows SIDs and Unix UIDs/GIDs is essential.

### Considerations

- **LDAP Schema Requirements**:  
  Your LDAP directory must support RFC 2307 attributes. In Active Directory, this may require additional configuration or schema extension.
  
- **Read vs. Write Behavior**:  
  Depending on your configuration, idmap_rfc2307 may be set to read-only mode (merely reading UID/GID values from LDAP) or to allow updates. Decide based on your environment’s needs.
  
- **Range Overlap**:  
  Ensure that the ranges specified for idmap_rfc2307 do not conflict with those used for local accounts or other mapping backends.

### Troubleshooting and Community Insights

- **Common Issues**:  
  - Misconfigured LDAP connections or missing RFC 2307 attributes can lead to mapping failures.  
  - Logs from Samba (in `/var/log/samba/`) can provide insights if UID/GID mappings are not being applied correctly.

- **Community Feedback**:  
  - Administrators on forums and mailing lists often recommend idmap_rfc2307 for environments where central management of Unix attributes is required.  
  - It’s well-documented in the Samba Wiki and various technical guides as a robust solution for AD integration when Unix interoperability is needed.

### Conclusion

**idmap_rfc2307** is an essential backend for Samba environments that require integration with LDAP directories using the RFC 2307 schema. It enables the deterministic mapping of Windows SIDs to Unix UIDs and GIDs by utilizing existing Unix attributes in your directory, ensuring consistent identity management across both Windows and Unix systems. Proper configuration of both Samba (`smb.conf`) and your LDAP directory is critical for leveraging the benefits of idmap_rfc2307 in a mixed-OS environment.
