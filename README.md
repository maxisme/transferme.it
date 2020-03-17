<p align="center"><img height="150px" src="https://github.com/maxisme/transferme.it/raw/master/images/og_logo.png"></p>

# [transferme.it](https://transferme.it/)

## [Mac App](https://github.com/maxisme/transfermeit) | Website | [Backend](https://github.com/maxisme/transfermeit-backend)
![Push](https://github.com/maxisme/transferme.it/workflows/Push/badge.svg)

Create binary:
```
cd /root/
git clone https://github.com/maxisme/transferme.it
go build -o /usr/local/bin/transfermeit .
```
Create socket service
```
$ cp transfermeit.service /etc/systemd/system/
$ cp transfermeit.socket /etc/systemd/system/
```
```
$ systemctl daemon-reload
$ systemctl enable transfermeit.socket
$ systemctl start transfermeit.socket
$ systemctl status transfermeit.socket
$ curl http://127.0.0.1:8082
```