#!/usr/bin/python3

import subprocess

def getAllPCI():
    result = subprocess.run(['lspci'], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    output = result.stdout.decode('utf-8')
    all = output.splitlines()
    return all

def isGPU(line):
    if line.find("VGA") > 0:
        return True
    else:
        return False

def getGPU():
    all = getAllPCI()
    gpus = []
    for line in all:
        if isGPU(line): gpus.append(line)
    return gpus

def countGPUs():
    return len(getGPU())

def getBus(gpu_num=0):
    gpu = getGPU()
    if countGPUs() >= gpu_num:
        print(gpu[gpu_num][0:7])
    else:
        try:
            print(x)
        except:
            print("GPU Not found!")

def main():
    getBus(1)

main()

# from:
# - https://www.heiko-sieger.info/blacklisting-graphics-driver/
# - https://www.w3docs.com/snippets/python/running-shell-command-and-capturing-the-output.html
# - https://stackoverflow.com/questions/13332268/how-to-use-subprocess-command-with-pipes
# quote: "Instead, create the ps and grep processes separately, and pipe the output from one into the other, like so:"