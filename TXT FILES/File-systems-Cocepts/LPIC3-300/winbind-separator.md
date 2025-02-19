# winbind-separator

`winbind separator` is a Samba configuration parameter that defines the character used to separate the domain name from the username in Winbind’s output. By customizing this separator, administrators can control how domain user names appear on Unix systems, ensuring they fit the conventions or preferences of their environment.

## Overview

- **Purpose**:  
  The parameter specifies the delimiter placed between the Windows domain and the username when Winbind returns user and group information. For example, if the domain is `EXAMPLE` and the username is `jdoe`, the default output might be `EXAMPLE+jdoe`.

- **Default Behavior**:  
  By default, Samba sets the separator to the plus sign (`+`), so without any changes, domain users appear in the format `DOMAIN+username`.

## Configuration

To customize the Winbind separator, add or modify the directive in your Samba configuration file (`smb.conf`) within the `[global]` section. For example:

```ini
[global]
   winbind separator = +
```

### Changing the Separator

If you prefer a different character, such as a hyphen or even a backslash, you can specify it here:

```ini
[global]
   winbind separator = -
```

This configuration changes the display so that a user from the `EXAMPLE` domain named `jdoe` would appear as `EXAMPLE-jdoe`.

> **Note**: When using special characters like a backslash (`\`), ensure proper escaping if required by your configuration syntax.

## Impact on Identity Resolution

- **Consistency**:  
  The separator setting affects how usernames are formatted across all applications that use Winbind for domain user resolution. Consistent use of a chosen separator helps avoid ambiguity.

- **Scripting and Automation**:  
  Scripts or tools that parse Winbind output must be aware of the separator. Changing it from the default may require updates to these scripts to ensure correct processing of domain-qualified usernames.

## Example

Assuming you want to use a hyphen as the separator, update your `smb.conf` as follows:

```ini
[global]
   winbind separator = -
```

After restarting Samba, a domain user from the `EXAMPLE` domain with the username `jdoe` would be shown as:

```
EXAMPLE-jdoe
```

## Conclusion

The `winbind separator` parameter in Samba allows you to control the formatting of domain user names by specifying the character that separates the domain from the username. Adjusting this parameter can help maintain consistency across your environment and ensure that user identities are easily parsed and recognized by client applications and automated systems.
