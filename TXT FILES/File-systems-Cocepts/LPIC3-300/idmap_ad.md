# idmap_ad

`idmap_ad` is a Samba ID mapping backend that uses Active Directory (AD) as the source for Unix identity attributes. It directly retrieves values such as `uidNumber`, `gidNumber`, `loginShell`, and `unixHomeDirectory` from AD objects, allowing Samba to map Windows SIDs to Unix UIDs and GIDs based on data stored in the directory. This is particularly useful in environments where AD is the centralized source of user and group information.

## Overview

- **Active Directory Integration**:  
  Retrieves Unix attributes from AD, enabling Samba to use centrally managed values for user and group identities.

- **Centralized Identity Management**:  
  Ensures that updates made in AD (e.g., changes to UID/GID) are automatically reflected in Samba, promoting consistency across Windows and Unix systems.

- **Direct Mapping**:  
  Unlike file-based or algorithmic backends, `idmap_ad` reads Unix attributes directly from AD, simplifying administration in domains that support Unix extensions.

## Key Features

- **Centralized Source of Truth**:  
  Unix identity information is maintained in AD, so administrators can manage user and group attributes using AD tools.
  
- **Consistency**:  
  Provides consistent mapping of Windows SIDs to Unix IDs across all systems in the domain.

- **Simplified Administration**:  
  Eliminates the need for maintaining separate local mapping databases—changes in AD are automatically propagated to Samba.

## Configuration

To configure `idmap_ad`, add the following to your Samba configuration file (`smb.conf`) in the `[global]` section:

```ini
[global]
    # Default backend for non-domain-specific mapping
    idmap config * : backend = tdb2
    idmap config * : range = 10000-20000

    # Use idmap_ad for your specific Active Directory domain (replace YOURDOMAIN)
    idmap config YOURDOMAIN : backend = ad
    idmap config YOURDOMAIN : range = 30000-40000
```

- **`backend = ad`**:  
  Directs Samba to use the AD-based backend for the specified domain.
  
- **`range`**:  
  Specifies the range of Unix IDs allocated for domain accounts. Ensure this range does not conflict with local Unix accounts.

## Use Cases

- **Enterprise Environments**:  
  Ideal for organizations where AD is the central user management system, and Unix attributes are stored within AD.

- **Mixed-OS Environments**:  
  Facilitates seamless integration between Windows and Unix systems by providing consistent UID/GID mapping across both platforms.

- **Simplified User Management**:  
  Administrators can manage Unix identity attributes via AD, reducing the need to update multiple systems.

## Troubleshooting

- **Attribute Availability**:  
  Verify that the necessary Unix attributes (e.g., `uidNumber`, `gidNumber`) are correctly set for each user in AD.

- **UID/GID Range Conflicts**:  
  Check that the configured range in `smb.conf` does not overlap with local system IDs.

- **Connectivity Issues**:  
  Ensure that the Samba server can communicate with the AD domain controller to retrieve the required attributes.

## Conclusion

`idmap_ad` leverages Active Directory as a centralized repository for Unix identity attributes, enabling consistent and reliable mapping of Windows SIDs to Unix UIDs and GIDs. This backend simplifies administration in mixed environments by ensuring that identity data is managed in a single location, reducing redundancy and promoting consistency across systems. With proper configuration and troubleshooting, `idmap_ad` offers a robust solution for integrating Samba with Active Directory.
