# drive ACL permission
https://github.com/jedi4ever/veewee/issues/996

## Solution:
changing owner off all parent path to root, and
changing owner of .iso to 1001 Fixes it!
`chown 1001:1001 kubuntu-20.04.1-desktop-amd64.iso`
