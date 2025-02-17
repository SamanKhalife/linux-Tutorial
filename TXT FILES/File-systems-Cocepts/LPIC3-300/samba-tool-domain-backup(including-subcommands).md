# samba-tool domain backup
The `samba-tool domain backup` is a command in the **Samba** suite used for backing up and restoring domain controllers, specifically Active Directory (AD) related data. It is an essential tool for administrators to ensure that a copy of the domain's AD data is safely stored in case of system failure or disaster recovery.

### Overview of `samba-tool domain backup`

This command allows for the backup and restore of the **Samba Active Directory (AD)** domain controller's database, configurations, and other essential files. It is primarily used in **Samba 4.x** environments where Samba serves as an AD domain controller. The tool facilitates both full and incremental backups, including system state information.

The subcommands available under `samba-tool domain backup` provide various options for creating and managing backups. You can perform backups of specific domain controllers or the entire domain. It’s recommended to back up regularly to ensure that domain data is recoverable.

### General Syntax:
```bash
samba-tool domain backup [subcommand] [options]
```

Where:
- `subcommand` is the operation to perform (e.g., `offline`, `online`, `restore`, etc.)
- `options` are various command-line options to customize the behavior of the tool.

### Subcommands for `samba-tool domain backup`

1. **`offline`**  
   This subcommand allows you to take an offline backup of the domain controller's data. It requires the domain controller to be offline (i.e., not actively serving requests) while the backup is in progress. It's typically used for a snapshot of the entire domain controller system, including the Active Directory database and system state.

   #### Usage:
   ```bash
   samba-tool domain backup offline --path=/path/to/backup
   ```

   #### Options:
   - `--path=/path/to/backup`: Path to where the backup should be stored.
   - `--targetdir=TARGET_DIR`: Specifies the target directory for the backup.
   - `--no-wait`: If specified, this option will not wait for the backup to complete before returning control to the shell.

2. **`online`**  
   The online backup subcommand allows you to back up a domain controller while it's still running. This backup does not require the domain controller to go offline and can be used for a live backup of AD data.

   #### Usage:
   ```bash
   samba-tool domain backup online --path=/path/to/backup
   ```

   #### Options:
   - `--path=/path/to/backup`: Path to where the backup will be stored.
   - `--targetdir=TARGET_DIR`: Specifies the directory to store backup data.
   - `--no-wait`: Like in the offline backup, the backup process will not block the terminal.

3. **`restore`**  
   The restore subcommand allows you to restore a previously backed-up domain. This is critical in recovery scenarios when a domain controller is corrupted or needs to be migrated to a new machine. Restoring from a backup typically involves restoring the AD database and associated system states to a previous point in time.

   #### Usage:
   ```bash
   samba-tool domain backup restore --path=/path/to/backup
   ```

   #### Options:
   - `--path=/path/to/backup`: Path to the backup data.
   - `--targetdir=TARGET_DIR`: Directory where the restore will be applied.

4. **`list`**  
   This subcommand lists all available backups in a specified backup directory. It is helpful for verifying backup files and ensuring that the backup was performed correctly.

   #### Usage:
   ```bash
   samba-tool domain backup list --path=/path/to/backup
   ```

   #### Options:
   - `--path=/path/to/backup`: Path to the backup directory where backups are stored.

5. **`delete`**  
   This subcommand deletes a previously created backup. It can be used for cleanup after a successful backup and restore process.

   #### Usage:
   ```bash
   samba-tool domain backup delete --path=/path/to/backup --backup-id=ID
   ```

   #### Options:
   - `--path=/path/to/backup`: Path to the backup location.
   - `--backup-id=ID`: ID of the backup you want to delete.

6. **`verify`**  
   The verify subcommand checks the integrity of a backup to ensure that the backup is complete and hasn't been corrupted.

   #### Usage:
   ```bash
   samba-tool domain backup verify --path=/path/to/backup
   ```

   #### Options:
   - `--path=/path/to/backup`: Path where the backup is stored.

### Example Use Cases

#### Example 1: Offline Backup
To create an offline backup of the domain:
```bash
samba-tool domain backup offline --path=/mnt/backup/2023-10-17
```

#### Example 2: Online Backup
To create an online backup (without taking the domain controller offline):
```bash
samba-tool domain backup online --path=/mnt/backup/online-backup
```

#### Example 3: Restoring from Backup
To restore a backup from a specific path:
```bash
samba-tool domain backup restore --path=/mnt/backup/2023-10-17
```

#### Example 4: Listing Backups
To list available backups in a directory:
```bash
samba-tool domain backup list --path=/mnt/backup
```

#### Example 5: Verifying a Backup
To verify the integrity of a backup:
```bash
samba-tool domain backup verify --path=/mnt/backup/2023-10-17
```

#### Example 6: Deleting a Backup
To delete a specific backup:
```bash
samba-tool domain backup delete --path=/mnt/backup --backup-id=backup-id-2023-10-17
```

### Key Considerations:
- **Backup Frequency**: It's important to back up the AD domain controller regularly. Typically, daily or weekly backups are recommended depending on the rate of change in your domain.
- **Backup Storage**: Ensure that backups are stored in a secure, separate location from the domain controllers to protect against physical damage or theft.
- **Restore Testing**: It is recommended to test backups and restore procedures to ensure that they work as expected during disaster recovery scenarios.
- **Backup Size**: Domain backups can grow large, depending on the size of the Active Directory database and associated configurations. Monitor backup storage space to ensure sufficient room is available.

### Conclusion
The `samba-tool domain backup` command provides administrators with tools for backing up and restoring Samba Active Directory domain controllers. Regular backups, particularly with online backup capabilities, ensure that the domain controller can be restored quickly and effectively in case of failure or disaster recovery scenarios. Properly managing backups and understanding the available subcommands is critical to maintaining a reliable and secure domain infrastructure.
