# vfs objects
**`vfs objects`** is a Samba configuration directive used in the `smb.conf` file that allows the administrator to specify Virtual File System (VFS) modules that should be loaded for a specific share or globally for all shares. VFS modules are extensions that Samba can use to modify the behavior of file operations, such as auditing, virus scanning, or snapshot integration.

### Key Points about `vfs objects`:

1. **Purpose**:
   - VFS objects allow you to enhance or alter the default functionality of Samba by adding extra features or customization layers, such as access logging, snapshot support, or file content analysis.

2. **Usage**:
   - The directive can be applied either globally or per-share to control how Samba interacts with the file system.

3. **Syntax in `smb.conf`**:
   - You define the `vfs objects` in your `smb.conf` file under a share or in the global section.
   - Multiple VFS modules can be chained by specifying them in a space-separated list.
   
   **Example**:
   ```ini
   [myshare]
   path = /srv/samba/myshare
   vfs objects = recycle full_audit
   ```

4. **Common VFS Modules**:
   - **recycle**: Implements a "Recycle Bin" functionality where deleted files are moved to a designated directory rather than being permanently deleted.
   - **full_audit**: Provides detailed logging of file access and operations, such as file opens, reads, writes, etc.
   - **shadow_copy2**: Enables integration with snapshot technologies like Volume Shadow Copy (on Windows) or Linux snapshots, allowing users to view and restore previous versions of files.
   - **acl_xattr**: Stores Windows NT ACLs (Access Control Lists) in extended attributes on the filesystem.
   - **virusfilter**: Scans files for viruses when they are opened, written to, or executed, useful for integrating with anti-virus software.

### Common Examples of VFS Modules:

1. **Recycle Bin Configuration**:
   ```ini
   [shared]
   path = /srv/samba/shared
   vfs objects = recycle
   recycle:repository = .recycle/%U
   recycle:keeptree = yes
   recycle:versions = yes
   recycle:touch = yes
   ```

   - This creates a hidden directory `.recycle` in the user's folder where deleted files are stored.
   - The `%U` is replaced by the username, creating a separate recycle bin for each user.

2. **Full Audit Logging**:
   ```ini
   [audit_share]
   path = /srv/samba/audit
   vfs objects = full_audit
   full_audit:prefix = %u|%I|%S
   full_audit:success = mkdir rmdir open read pread write pwrite rename unlink
   full_audit:failure = none
   full_audit:facility = LOCAL7
   full_audit:priority = NOTICE
   ```

   - This logs specific operations such as file reads, writes, and directory creation in the audit log.
   - The `full_audit:prefix` can include placeholders for the user (`%u`), IP address (`%I`), and share name (`%S`).

3. **Shadow Copy Support**:
   ```ini
   [share_with_snapshots]
   path = /srv/samba/data
   vfs objects = shadow_copy2
   shadow:snapdir = /srv/samba/snapshots
   shadow:sort = desc
   ```

   - This integrates snapshots into the share so users can view previous versions of files.

### Benefits of Using `vfs objects`:

- **Flexibility**: You can load specific functionality only when needed, like recycling bins for sensitive shares or virus scanning for high-security environments.
- **Security**: Modules like `virusfilter` and `acl_xattr` enhance security by scanning files or preserving file permission integrity across different platforms.
- **Auditing**: Modules like `full_audit` allow detailed logging of file operations, useful for compliance and troubleshooting.
- **Backup and Recovery**: Modules like `shadow_copy2` provide integration with snapshot technologies, helping users recover previous versions of files without administrator intervention.

### Conclusion:
The `vfs objects` directive in Samba allows administrators to add powerful extensions to Samba shares, enabling features like access auditing, virus scanning, file versioning, and more. Each module enhances the Samba server's functionality, making it highly adaptable to various business and security needs.
