# realm

The `realm` parameter in Samba specifies the Kerberos realm or Active Directory (AD) domain that the Samba server belongs to. It is a critical setting for environments where Samba operates as a domain member or domain controller, ensuring proper integration with Windows AD and Kerberos authentication.

## Purpose

- **Kerberos Integration**:  
  Defines the Kerberos realm that Samba uses for authentication. This is essential for ticket-based authentication and secure communication with domain controllers.

- **Active Directory Membership**:  
  Specifies the AD domain (typically in uppercase) for Samba to join, ensuring that domain membership and related services work correctly.

- **Single Sign-On (SSO)**:  
  Properly setting the `realm` enables seamless SSO, so users can access resources across the domain with a single set of credentials.

## Configuration

The `realm` parameter is set in the `[global]` section of your Samba configuration file (`smb.conf`).

### Example Configuration

```ini
[global]
   workgroup = EXAMPLE
   realm = EXAMPLE.COM
   security = ADS
   encrypt passwords = yes
```

- **`workgroup`**: Should match your Windows domain workgroup.
- **`realm`**: Specifies the Kerberos realm, usually the fully qualified domain name in uppercase.
- **`security = ADS`**: Instructs Samba to operate in Active Directory mode.
- **`encrypt passwords = yes`**: Ensures that passwords are encrypted as required by AD.

## Best Practices

- **Use Uppercase for the Realm**:  
  The realm is case-sensitive and is typically specified in uppercase (e.g., `EXAMPLE.COM`).

- **Ensure Time Synchronization**:  
  Accurate time settings are essential for Kerberos authentication. Make sure that the system clock is synchronized with the domain controllers.

- **Configure Kerberos Properly**:  
  Your `/etc/krb5.conf` file should be configured to include the correct realm and KDC (Key Distribution Center) entries that match the Samba `realm` setting.

## Troubleshooting

- **Authentication Issues**:  
  If users cannot authenticate or if Samba cannot join the domain, verify that the `realm` in `smb.conf` exactly matches the Kerberos realm and AD domain name.

- **Kerberos Tickets**:  
  Use `kinit` to obtain a ticket and `klist` to list the current tickets. Mismatches in the `realm` configuration can cause ticket acquisition to fail.

## Conclusion

The `realm` parameter is vital for integrating Samba with Windows Active Directory and Kerberos. By ensuring that the Samba server uses the correct Kerberos realm, administrators can achieve seamless authentication, domain membership, and single sign-on functionality in a mixed-OS environment.
