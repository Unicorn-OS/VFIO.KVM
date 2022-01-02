# Alter

## All *
docs.openstack.org/image-guide/modify-images.html

## Cloud-init
cloud-init.io/
wiki.archlinux.org/index.php/Cloud-init
cloudinit.readthedocs.io/en/latest/

## Guestfish
access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-guest_virtual_machine_disk_access_with_offline_tools-the_guestfish_shell

## libguestfs
libguestfs.org

## virt-customize

## Virt-resize
libguestfs.org/virt-resize.1.html
Lvm is supported


# Clone
## virt-sysprep
libguestfs.org/virt-sysprep.1.html
access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/sect-guest_virtual_machine_disk_access_with_offline_tools-using_virt_sysprep

# Builder
## nbdkit
rwmj.wordpress.com/tag/nbdkit/
rwmj.wordpress.com/2019/04/13/virt-install-nbdkit-live-install/


# People
rwmj.wordpress.com
rwmj.wordpress.com/tag/libguestfs/
rwmj.wordpress.com/2018/05/02/dockerfile-for-running-libguestfs-virt-tools-and-virt-v2v/

# Snapshot
## lvm_thin
access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/logical_volume_manager_administration/lv_overview#thinly-provisioned_snapshot_volumes

''A thin snapshot volume does not need to be activated with its origin, so a user may have only the origin active while there are many inactive snapshot volumes of the origin.
When you delete the origin of a thinly-provisioned snapshot volume, each snapshot of that origin volume becomes an independent thinly-provisioned volume. This means that instead of merging a snapshot with its origin volume, you may choose to delete the origin volume and then create a new thinly-provisioned snapshot using that independent volume as the origin volume for the new snapshot.''

# Store
## Cache-high-performance
https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/logical_volume_manager_administration/lv_overview#cache_volumes
