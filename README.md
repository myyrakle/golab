# golab

실험소

## memoryeather

OOM 관련 실험

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 memoryeater
```
