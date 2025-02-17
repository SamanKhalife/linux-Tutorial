# pdbedit

The **`pdbedit`** command is a utility in **Samba** used to manage and administer the Samba **Passdb** (Password Database). This command is primarily used for managing Samba user accounts and their attributes, such as adding, modifying, or deleting users, changing passwords, and viewing user details. It is particularly useful when Samba is being used as a **file server** or **domain controller**.

`pdbedit` operates directly on the Samba user database, which can either be stored in **TDB (Trivial Database)** files (such as `passdb.tdb`) or external backends like **LDAP** when Samba is configured for a domain controller.

### Key Features of `pdbedit`:
- **Add, Modify, and Delete Users**: It allows administrators to manage user accounts within the Samba database.
- **View User Information**: You can view various attributes of a Samba user account, such as their SID (Security Identifier), last password change time, and more.
- **Password Management**: You can use `pdbedit` to change or expire passwords, or set password policies.
- **Support for Multiple Backends**: It works with Samba's default TDB backend as well as external directories like LDAP when Samba is configured to use them.

### General Syntax:
```bash
pdbedit [OPTIONS] [USER]
```
Where:
- **`[OPTIONS]`**: Various options to specify the action you want to perform.
- **`[USER]`**: The username or user ID you want to modify or query.

### Common `pdbedit` Commands:

1. **Listing Samba Users**:
   To list all users in the Samba password database, simply use the following command:
   ```bash
   pdbedit -L
   ```
   - **`-L`**: Lists all Samba users in the database.
   
2. **Viewing User Information**:
   To view detailed information about a specific user:
   ```bash
   pdbedit -v username
   ```
   - **`-v`**: Provides detailed information about the specified user, including their SID, password expiration time, last password change, and more.

3. **Adding a New User**:
   To add a new user to Samba’s password database:
   ```bash
   pdbedit -a username
   ```
   - **`-a`**: Adds a new user.
   - After running this command, you'll typically be prompted to set a password for the new user.

4. **Changing a User’s Password**:
   To change the password of a user:
   ```bash
   pdbedit -r username
   ```
   - **`-r`**: Resets the password for the specified user.
   - You'll be prompted to enter the new password.

5. **Deleting a User**:
   To delete a user from the Samba password database:
   ```bash
   pdbedit -x username
   ```
   - **`-x`**: Deletes the specified user.

6. **Viewing Specific User’s SID**:
   To view the **SID (Security Identifier)** for a specific user:
   ```bash
   pdbedit -s username
   ```
   - **`-s`**: Displays the SID for the specified user.

7. **Exporting User Data**:
   To export the entire user database to an LDIF file or another format:
   ```bash
   pdbedit -e ldif:/path/to/exportfile.ldif
   ```
   - **`-e`**: Specifies the export format and target file.

8. **Setting Password Expiry for a User**:
   To set the password expiry for a user:
   ```bash
   pdbedit -P username
   ```
   - **`-P`**: Marks the user’s password to expire.

### Options:

- **`-h`** or **`--help`**: Displays help information for the `pdbedit` command.
   ```bash
   pdbedit --help
   ```

- **`-L`**: Lists all users in the password database.
   ```bash
   pdbedit -L
   ```

- **`-v`**: Provides verbose output, showing detailed user information.
   ```bash
   pdbedit -v username
   ```

- **`-a`**: Adds a new user to the database.
   ```bash
   pdbedit -a newuser
   ```

- **`-x`**: Deletes a user from the database.
   ```bash
   pdbedit -x username
   ```

- **`-s`**: Displays the SID (Security Identifier) of a user.
   ```bash
   pdbedit -s username
   ```

- **`-r`**: Resets a user’s password.
   ```bash
   pdbedit -r username
   ```

- **`-P`**: Sets a user’s password to expire.
   ```bash
   pdbedit -P username
   ```

### Example Usage:

1. **List all Samba users**:
   ```bash
   pdbedit -L
   ```
   This will list all the users in the Samba password database.

2. **View detailed information for a specific user (`john.doe`)**:
   ```bash
   pdbedit -v john.doe
   ```

3. **Add a new user (`jane.doe`)**:
   ```bash
   pdbedit -a jane.doe
   ```
   This will add `jane.doe` to the Samba database, prompting you to set a password.

4. **Change the password for `john.doe`**:
   ```bash
   pdbedit -r john.doe
   ```
   This will allow you to reset the password for `john.doe`.

5. **Delete the user `john.doe` from the database**:
   ```bash
   pdbedit -x john.doe
   ```
   This will permanently delete `john.doe` from the Samba password database.

6. **Export the user database to an LDIF file**:
   ```bash
   pdbedit -e ldif:/path/to/output.ldif
   ```

### Practical Use Cases:

1. **Managing User Accounts**: 
   When Samba is used in a network environment, managing user accounts through `pdbedit` is crucial for maintaining access control. You can create new users, modify their information, or delete accounts.

2. **Domain Controller Management**: 
   In Samba's **Active Directory Domain Controller** mode, `pdbedit` can be used to administer domain user accounts, set password policies, and manage group memberships.

3. **Password and Security Management**: 
   You can use `pdbedit` to reset user passwords, enforce password expiration, and view the account details for troubleshooting or auditing purposes.

4. **Bulk User Creation**: 
   For bulk user account creation or updates, `pdbedit` can be scripted to automate the process, especially useful in larger environments or when migrating from other systems.

### Conclusion:

The **`pdbedit`** command is a powerful utility for managing Samba's password database. It allows administrators to add, modify, and delete users, as well as perform various administrative tasks related to Samba user accounts. Whether managing a Samba file server or a domain controller, `pdbedit` is an essential tool for user and password management, making it an important part of any Samba administrator's toolkit.
