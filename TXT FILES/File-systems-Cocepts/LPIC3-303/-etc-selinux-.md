# /etc/selinux/

The `/etc/selinux/` directory is a crucial component of SELinux (Security-Enhanced Linux), housing configuration files that determine the behavior and policies of SELinux on a system. This directory contains several important files and subdirectories that help in managing SELinux policies, settings, and modes. Below is an explanation of the key components within the `/etc/selinux/` directory and their purposes.

### Key Components of `/etc/selinux/`

1. **config**
   - **Description**: The primary configuration file for SELinux, specifying the SELinux policy type and mode.
   - **Key Parameters**:
     - `SELINUX`: Defines the mode of SELinux. It can be set to `enforcing`, `permissive`, or `disabled`.
     - `SELINUXTYPE`: Specifies the policy type, typically `targeted` or `mls`.
   - **Example Content**:
     ```plaintext
     # This file controls the state of SELinux on the system.
     # SELINUX= can take one of these three values:
     #     enforcing - SELinux security policy is enforced.
     #     permissive - SELinux prints warnings instead of enforcing.
     #     disabled - No SELinux policy is loaded.
     SELINUX=enforcing
     # SELINUXTYPE= can take one of these two values:
     #     targeted - Targeted processes are protected,
     #     mls - Multi Level Security protection.
     SELINUXTYPE=targeted
     ```

2. **policies/**
   - **Description**: Contains directories for different SELinux policy versions. Each directory includes compiled policy files and modules.
   - **Typical Subdirectories**:
     - `/etc/selinux/targeted/`
     - `/etc/selinux/mls/`

3. **targeted/contexts/**
   - **Description**: Holds context files that define the default security contexts for various system objects.
   - **Important Files**:
     - `files/`: Defines the contexts for files.
     - `user/`: Defines contexts for SELinux users.
     - `default_contexts`: Specifies default contexts for login.

4. **policy.###**
   - **Description**: A binary representation of the SELinux policy. The number (`###`) corresponds to the policy version.
   - **Location Example**: `/etc/selinux/targeted/policy/policy.32`

5. **modules/active/**
   - **Description**: Contains active policy modules. Each module directory includes the necessary files for a specific policy module.
   - **Location Example**: `/etc/selinux/targeted/modules/active/modules`

6. **seusers**
   - **Description**: Maps Linux users to SELinux users, defining the security contexts for users upon login.
   - **Example Content**:
     ```plaintext
     # Map linux user to SELinux user
     root:system_u:s0-s0:c0.c1023
     user_u:guest_u:s0
     ```

### Managing SELinux Configuration

**Checking SELinux Status**
To check the current status and configuration of SELinux, use:
```bash
sestatus
```

**Changing SELinux Mode**
To change the mode of SELinux (temporarily), use:
```bash
setenforce 0   # Switches to permissive mode
setenforce 1   # Switches to enforcing mode
```

For a permanent change, edit the `/etc/selinux/config` file and set the `SELINUX` parameter to `enforcing`, `permissive`, or `disabled`.

**Changing SELinux Policy Type**
To change the policy type, edit the `/etc/selinux/config` file and set the `SELINUXTYPE` parameter to `targeted` or `mls`.

**Relabeling the File System**
After changing SELinux configurations or policies, it may be necessary to relabel the file system to ensure all files have the correct SELinux contexts. To do this, create an empty file named `.autorelabel` at the root directory and reboot the system:
```bash
touch /.autorelabel
reboot
```

### Security Considerations

- **Backup Configuration Files**: Always back up configuration files before making significant changes.
- **Review Policy Changes**: Carefully review and understand policy changes to avoid inadvertently compromising system security.
- **Testing**: Test changes in a controlled environment before applying them to production systems.

### Conclusion

The `/etc/selinux/` directory and its contents play a crucial role in configuring and managing SELinux on a Linux system. Understanding the key files and their purposes helps administrators effectively control SELinux policies and ensure system security. By managing these configurations correctly, you can maintain a secure and compliant SELinux environment.
