import re
import subprocess
def getAllPCI():
    result = subprocess.run(['lspci', '-nn'], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    output = result.stdout.decode('utf-8')
    all = output.splitlines()
    return all

def isAudio(device):
    if device.find("Audio device") > 0:
        return True
    else:
        return False

def isGPU(device):
    if device.find("VGA") > 0:
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
        if isGPU(line): gpus.append(line)
    return gpus

def countGPUs():
    return len(getGPU())

def getBus(gpu_num = 0):
    gpu = getGPU()
    if countGPUs() >= gpu_num:
        return gpu[gpu_num][0:7]
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

def getAudioBus(gpu_num = 0):
    gpu_bus = getBus(gpu_num)
    audio_bus = gpu_bus[:6] + str(int(gpu_bus[6]) + 1)

    device = getPCIDevice(audio_bus)
    if isAudio(device):
        return audio_bus
    else:
        return False

def gpuString(gpu_num):
    gpu = toHex(getBus(gpu_num))
    audio = toHex(getAudioBus(gpu_num))
    return f"{gpu},{audio}"

def _test():
    print(getBus(1))

def main():
    print(gpuString(1))

main()

# from:
# - https://www.heiko-sieger.info/blacklisting-graphics-driver/
# - https://www.w3docs.com/snippets/python/running-shell-command-and-capturing-the-output.html
# - https://stackoverflow.com/questions/13332268/how-to-use-subprocess-command-with-pipes
# - https://youtu.be/AEE9ecgLgdQ