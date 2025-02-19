# winbind-offline-logon

**`winbind offline logon`** is a Samba configuration parameter that enables domain users to log on using cached credentials when a connection to the Active Directory Domain Controller is unavailable. This functionality provides a fallback authentication mechanism, similar to the cached credentials feature in Windows environments, ensuring that users can still access their accounts even when network connectivity is intermittent.

## Purpose

- **Offline Authentication**:  
  Allows domain users to log on to a Samba-integrated system using previously cached credentials if the AD domain controller is unreachable.

- **User Convenience**:  
  Particularly beneficial for laptops and mobile devices, ensuring that users can access their systems while on the go or during network outages.

- **Continuous Availability**:  
  Enhances system resilience by permitting logons even in environments with unreliable or temporary network connectivity.

## How It Works

- **Credential Caching**:  
  When a user successfully logs on while the domain controller is available, their credentials are cached locally.
  
- **Fallback Mechanism**:  
  On subsequent logon attempts, if the system cannot contact a domain controller, Samba uses the cached credentials to authenticate the user.

## Configuration

To enable offline logon functionality, add the following line to the `[global]` section of your Samba configuration file (`smb.conf`):

```ini
[global]
   winbind offline logon = yes
```

This setting instructs Samba to cache domain user credentials, allowing for offline authentication.

## Use Cases

- **Mobile and Remote Work**:  
  Users who frequently disconnect from the corporate network can still access their systems using cached credentials.

- **Intermittent Connectivity**:  
  In environments with unreliable network connections, offline logon ensures continuity of user access.

- **Maintenance and Outages**:  
  Provides a safeguard during planned maintenance or unexpected network outages, allowing users to log on until connectivity is restored.

## Considerations

- **Security**:  
  Caching credentials locally introduces potential security risks. Ensure that proper security measures (e.g., disk encryption, strong local authentication policies) are in place to protect cached data.

- **Cache Expiry**:  
  Be aware of how long cached credentials are retained and consider policies for cache expiration or forced re-authentication when the network is restored.

- **Compatibility**:  
  Confirm that your Samba version supports offline logon caching and that your domain configuration is set up to permit this functionality.

## Troubleshooting

- **Offline Logon Failures**:  
  If users are unable to log on offline, check the Samba log files (e.g., `/var/log/samba/log.winbindd`) for any error messages related to credential caching.

- **Verification Tools**:  
  Use tools like `wbinfo -u` and `getent passwd` to verify that domain user data is being correctly cached and resolved.

- **Network Issues**:  
  Ensure that during normal operations, the system can contact the domain controller so that credentials are properly cached.

## Conclusion

Enabling **`winbind offline logon`** in Samba is an essential feature for environments where network connectivity to the domain controller may be unreliable. By caching domain credentials, this feature allows users to continue accessing their systems even when the AD domain controller is temporarily unavailable, ensuring a smooth and uninterrupted user experience.

