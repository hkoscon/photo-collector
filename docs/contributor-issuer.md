# Contributor Issuer Guide

A GUI is recommended

## 1. Download Golang

Please download Golang 1.10.x version (Remark: other version might work but I didn't test on them).
Extract the tarball somewhere and make the bin directory in the $PATH

## 2. Config GOPATH

By default, if `$GOPATH` is not set, `$HOME/go` is used.
You may pick another directory and set it at `$GOPATH`.
And run these commands

```bash
mkdir -p $GOPATH/src/hkoscon.org $GOPATH/bin $GOPATH/pkg
export PATH=$GOPATH/bin:$PATH
```

## 3. Download dep
Download the dep binary from [dep release page](https://github.com/golang/dep/releases)
And put it in `$GOPATH/bin` as `dep`

## 4. Clone the source
```bash
cd $GOPATH/src/hkoscon.org
git clone https://github.com/hkoscon/photo-collector photo
```
You may checkout the release tag if you preferred

## 5. Build it
```bash
dep ensure -v -vendor-only
go build -o [/yout/path/for/issuer] ./cmd/issuer/main.go
```
Change the binary path `/yout/path/for/issuer` for your choice

## 6. Config Issuer
```bash
export PUBLIC_KEY_PATH=/path/to/key
export KEY_LABEL=foobar
```

## 7. Generate QR Code
```bash
./issuer "The July Jasmin" /path/to/qrcode.png
```

