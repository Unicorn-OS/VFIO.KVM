# Hyperion1

Hyperion module for all OS's!

Make this dynamic using Go script!


# Keep!
[Sch:](https://www.google.com/search?q=dracut+vfio_pci)
- https://laketide.com/setting-up-gpu-passthrough-with-kvm-on-fedora/
- https://github.com/p7cq/VFIO-GPU-Configuration
- https://qubitrenegade.com/virtualization/kvm/vfio/2019/07/17/VFIO-Fedora-Notes.html


# Solution
- https://www.reddit.com/r/VFIO/comments/amqqlz/passthrough_guide_for_fedora_29/

'''rLinks234
IIRC, Fedora recently stopped including initramfs to generate the initial ramdisk. If the guide mentions it, just use dracut -f --kver $(uname -r) in its place.

This was at least one of the differences I've seen from Arch's guide on passthrough when setting up passthrough for my Fedora 29 desktop. While I don't use Arch, I think its documentation is far better than any other distros.
'''

## Works
'''cereal7802
So here is what I did to passthrough my GTX 970. you would need to consult the online guides to tweak to your setup.

I edited /etc/default/grub to add rd.driver.pre=vfio-pci and intel_iommu=on to the GRUB_CMDLINE_LINUX line:
'''

## Dynamic - Best!
'''qubitrenegade
These are my setup notes for Fedora 30. Fedora 29 was the same. I create a vfio-pci-override.sh to dynamically select the card to be passed through because I have two of the same card. I believe that this will work for setups with different cards... While I can't seem to find the documentation I believe the boot_vga flag is set for cards that initialize their display on boot... e.g.: if I had monitors plugged into both cards, neither would be passed through...

I would love any feedback. The "source" can be found here.
'''
- discuss: https://www.reddit.com/r/VFIO/comments/amqqlz/passthrough_guide_for_fedora_29/etznv9r/
prime.Ark https://github.com/primesrc/qubitrenegade.github.io
