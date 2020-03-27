# Mattermost Jitsi Plugin

Mattemost slash command to create [Jitsi](https://meet.jit.si/) Meeting Room. Usage:
```
/jitsi
```
It will be created room named team-name_channel-name
Or you can provide custom room name:
```
/jitsi testin123
```
![Alt text](https://cloud.appmodule.net/s/HCt7ExBLnJKonHm/preview "Showroom")

You can also set Jitsi URL in case that you have self-hosted Jitsi:
![Alt text](https://cloud.appmodule.net/s/HkkQbYnNebcy5Rn/preview "Settings")
## Download and install
Download plugin [HERE](https://github.com/appmodule/mattermost-plugin-jitsi/releases)
You can install it by following instructions [HERE](https://docs.mattermost.com/administration/plugins.html#custom-plugins)

## Build
You can build plugin with:
```
make
```
You need to have Go build tools.
