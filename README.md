![AutoMover](https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/microsoft/209/mouse-face_1f42d.png)
<img src="https://img.shields.io/badge/language-go-green?logo=go"/> <img src="https://img.shields.io/badge/Version-0.5-green">

# Auto Mover

보안프로그램을 이용하다보면, 강제로 N분 이후에 이벤트가 없으면 절전모드로 변환되는 경우가 발생하는데 이를 방지하기 위해 5분동안 키,
마우스 이벤트가 발생하지 않으면 강제로 (스크롤 이벤트를) 발생시켜주는 프로그램이다

## Config (go.automover.config)

앱의 상세설정은 go.automover.config 파일을 통해 설정한다. (구분자는 |)

| Index | Name         | Description                    | Default |
| ----- | ------------ | ------------------------------ | ------- |
| 0     | Version      | Automover Version              | 0.1     |
| 1     | TickTimeout  | Ticker Timeout                 | 10      |
| 2     | TickMaxCount | Count Ticker during not moving | 5 Min   |

## Run Go script (test)

```
go run go.automover.go
```

## Add Icon

```
rsrc -manifest ./resource/automover.exe.manifest -ico ./resource/mouse.ico
go build
```

## Build Go script

```
go build go.automover.go
```
