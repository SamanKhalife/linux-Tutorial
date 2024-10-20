# LPIC-3 High Availability and Storage Clusters 306 Objectives
## Topic 361: High Availability Cluster Management
### 361.1 High Availability Concepts and Theory (weight: 6)
#### Description:	Candidates should understand the properties and design approaches of high availability clusters.
Key Knowledge Areas:

Understand the goals of High Availability and Site Reliability Engineering
Understand common cluster architectures
Understand recovery and cluster reorganization mechanisms
Design an appropriate cluster architecture for a given purpose
Understand application aspects of high availability
Understand operational considerations of high availability
Partial list of the used files, terms and utilities:

Active/Passive Cluster
Active/Active Cluster
Failover Cluster
Load Balanced Cluster
Shared-Nothing Cluster
Shared-Disk Cluster
Cluster resources
Cluster services
Quorum
Fencing (Node and Resource Level Fencing)
Split brain
Redundancy
Mean Time Before Failure (MTBF)
Mean Time To Repair (MTTR)
Service Level Agreement (SLA)
Disaster Recovery
State Handling
Replication
Session handling


### 361.2 Load Balanced Clusters (weight: 8)
#### Description:	Candidates should know how to install, configure, maintain and troubleshoot LVS. This includes the configuration and use of keepalived and ldirectord. Candidates should further be able to install, configure, maintain and troubleshoot HAProxy.
Key Knowledge Areas:

Understand the concepts of LVS / IPVS
Understand the basics of VRRP
Configure keepalived
Configure ldirectord
Configure backend server networking
Understand HAProxy
Configure HAProxy
Partial list of the used files, terms and utilities:

ipvsadm
syncd
LVS Forwarding (NAT, Direct Routing, Tunneling, Local Node)
connection scheduling algorithms
keepalived configuration file
ldirectord configuration file
genhash
HAProxy configuration file
load balancing algorithms
ACLs


### 361.3 Failover Clusters (weight: 8)
#### Description:	Candidates should have experience in the installation, configuration, maintenance and troubleshooting of a Pacemaker cluster. This includes the use of Corosync. The focus is on Pacemaker 2.x for Corosync 2.x.
Key Knowledge Areas:

Understand the architecture and components of Pacemaker (CIB, CRMd, PEngine, LRMd, DC, STONITHd)
Manage Pacemaker cluster configurations
Understand Pacemaker resource classes (OCF, LSB, Systemd, Service, STONITH, Nagios)
Manage Pacemaker resources
Manage resource rules and constraints (location, order, colocation).
Manage advanced resource features (templates, groups, clone resources, multi-state resources)
Obtain node information and manage node health
Manage quorum and fencing in a Pacemaker cluster
Configure the Split Brain Detector on shared storage
Manage Pacemaker using pcs
Manage Pacemaker using crmsh
Configure and management of corosync in conjunction with Pacemaker
Awareness of Pacemaker ACLs
Awareness of other cluster engines (OpenAIS, Heartbeat, CMAN)
Partial list of the used files, terms and utilities:

pcs
crm
crm_mon
crm_verify
crm_simulate
crm_shadow
crm_resource
crm_attribute
crm_node
crm_standby
cibadmin
corosync.conf
authkey
corosync-cfgtool
corosync-cmapctl
corosync-quorumtool
stonith_admin
stonith
ocf:pacemaker:ping
ocf:pacemaker:NodeUtilization
ocf:pacemaker:ocf:SysInfo
ocf:pacemaker:HealthCPU
ocf:pacemaker:HealthSMART
sbd


## Topic 362: High Availability Cluster Storage
### 362.1 DRBD (weight: 6)
#### Description:	Candidates are expected to have the experience and knowledge to install, configure, maintain and troubleshoot DRBD devices. This includes integration with Pacemaker. DRBD configuration of version 9.0.x is covered.
Key Knowledge Areas:

Understand the DRBD architecture
Understand DRBD resources, states and replication modes
Configure DRBD disks and devices
Configure DRBD networking connections and meshes
Configure DRBD automatic recovery and error handling
Configure DRBD quorum and handlers for split brain and fencing
Manage DRBD using drbdadm
Understand the principles of drbdsetup and drbdmeta
Restore and verify the integrity of a DRBD device after an outage
Integrate DRBD with Pacemaker
Understand the architecture and features of LINSTOR
Partial list of the used files, terms and utilities:

Protocol A, B and C
Primary, Secondary
Three-way replication
drbd kernel module
drbdadm
drbdmon
drbdsetup
drbdmeta
/etc/drbd.conf
/etc/drbd.d/
/proc/drbd


### 362.2 Cluster Storage Access (weight: 3)
#### Description:	Candidates should be able to connect a Linux node to remote block storage. This includes understanding common SAN technology and architectures, including management of iSCSI, as well as configuring multipathing for high availability and using LVM on a clustered storage.
Key Knowledge Areas:

Understand the concepts of Storage Area Networks
Understand the concepts of Fibre Channel, including Fibre Channel Topologies
Understand and manage iSCSI targets and initiators
Understand and configure Device Mapper Multipath I/O (DM-MPIO)
Understand the concept of a Distributed Lock Manager (DLM)
Understand and manage clustered LVM
Manage DLM and LVM with Pacemaker
Partial list of the used files, terms and utilities:

tgtadm
targets.conf
iscsiadm
iscsid.conf
/etc/multipath.conf
multipath
kpartx
pvmove
vgchange
lvchange



### 362.3 Clustered File Systems (weight: 4)
#### Description:	Candidates should be able to install, maintain and troubleshoot GFS2 and OCFS2 filesystems. This includes awareness of other clustered filesystems available on Linux.
Key Knowledge Areas:

