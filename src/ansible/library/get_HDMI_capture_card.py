#!/usr/bin/python

from __future__ import (absolute_import, division, print_function)
__metaclass__ = type

RETURN = r'''
# These are examples of possible return values, and in general should use other names for return values.
original_message:
    description: The original name param that was passed in.
    type: str
    returned: always
    sample: 'hello world'
message:
    description: The output message that the test module generates.
    type: str
    returned: always
    sample: 'goodbye'
'''

import re
import subprocess

from ansible.module_utils.basic import AnsibleModule

def getAllPCI():
    result = subprocess.run(['lspci', '-nn'], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    output = result.stdout.decode('utf-8')
    all = output.splitlines()
    return all

def isElgato(device):
    if device.find("YUAN High-Tech Development Co.") > 0:
        return True
    else:
        return False

def getPCIDevice(bus):
    all = getAllPCI()
    device = ""
    for line in all:
        if line.find(bus) >= 0:
            device = line
    return device

def getCaptureCard():
    all = getAllPCI()
    cards = []
    for line in all:
        if isElgato(line): cards.append(line)
    return cards

def countCards():
    return len(getCaptureCard())   

def getBus(card_num = 0):
    card = getCaptureCard()
    if countCards() >= card_num:
        return card[card_num][0:7]
    else:
        try:
            print(x)
        except:
            print("Card Not found!")

def toHex(bus):
    # regex = "(?<=\[)....:....(?=\])"
    regex = r"(?<=\[)....:....(?=\])"
    pattern = re.compile(rf"{regex}")

    lspci = getPCIDevice(bus)
    bushex = pattern.findall(lspci)
    return bushex[0]

def pciString(card_num):
    pci = toHex(getBus(card_num))
    return f"{pci}"

def cardType(card_num = 0):
    elgato_4k = "YUAN High-Tech Development Co."
    todo = "todo"
    intel = "todo1"
    
    pci_bus = getBus(card_num)
    lspci = getPCIDevice(pci_bus)
    
    if lspci.find(elgato_4k) >= 0:
        card_type = "Elgato"
    elif lspci.find(todo) >= 0:
        card_type = "Nvidia"
    else:
        card_type = "none"
    return card_type    

def _test():
    print(getBus(1))

# def main():
#     print(pciString(1))

#----------------------------------------------------
def run_module():
    module_args = dict(
        card_num=dict(type='int', required=False, default=0),
        make=dict(type='str', required=True)
    )

    # seed the result dict in the object
    result = dict(
        changed=False,
        pci_bus='',
        pci_hex='',
        pci_string='',
        card_type=''
    )

    # the AnsibleModule object will be our abstraction working with Ansible
    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    # if the user is working with this module in only check mode we do not
    if module.check_mode:
        module.exit_json(**result)

    # manipulate or modify the state as needed (this is going to be the
    card_num = module.params['card_num']

    result['pci_bus'] = getBus(card_num)
    result['pci_hex'] = toHex(getBus(card_num))
    result['pci_string'] = pciString(card_num)
    result['card_type'] = cardType(card_num)

    # use whatever logic you need to determine whether or not this module
    result['changed'] = False

    # in the event of a successful module execution, you will want to
    module.exit_json(**result)


def main():
    run_module()


if __name__ == '__main__':
    main()
