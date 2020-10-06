start-udp-server:
	nc -l -k 7000 > /dev/null

start-tcp-server:
	nc -l -u 7001 > /dev/null

benchmark:
	go test -benchmem -bench .