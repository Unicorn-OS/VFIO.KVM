# For adding another user! fixes libvirt permissions
# works: Ubuntu 20.04

https://www.google.com/search?q=virt-manager+as+another+user
https://computingforgeeks.com/use-virt-manager-as-non-root-user/
https://www.poftut.com/use-virt-manager-libvirt-normal-user-without-root-privileges-without-asking-password/

# Solution:
## Comment:
'With Fedora 27 there is already a built-in “libvirt” group and a polkit rule in /usr/share/polkit-1/rules.d/50-libvirt.rules, so all you have to do is to add the user to the libvirt group.'
- www.poftut.com/use-virt-manager-libvirt-normal-user-without-root-privileges-without-asking-password/#comment-309

copy this file:
`cp /usr/share/polkit-1/rules.d/*-libvirt.rules /etc/polkit-1/rules.d/80-libvirt.rules`
