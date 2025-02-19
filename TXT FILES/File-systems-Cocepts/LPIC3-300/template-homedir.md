# template-homedir

`template-homedir` is a Samba configuration parameter used to define a default template for user home directories. It is typically set within the `[homes]` share in your `smb.conf` file. By using placeholders (such as `%U` for the username), it automatically constructs the home directory path for each user, ensuring a consistent and standardized directory structure across the domain.

## Purpose

- **Default Home Directory Path**:  
  Specifies the default location for user home directories. For example, setting `template homedir = /home/%U` means that the home directory for a user named `alice` will be `/home/alice`.

- **Consistency**:  
  Ensures that all user home directories follow the same structure, which simplifies administration and support.

- **Automation**:  
  Allows Samba to automatically create or reference the correct home directory for a user during login, without requiring manual intervention.

## Typical Usage

The `template-homedir` parameter is usually configured within the `[homes]` share in your Samba configuration file (`smb.conf`).

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

- **`template homedir = /home/%U`**:  
  Sets the home directory to `/home/username` (where `%U` is replaced by the actual username).

- **`template shell = /bin/bash`**:  
  Optionally specifies the default login shell for users.

## How It Works

- **Dynamic Substitution**:  
  When a user logs in, Samba replaces `%U` in the `template-homedir` value with the user's actual username, generating the proper home directory path.
  
- **Automatic Home Directory Creation**:  
  If the home directory does not exist, Samba (along with the appropriate PAM configuration) can automatically create it based on this template.

## Benefits

- **Standardization**:  
  Provides a consistent layout for home directories across all users.
  
- **Ease of Management**:  
  Simplifies user environment setup by automating the creation of home directories with predefined paths.
  
- **Integration**:  
  Works seamlessly with other Samba settings (such as `template shell`) to fully define the user’s login environment.

## Considerations

- **Directory Permissions**:  
  Ensure that the parent directory (e.g., `/home`) has the correct permissions to allow Samba (and PAM, if used) to create home directories.
  
- **PAM Integration**:  
  Verify that your PAM configuration supports automatic home directory creation if you rely on this feature.
  
- **Customization**:  
  You can adjust the template to meet your organization's requirements (e.g., different base paths or additional placeholders).

## Conclusion

The `template-homedir` parameter is a powerful tool in Samba for standardizing and automating the creation of user home directories. By defining a template path that uses placeholders like `%U`, administrators can ensure that all users receive a consistent environment, simplifying management in a Samba domain.
