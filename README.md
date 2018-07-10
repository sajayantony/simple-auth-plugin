# simple-auth-plugin

A very simple authorization plugin to logs the docker cli client requests. 

The authorization plugin needs to have a socket or TCP server which satisfies the [authz plugin API](https://docs.docker.com/engine/extend/plugins_authorization/#api-schema-and-implementation). 

One of the ways to create and enable a plugin is using a [rootfs and config.json](#using-docker-plugin-commands) and using the `docker plugin` commands. **NOT working yet**


## Steps to run and debug the plugin 

1. Build the container  

```
docker build -t simple-auth-plugin .
```

2. Run the plugin with a restart policy and ensure that the `/run/docker/plugins` directory is volume mounted. 

```
 docker run -d --restart=always -v /run/docker/plugins/:/run/docker/plugins simple-auth-plugin
 ```
 Make sure you have the container running 

```bash
 $ docker ps --filter ancestor=simple-auth-plugin
CONTAINER ID        IMAGE                COMMAND                  CREATED             STATUS              PORTS               NAMES
82e77369baad        simple-auth-plugin   "/bin/simple-auth-pl…"   2 minutes ago       Up 2 minutes                            romantic_sammet
 ```
 3. Restart the daemon with the plugin enabled 
 
 ```bash
 sudo  dockerd --authorization-plugin simple-auth-plugin
 ```
 Or you can use the dev build using `dockerd-dev`

```
sudo ./dockerd-dev --data-root /home/sajay/temp/docker-data-root/ --config-file /home/sajay/temp/docker-config/daemon.json  --authorization-plugin simple-auth-plugin
```
4. View the logs by following the docker logs. 
```
docker logs -f $(docker ps -q --filter ancestor=simple-auth-plugin)
```
