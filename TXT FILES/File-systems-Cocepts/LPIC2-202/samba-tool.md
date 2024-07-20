# samba-tool

`samba-tool` is a comprehensive command-line utility that allows administrators to manage and configure various aspects of Samba. It is particularly useful for tasks related to Active Directory (AD) management, domain controllers, and other advanced configurations.

#### General Usage

```bash
samba-tool [command] [subcommand] [options]
```

`samba-tool` encompasses a wide range of commands and subcommands for different purposes. Below are some of the most commonly used categories and examples.

#### Active Directory (AD) Management

1. **Creating a Domain**

   ```bash
   samba-tool domain provision --use-rfc2307 --interactive
   ```

   This command guides you through the process of creating a new Samba AD domain controller interactively.

2. **Joining a Domain**

   ```bash
   samba-tool domain join example.com DC -U"EXAMPLE\administrator"
   ```

   This command joins an existing domain as an additional domain controller.

3. **Managing Users**

   - **Creating a User**

     ```bash
     samba-tool user create username password --given-name=John --surname=Doe
     ```

     This command creates a new AD user with the specified attributes.

   - **Deleting a User**

     ```bash
     samba-tool user delete username
     ```

     This command deletes an AD user.

   - **Changing a User's Password**

     ```bash
     samba-tool user setpassword username
     ```

     This command prompts to set a new password for the specified user.

4. **Managing Groups**

   - **Creating a Group**

     ```bash
     samba-tool group add groupname
     ```

     This command creates a new AD group.

   - **Adding a User to a Group**

     ```bash
     samba-tool group addmembers groupname username
     ```

     This command adds a user to a specified group.

   - **Removing a User from a Group**

     ```bash
     samba-tool group removemembers groupname username
     ```

     This command removes a user from a specified group.

#### Domain and DNS Management

1. **Listing Domain Controllers**

   ```bash
   samba-tool domain listdcs
   ```

   This command lists all domain controllers in the domain.

2. **Managing DNS Records**

   - **Adding a DNS Record**

     ```bash
     samba-tool dns add [dns-server-ip] [zone] [name] [A|AAAA|CNAME|etc.] [ip-address]
     ```

     Example:

     ```bash
     samba-tool dns add 192.168.1.1 example.com www A 192.168.1.100
     ```

   - **Deleting a DNS Record**

     ```bash
     samba-tool dns delete [dns-server-ip] [zone] [name] [A|AAAA|CNAME|etc.] [ip-address]
     ```

     Example:

     ```bash
     samba-tool dns delete 192.168.1.1 example.com www A 192.168.1.100
     ```

3. **Forcing a Replication**

   ```bash
   samba-tool drs replicate <srcDC> <dstDC> [options]
   ```

   Example:

   ```bash
   samba-tool drs replicate DC1 DC2 --full-sync
   ```

#### Schema Management

1. **Extending the Schema**

   ```bash
   samba-tool schema upgrade
   ```

   This command upgrades the AD schema to match the schema of a given Windows AD.

2. **Checking Schema**

   ```bash
   samba-tool dbcheck --cross-ncs --fix
   ```

   This command performs various consistency checks on the AD database.

#### Troubleshooting and Maintenance

1. **Checking the Database**

   ```bash
   samba-tool dbcheck
   ```

   This command checks the AD database for consistency issues.

2. **Resetting the SYSVOL ACLs**

   ```bash
   samba-tool ntacl sysvolreset
   ```

   This command resets the Access Control Lists (ACLs) on the SYSVOL share.

3. **Inspecting Logs**

   ```bash
   samba-tool logs <command> [options]
   ```

   Example:

   ```bash
   samba-tool logs level 3
   ```

   This command sets the log level for Samba components to aid in troubleshooting.

### Example Usage Scenarios

1. **Provisioning a New Domain**

   ```bash
   samba-tool domain provision --use-rfc2307 --interactive
   ```

2. **Creating a New User**

   ```bash
   samba-tool user create jane.doe StrongPassword123 --given-name=Jane --surname=Doe
   ```

3. **Adding a DNS Record**

   ```bash
   samba-tool dns add 192.168.1.1 example.com mail A 192.168.1.200
   ```

4. **Forcing Replication Between Domain Controllers**

   ```bash
   samba-tool drs replicate DC1 DC2 --full-sync
   ```

5. **Resetting SYSVOL ACLs**

   ```bash
   samba-tool ntacl sysvolreset
   ```

### Conclusion

`samba-tool` is a versatile and powerful command-line utility that simplifies the management and administration of Samba, particularly in an Active Directory environment. Whether provisioning domains, managing users and groups, configuring DNS, or troubleshooting, `samba-tool` provides a comprehensive set of commands to effectively manage Samba servers. Understanding and utilizing these commands can greatly enhance an administrator’s ability to maintain and optimize a Samba-based network infrastructure.
