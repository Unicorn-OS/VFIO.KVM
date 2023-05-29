# Works!
sch: https://www.google.com/search?q=ubuntu+vfio+gpu, https://www.google.com/search?q=vfio+blacklist+amdgpu

Steps:
1. https://askubuntu.com/questions/1406888/ubuntu-22-04-gpu-passthrough-qemu
2. Blacklist amdgpu https://forum.proxmox.com/threads/amd-gpu-passtrought-dont-work-anymore-since-upgrade-to-7-2-11.117377/#post-508095

# src:
```
echo '''## VFIO
blacklist snd_hda_intel
blacklist snd_hda_codec_hdmi
blacklist radeon
blacklist nouveau
blacklist nvidia
blacklist i2c-nvidia-gpu
blacklist amdgpu
blacklist nvidiafb
blacklist snd_hda_codec
blacklist snd_hda_core''' | sudo tee -a /etc/modprobe.d/blacklist.conf
```
