# exportfs

The `exportfs` command is used to maintain the NFS table of exported file systems. It is a crucial tool for managing NFS shares defined in the `/etc/exports` file. This guide explains the usage of `exportfs`, including its options and typical use cases.

### Usage

```bash
exportfs [options] [client:/directory ...]
```

### Common Options

- **`-a`**: Export or unexport all directories listed in `/etc/exports`.
- **`-r`**: Reexport all directories. This is useful after modifying `/etc/exports`.
- **`-u`**: Unexport one or more directories.
- **`-v`**: Verbose output. Useful for debugging and seeing detailed information.
- **`-o`**: Specify export options (overrides `/etc/exports`).

### Examples

#### Export All NFS Shares

To export all directories listed in `/etc/exports`:

```bash
sudo exportfs -a
```

#### Reexport All NFS Shares

To reexport all directories (e.g., after modifying `/etc/exports`):

```bash
sudo exportfs -r
```

#### Unexport All NFS Shares

To unexport all directories listed in `/etc/exports`:

```bash
sudo exportfs -ua
```

#### Export a Specific Directory

To export a specific directory to a specific client:

```bash
sudo exportfs client:/directory
```

Example:

```bash
sudo exportfs 192.168.1.100:/srv/nfs
```

#### Unexport a Specific Directory

To unexport a specific directory:

```bash
sudo exportfs -u client:/directory
```

Example:

```bash
sudo exportfs -u 192.168.1.100:/srv/nfs
```

#### Export with Specific Options

To export a directory with specific options (overriding `/etc/exports`):

```bash
sudo exportfs -o rw,sync,no_subtree_check client:/directory
```

Example:

```bash
sudo exportfs -o rw,sync,no_subtree_check 192.168.1.100:/srv/nfs
```

### Display Currently Exported File Systems

To see the list of currently exported file systems:

```bash
sudo exportfs -v
```

### Practical Use Cases

#### Adding a New Export

1. **Edit `/etc/exports`**:
   ```bash
   sudo nano /etc/exports
   ```
   Add a line for the new export:
   ```plaintext
   /new/export 192.168.1.0/24(rw,sync,no_subtree_check)
   ```

2. **Reexport NFS Shares**:
   ```bash
   sudo exportfs -r
   ```

#### Temporarily Export a Directory

To temporarily export a directory without modifying `/etc/exports`:

```bash
sudo exportfs -o rw,sync,no_subtree_check 192.168.1.100:/temporary/export
```

#### Unexporting a Directory for Maintenance

1. **Unexport the Directory**:
   ```bash
   sudo exportfs -u 192.168.1.100:/srv/nfs
   ```

2. **Perform Maintenance**.

3. **Reexport the Directory**:
   ```bash
   sudo exportfs 192.168.1.100:/srv/nfs
   ```

### Security Considerations

- **Restrict Access**: Always specify specific clients or networks to minimize unauthorized access.
- **Monitor Exports**: Regularly check the list of exported directories to ensure only intended shares are available.
- **Use Secure Options**: Utilize options like `root_squash` to mitigate risks associated with privileged access from clients.

### Conclusion

The `exportfs` command is a powerful tool for managing NFS shares, allowing administrators to export, unexport, and reexport directories efficiently. Proper use of this command, along with careful configuration of `/etc/exports`, ensures secure and reliable NFS file sharing.
