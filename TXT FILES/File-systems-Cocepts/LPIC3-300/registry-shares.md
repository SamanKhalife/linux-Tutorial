# Registry shares
**Registry shares** in Samba allow you to manage share configurations in the Windows registry style, rather than using the traditional `smb.conf` configuration file. This feature is especially useful in environments where integration with Windows systems is important, as it mimics the way Windows handles shared resources through its registry.

### Key Concepts of Registry Shares

1. **Storage in Samba's `registry.tdb`**:
   - Samba stores the configuration of registry shares in a database file called `registry.tdb`.
   - The shares are managed via the Samba registry, similar to how Windows stores its shared folders in the registry.
   - Using registry shares provides flexibility for managing Samba shares via tools like `net registry` or `samba-regedit`.

2. **Creation and Management**:
   - You can create, edit, and delete registry shares via the command-line tool `net registry` or using a graphical registry editor (`samba-regedit`).
   - These shares can be manipulated in a manner similar to how one would edit the Windows registry with `regedit.exe`.
   
3. **Access via `net registry` Command**:
   - The `net registry` command provides the ability to list, add, and manage registry shares.
   
   **Example: Creating a Share Using `net registry`**:
   ```bash
   net registry addshare ShareName /path/to/share
   ```

4. **Location of Registry Shares**:
   - Samba registry shares are stored in the database, not in the `smb.conf` file.
   - If you want to view or manipulate registry-based shares, you need to use registry management tools (`net registry`, `samba-regedit`).

5. **Advantages of Registry Shares**:
   - **Centralized Management**: Especially useful in Active Directory environments or when trying to match Windows server behavior.
   - **Dynamic**: Shares can be updated without restarting the Samba service, as registry changes are applied dynamically.
   - **Ease of Use**: Management of shares can be handled in a Windows-like fashion, which can simplify integration in mixed Windows/Linux environments.

6. **Migrating from smb.conf to Registry Shares**:
   - It is possible to migrate shares defined in `smb.conf` to the registry using the `net conf import` command.
   
   **Example**:
   ```bash
   net conf import smb.conf
   ```

7. **Viewing Registry Shares**:
   - To view the existing registry shares, you can use `net registry` or `samba-regedit`.

   **Example**:
   ```bash
   net registry listshares
   ```

8. **Editing Registry Shares**:
   - Use `samba-regedit` to graphically manage Samba registry shares.
   - This tool allows you to add, delete, or modify registry shares through a graphical interface, similar to the Windows `regedit` tool.

### Basic Workflow for Managing Registry Shares

1. **Add a New Share**:
   - To add a share directly to the registry:
     ```bash
     net registry addshare "MyShare" /srv/samba/myshare
     ```

2. **Delete a Share**:
   - To delete a share:
     ```bash
     net registry delshare "MyShare"
     ```

3. **List All Shares**:
   - To list all shares:
     ```bash
     net registry listshares
     ```

4. **Edit Shares**:
   - You can modify a share using `net registry` or via `samba-regedit`.

5. **View Share Properties**:
   - To view properties of a specific share:
     ```bash
     net registry getshareinfo "MyShare"
     ```

### Example Setup for a Registry-Based Share

Let’s say we want to create a new share called `public` with the path `/srv/samba/public`.

1. **Create the Share**:
   ```bash
   net registry addshare "public" /srv/samba/public
   ```

2. **Set Permissions**:
   You can set the share to be read-only or writable:
   ```bash
   net registry setshareinfo "public" writeable=true
   ```

3. **List Shares**:
   To verify the share has been created:
   ```bash
   net registry listshares
   ```

4. **Edit Share**:
   Use `samba-regedit` to further modify or configure advanced options of the share, such as permissions or access controls.

### Conclusion

Registry shares in Samba offer a more Windows-like approach to managing file shares, particularly useful in hybrid or AD environments. Using tools like `net registry` and `samba-regedit`, you can easily add, edit, and manage shares without directly modifying `smb.conf`. This makes Samba more dynamic and compatible with enterprise Windows infrastructures, enabling easier integration and centralized management.
