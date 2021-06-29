![AutoMouse](https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/microsoft/209/mouse-face_1f42d.png)
<img src="https://img.shields.io/badge/language-go-green?logo=go"/> <img src="https://img.shields.io/badge/Version-0.5-green">

# Auto Mouse
보안프로그램을 이용하다보면, 강제로 N분 이후에 이벤트가 없으면 절전모드로 변환되는 경우가 발생하는데 이를 방지하기 위해 5분동안 키, 
마우스 이벤트가 발생하지 않으면 강제로 (스크롤 이벤트를) 발생시켜주는 프로그램이다

When using a security program, if there is no keyboard or mouse event after N minutes by force, it may switch to power saving mode.
To prevent this, press the key for 5 minutes, This is a program that forcibly generates (scroll event) when no mouse event occurs.

## Run Go script (test)
```
go run mouse.go
```

## Add Icon
```
rsrc -manifest ./resource/mouse.exe.manifest -ico ./resource/mouse.ico 
go build
```

## Build Go script
```
go build mouse.go
```
