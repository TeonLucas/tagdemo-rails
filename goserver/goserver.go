package main

import (
	"os"

	"github.com/DavidSantia/tag_api"
	"github.com/DavidHuie/quartz/go/quartz"
)

type RpcFunc int

type Args struct{
        Id int
}

type Response struct {
	Images []tag_api.Image
}

func (t *RpcFunc) GetImages(args Args, response *Response) error {
	var resp Response

	for _, image := range data.ImageMap {
		resp.Images = append(resp.Images, image)
	}
	*response = resp
        tag_api.Log.Info.Printf("GetImages returned %d images", len(resp.Images))
	return nil
}

var data *tag_api.ApiData

func main() {
	var err error
        var rpcFunc RpcFunc

	quartz.RegisterName("goserver", &rpcFunc)

	// Initialize log
	var level tag_api.Level = tag_api.LogINFO
	tag_api.NewLog(level, "")

        tag_api.Log.Info.Println("Starting Go Server")
	data = tag_api.NewData()
	data.DbHost = "127.0.0.1"
	data.DbPort = "3306"
	err = data.ConnectDB()
	if err != nil {
		tag_api.Log.Error.Println(err)
		os.Exit(1)
	}

	data.LoadImages()
        tag_api.Log.Info.Println("Listening for RPC requests")
	quartz.Start()
}
