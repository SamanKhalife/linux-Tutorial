To clear the cache memory in Linux, you can use the following methods:

1. **Clear RAM Cache:**
   To clear the RAM cache, you can use the following command:
   ```bash
   sudo sync; echo 3 > /proc/sys/vm/drop_caches
   ```
   This command first synchronizes all current data to the disk and then clears the cache memory.

2. **Clear Page Cache:**
   To clear only the page cache, execute the following command:
   ```bash
   sudo sync; echo 1 > /proc/sys/vm/drop_caches
   ```

3. **Clear Inode and Dentry Cache:**
   If you want to clear the inode and dentry caches as well, run this command:
   ```bash
   sudo sync; echo 2 > /proc/sys/vm/drop_caches
   ```

4. **Use the sync Command:**
   The `sync` command on its own can also synchronize all data to the disk and clear the cache memory to a certain extent:
   ```bash
   sudo sync
   ```

Please note that using these commands should be done carefully. Additionally, these commands should be executed as the root user or with sudo privileges. Clearing the cache memory may have an impact on system performance, so use it as needed and be aware of the importance of the data in the cache.