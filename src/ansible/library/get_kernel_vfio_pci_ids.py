# Gets currently active 'vfio-pci.ids'!

def rpmOstreeKargs():
    cmd = "rpm-ostree kargs"

def getProcCmdline():
    cmd = "cat /proc/cmdline"

if distro = "Silverblue":

else:
    getProcCmdline()