# samba-tool domain passwordsettings

`samba-tool domain passwordsettings` is a Samba utility that allows administrators to view and modify domain-wide password policy settings for a Samba Active Directory domain. It lets you enforce security measures such as minimum and maximum password age, minimum password length, password history, and account lockout policies.

## Overview

- **Purpose**:  
  Manage and enforce password policies across your domain, ensuring that user credentials adhere to your organization’s security requirements.

- **Key Functions**:  
  - Display current domain password policy settings.
  - Modify settings such as:
    - Minimum and maximum password age.
    - Minimum password length.
    - Password history length.
    - Account lockout thresholds and duration.
    - (Additional options may include password complexity settings and reversible encryption.)

## Common Subcommands and Options

### Viewing Current Password Settings

To display the current password policy, run:

```bash
samba-tool domain passwordsettings show
```

This outputs details like:
- Minimum password age
- Maximum password age
- Minimum password length
- Password history length
- Lockout threshold and duration
- Password complexity requirements

### Setting Password Policies

Use the `set` subcommand with the appropriate options to modify a policy.

#### Set Minimum Password Age
```bash
samba-tool domain passwordsettings set --min-pwd-age=1
```
*Sets the minimum password age to 1 day.*

#### Set Maximum Password Age
```bash
samba-tool domain passwordsettings set --max-pwd-age=90
```
*Sets the maximum password age to 90 days.*

#### Set Minimum Password Length
```bash
samba-tool domain passwordsettings set --min-pwd-length=8
```
*Requires a minimum of 8 characters for passwords.*

#### Set Password History Length
```bash
samba-tool domain passwordsettings set --pwd-history-length=5
```
*Keeps a history of the last 5 passwords to prevent reuse.*

#### Configure Account Lockout Policy
```bash
samba-tool domain passwordsettings set --lockout-threshold=5 --lockout-duration=30
```
*Locks an account after 5 failed login attempts for 30 minutes.*

### Additional Options

Depending on your Samba version, you might have additional options to:
- Configure password complexity settings.
- Enable or disable reversible encryption.
- Manage notification settings for expiring passwords.

## Example Usage

**Viewing the Current Password Policy:**
```bash
samba-tool domain passwordsettings show
```
_Example output:_
```
Minimum password age: 1 day
Maximum password age: 90 days
Minimum password length: 8 characters
Password history length: 5 passwords
Lockout threshold: 5 invalid attempts
Lockout duration: 30 minutes
Password complexity: enabled
```

**Enforcing a New Policy:**
To require passwords to be at least 10 characters long and to remember the last 7 passwords:
```bash
samba-tool domain passwordsettings set --min-pwd-length=10 --pwd-history-length=7
```

## Best Practices

- **Ensure Compliance**:  
  Align password settings with your organization's security policies to mitigate risks like brute-force attacks and password reuse.

- **Test Changes**:  
  Always test policy changes in a non-production environment first, as overly strict settings might inadvertently lock out users.

- **Document and Backup**:  
  Keep a record of your current settings and back up your Samba configuration, so you can restore policies if needed.

## Conclusion

The `samba-tool domain passwordsettings` command is essential for administering password policies in a Samba Active Directory environment. It provides a straightforward way to view and enforce critical security measures, ensuring that domain-wide credentials meet the required standards and protecting against common security threats.
