echo '''
blacklist xhci_hcd
blacklist intel_pch_thermal
''' | sudo tee /etc/modprobe.d/blacklist-usb.conf
