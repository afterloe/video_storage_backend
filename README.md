# video_storage_backend

```
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

```bash
ffprobe -select_streams v \
-show_entries format=duration,size,bit_rate,filename \
-show_streams \
-v quiet \
-of csv="p=0" \
-of json \
-i ./dde-introduction.mp4
```