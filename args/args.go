package args

import "flag"

var (
	Port = flag.Int("port", 8080, "Port number to run the server on (default: 8080)")
	Dir  = flag.String("directory", ".", "Directory path to serve files (default: current directory)")
)

func Init() {
	flag.Parse()
}
