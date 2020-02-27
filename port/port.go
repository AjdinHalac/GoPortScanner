package port

import (
	"net"
	"strconv"
	"time"
)

var KnownTcpPorts = map[int]string{
	7:     "Echo",
	21:    "FTP",
	22:    "SSH",
	23:    "telnet",
	25:    "SMTP",
	53:    "DNS",
	66:    "Oracle SQL*NET?",
	69:    "tftp",
	80:    "HTTP",
	88:    "kerberos",
	110:   "POP3",
	123:   "NTP",
	137:   "netbios",
	139:   "netbios",
	194:   "IRC",
	118:   "SQL service?",
	150:   "SQL-net?",
	443:   "HTTP w/TLS",
	445:   "Samba",
	554:   "RTSP",
	631:   "CUPS",
	1433:  "Microsoft SQL server",
	1434:  "Miocrosoft SQL monitor",
	1883:  "MQTT",
	3306:  "MySQL/MariaDB ",
	3535:  "SMTP (alternate)",
	5000:  "Heroku, Docker, UPnP, Flask",
	5672:  "RabbitMQ",
	5800:  "VNC remote desktop",
	6000:  "lixo",
	8080:  "HTTP",
	9160:  "Cassandra [ http://cassandra.apache.org/ ]",
	15672: "RabbitMQ Management console",
	27017: "mongodb [ http://www.mongodb.org/ ]",
	28017: "mongodb web admin [ http://www.mongodb.org/ ]",
}

const MaxTcpPort = 32728

type ScanResult struct {
	Port     int
	Open     bool
	Protocol string
}

func ScanTcpPort(hostname string, port int) (result ScanResult) {
	result := ScanResult{Port: port, Open: false, Protocol: "tcp"}

	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := net.DialTimeout(result.Protocol, address, 60*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	result.State = true
	return
}

func ScanUdpPort(hostname string, port int) (result ScanResult) {
	result := ScanResult{Port: port, Open: false, Protocol: "udp"}

	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := net.DialTimeout(result.Protocol, address, 60*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	result.State = true
	return
}

func ScanHost(hostname string, lastPort int) (results []ScanResult) {
	for i := 1; i <= lastPort; i++ {
		results = append(results, ScanTcpPort(hostname, i))
	}
	return
}

func ScanHost(hostname string) (results []ScanResult) {
	for i := 1; i <= MaxTcpPort; i++ {
		results = append(results, ScanTcpPort(hostname, i))
	}
	return
}
