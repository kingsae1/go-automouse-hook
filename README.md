# go.automover.hook
회사 프로그램 보안프로그램을 이용하다보면, 강제로 N분 이후에 이벤트가 없으면
절전모드로 변환되는 경우가 발생하는데 이를 방지하기 위해 5분동안 키, 마우스 이벤트가
발생하지 않으면 강제로 발생시켜주는 프로그램이다

## Run Go script (test)
```
go run mouse.go
```

## Build Go script
```
rsrc -manifest ./resource/mouse.exe.manifest -ico ./resource/mouse.ico 
go build mouse.go
```