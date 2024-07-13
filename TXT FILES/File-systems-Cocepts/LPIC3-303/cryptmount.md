# cryptmount

**cryptmount** is a utility for managing encrypted filesystems on Linux systems. It simplifies the process of creating, mounting, and unmounting encrypted directories or devices, providing a user-friendly interface to handle encrypted storage.

#### Installation
Before using **cryptmount**, ensure it is installed on your Linux distribution. Installation methods may vary:

For Debian/Ubuntu:
```bash
sudo apt-get install cryptmount
```

For RHEL/CentOS:
```bash
sudo yum install cryptmount
```

#### Basic Usage

**1. Creating an Encrypted Filesystem**

To create an encrypted filesystem using **cryptmount**, follow these steps:

**Step 1: Create a Configuration File**

Create a configuration file (e.g., `/etc/cryptmount/my_encrypted_volume.conf`) with the details of the encrypted volume:

```plaintext
<your_encrypted_volume>
    keyfile=/path/to/keyfile
    cipher=aes
    size=256
    filesystem=ext4
    device=/dev/sdX1
```

Replace `<your_encrypted_volume>` with your chosen name and adjust the parameters (`keyfile`, `cipher`, `size`, `filesystem`, `device`) accordingly.

**Step 2: Initialize the Encrypted Volume**

```bash
sudo cryptmount --generate <your_encrypted_volume>
```

This command initializes the encrypted volume based on the configuration file.

**2. Mounting and Unmounting Encrypted Volumes**

**Mounting:**

```bash
sudo cryptmount --mount <your_encrypted_volume>
```

This command mounts the encrypted volume specified in the configuration file.

**Unmounting:**

```bash
sudo cryptmount --umount <your_encrypted_volume>
```

This command unmounts the encrypted volume.

#### Additional Commands and Options

- **Listing Encrypted Volumes:**

```bash
cryptmount --show-active
```

This command displays a list of currently active encrypted volumes.

- **Changing Volume Passphrase:**

```bash
sudo cryptmount --change-passphrase <your_encrypted_volume>
```

Allows you to change the passphrase used to unlock the encrypted volume.

- **Configuring Automatic Mounting:**

```plaintext
<your_encrypted_volume>
    keyfile=/path/to/keyfile
    cipher=aes
    size=256
    filesystem=ext4
    device=/dev/sdX1
    mountoptions=defaults,noatime
```

Add `mountoptions=defaults,noatime` to specify mount options for the encrypted volume.

- **Using Keyfiles:**

You can specify a keyfile (`keyfile=/path/to/keyfile`) in the configuration to automatically unlock the volume during mount.

#### Use Cases

**1. Data Encryption:**

cryptmount is ideal for securely storing sensitive data, ensuring it remains protected even if the storage device is compromised.

**2. Secure Backups:**

Encrypting backup disks or directories using cryptmount adds an extra layer of security to critical data.

**3. Compliance Requirements:**

For organizations needing to comply with data protection regulations, cryptmount provides robust encryption capabilities.

**4. Personal Privacy:**

Individual users can use cryptmount to encrypt personal data stored on external drives or cloud storage.

#### Conclusion

cryptmount simplifies the management of encrypted filesystems on Linux, providing an intuitive command-line interface for creating, mounting, and managing encrypted volumes. Whether for securing sensitive data or meeting compliance requirements, cryptmount offers a flexible and powerful solution for encrypted storage needs.
