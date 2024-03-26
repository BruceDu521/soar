## Get Released Binary

```bash
wget https://github.com/BruceDu521/soar/releases/download/${tag}/soar.${OS}-amd64 -O soar
chmod a+x soar
eg.
wget https://github.com/BruceDu521/soar/releases/download/0.9.0/soar.linux-amd64 -O soar
chmod a+x soar
```

## Build From Source

```bash
go get -d github.com/BruceDu521/soar
cd ${GOPATH}/src/github.com/BruceDu521/soar && make
```

## Simple Test Case

```bash
echo 'select * from film' | ./soar
```
