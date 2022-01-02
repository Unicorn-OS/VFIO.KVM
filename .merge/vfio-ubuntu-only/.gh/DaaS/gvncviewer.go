inst(){
  sudo apt -y install gvncviewer
}

gvncXen(display){
// Works on xl create

if display = 0: then
  gvncviewer localhost
else
  gvncviewer localhost:display
fi
}
