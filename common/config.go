package common

var Userdict = map[string][]string{
	"ftp":        {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data"},
	"mysql":      {"root", "mysql"},
	"mssql":      {"sa", "sql"},
	"smb":        {"administrator", "admin", "guest"},
	"rdp":        {"administrator", "admin", "guest"},
	"postgresql": {"postgres", "admin"},
	"ssh":        {"root", "admin"},
	"mongodb":    {"root", "admin"},
	"oracle":     {"sys", "system", "admin", "test", "web", "orcl"},
}

var Passwords = []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "{user}@123#4", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e", "Charge123", "Aa123456789"}
var PORTList = map[string]int{
	"ftp":         21,
	"ssh":         22,
	"findnet":     135,
	"netbios":     139,
	"smb":         445,
	"mssql":       1433,
	"oracle":      1521,
	"mysql":       3306,
	"rdp":         3389,
	"psql":        5432,
	"redis":       6379,
	"fcgi":        9000,
	"mem":         11211,
	"mgo":         27017,
	"ms17010":     1000001,
	"cve20200796": 1000002,
	"web":         1000003,
	"webonly":     1000003,
	"webpoc":      1000003,
	//"smb2":        1000004,
	//"wmiexec":     1000005,
	"all":      0,
	"portscan": 0,
	"icmp":     0,
	"main":     0,
}
var PortGroup = map[string]string{
	"ftp":         "21",
	"ssh":         "22",
	"findnet":     "135",
	"netbios":     "139",
	"smb":         "445",
	"mssql":       "1433",
	"oracle":      "1521",
	"mysql":       "3306",
	"rdp":         "3389",
	"psql":        "5432",
	"redis":       "6379",
	"fcgi":        "9000",
	"mem":         "11211",
	"mgo":         "27017",
	"ms17010":     "445",
	"cve20200796": "445",
	"service":     "21,22,135,139,445,1433,1521,3306,3389,5432,6379,9000,11211,27017",
	"db":          "1433,1521,3306,5432,6379,11211,27017",
	"web":         "80,81,82,83,84,85,86,87,88,89,90,91,92,98,99,443,800,801,808,880,888,889,1000,1010,1080,1081,1082,1099,1118,1888,2008,2020,2100,2375,2379,3000,3008,3128,3505,5555,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8011,8012,8016,8018,8020,8028,8030,8038,8042,8044,8046,8048,8053,8060,8069,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8108,8118,8161,8172,8180,8181,8200,8222,8244,8258,8280,8288,8300,8360,8443,8448,8484,8800,8834,8838,8848,8858,8868,8879,8880,8881,8888,8899,8983,8989,9000,9001,9002,9008,9010,9043,9060,9080,9081,9082,9083,9084,9085,9086,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9443,9448,9800,9981,9986,9988,9998,9999,10000,10001,10002,10004,10008,10010,10250,12018,12443,14000,16080,18000,18001,18002,18004,18008,18080,18082,18088,18090,18098,19001,20000,20720,21000,21501,21502,28018,20880",
	"all":         "1-65535",
	"main":        "21,22,80,81,135,139,443,445,1433,1521,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017",
}
var Outputfile = "result.txt"
var IsSave = true
var Webport = "80,81,82,83,84,85,86,87,88,89,90,91,92,98,99,443,800,801,808,880,888,889,1000,1010,1080,1081,1082,1099,1118,1888,2008,2020,2100,2375,2379,3000,3008,3128,3505,5555,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8011,8012,8016,8018,8020,8028,8030,8038,8042,8044,8046,8048,8053,8060,8069,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8108,8118,8161,8172,8180,8181,8200,8222,8244,8258,8280,8288,8300,8360,8443,8448,8484,8800,8834,8838,8848,8858,8868,8879,8880,8881,8888,8899,8983,8989,9000,9001,9002,9008,9010,9043,9060,9080,9081,9082,9083,9084,9085,9086,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9443,9448,9800,9981,9986,9988,9998,9999,10000,10001,10002,10004,10008,10010,10250,12018,12443,14000,16080,18000,18001,18002,18004,18008,18080,18082,18088,18090,18098,19001,20000,20720,21000,21501,21502,28018,20880"
var DefaultPorts = "21,22,80,81,135,139,443,445,1433,1521,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017"

type HostInfo struct {
	Host    string
	Ports   string
	Url     string
	Infostr []string
}

type PocInfo struct {
	Target  string
	PocName string
}

var (
	Ports       string
	Path        string
	Scantype    string
	Command     string
	SshKey      string
	Domain      string
	Username    string
	Password    string
	Proxy       string
	Timeout     int64 = 3
	WebTimeout  int64 = 5
	TmpSave     bool
	NoPing      bool
	Ping        bool
	Pocinfo     PocInfo
	NoPoc       bool
	IsBrute     bool
	RedisFile   string
	RedisShell  string
	Userfile    string
	Passfile    string
	HostFile    string
	PortFile    string
	PocPath     string
	Threads     int
	URL         string
	UrlFile     string
	Urls        []string
	NoPorts     string
	NoHosts     string
	SC          string
	PortAdd     string
	UserAdd     string
	PassAdd     string
	BruteThread int
	LiveTop     int
	//Socks5Proxy string
	Hash string
	//HashBytes   []byte
	HostPort    []string
	IsWmi       bool
	Noredistest bool
)

var (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
	Accept    = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
	DnsLog    bool
	PocNum    int
	PocFull   bool
	Cookie    string
)

/*
用户名和密码字典
Userdict: 定义了针对不同服务（如FTP、MySQL、MSSQL等）的默认用户名列表。这些用户名在尝试进行弱口令或暴力破解攻击时会被使用。
Passwords: 包含了一系列常见的弱密码，同样用于暴力破解尝试。{user}占位符表示将尝试使用用户名作为密码的一部分。
端口列表和分组
PORTList: 映射了服务名到其默认端口的关系，例如MySQL默认为3306端口。这对于针对特定服务的扫描行为非常重要。
PortGroup: 类似于PORTList，但是它以字符串形式提供了一些服务组的端口列表，便于扫描时指定多个端口。
扫描配置
Outputfile, IsSave, Webport: 分别定义了输出文件名、是否保存扫描结果以及用于Web服务扫描的端口列表。
HostInfo结构体: 包含了单个扫描目标的信息，例如主机名、端口、URL和其他特定信息。
PocInfo结构体: 定义了用于漏洞验证的目标信息和PoC名称。
扫描参数
Ports, Path, Scantype, Command等: 定义了用于扫描的端口、路径、扫描类型、执行命令等参数。
Timeout, WebTimeout: 分别指定了扫描和Web服务扫描的超时时间。
Threads, BruteThread: 定义了用于扫描的线程数和暴力破解时的线程数。
其他配置
UserAgent, Accept, Cookie: 用于HTTP请求时的用户代理、接受类型和Cookie。
DnsLog, PocNum, PocFull: 控制DNS日志记录、PoC验证数量和是否进行完整PoC验证的标志。
功能和作用
这段代码的主要功能是作为漏洞扫描工具的配置中心，提供了针对不同服务的用户名和密码字典、服务端口映射、扫描参数和其他相关设置。这样做的目的是为了使扫描工具能够灵活地针对不同的目标和服务执行扫描和漏洞验证操作。

实现的功能包括但不限于：

支持对多种服务进行弱口令检测。
允许用户自定义扫描参数，如端口、路径和扫描类型。
提供了一个结构化的方式来配置和执行漏洞扫描和验证。
通过这样的配置，扫描工具可以根据用户的需求和目标环境的特点灵活地进行扫描，帮助发现和修复潜在的安全漏洞。

*/
