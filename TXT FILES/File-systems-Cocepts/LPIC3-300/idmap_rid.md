# idmap_rid

**idmap_rid** is one of the ID mapping backends provided by Samba for mapping Windows Security Identifiers (SIDs) to Unix user IDs (UIDs) and group IDs (GIDs). Unlike file-based backends (such as idmap_tdb or idmap_tdb2) that use a local database, **idmap_rid** calculates UID/GID values directly from the Relative Identifier (RID) portion of a Windows SID. This method provides a deterministic mapping that is especially useful in environments where Samba acts as an Active Directory Domain Controller (AD DC) or a domain member server.

### Key Features and Functionality

- **Deterministic Mapping**:  
  - **RID-Based Calculation**: The backend computes the Unix ID for a given Windows SID by applying a mathematical formula to the RID (the last part of the SID).  
  - **Consistency**: This ensures that the same Windows account always maps to the same UID/GID, provided the configuration remains unchanged.

- **No External Database Required**:  
  - **Lightweight**: Since it derives UID/GID values from the SID itself, there’s no need for an external database (such as a TDB file) to store mappings.
  - **Simplicity**: This makes idmap_rid simpler to manage, with fewer moving parts, particularly in environments where ease of configuration is important.

- **Configurable Range**:  
  - Administrators can define a numerical range (using the `range` parameter) within which Unix IDs will be allocated. This helps prevent conflicts with local system accounts.
  
- **Use Case in AD Environments**:  
  - idmap_rid is commonly used in Samba Active Directory setups, where a deterministic and scalable mapping of thousands of Windows accounts is required.

### Configuration

To configure **idmap_rid** in your Samba configuration file (`smb.conf`), you typically specify it for your Active Directory domain. Here’s an example configuration snippet:

```ini
[global]
   # Set the default backend for all domains (if not overridden by domain-specific settings)
   idmap config * : backend = tdb2
   idmap config * : range = 10000-20000

   # For the specific domain (e.g., EXAMPLE)
   idmap config EXAMPLE : backend = rid
   idmap config EXAMPLE : range = 20000-30000
```

- **`backend = rid`**:  
  This setting tells Samba to use the idmap_rid backend for mapping SIDs for the specified domain (here, "EXAMPLE").

- **`range`**:  
  Specifies the range of Unix IDs that can be assigned to Windows accounts for that domain. It is important to choose a range that does not overlap with the local system accounts or other id mapping backends.


### Advantages

- **Efficiency and Scalability**:  
  idmap_rid is particularly efficient in environments with a large number of Windows accounts because it does not require the overhead of maintaining a separate mapping database.

- **Consistency Across Restarts**:  
  Because the mapping is computed from the SID, it remains consistent even if the Samba service restarts, as long as the configuration parameters (especially the UID/GID range) remain the same.

- **Simplified Administration**:  
  Administrators do not need to worry about synchronizing an external database; the mapping is automatically derived.

### Considerations and Best Practices

- **Fixed Range**:  
  Ensure that the UID/GID range specified for idmap_rid is carefully chosen to avoid conflicts with other local users and groups.

- **Domain Specificity**:  
  idmap_rid is best applied on a per-domain basis in a Samba Active Directory environment. For non-domain scenarios, other backends (like idmap_tdb2) might be more appropriate.

- **Backup Configuration**:  
  While idmap_rid doesn’t maintain a separate database file, it’s important to back up your `smb.conf` and document your UID/GID range settings. Changes to these settings later could potentially alter the mapping.

- **Testing**:  
  Before deploying changes in a production environment, test the mappings on a staging system. Use commands like `wbinfo -u` and `getent passwd` to verify that Windows domain users are correctly mapped to Unix IDs.

### Conclusion

**idmap_rid** offers a deterministic and lightweight method for mapping Windows SIDs to Unix UIDs and GIDs by leveraging the RID component of the SID. It is particularly well-suited for Samba Active Directory environments where scalability and consistency are crucial. With careful configuration of the UID/GID range in `smb.conf`, idmap_rid can provide a reliable mapping solution that integrates seamlessly with both Windows and Unix-based systems.
