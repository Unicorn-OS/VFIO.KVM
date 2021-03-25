# This will create inventory & ansible.cfg to match your uniNetwork! You can use the Default - Dyanmic Inventory - or edit static files.
# You can add inventory files to .net/hosts and it'll be ignored by git.
dynamic(){
  echo todo
}

static(){
  cp net.e hosts/net
  ln -srf ansible.cfg.e ../ansible.cfg
}

main(){
  static
}
main
