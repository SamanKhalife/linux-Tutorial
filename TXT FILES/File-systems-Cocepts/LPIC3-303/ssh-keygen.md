# ssh-keygen

The `ssh-keygen` command is a fundamental utility in the SSH (Secure Shell) protocol suite for managing SSH keys, including generating new keys, converting key formats, and managing key pairs for authentication.

### Overview of `ssh-keygen`

#### Purpose

`ssh-keygen` is primarily used for:
- Generating SSH key pairs (public and private keys).
- Converting between different key formats.
- Managing SSH keys, including adding or removing passphrase protection.

### Basic Usage

The general syntax for `ssh-keygen` is:

```bash
ssh-keygen [options]
```

### Common Options and Usage

1. **Generating a New SSH Key Pair**

   To generate a new SSH key pair (public and private key), use:

   ```bash
   ssh-keygen -t rsa -b 2048 -f ~/.ssh/id_rsa
   ```

   - `-t rsa`: Specifies the type of key to create (RSA).
   - `-b 2048`: Specifies the number of bits in the key (2048 bits).
   - `-f ~/.ssh/id_rsa`: Specifies the filename of the generated key (default is `id_rsa`).

   This command creates `id_rsa` (private key) and `id_rsa.pub` (public key) in the `~/.ssh/` directory.

2. **Specifying Key Type and Length**

   You can choose different key types (`-t`) and lengths (`-b`):
   - `rsa`: RSA key type.
   - `dsa`: DSA (Digital Signature Algorithm) key type.
   - `ecdsa`: ECDSA (Elliptic Curve Digital Signature Algorithm) key type.
   - `ed25519`: Ed25519 key type.

3. **Changing the Type or Size of an Existing Key**

   You can change the type or size of an existing key (e.g., convert from RSA to Ed25519) using `-t` and `-b`, followed by `-i` for input and `-o` for output.

4. **Managing Passphrases**

   To add or change the passphrase of a private key:
   ```bash
   ssh-keygen -p -f ~/.ssh/id_rsa
   ```

   To remove the passphrase (note the danger in doing so, as it removes a layer of security):
   ```bash
   ssh-keygen -p -P old_passphrase -N '' -f ~/.ssh/id_rsa
   ```

5. **Converting Between Key Formats**

   To convert a key to OpenSSH format from other formats:
   ```bash
   ssh-keygen -i -f input_key.pem > output_key.pub
   ```

   To convert an OpenSSH format key to other formats:
   ```bash
   ssh-keygen -e -f ~/.ssh/id_rsa.pub -m pem > id_rsa.pub.pem
   ```

6. **Viewing Key Fingerprints**

   To display the fingerprint of a key:
   ```bash
   ssh-keygen -lf ~/.ssh/id_rsa.pub
   ```

### Example Use Case

Generate an Ed25519 SSH key pair and display the public key fingerprint:
```bash
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519 -C "My Ed25519 Key"
ssh-keygen -lf ~/.ssh/id_ed25519.pub
```

### Conclusion

`ssh-keygen` is a versatile command-line tool for managing SSH keys, essential for secure authentication and encrypted communication in SSH environments. Understanding its options and usage allows administrators and users to effectively manage SSH keys for secure access to remote systems.
