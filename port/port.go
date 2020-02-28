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

const MaxRoutines = 768
const MaxTcpPort = 65535
const DefaultTimeoutSecs = 5

type ScanResult struct {
	Port        int
	Open        bool
	Protocol    string
	ServiceName string
}

func ScanTcpPort(hostname string, port int) (result ScanResult) {
	result := ScanResult{Port: port, Open: false, Protocol: "tcp", ServiceName: KnownTcpPorts[port]}

	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := net.DialTimeout(result.Protocol, address, DefaultTimeoutSecs*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	result.State = true
	return
}

func ScanUdpPort(hostname string, port int) (result ScanResult) {
	result := ScanResult{Port: port, Open: false, Protocol: "udp", ServiceName: KnownTcpPorts[port]}

	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := net.DialTimeout(result.Protocol, address, DefaultTimeoutSecs*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	result.State = true
	return
}

func scanMultipleTcpPorts(wg *sync.WaitGroup, ipAddr string, portRange []int, results []ScanResult) {
	defer wg.Done()
	for _, port := range portRange {
		results = append(results, ScanTcpPort(hostname, port))
	}
}

func ScanHost(hostname string) (results []ScanResult) {
	ports = make([]int, len(KnownTcpPorts))
	curr := 0
	for k, _ := range KnownTcpPorts {
		ports[curr] = k
		curr++
	}
	numRoutines = len(ports)

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	portsPerRoutine := len(ports) / numRoutines

	currStart := 0
	currEnd := portsPerRoutine + len(ports)%numRoutines

	for currRoutine := 0; currRoutine < numRoutines; currRoutine++ {
		go scanMultipleTcpPorts(&wg, *targetPtr, ports[currStart:currEnd], results)
		currStart = currEnd
		currEnd += portsPerRoutine
	}

	wg.Wait()
	return
}

func SweepHost(hostname string) (results []ScanResult) {
	ports = make([]int, MaxTcpPort)
	for i := 0; i < MaxTcpPort; i++ {
		ports[i] = i + 1
	}

	var wg sync.WaitGroup
	wg.Add(MaxRoutines)

	portsPerRoutine := len(ports) / MaxRoutines

	currStart := 0
	currEnd := portsPerRoutine + len(ports)%MaxRoutines

	for currRoutine := 0; currRoutine < MaxRoutines; currRoutine++ {
		go scanMultipleTcpPorts(&wg, *targetPtr, ports[currStart:currEnd], results)
		currStart = currEnd
		currEnd += portsPerRoutine
	}

	wg.Wait()
	return
}

func SweepHostRange(hostname string, rng int) (results []ScanResult) {
	ports = make([]int, rng)
	for i := 0; i < rng; i++ {
		ports[i] = i + 1
	}

	var wg sync.WaitGroup
	wg.Add(MaxRoutines)

	portsPerRoutine := len(ports) / MaxRoutines

	currStart := 0
	currEnd := portsPerRoutine + len(ports)%MaxRoutines

	for currRoutine := 0; currRoutine < MaxRoutines; currRoutine++ {
		go scanMultipleTcpPorts(&wg, *targetPtr, ports[currStart:currEnd], results)
		currStart = currEnd
		currEnd += portsPerRoutine
	}

	wg.Wait()
	return
}
