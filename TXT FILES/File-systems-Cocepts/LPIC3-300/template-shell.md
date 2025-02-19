# template-shell

`template-shell` is a Samba configuration parameter used to define the default login shell for users accessing their home directories via Samba. When a user logs in, Samba uses this template value to set or override the user's shell, ensuring a consistent command-line environment across the domain.

## Purpose

- **Default Shell Assignment**:  
  Automatically assigns a default shell to users when no specific shell is set in their account. For example, setting `template shell = /bin/bash` ensures that every user gets `/bin/bash` as their login shell.

- **Consistency**:  
  Helps maintain a uniform user environment across all systems integrated with Samba, especially in a domain controller or networked file server scenario.

- **Simplified Administration**:  
  By providing a default shell via a template, administrators can avoid having to manually configure the shell for each user.

## Typical Usage

The `template-shell` parameter is usually configured within the `[homes]` share section in your Samba configuration file (`smb.conf`). This ensures that when users access their home directories, they are assigned the specified shell.

### Example Configuration

```ini
[homes]
   comment = Home Directories
   browseable = no
   writable = yes
   valid users = %S
   template homedir = /home/%U
   template shell = /bin/bash
```

- **`template shell = /bin/bash`**:  
  Sets the default login shell for users to `/bin/bash`.

## How It Works

- **Dynamic Assignment**:  
  When a user logs in to the Samba share for their home directory, the `template-shell` value is used to set the login shell for that session if a shell isn’t explicitly defined elsewhere.

- **Integration with PAM**:  
  In many environments, PAM (Pluggable Authentication Modules) works in conjunction with Samba to create home directories and set user environments. The `template shell` parameter ensures that the user's shell is set consistently during this process.

## Benefits

- **Uniform User Environment**:  
  Provides all users with a consistent shell environment, which simplifies training and support.
  
- **Ease of Management**:  
  Administrators can centrally define and enforce a standard shell for all users without needing to modify individual user profiles.

- **Default Fallback**:  
  Acts as a fallback option when user-specific shell settings are not provided in other account management systems.

## Considerations

- **Customization**:  
  The default shell specified by `template shell` can be customized to meet your organization’s requirements. For instance, some environments might prefer `/bin/zsh` or `/usr/bin/fish` as the default shell.

- **Overriding Settings**:  
  If a user’s account already specifies a shell (for example, through an LDAP attribute or local configuration), that value will typically take precedence over the template setting.

- **Integration with User Management**:  
  Ensure that your user account management system (e.g., LDAP or local user files) is configured in harmony with the `template shell` setting to avoid conflicts.

## Conclusion

The `template-shell` parameter in Samba is a valuable tool for ensuring that users receive a consistent default login shell when accessing their home directories via Samba. By defining a default shell (such as `/bin/bash`), administrators can streamline user environment setup, reduce configuration overhead, and maintain uniformity across a networked environment. This is particularly beneficial in Samba AD domain controller setups or shared file server environments.
