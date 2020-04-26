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

____
Write custom .env first then run:
```
docker stack deploy transfermeit -c transfermeit.yml 
```

