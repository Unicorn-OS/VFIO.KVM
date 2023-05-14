# Favorite:
https://mathiashueber.com/fighting-error-43-nvidia-gpu-virtual-machine/

# wiki:
https://wiki.archlinux.org/title/PCI_passthrough_via_OVMF#.22Error_43:_Driver_failed_to_load.22_on_Nvidia_GPUs_passed_to_Windows_VMs

section: "Video card driver virtualisation detection
Video card drivers by AMD incorporate very basic virtual machine detection targeting Hyper-V extensions. Should this detection mechanism trigger, the drivers will refuse to run, resulting in a black screen.

If this is the case, it is required to modify the reported Hyper-V vendor ID:
"
