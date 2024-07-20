# net

The `net` command is a versatile utility in Samba that provides a range of network and administrative operations. It allows you to manage various aspects of Samba, including domain operations, user and group management, and more. Below are the main uses and commands of `net`.

#### General Usage

```bash
net [SUBCOMMAND] [OPTIONS]
```

### Common Subcommands and Their Usage

#### Domain Management

1. **Joining a Domain**

   ```bash
   net ads join -U administrator
   ```

   This command joins a Samba server to an Active Directory domain.

2. **Leaving a Domain**

   ```bash
   net ads leave -U administrator
   ```

   This command removes a Samba server from an Active Directory domain.

3. **Checking Domain Trust**

   ```bash
   net rpc trustdom list -U administrator
   ```

   This command lists the trusted domains for the Samba server.

#### User and Group Management

1. **Adding a User**

   ```bash
   net rpc user add username -U administrator
   ```

   This command adds a new user to the Samba server.

2. **Deleting a User**

   ```bash
   net rpc user delete username -U administrator
   ```

   This command deletes a user from the Samba server.

3. **Listing Users**

   ```bash
   net rpc user -U administrator
   ```

   This command lists all users on the Samba server.

4. **Adding a Group**

   ```bash
   net rpc group add groupname -U administrator
   ```

   This command adds a new group to the Samba server.

5. **Deleting a Group**

   ```bash
   net rpc group delete groupname -U administrator
   ```

   This command deletes a group from the Samba server.

6. **Listing Groups**

   ```bash
   net rpc group -U administrator
   ```

   This command lists all groups on the Samba server.

7. **Adding a User to a Group**

   ```bash
   net rpc group addmem groupname username -U administrator
   ```

   This command adds a user to a specified group.

8. **Removing a User from a Group**

   ```bash
   net rpc group delmem groupname username -U administrator
   ```

   This command removes a user from a specified group.

#### Share Management

1. **Creating a Share**

   ```bash
   net usershare add sharename path "Comment" everyone:F guest_ok=n
   ```

   This command creates a new share on the Samba server.

2. **Deleting a Share**

   ```bash
   net usershare delete sharename
   ```

   This command deletes a share from the Samba server.

3. **Listing Shares**

   ```bash
   net usershare list
   ```

   This command lists all shares on the Samba server.

#### Print Management

1. **Listing Printers**

   ```bash
   net rpc printer list -U administrator
   ```

   This command lists all printers on the Samba server.

2. **Adding a Printer**

   ```bash
   net rpc printer add printername -U administrator
   ```

   This command adds a new printer to the Samba server.

3. **Deleting a Printer**

   ```bash
   net rpc printer delete printername -U administrator
   ```

   This command deletes a printer from the Samba server.

#### Password Management

1. **Setting a User's Password**

   ```bash
   net rpc password username -U administrator
   ```

   This command sets the password for a specified user.

2. **Changing Own Password**

   ```bash
   net rpc password -U username
   ```

   This command allows a user to change their own password.

#### Examples

1. **Join a Samba Server to a Domain**

   ```bash
   net ads join -U administrator
   ```

   Example output:

   ```plaintext
   Enter administrator's password:
   Using short domain name -- EXAMPLE
   Joined 'SAMBA-SERVER' to dns domain 'example.com'
   ```

2. **List All Users**

   ```bash
   net rpc user -U administrator
   ```

   Example output:

   ```plaintext
   Enter administrator's password:
   User accounts for \\SAMBA-SERVER
   ---------------------------------
   administrator
   guest
   john.doe
   ```

3. **Create a New Share**

   ```bash
   net usershare add public /srv/samba/public "Public Share" everyone:F guest_ok=y
   ```

   Example output:

   ```plaintext
   Added share public.
   ```

4. **List All Shares**

   ```bash
   net usershare list
   ```

   Example output:

   ```plaintext
   public
   ```

5. **Set User Password**

   ```bash
   net rpc password john.doe -U administrator
   ```

   Example interaction:

   ```plaintext
   Enter administrator's password:
   Enter new password for user john.doe:
   ```

### Summary

| Command                | Purpose                                      | Common Usage                                              | Example Command                                      |
|------------------------|----------------------------------------------|-----------------------------------------------------------|------------------------------------------------------|
| `net ads join`         | Joins a Samba server to an AD domain         | `net ads join -U administrator`                           | `net ads join -U administrator`                      |
| `net rpc user add`     | Adds a new user                              | `net rpc user add username -U administrator`              | `net rpc user add john.doe -U administrator`         |
| `net rpc group add`    | Adds a new group                             | `net rpc group add groupname -U administrator`            | `net rpc group add developers -U administrator`      |
| `net usershare add`    | Creates a new share                          | `net usershare add sharename path "Comment" everyone:F`   | `net usershare add public /srv/samba/public "Public Share" everyone:F guest_ok=y` |
| `net rpc password`     | Sets a user's password                       | `net rpc password username -U administrator`              | `net rpc password john.doe -U administrator`         |

### Conclusion

The `net` command in Samba provides extensive functionality for managing domains, users, groups, shares, printers, and passwords. It is a crucial tool for Samba administrators to perform a wide variety of administrative tasks efficiently. Understanding the different subcommands and their options allows administrators to effectively manage and maintain a Samba server in a network environment.
