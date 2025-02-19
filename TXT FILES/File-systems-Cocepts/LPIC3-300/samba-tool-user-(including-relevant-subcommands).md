# samba-tool user

The `samba-tool user` command is a powerful utility provided by Samba for managing user accounts in an Active Directory domain. It offers a wide range of subcommands that allow administrators to add, delete, modify, and query user information, thereby simplifying user administration in mixed Windows/Unix environments.

## Overview

Using `samba-tool user`, you can:

- **Create new user accounts**
- **Delete existing accounts**
- **Modify user attributes**
- **Rename user accounts**
- **Manage passwords**
- **Unlock, disable, or enable accounts**
- **Display detailed user information**
- **List all domain users**

Below are the common subcommands along with usage examples.

## Common Subcommands

### 1. Add a New User
**Purpose**: Create a new user account in the domain.

**Usage**:
```bash
samba-tool user add <username> [password] [options]
```

**Example**:
```bash
samba-tool user add jdoe P@ssw0rd --given-name="John" --surname="Doe" --mail-address=jdoe@example.com
```
*This command creates a new user `jdoe` with the specified password and additional attributes.*

### 2. Delete a User
**Purpose**: Remove an existing user account from the domain.

**Usage**:
```bash
samba-tool user delete <username>
```

**Example**:
```bash
samba-tool user delete jdoe
```
*This removes the user `jdoe` from the Active Directory.*

### 3. Modify User Attributes
**Purpose**: Update properties of an existing user account.

**Usage**:
```bash
samba-tool user modify <username> [attribute options]
```

**Example** (setting an expiry date):
```bash
samba-tool user setexpiry jdoe 2023-12-31
```
*This sets the password expiry for `jdoe` to December 31, 2023.*

### 4. Rename a User
**Purpose**: Change the username of an existing account.

**Usage**:
```bash
samba-tool user rename <oldusername> <newusername>
```

**Example**:
```bash
samba-tool user rename jdoe john.doe
```
*This renames the user account from `jdoe` to `john.doe`.*

### 5. Set or Reset a User’s Password
**Purpose**: Assign or update a user's password.

**Usage**:
```bash
samba-tool user setpassword <username>
```

**Example**:
```bash
samba-tool user setpassword john.doe
```
*This command prompts for a new password for `john.doe`.*

### 6. Unlock a User Account
**Purpose**: Remove a lock from a user account that was locked due to multiple failed logins.

**Usage**:
```bash
samba-tool user unlock <username>
```

**Example**:
```bash
samba-tool user unlock john.doe
```
*This unlocks the user account `john.doe`.*


### 7. Enable or Disable a User Account
**Purpose**: Control the active status of a user account.

**Usage**:
- **Disable a user**:
  ```bash
  samba-tool user disable <username>
  ```
- **Enable a user**:
  ```bash
  samba-tool user enable <username>
  ```

**Example**:
```bash
samba-tool user disable john.doe
samba-tool user enable john.doe
```
*These commands disable and then enable the user `john.doe`.*

### 8. Show User Details
**Purpose**: Display detailed information about a specific user.

**Usage**:
```bash
samba-tool user show <username>
```

**Example**:
```bash
samba-tool user show john.doe
```
*This displays detailed attributes of the user `john.doe`, such as SID, account flags, and other properties.*

---

### 9. List All Users
**Purpose**: Retrieve a list of all user accounts in the domain.

**Usage**:
```bash
samba-tool user list
```

**Example**:
```bash
samba-tool user list
```
*This command lists all the user accounts available in the domain.*

## Additional Options and Considerations

- **Attribute Options**:  
  When adding or modifying a user, you can supply additional attributes (e.g., `--given-name`, `--surname`, `--mail-address`).

- **Scripting and Automation**:  
  The commands in `samba-tool user` are script-friendly, enabling batch operations for large-scale user management.

- **Privileges**:  
  Operations typically require domain administrator privileges. Ensure you run these commands with appropriate credentials.

- **Verification**:  
  After performing changes, use `samba-tool user show <username>` to verify that the modifications are applied as expected.

## Conclusion

The `samba-tool user` command is an essential utility for managing user accounts in a Samba Active Directory environment. It provides comprehensive functionality—from creating and deleting users to managing passwords and modifying account properties—enabling administrators to maintain a secure and well-organized domain. By integrating these subcommands into your administrative workflows, you can efficiently manage user identities in mixed-OS environments.
