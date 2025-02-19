# delete veto files

`delete veto files` is a parameter used in conjunction with Samba’s veto_files VFS module. It determines whether files matching the veto file patterns are actually removed from the filesystem when a client attempts to delete them. This setting helps manage unwanted or temporary files by ensuring they do not persist on the server.

## Purpose

- **File Cleanup**:  
  When enabled, if a client attempts to delete a file that matches one of the veto patterns, Samba will remove the file from the share rather than simply denying access.
  
- **Automated Maintenance**:  
  Helps in automatically cleaning up files that are not meant to be stored long-term (such as temporary or system files like `Thumbs.db` or `desktop.ini`).

## How It Works

- **Veto Files Module**:  
  Samba’s veto_files module allows administrators to specify file patterns (using the `veto files` parameter) that should be hidden from users. These files are not normally visible in the share.
  
- **Deletion Behavior**:  
  With `delete veto files = yes`, if a user attempts to delete a file that is vetoed, Samba will perform the deletion on the underlying filesystem instead of returning an error or ignoring the command.

## Configuration Example

Below is an example of how to configure the veto_files module along with the deletion option in your `smb.conf`:

```ini
[public]
   path = /srv/samba/public
   read only = no
   vfs objects = veto_files
   veto files = /Thumbs.db/desktop.ini/
   delete veto files = yes
```

- **`vfs objects = veto_files`**:  
  Loads the veto_files module.
  
- **`veto files = /Thumbs.db/desktop.ini/`**:  
  Specifies that files matching `Thumbs.db` and `desktop.ini` are vetoed (hidden).
  
- **`delete veto files = yes`**:  
  Ensures that when a vetoed file is deleted by a client, it is actually removed from the filesystem.

## Use Cases

- **Automatic Removal of Unwanted Files**:  
  In environments where temporary or system files (e.g., `Thumbs.db`, `desktop.ini`) are created, this setting ensures these files do not accumulate and clutter the share.

- **Simplified File Management**:  
  Prevents orphaned or unnecessary files from remaining on the server, reducing administrative overhead.

## Considerations

- **Irreversible Deletion**:  
  Enabling this option means that once a vetoed file is deleted, it is permanently removed. Ensure that your veto file patterns are correctly defined to avoid accidental deletion of important files.
  
- **Appropriate Use of Veto Patterns**:  
  Carefully tailor your `veto files` patterns to target only those files that should not persist in the share.

## Conclusion

The `delete veto files` parameter, when used with Samba’s veto_files module, provides an effective way to automatically clean up unwanted files from shares. This not only helps maintain a tidy file system but also reduces the manual effort required to manage temporary or system-generated files.
