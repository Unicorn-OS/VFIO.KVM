# pcie_acs_override=downstream
https://pve.proxmox.com/wiki/Pci_passthrough

'''If you don't have dedicated IOMMU groups, you can try:

1) moving the card to another pci slot

2) adding "pcie_acs_override=downstream" to kernel boot commandline (grub or systemd-boot) options, which can help on some setup with bad ACS implementation.

Checkout the documentation about Editing the kernel commandline
More infos:

http://vfio.blogspot.be/2015/10/intel-processors-with-acs-support.html http://vfio.blogspot.be/2014/08/iommu-groups-inside-and-out.html
'''
