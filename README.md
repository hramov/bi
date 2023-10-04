# My test bi system

## How to start the system
1) Open a terminal (preferably GitBash or another UNIX-like terminal)
2) Navigate to deploy folder ```cd ./deploy```
3) Set executable flag to deploy.sh ```chmod +x ./deploy.sh```
4) Start the script ```./deploy.sh```

The script then pulls latest changes from ```dev``` branch and restarts containers. 

Script also creates external docker network ```bi``` to have access to host network.