resx = 1440
resy = 900


pvfb_Resolution(){
  /* Works in Ubuntu!
  https://wiki.xenproject.org/wiki/Xen_Common_Problems#How_do_I_change_the_resolution_of_Xen_PV_domU_vfb_graphical_VNC_console.3F
  */
  echo "options xen-fbfront video=32,$resx,$resy" | sudo tee /etc/modprobe.d/pvfb.conf
  sudo update-initramfs -u
}

Mouse_fix(){
  // http://www.exonotes.com/node/76

  xinput list

  mouse=7

  exec xinput set-prop $mouse 'Evdev Axis Calibration' 0 $resx 0 $resy
}
