package etcd

import (
    "fmt"

    "golang.org/x/net/context"
    "github.com/coreos/etcd/client"
)

func Watch() {

    // need to make these command line args
    cfg := client.Config{
    	Endpoints: []string{
            "http://192.168.99.100:2379",
        },
    }

    c, err := client.New(cfg)
    if err != nil {
    	panic(err.Error())
    }

    kAPI := client.NewKeysAPI(c)

    watcherOpts := client.WatcherOptions{AfterIndex: 0, Recursive: true}

    watcher := kAPI.Watcher("/namespaces", &watcherOpts)

    for {
        fmt.Println("waiting for etcd event......")
        next, err := watcher.Next(context.Background())
        if err != nil {
            panic(err.Error())
        }

        fmt.Println(fmt.Sprintf("something happened: %s", next.Action))
    }

}
