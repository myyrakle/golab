# golab

실험소

## memoryeather

OOM 관련 실험

기본 실행 (OOM으로 터짐)

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 memoryeater
```

GOGC 조정 (의미없음)

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 -e GOGC=50 memoryeater
```
