sch: https://www.google.com/search?q=linux+kernel+6+vfio-pci https://www.google.com/search?q=linux+kernel+6+vfio+blacklist

# discuss:
https://forum.level1techs.com/t/linux-kernel-6-seems-to-be-incompatible-with-the-vfio-pci-module-needed-for-pci-passthrough/190039

Quote:
- With it working for @Janos, that looks like the approach that I had to do when I ran debian. Maybe I’ll give that a try just with nvidia instead of AMD.
- UPDATE - that worked. I added “amdgpu” before the vfio hooks and the prompt came back (after a second long black screen as it switched to using the igpu). Seems like a decent workaround for the time being https://forum.level1techs.com/t/linux-kernel-6-seems-to-be-incompatible-with-the-vfio-pci-module-needed-for-pci-passthrough/190039/11
- I have managed to workaround this issue and posted the solution on my blog, hope this helps. hardcoded.info/post/2023/03/04/pci-passthrough-using-vfio-pci-for-linux-kernel-version-6-solution/ -https://forum.level1techs.com/t/linux-kernel-6-seems-to-be-incompatible-with-the-vfio-pci-module-needed-for-pci-passthrough/190039/15
- The reason for this behavior is a change in Ubuntu kernel configuration. vfio-pci is built as a module from 6.2 on, like ArchLinux: https://forum.level1techs.com/t/linux-kernel-6-seems-to-be-incompatible-with-the-vfio-pci-module-needed-for-pci-passthrough/190039/37


# grub method
guide: https://www.heiko-sieger.info/vfio-grub-vfio-pci-ids-doesnt-work-with-kernel-6-try-driver-verride-feature/
