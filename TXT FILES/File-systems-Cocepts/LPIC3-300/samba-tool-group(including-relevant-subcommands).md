# samba-tool group

The `samba-tool group` command is a versatile utility within the Samba suite used to manage groups in a Samba Active Directory domain. It allows administrators to create, modify, delete, and query group information, making it an essential tool for managing group-based permissions and access control in mixed Windows/Unix environments.

## Overview

`samba-tool group` helps administrators:

- **Create and Delete Groups**: Add new groups to the domain or remove obsolete ones.
- **Modify Group Attributes**: Change group properties such as descriptions or membership.
- **List and Query Groups**: Retrieve information on groups, including listing all groups or showing details for a specific group.
- **Manage Group Memberships**: Add or remove members from a group.

## Common Subcommands

Below are some of the most frequently used subcommands with examples.

### 1. **add**
- **Purpose**: Create a new group in the domain.
- **Usage**:
  ```bash
  samba-tool group add <GroupName>
  ```
- **Example**:
  ```bash
  samba-tool group add Marketing
  ```
  This command creates a new group called "Marketing".

### 2. **delete**
- **Purpose**: Delete an existing group from the domain.
- **Usage**:
  ```bash
  samba-tool group delete <GroupName>
  ```
- **Example**:
  ```bash
  samba-tool group delete OldGroup
  ```
  This removes the group "OldGroup" from the domain.

### 3. **list**
- **Purpose**: List all groups in the domain.
- **Usage**:
  ```bash
  samba-tool group list
  ```
- **Example**:
  ```bash
  samba-tool group list
  ```
  This command outputs a list of all groups defined in the domain.

### 4. **show**
- **Purpose**: Display detailed information about a specific group.
- **Usage**:
  ```bash
  samba-tool group show <GroupName>
  ```
- **Example**:
  ```bash
  samba-tool group show Marketing
  ```
  This displays detailed attributes of the "Marketing" group, such as its SID and description.

### 5. **addmembers**
- **Purpose**: Add one or more members to a group.
- **Usage**:
  ```bash
  samba-tool group addmembers <GroupName> <User1> [<User2> ...]
  ```
- **Example**:
  ```bash
  samba-tool group addmembers Marketing alice bob charlie
  ```
  This command adds the users "alice", "bob", and "charlie" to the "Marketing" group.

### 6. **removemembers**
- **Purpose**: Remove one or more members from a group.
- **Usage**:
  ```bash
  samba-tool group removemembers <GroupName> <User1> [<User2> ...]
  ```
- **Example**:
  ```bash
  samba-tool group removemembers Marketing alice
  ```
  This removes "alice" from the "Marketing" group.

### 7. **rename**
- **Purpose**: Rename an existing group.
- **Usage**:
  ```bash
  samba-tool group rename <OldGroupName> <NewGroupName>
  ```
- **Example**:
  ```bash
  samba-tool group rename Marketing NewMarketing
  ```
  This command renames the "Marketing" group to "NewMarketing".

## Practical Considerations

- **Administrative Privileges**:  
  Most operations require domain administrator credentials. Ensure you run these commands with appropriate privileges.

- **Consistency and Auditing**:  
  Regularly review group memberships using `samba-tool group list` and `samba-tool group show` to ensure that group policies and access controls remain consistent with your organizational needs.

- **Scripting and Automation**:  
  `samba-tool group` commands are scriptable, making it possible to automate common group management tasks in large deployments.

## Conclusion

The `samba-tool group` command is an essential tool for managing group objects in a Samba Active Directory environment. With subcommands for adding, deleting, listing, modifying, and managing group memberships, it provides a comprehensive interface for administrators to maintain and secure group-based access control. By integrating these commands into your administrative routines, you can ensure that group policies are consistently applied and that your domain remains well-organized and secure.
