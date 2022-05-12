# sshserver



# build

```

mkdir ./pluginBin
go build -o pluginBin/    sshserver.go 
go build -o pluginBin/    controller/controller.go 
```


```
ls pluginBin/
controller	sshserver
```


install plugin


```
git clone https://github.com/yimtun/cf.git
cd cf
go build cf.go
cp cf pluginBin dir
```

