import time
from pylxd import Client

app = 'Citadel'


client = Client()

def ifExists( app=app ):
    #  Returns a handle if Container exiss! else return false
    if client.containers.exists(app):
        container = client.containers.get(app)
        return container
    else:
        return false

def createNode( app=app ):
    config = {'name': app, 'source': {'type': 'image', 'alias': 'ubuntu/bionic'}}
    container = client.containers.create(config, wait=True)
    container

def start( app=app ):
    if container.status == 'Stopped' :
        container.start(wait=True)
        print("starting container '" + app + "'")
        
def getNetwork( app=app ):
    if container.status == "Running" :
        addresses = container.state().network['eth0']['addresses']
        print(addresses[0])
    else:
        print(container.status)

def isRunning():
    if container.status == 'Stopped' :
        print("Offline")
        return False

    if container.status == 'Running' :
        print("Running normally")
        return True


container = ifExists()
#getNet()

#start()

#print(container.status)

isRunning()
start()

time.sleep(2)
isRunning() 
