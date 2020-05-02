# raspi-gobot
go get github.com/stianeikeland/go-rpio
go get -d -u gobot.io/x/gobot/...

# build project for raspi 3 (GARM=7)
env GOOS=linux GOARCH=arm GOARM=7 go build

# deploy to raspi 3
scp buzzer_robot pi@192.168.1.226:ledblink
