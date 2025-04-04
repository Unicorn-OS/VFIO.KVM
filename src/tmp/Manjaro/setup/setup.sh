# Grub setting:
sudo nano /etc/default/grub

line="""
GRUB_CMDLINE_LINUX_DEFAULT='quiet apparmor=1 security=apparmor udev.log_priority=3 intel_iommu=on iommu=pt video=efifb:off vfio-pci.ids=1002:73ff,1002:ab28'
"""

sudo grub-mkconfig -o /boot/grub/grub.cfg

# Add Modules
sudo nano /etc/mkinitcpio.conf

line="""
MODULES=(crc32c-intel vfio_pci vfio vfio_iommu_type1)
"""


# Get Kernel version to use:
ls /etc/mkinitcpio.d/

version=66

sudo mkinitcpio -p linux${version}
