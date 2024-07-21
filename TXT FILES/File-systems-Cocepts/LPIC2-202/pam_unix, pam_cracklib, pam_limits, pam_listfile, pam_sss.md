# pam_unix, pam_cracklib, pam_limits, pam_listfile, pam_sss
### Commonly Used PAM Modules

#### 1. `pam_unix`

The `pam_unix` module provides standard Unix authentication, integrating traditional password and account management. It handles authentication, account management, password updates, and session management.

**Common Usage in PAM Configuration:**

```plaintext
# Authentication management
auth       required   pam_unix.so

# Account management
account    required   pam_unix.so

# Password management
password   required   pam_unix.so

# Session management
session    required   pam_unix.so
```

**Common Options:**
- `nullok`: Allow null passwords.
- `try_first_pass`: Use the previously typed password.
- `use_first_pass`: Use the previously typed password, and don't prompt again.
- `shadow`: Use the shadow password file.

#### 2. `pam_cracklib`

The `pam_cracklib` module is used to enforce password strength requirements. It checks the complexity of passwords and ensures they meet certain criteria, such as length, the inclusion of different character types, and avoiding dictionary words.

**Common Usage in PAM Configuration:**

```plaintext
# Password management
password    required    pam_cracklib.so retry=3 minlen=8 difok=3
```

**Common Options:**
- `retry=N`: Prompt the user up to N times before returning an error.
- `minlen=N`: Set the minimum length of the password.
- `difok=N`: Require at least N different characters between the new and old passwords.
- `ucredit=N`, `lcredit=N`, `dcredit=N`, `ocredit=N`: Enforce the inclusion of uppercase, lowercase, digits, and other characters.

#### 3. `pam_limits`

The `pam_limits` module sets resource limits for user sessions. It enforces the limits specified in the `/etc/security/limits.conf` file.

**Common Usage in PAM Configuration:**

```plaintext
# Session management
session    required    pam_limits.so
```

**Configuration in `/etc/security/limits.conf`:**

```plaintext
# Limit core file size to 0 for all users
*      hard    core       0

# Allow unlimited number of processes for user "admin"
admin  hard    nproc      unlimited
```

#### 4. `pam_listfile`

The `pam_listfile` module is used to allow or deny access based on the contents of a file. This module can check usernames, user groups, terminal names, host names, and more against a list in a specified file.

**Common Usage in PAM Configuration:**

```plaintext
# Authentication management
auth      required    pam_listfile.so onerr=succeed item=user sense=allow file=/etc/security/allowed_users
```

**Common Options:**
- `onerr=succeed|fail`: Determine the action on error.
- `item=user|tty|rhost|group`: Specify the item to check.
- `sense=allow|deny`: Specify whether to allow or deny access.
- `file=/path/to/file`: Specify the file containing the list.

#### 5. `pam_sss`

The `pam_sss` module is used for integrating PAM with the System Security Services Daemon (SSSD). It provides access to remote identity and authentication providers.

**Common Usage in PAM Configuration:**

```plaintext
# Authentication management
auth       required    pam_sss.so

# Account management
account    required    pam_sss.so

# Password management
password   required    pam_sss.so

# Session management
session    required    pam_sss.so
```

**Common Options:**
- `forward_pass`: Forward the password to the next PAM module.
- `use_first_pass`: Use the password provided by the first PAM module that required a password.
- `use_authtok`: Use the authentication token provided by a previously successful authentication.

### Conclusion

PAM (Pluggable Authentication Modules) provides a highly flexible and configurable framework for user authentication on Unix-like systems. The modules discussed—`pam_unix`, `pam_cracklib`, `pam_limits`, `pam_listfile`, and `pam_sss`—each serve distinct purposes ranging from traditional Unix authentication to enforcing password policies, setting resource limits, controlling access, and integrating with remote identity services.

Understanding and configuring these modules allows administrators to implement robust authentication mechanisms tailored to the specific needs and security policies of their systems. Properly configured PAM ensures a secure and manageable authentication process, crucial for maintaining system integrity and security.
