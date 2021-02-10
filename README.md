

[Unit]
Description=goapi

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/arh/go/src/github.com/adhityarachmanh/go-api-arh/app

[Install]
WantedBy=multi-user.target

{"db":"cv","uri":""}
{"db":"cv","uri":"mongodb://localhost:27017"}

