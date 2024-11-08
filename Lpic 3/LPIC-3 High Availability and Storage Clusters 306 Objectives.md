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

* [Active/Passive Cluster](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/active-passive-cluster.md)

* [Active/Active Cluster](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Active-Active-Cluster.md)

* [Failover Cluster](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Failover-Cluster.md)

* [Load Balanced Cluster](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Load-Balanced-Cluster.md)

* [Shared-Nothing Cluster](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Shared-Nothing%20Cluster.md)

* [Shared-Disk Cluster](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Shared-Disk%20Cluster.md)

* [Cluster resources](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Cluster-resources.md)

* [Cluster services](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Cluster-services.md)

* [Quorum](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Quorum.md)

* [Fencing (Node and Resource Level Fencing)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Fencing(Node-and-Resource-Level-Fencing).md)

* [Split brain](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Split-brain.md)

* [Redundancy](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Redundancy.md)

* [Mean Time Before Failure (MTBF)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Mean-Time-Before-Failure(MTBF).md)

* [Mean Time To Repair (MTTR)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Mean-Time-To-Repair(MTTR).md)

* [Service Level Agreement (SLA)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Service-Level-Agreement(SLA).md)

* [Disaster Recovery](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Disaster-Recovery.md)

* [State Handling](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/State-Handling.md)

* [Replication](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Replication.md)

* [Session handling](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Session-handling.md)

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

* [ipvsadm](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ipvsadm.md)

* [syncd](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/syncd.md)

* [LVS Forwarding (NAT, Direct Routing, Tunneling, Local Node)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/LVS-Forwarding(NAT%2CDirect%20Routing%2CTunneling%2CLocal%20Node).md)

* [connection scheduling algorithms](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/connection-scheduling-algorithms.md)

* [keepalived configuration file](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/keepalived-configuration-file.md)

* [ldirectord configuration file](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ldirectord-configuration-file.md)

* [genhash](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/genhash.md)

* [HAProxy configuration file](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/HAProxy-configuration-file.md)

* [load balancing algorithms](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/loadbalancing-algorithms.md)

* [ACLs](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ACLs.md)

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

* [pcs](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pcs.md)

* [crm](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm.md)

* [crm_mon](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_mon.md)

* [crm_verify](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_verify.md)

* [crm_simulate](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_simulate.md)

* [crm_shadow](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_shadow.md)

* [crm_resource](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_resource.md)

* [crm_attribute](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_attribute.md)

* [crm_node](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_node.md)

* [crm_standby](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crm_standby.md)

* [cibadmin](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/cibadmin.md)

* [corosync.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/corosync.conf.md)

* [authkey](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/authkey.md)

* [corosync-cfgtool](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/corosync-cfgtool.md)

* [corosync-cmapctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/corosync-cmapctl.md)

* [corosync-quorumtool](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/corosync-quorumtool.md)

* [stonith_admin](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/stonith_admin.md)

* [stonith](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/stonith.md)

* [ocf:pacemaker:ping]()

* [ocf:pacemaker:NodeUtilization](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ocfpacemakerping.md)

* [ocf:pacemaker:ocf:SysInfo](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ocfpacemakerNodeUtilization.md)

* [ocf:pacemaker:HealthCPU](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ocfpacemakerocfSysInfo.md)

* [ocf:pacemaker:HealthSMART](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ocfpacemakerHealthCPU.md)

* [sbd](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ocfpacemakerHealthSMART.md)

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

* [Protocol A, B and C](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ProtocolA%2CBandC.md)

* [Primary, Secondary](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Primary%2C%20Secondary.md)

* [Three-way replication](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/Three-way%20replication.md)

* [drbd kernel module](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/drbd-kernel-module.md)

* [drbdadm]()

* [drbdmon](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/drbdadm.md)

* [drbdsetup](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/drbdsetup.md)

* [drbdmeta](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/drbdmeta.md)

* [/etc/drbd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-drbd.conf.md)

* [/etc/drbd.d/](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-drbd.d-.md)

* [/proc/drbd](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-proc-drbd.md)

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

* [tgtadm]()

* [targets.conf]()

* [iscsiadm]()

* [iscsid.conf]()

* [/etc/multipath.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-multipath.conf.md)

* [multipath](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/multipath.md)

* [kpartx](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/kpartx.md)

* [pvmove]()

* [vgchange](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pvmove.md)

* [lvchange](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/vgchange.md)

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

* [mkfs.gfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvchange.md)

