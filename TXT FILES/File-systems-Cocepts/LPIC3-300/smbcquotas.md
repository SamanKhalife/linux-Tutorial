# smbcquotas

`smbcquotas` is a Samba Virtual File System (VFS) module that provides disk quota support for SMB/CIFS shares. It allows administrators to enforce and monitor disk usage limits on shared resources, ensuring that no user or group exceeds their allocated storage quota. This module integrates with the underlying filesystem quota system, enabling transparent quota management for Samba clients.

## Overview

- **Purpose**:  
  To enforce disk usage limits on Samba shares by monitoring file operations and applying quota restrictions based on the underlying filesystem's quota settings.

- **Integration**:  
  Works in conjunction with the native filesystem quota system, so disk quotas must be enabled and configured on the server.

- **Monitoring and Reporting**:  
  Provides logging and reporting features that help administrators track quota usage and receive alerts when limits are approached or exceeded.

## Key Features

- **Quota Enforcement**:  
  Prevents users from using more disk space than allocated, ensuring fair resource distribution.
  
- **Filesystem Integration**:  
  Leverages the existing filesystem quota infrastructure (set up with tools like `edquota` and `quotaon`), making it a seamless extension of the local quota system.

- **Transparent Operation**:  
  Functions as part of Samba's VFS stack, requiring minimal user intervention once configured.

- **Logging**:  
  Logs quota-related events to a specified log file for auditing and troubleshooting.

## Configuration

To enable `smbcquotas`, add it to the list of VFS objects in your Samba configuration file (`smb.conf`). You may also need to specify additional options to tailor its behavior.

### Example Configuration

```ini
[global]
   # Load the smbcquotas VFS module for quota management
   vfs objects = smbcquotas

   # Configuration options for smbcquotas (options may vary by Samba version)
   smbcquotas:quota backend = native
   smbcquotas:quota path = /quota/samba
   smbcquotas:log file = /var/log/samba/quotas.log
```

- **`vfs objects = smbcquotas`**:  
  Enables the quota module for the Samba server or specific shares.

- **Quota Options**:  
  - `smbcquotas:quota backend = native`: Uses the native filesystem quota system.
  - `smbcquotas:quota path = /quota/samba`: Specifies the path where quota information is stored or managed.
  - `smbcquotas:log file = /var/log/samba/quotas.log`: Designates a log file for quota-related events.

## Use Cases

- **Enterprise File Sharing**:  
  Enforce storage policies to prevent any single user or department from consuming excessive disk space, which is critical for maintaining system performance and ensuring effective backup strategies.

- **Home and Small Office Networks**:  
  Manage limited storage resources by setting quotas, ensuring fair usage among multiple users sharing the same server.

- **Quota Monitoring and Reporting**:  
  Administrators can monitor quota usage through logs and reports, making it easier to plan for future storage needs and identify potential misuse.

## Considerations

- **Filesystem Quotas Must Be Enabled**:  
  Before using `smbcquotas`, ensure that your filesystem supports and has quotas enabled. Configure these using standard Linux quota tools.

- **Testing and Validation**:  
  Test the configuration in a controlled environment to verify that quotas are enforced correctly and that normal file operations are not negatively impacted.

- **Logging and Maintenance**:  
  Regularly review the log file specified for quota events to catch any potential issues early and adjust quota policies as needed.

## Conclusion

The `smbcquotas` module is a valuable addition to Samba for enforcing disk quotas on SMB/CIFS shares. By integrating with the native filesystem quota system, it provides transparent, efficient, and manageable quota control that helps prevent overuse of storage resources, ensuring fair and secure file sharing in both small and large environments.
