package main

import (
	"os"

	"github.com/DavidHuie/quartz/go/quartz"
)

type RpcFunc int

type Args struct {
	Id int
}

type Response struct {
	Images []Image
}

func (t *RpcFunc) GetImages(args Args, response *Response) error {
	var resp Response

	for _, image := range ds.ImageMap {
		resp.Images = append(resp.Images, image)
	}
	*response = resp
	Log.Info.Printf("GetImages returned %d images", len(resp.Images))
	return nil
}

var ds *DbService

const (
	// Default Bolt DB file
	BoltDB = "./content.db"

	// Retries to wait for docker DB instance
	DbConnectRetries = 5

	// MySQL DB info
	DbUser = "demo"
	DbPass = "welcome1"
	DbName = "tagdemo"
	DbHost = "127.0.0.1"
	DbPort = "3306"

	// NATS server
	NSub = "update"
)

func main() {
	var err error
	var rpcFunc RpcFunc

	quartz.RegisterName("goserver", &rpcFunc)

	// Initialize log
	var level Level = LogINFO
	NewLog(level, "")
	Log.Info.Println("Starting Go Server")

	// Load images from DB
	ds = NewDbService(DbUser, DbPass, DbName, DbHost, DbPort)
	err = ds.Connect()
	if err != nil {
		Log.Error.Println(err)
		os.Exit(1)
	}
	ds.loadImages(nil)

	Log.Info.Println("Listening for RPC requests")
	quartz.Start()
}
