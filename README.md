# doscg
## scg assignment

### Test
```
go test -v ./...
```

### Build
```
cd cmd/doscg-app
go build
```

### Run
```
cd cmd/doscg-app

# set up env
GOOGLE_KEY=your_google_api_key
LINE_NOTI_TOKEN=your_noti_token
LINE_CHAN_SECRET=your_line_message_api_secret
LINE_CHAN_TOKEN=your_line_message_api_token
PORT=8080

# actual run
./doscg-app
# or
go run main.go

```