* [mount.gfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/mkfs.gfs2.md)

* [fsck.gfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/mount.gfs2.md)

* [gfs2_grow](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/gfs2_grow.md)

* [gfs2_edit](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/gfs2_edit.md)

* [gfs2_jadd](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/gfs2_jadd.md)

* [mkfs.ocfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/mkfs.ocfs2.md)

* [mount.ocfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/mount.ocfs2.md)

* [fsck.ocfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/fsck.ocfs2.md)

* [tunefs.ocfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/tunefs.ocfs2.md)

* [mounted.ocfs2](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/mounted.ocfs2.md)

* [o2info](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/o2info.md)

* [o2image](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/o2image.md)

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

* [gluster (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/gluster(including-relevant-subcommands).md)

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

* [ceph-deploy (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ceph-deploy%20(including%20relevant%20subcommands).md)

* [ceph.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ceph.conf.md)

* [ceph (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ceph(including-relevant-subcommands).md)

* [rados (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/rados-(including-relevant-subcommands).md)

* [rdb (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/rdb(including-relevant-subcommands).md)

* [cephfs (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/cephfs-(including%20relevant%20subcommands).md)

* [ceph-volume (including relevant subcommands)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ceph-volume%20(including%20relevant%20subcommands).md)

* [ceph-authtool](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ceph-authtool.md)

* [ceph-bluestore-tool](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ceph-bluestore-tool.md)

* [crushtool](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/crushtool.md)

## Topic 364: Single Node High Availability
### 364.1 Hardware and Resource High Availability (weight: 2)
#### Description:	Candidates should be able to monitor a local node for potential hardware failures and resource shortages.
Key Knowledge Areas:

Understand and monitor S.M.A.R.T values using smartmontools, including triggering frequent disk checks
Configure system shutdown at specific UPS events
Configure monit for alerts in case of resource exhaustion
Partial list of the used files, terms and utilities:

* [smartctl](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/smartctl.md)

* [/etc/smartd.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-smartd.conf.md)

* [smartd](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/smartd.md)

* [nvme-cli](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/nvme-cli.md)

* [apcupsd](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/apcupsd.md)

* [apctest](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/apctest.md)

* [monit](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/monit.md)

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

* [mdadm](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/mdadm.md)

* [/proc/mdstat](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-proc-mdstat.md)

* [/proc/sys/dev/raid/*](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-proc-sys-dev-raid.md)

### 364.3 Advanced LVM (weight: 3)
#### Description:	Candidates should be able to configure LVM volumes. This includes managing LVM snapshot, pools and RAIDs.
Key Knowledge Areas:

Understand and manage LVM, including linear and striped volumes
Extend, grow, shrink and move LVM volumes
Understand and manage LVM snapshots
Understand and manage LVM thin and thick pools
Understand and manage LVM RAIDs
Partial list of the used files, terms and utilities:

* [/etc/lvm/lvm.conf](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-lvm-lvm.conf.md)

* [pvcreate](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pvcreate.md)

* [pvdisplay](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pvdisplay.md)

* [pvmove](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pvmove.md)

* [pvremove](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pvremove.md)

* [pvresize](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/pvresize.md)

* [vgcreate](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/vgcreate.md)

* [vgdisplay](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/vgdisplay.md)

* [vgreduce](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/vgreduce.md)

* [lvconvert](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvconvert.md)

* [lvcreate](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvcreate.md)

* [lvdisplay](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvdisplay.md)

* [lvextend](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvextend.md)

* [lvreduce](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvreduce.md)

* [lvresize](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/lvresize.md)

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

* [bonding.ko (including relevant module options)](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/bonding.ko-(including-relevant-module-options).md)

* [/etc/network/interfaces](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-network-interfaces.md)

* [/etc/sysconfig/networking-scripts/ifcfg-*](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-sysconfig-networking-scripts-ifcfg-.md)

* [/etc/systemd/network/*.network]()

* [/etc/systemd/network/*.netdev](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-etc-systemd-network-.netdev.md)

* [nmcli](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/nmcli.md)

* [/sys/class/net/bonding_masters](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-sys-class-net-bond-bonding-miimon.md)

* [/sys/class/net/bond*/bonding/miimon](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-sys-class-net-bonding_masters.md)

* [/sys/class/net/bond*/bonding/slaves](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/-sys-class-net-bond-bonding-slaves.md)

* [ifenslave](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ifenslave.md)

* [ip](../TXT%20FILES/File-systems-Cocepts/LPIC3-306/ip.md)
