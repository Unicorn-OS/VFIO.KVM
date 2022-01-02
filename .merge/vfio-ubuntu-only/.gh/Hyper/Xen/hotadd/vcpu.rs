/* Works!
https://backdrift.org/how-to-hot-addremove-vcpus-from-a-xen-domain
fix:    https://wiki.xen.org/wiki/Paravirt_Linux_CPU_Hotplug
*/


fn Works_cfg(){
    // https://ycnrg.org/xen-install-os-from-iso-pv/

    xencfg = '''
    # We will be using PyGrub as the bootloader
    bootloader = '/usr/lib/xen-4.9/bin/pygrub'

    # Set hostname, memory, vpcus, etc.
    name = "my"
    memory = 1024
    maxmem = 6144

    vcpus = 6
    maxvcpus = 8

    # Use the same disk as used previously, with the same device name
    disk = [ 'phy:/dev/vg0/my,xvda,rw' ]

    # Set up a proper routed network connection
    vif = [ 'ip=192.168.XXX.XXX, mac=00:16:3E:29:33:14, gatewaydev=xenbr0:1' ]
    '''
}