Understand the principles of cluster file systems and distributed file systems
Understand the Distributed Lock Manager
Create, maintain and troubleshoot GFS2 file systems in a cluster
Create, maintain and troubleshoot OCFS2 file systems in a cluster
Awareness of the O2CB cluster stack
Awareness of other commonly used clustered file systems, such as AFS and Lustre
Partial list of the used files, terms and utilities:

mkfs.gfs2
mount.gfs2
fsck.gfs2
gfs2_grow
gfs2_edit
gfs2_jadd
mkfs.ocfs2
mount.ocfs2
fsck.ocfs2
tunefs.ocfs2
mounted.ocfs2
o2info
o2image



## Topic 363: High Availability Distributed Storage
### 363.1 GlusterFS Storage Clusters (weight: 5)
#### Description:	Candidates should be able to manage and maintain a GlusterFS storage cluster.
Key Knowledge Areas:

Understand the architecture and components of GlusterFS
Manage GlusterFS peers, trusted storge pools, bricks and volumes
Mount and use an existing GlusterFS
Configure high availability aspects of GlusterFS
Scale up a GlusterFS cluster
Replace failed bricks
Recover GlusterFS from a physical media failure
Restore and verify the integrity of a GlusterFS cluster after an outage
Awareness of GNFS
Partial list of the used files, terms and utilities:

gluster (including relevant subcommands)


### 363.2 Ceph Storage Clusters (weight: 8)
#### Description:	Candidates should be able to manage and maintain a Ceph Cluster. This includes the configuration of RGW, RDB devices and CephFS.
Key Knowledge Areas:

Understand the architecture and components of Ceph
Manage OSD, MGR, MON and MDS
Understand and manage placement groups and pools
Understand storage backends (FileStore and BlueStore)
Initialize a Ceph cluster
Create and manage Rados Block Devices
Create and manage CephFS volumes, including snapshots
Mount and use an existing CephFS
Understand and adjust CRUSH maps
Configure high availability aspects of Ceph
Scale up a Ceph cluster
Restore and verify the integrity of a Ceph cluster after an outage
Understand key concepts of Ceph updates, including update order, tunables and features
Partial list of the used files, terms and utilities:

ceph-deploy (including relevant subcommands)
ceph.conf
ceph (including relevant subcommands)
rados (including relevant subcommands)
rdb (including relevant subcommands)
cephfs (including relevant subcommands)
ceph-volume (including relevant subcommands)
ceph-authtool
ceph-bluestore-tool
crushtool



## Topic 364: Single Node High Availability
### 364.1 Hardware and Resource High Availability (weight: 2)
#### Description:	Candidates should be able to monitor a local node for potential hardware failures and resource shortages.
Key Knowledge Areas:

Understand and monitor S.M.A.R.T values using smartmontools, including triggering frequent disk checks
Configure system shutdown at specific UPS events
Configure monit for alerts in case of resource exhaustion
Partial list of the used files, terms and utilities:

smartctl
/etc/smartd.conf
smartd
nvme-cli
apcupsd
apctest
monit


### 364.2 Advanced RAID (weight: 2)
#### Description:	Candidates should be able to manage software raid devices on Linux. This includes advanced features such as partitonable RAIDs and RAID containers as well as recovering RAID arrays after a failure.
Key Knowledge Areas:

Manage RAID devices using various raid levels, including hot spare discs, partitionable RAIDs and RAID containers
Add and remove devices from an existing RAID
Change the RAID level of an existing device
Recover a RAID device after a failure
Understand various metadata formats and RAID geometries
Understand availability and performance properties of various raid levels
Configure mdadm monitoring and reporting
Partial list of the used files, terms and utilities:

mdadm
/proc/mdstat
/proc/sys/dev/raid/*


### 364.3 Advanced LVM (weight: 3)
#### Description:	Candidates should be able to configure LVM volumes. This includes managing LVM snapshot, pools and RAIDs.
Key Knowledge Areas:

Understand and manage LVM, including linear and striped volumes
Extend, grow, shrink and move LVM volumes
Understand and manage LVM snapshots
Understand and manage LVM thin and thick pools
Understand and manage LVM RAIDs
Partial list of the used files, terms and utilities:

/etc/lvm/lvm.conf
pvcreate
pvdisplay
pvmove
pvremove
pvresize
vgcreate
vgdisplay
vgreduce
lvconvert
lvcreate
lvdisplay
lvextend
lvreduce
lvresize


### 364.4 Network High Availability (weight: 5)
#### Description:	Candidates should be able to configure redundant networking connections and manage VLANs. Furthermore, candidates should have a basic understanding of BGP.
Key Knowledge Areas:

Understand and configure bonding network interface
Network bond modes and algorithms (active-backup, balance-tlb, balance-alb, 802.3ad, balance-rr, balance-xor, broadcast)
Configure switch configuration for high availability, including RSTP
Configure VLANs on regular and bonded network interfaces
Persist bonding and VLAN configuration
Understand the principle of autonomous systems and BGP to manage external redundant uplinks
Awareness of traffic shaping and control capabilities of Linux
Partial list of the used files, terms and utilities:

bonding.ko (including relevant module options)
/etc/network/interfaces
/etc/sysconfig/networking-scripts/ifcfg-*
/etc/systemd/network/*.network
/etc/systemd/network/*.netdev
nmcli
/sys/class/net/bonding_masters
/sys/class/net/bond*/bonding/miimon
/sys/class/net/bond*/bonding/slaves
ifenslave
ip
