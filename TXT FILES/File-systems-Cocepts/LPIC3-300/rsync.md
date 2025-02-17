# rsync

**`rsync`** (remote sync) is a fast, versatile, and widely-used command-line utility for synchronizing and transferring files and directories between different systems. It is commonly used for backup, mirroring, and data migration due to its efficiency and flexibility.

### **Key Features of `rsync`**:
1. **Incremental File Transfer**:  
   `rsync` only transfers the changes made to files (based on their timestamps or checksums), reducing the amount of data transferred, which makes it ideal for backups or synchronizations.
   
2. **Compression**:  
   `rsync` supports file compression during transfer using the `-z` option, which reduces network usage.
   
3. **Preserves File Attributes**:  
   It can preserve file permissions, timestamps, symbolic links, and other metadata with options like `-a` (archive mode).
   
4. **Efficient Network Usage**:  
   It minimizes data transfer by using algorithms that detect file changes efficiently. This makes `rsync` ideal for both local and remote file transfers.
   
5. **Secure Transfers**:  
   `rsync` can work over SSH (secure shell) for encrypted transfers, making it a secure option for transferring files over untrusted networks.

6. **Flexible Synchronization**:  
   `rsync` can synchronize data between local directories, between a local and remote directory, or even between remote systems (remote-to-remote), all with various filtering options.

---

### **Basic Syntax**:
```bash
rsync [options] <source> <destination>
```
- **`<source>`**: The path to the source file or directory.
- **`<destination>`**: The path to the destination file or directory.

---

### **Commonly Used `rsync` Options**

1. **`-a` (archive mode)**:  
   This is a shorthand for a combination of options that preserve symbolic links, file permissions, timestamps, and recursively copies directories. It's often used for backups and syncing directories.
   ```bash
   rsync -a /source/directory/ /destination/directory/
   ```

2. **`-v` (verbose)**:  
   Displays detailed information about the file transfer process.
   ```bash
   rsync -av /source/directory/ /destination/directory/
   ```

3. **`-z` (compress)**:  
   Compresses file data during the transfer to reduce the amount of data sent over the network.
   ```bash
   rsync -avz /source/directory/ user@remote:/destination/directory/
   ```

4. **`-r` (recursive)**:  
   Recursively copies directories. This is often implied with the `-a` flag.
   ```bash
   rsync -r /source/directory/ /destination/directory/
   ```

5. **`-e` (remote shell)**:  
   Specifies the remote shell to use (commonly used with SSH).
   ```bash
   rsync -avz -e ssh /source/directory/ user@remote:/destination/directory/
   ```

6. **`--delete`**:  
   Deletes files from the destination directory that are no longer present in the source directory. This is useful for making the destination an exact mirror of the source.
   ```bash
   rsync -av --delete /source/directory/ /destination/directory/
   ```

7. **`--dry-run`**:  
   Simulates the operation without making any actual changes. This is useful for testing what will happen before actually performing the sync.
   ```bash
   rsync -av --dry-run /source/directory/ /destination/directory/
   ```

8. **`-u` (skip files that are newer in the destination)**:  
   This option skips files that are already up-to-date in the destination.
   ```bash
   rsync -avu /source/directory/ /destination/directory/
   ```

9. **`-P` (partial + progress)**:  
   This option shows progress during the transfer and keeps partially transferred files if the transfer is interrupted.
   ```bash
   rsync -avP /source/directory/ /destination/directory/
   ```

10. **`--exclude` and `--include`**:  
    You can exclude or include specific files or directories during the sync. This is useful for excluding temporary files or system files.
    ```bash
    rsync -av --exclude "*.log" /source/directory/ /destination/directory/
    ```

11. **`--bwlimit=KBPS`**:  
    Limits the transfer rate in kilobytes per second. This can be useful when you need to throttle the bandwidth usage.
    ```bash
    rsync -av --bwlimit=1000 /source/directory/ /destination/directory/
    ```

---

### **Example Commands**

1. **Copying Files from Local to Remote**:
   ```bash
   rsync -avz /local/directory/ user@remote:/remote/directory/
   ```

2. **Copying Files from Remote to Local**:
   ```bash
   rsync -avz user@remote:/remote/directory/ /local/directory/
   ```

3. **Synchronizing Directories Locally (Two-Way Sync)**:
   This will sync changes between two local directories.
   ```bash
   rsync -av --delete /source/directory/ /destination/directory/
   ```

4. **Backup with Compression and Exclusion of Certain Files**:
   Copy files to a backup directory while excluding log files.
   ```bash
   rsync -avz --exclude "*.log" /source/directory/ /backup/directory/
   ```

5. **Remote-to-Remote Sync**:
   Sync data directly between two remote servers without needing to transfer data through the local machine.
   ```bash
   rsync -avz -e ssh user@remote1:/source/directory/ user@remote2:/destination/directory/
   ```

---

### **Using `rsync` as a Daemon**

In addition to being used from the command line for one-time transfers, `rsync` can run as a daemon for continuous synchronization or for centralized file sharing. The daemon is typically configured via the `rsyncd.conf` file.

1. **Starting the `rsync` Daemon**:
   ```bash
   rsync --daemon
   ```

2. **Example of an `rsyncd.conf` file**:
   ```ini
   [backup]
   path = /srv/backup
   comment = Backup directory
   read only = no
   auth users = backupuser
   secrets file = /etc/rsyncd.secrets
   ```

3. **Connecting to the `rsync` Daemon**:
   ```bash
   rsync -avz /local/directory/ rsync://user@remote/backup
   ```

---

### **Common Use Cases**

1. **Backups**:  
   Regularly sync files and directories to an external or remote server as part of a backup strategy.

2. **Mirroring**:  
   Used to create mirrors of remote directories, websites, or file servers.

3. **Data Migration**:  
   Move or replicate data between servers, especially for large-scale migrations where minimal downtime is crucial.

4. **File Synchronization**:  
   Keep files in sync across multiple systems, for example, in a multi-server setup where data consistency across machines is required.

---

### **Conclusion**

`rsync` is a powerful and versatile tool for copying, synchronizing, and backing up files across different systems. With options like incremental copying, bandwidth control, secure transfers, and detailed logging, it is one of the most popular utilities for system administrators, especially for tasks like backups, migrations, and file synchronization. Whether you're working with local or remote files, `rsync` provides robust performance with minimal overhead, making it an indispensable tool for managing files on both individual systems and large-scale environments.
