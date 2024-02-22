package common

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func Parse(Info *HostInfo) {
	ParseUser()
	ParsePass(Info)
	ParseInput(Info)
	ParseScantype(Info)
}

func ParseUser() {
	if Username == "" && Userfile == "" {
		return
	}
	var Usernames []string
	if Username != "" {
		Usernames = strings.Split(Username, ",")
	}

	if Userfile != "" {
		users, err := Readfile(Userfile)
		if err == nil {
			for _, user := range users {
				if user != "" {
					Usernames = append(Usernames, user)
				}
			}
		}
	}

	Usernames = RemoveDuplicate(Usernames)
	for name := range Userdict {
		Userdict[name] = Usernames
	}
}

func ParsePass(Info *HostInfo) {
	var PwdList []string
	if Password != "" {
		passs := strings.Split(Password, ",")
		for _, pass := range passs {
			if pass != "" {
				PwdList = append(PwdList, pass)
			}
		}
		Passwords = PwdList
	}
	if Passfile != "" {
		passs, err := Readfile(Passfile)
		if err == nil {
			for _, pass := range passs {
				if pass != "" {
					PwdList = append(PwdList, pass)
				}
			}
			Passwords = PwdList
		}
	}
	if URL != "" {
		urls := strings.Split(URL, ",")
		TmpUrls := make(map[string]struct{})
		for _, url := range urls {
			if _, ok := TmpUrls[url]; !ok {
				TmpUrls[url] = struct{}{}
				if url != "" {
					Urls = append(Urls, url)
				}
			}
		}
	}
	if UrlFile != "" {
		urls, err := Readfile(UrlFile)
		if err == nil {
			TmpUrls := make(map[string]struct{})
			for _, url := range urls {
				if _, ok := TmpUrls[url]; !ok {
					TmpUrls[url] = struct{}{}
					if url != "" {
						Urls = append(Urls, url)
					}
				}
			}
		}
	}
	if PortFile != "" {
		ports, err := Readfile(PortFile)
		if err == nil {
			newport := ""
			for _, port := range ports {
				if port != "" {
					newport += port + ","
				}
			}
			Ports = newport
		}
	}
}

func Readfile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open %s error, %v\n", filename, err)
		os.Exit(0)
	}
	defer file.Close()
	var content []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			content = append(content, scanner.Text())
		}
	}
	return content, nil
}

func ParseInput(Info *HostInfo) {
	if Info.Host == "" && HostFile == "" && URL == "" && UrlFile == "" {
		fmt.Println("Host is none")
		flag.Usage()
		os.Exit(0)
	}

	if BruteThread <= 0 {
		BruteThread = 1
	}

	if TmpSave == true {
		IsSave = false
	}

	if Ports == DefaultPorts {
		Ports += "," + Webport
	}

	if PortAdd != "" {
		if strings.HasSuffix(Ports, ",") {
			Ports += PortAdd
		} else {
			Ports += "," + PortAdd
		}
	}

	if UserAdd != "" {
		user := strings.Split(UserAdd, ",")
		for a := range Userdict {
			Userdict[a] = append(Userdict[a], user...)
			Userdict[a] = RemoveDuplicate(Userdict[a])
		}
	}

	if PassAdd != "" {
		pass := strings.Split(PassAdd, ",")
		Passwords = append(Passwords, pass...)
		Passwords = RemoveDuplicate(Passwords)
	}

}

func ParseScantype(Info *HostInfo) {
	_, ok := PORTList[Scantype]
	if !ok {
		showmode()
	}
	if Scantype != "all" && Ports == DefaultPorts+","+Webport {
		switch Scantype {
		case "wmiexec":
			Ports = "135"
		case "wmiinfo":
			Ports = "135"
		case "smbinfo":
			Ports = "445"
		case "hostname":
			Ports = "135,137,139,445"
		case "smb2":
			Ports = "445"
		case "web":
			Ports = Webport
		case "webonly":
			Ports = Webport
		case "ms17010":
			Ports = "445"
		case "cve20200796":
			Ports = "445"
		case "portscan":
			Ports = DefaultPorts + "," + Webport
		case "main":
			Ports = DefaultPorts
		default:
			port, _ := PORTList[Scantype]
			Ports = strconv.Itoa(port)
		}
		fmt.Println("-m ", Scantype, " start scan the port:", Ports)
	}
}

func CheckErr(text string, err error, flag bool) {
	if err != nil {
		fmt.Println("Parse", text, "error: ", err.Error())
		if flag {
			if err != ParseIPErr {
				fmt.Println(ParseIPErr)
			}
			os.Exit(0)
		}
	}
}

func showmode() {
	fmt.Println("The specified scan type does not exist")
	fmt.Println("-m")
	for name := range PORTList {
		fmt.Println("   [" + name + "]")
	}
	os.Exit(0)
}

// WrapperTcpWithTimeout 创建一个带有超时设置的TCP连接
func WrapperTcpWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	// 使用net.DialTimeout来建立一个带有超时的连接
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
