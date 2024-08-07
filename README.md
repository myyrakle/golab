# golab

실험소

## memoryeather

OOM 관련 실험

1. 기본 실행 (OOM으로 터짐)

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 memoryeater
```

2. GOGC 조정 (효과는 있음)

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 -e GOGC=50 memoryeater
```

3. GOMEMLIMIT 조정 (효과있음)

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 -e GOMEMLIMIT=1000MiB memoryeater
```

4. 직접 Free (효과있음)

```go
    // 직접 해제
    debug.FreeOSMemory()
```

```bash
docker build -f docker/Dockerfile.memory -t memoryeater .
docker run -m 2048m --memory-swap 0 --memory-reservation 0 memoryeater
```
