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

def getGPU():
    all = getAllPCI()
    gpus = []
    for line in all:
        if isElgato(line): gpus.append(line)
    return gpus

def countGPUs():
    return len(getGPU())   

def getBus(card_num = 0):
    gpu = getGPU()
    if countGPUs() >= card_num:
        return gpu[card_num][0:7]
    else:
        try:
            print(x)
        except:
            print("GPU Not found!")

def toHex(bus):
    regex = "(?<=\[)....:....(?=\])"
    pattern = re.compile(rf"{regex}")

    lspci = getPCIDevice(bus)
    bushex = pattern.findall(lspci)
    return bushex[0]

def pciString(card_num):
    pci = toHex(getBus(card_num))
    return f"{pci}"

def gpuType(card_num = 0):
    elgato_4k = "Elgato"
    nvidia = "GeForce"
    intel = "Intel"
    
    gpu_bus = getBus(card_num)
    lspci = getPCIDevice(gpu_bus)
    
    if lspci.find(elgato_4k) >= 0:
        card_type = "Elgato"
    elif lspci.find(nvidia) >= 0:
        card_type = "Nvidia"
    else:
        card_type = "none"
    return card_type    

def _test():
    print(getBus(1))

# def main():
#     print(pciString(1))

# main()

#----------------------------------------------------
def run_module():
    module_args = dict(
        pci_num=dict(type='int', required=False, default=1),
        make=dict(type='str', required=True),
    )

    # seed the result dict in the object
    # we primarily care about changed and state
    # changed is if this module effectively modified the target
    # state will include any data that you want your module to pass back
    # for consumption, for example, in a subsequent task
    result = dict(
        changed=False,
        gpu='',
        pci_string='',
    )

    # the AnsibleModule object will be our abstraction working with Ansible
    # this includes instantiation, a couple of common attr would be the
    # args/params passed to the execution, as well as if the module
    # supports check mode
    module = AnsibleModule(
        argument_spec=module_args,
        supports_check_mode=True
    )

    # if the user is working with this module in only check mode we do not
    # want to make any changes to the environment, just return the current
    # state with no modifications
    if module.check_mode:
        module.exit_json(**result)

    # manipulate or modify the state as needed (this is going to be the
    # part where your module will do what it needs to do)
    card_num = module.params['card_num']

    result['pci'] = toHex(getBus(card_num))
    result['audio'] = toHex(getAudioBus(card_num))
    result['pci_string'] = pciString(card_num)
    result['card_type'] = gpuType(card_num)

    # use whatever logic you need to determine whether or not this module
    # made any modifications to your target
    result['changed'] = False

    # during the execution of the module, if there is an exception or a
    # conditional state that effectively causes a failure, run
    # AnsibleModule.fail_json() to pass in the message and the result
    if module.params['name'] == 'fail me':
        module.fail_json(msg='You requested this to fail', **result)

    # in the event of a successful module execution, you will want to
    # simple AnsibleModule.exit_json(), passing the key/value results
    module.exit_json(**result)


def main():
    run_module()


if __name__ == '__main__':
    main()
