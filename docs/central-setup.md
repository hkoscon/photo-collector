# Central Server Setup Guide

## 1. Download Minio
Minio is a simple S3 compatible server.
Download the latest release of minio.
https://docs.minio.io/docs/minio-quickstart-guide

Make sure the download binary is runnable.

## 2. Config Minio
```bash
export MINIO_ACCESS_KEY=[admin]
export MINIO_SECRET_KEY=[password]
```
The word [admin] and [password] is API Access Key pair and should be replaced.

## 3. Mount the storing hard disk
Please mount the disk to the server

## 4. Start minio
```bash
./minio server /data
```
Replace `/data` to the root path of the data destination

## Remark
Make sure the collector are able to access 9000 port via TCP

It is a good idea to have multiple disk for data store
```bash
./minio server [file paths...]
``` 

e.g.
```bash
./minio server /disk1 /disk2
```