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

const MaxTcpPort = 65535
const DefaultTimeoutSecs = 5

type ScanResult struct {
	Port        int
	Open        bool
	Protocol    string
	ServiceName string
}

func ScanTcpPort(hostname string, port int) (result ScanResult) {
	result := ScanResult{Port: port, Open: false, Protocol: "tcp"}

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
	result := ScanResult{Port: port, Open: false, Protocol: "udp"}

	address := net.JoinHostPort(hostname, strconv.Itoa(port))
	conn, err := net.DialTimeout(result.Protocol, address, DefaultTimeoutSecs*time.Second)

	if err != nil {
		return
	}
	defer conn.Close()

	result.State = true
	return
}

func scanMultipleTcpPorts(wg *sync.WaitGroup, ipAddr string, portRange []int) {
	defer wg.Done()
	for _, port := range portRange {
		isOpen := scanTcpPort(ipAddr, port)

		serviceName := KnownTcpPorts[port]

		if isOpen && serviceName != "" {
			fmt.Println("Target is likely running", KnownTcpPorts[port], "on port", port)
		} else if isOpen && serviceName == "" {
			fmt.Println("Port", port, "is open, but no match for known services")
		}
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
		go scanMultipleTcpPorts(&wg, *targetPtr, ports[currStart:currEnd])
		currStart = currEnd
		currEnd += portsPerRoutine
	}

	wg.Wait()
	for i := 1; i <= lastPort; i++ {
		results = append(results, ScanTcpPort(hostname, i))
	}
	return
}

func SweepHost(hostname string) (results []ScanResult) {
	numRoutines = MaxOpenFileDescriptors
	ports = make([]int, MaxTcpPort)
	for i := 0; i < MaxTcpPort; i++ {
		ports[i] = i + 1
	}

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	portsPerRoutine := len(ports) / numRoutines

	currStart := 0
	currEnd := portsPerRoutine + len(ports)%numRoutines

	for currRoutine := 0; currRoutine < numRoutines; currRoutine++ {
		go scanMultipleTcpPorts(&wg, *targetPtr, ports[currStart:currEnd])
		currStart = currEnd
		currEnd += portsPerRoutine
	}

	wg.Wait()
	for i := 1; i <= MaxTcpPort; i++ {
		results = append(results, ScanTcpPort(hostname, i))
	}
	return
}

func SweepHostRange(hostname string, rng int) (results []ScanResult) {
	numRoutines = MaxOpenFileDescriptors
	ports = make([]int, rng)
	for i := 0; i < MaxTcpPort; i++ {
		ports[i] = i + 1
	}

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	portsPerRoutine := len(ports) / numRoutines

	currStart := 0
	currEnd := portsPerRoutine + len(ports)%numRoutines

	for currRoutine := 0; currRoutine < numRoutines; currRoutine++ {
		go scanMultipleTcpPorts(&wg, *targetPtr, ports[currStart:currEnd])
		currStart = currEnd
		currEnd += portsPerRoutine
	}

	wg.Wait()
	for i := 1; i <= MaxTcpPort; i++ {
		results = append(results, ScanTcpPort(hostname, i))
	}
	return
}
