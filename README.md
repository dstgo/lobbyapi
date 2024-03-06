# lobby
**lobby** is a http client of dst lobby server, help you more convenient to get information from lobby.
> Note: The HTTP API maybe changed by klei official in the future

## install
```bash
go get -u github.com/dstgo/lobby
```

## use
Here are some examples as follows

### regions
get all the available regions
```go
package main

import (
	"fmt"
	"github.com/dstgo/lobby"
)

func main() {
	lobbyclient := lobby.Open("no need klei token")
	capableRegions, err := lobbyclient.GetCapableRegions()
	if err != nil {
		panic(err)
	}
	fmt.Println(capableRegions)
}
```
output
```bash
{[{us-east-1} {eu-central-1} {ap-southeast-1} {ap-east-1}]}
```

### servers
get server list with region and platform, available platforms as follows:
* Steam
* PSN
* Rail
* XBone
* Switch

here is an example
```go
package main

import (
	"fmt"
	"github.com/dstgo/lobby"
)

func main() {
	lobbyclient := lobby.Open("no need klei token")
	list, err := lobbyclient.GetLobbyServers("ap-east-1","Steam")
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
```

### details
get details for specific server
> must be used with klei token
```go
package main

import (
	"fmt"
	"github.com/dstgo/lobby"
)

func main() {
	lobbyclient := lobby.Open("klei token")
	details, err := lobbyclient.GetServerDetails("ap-east-1","xxxxx")
	if err != nil {
		panic(err)
	}
	fmt.Println(details)
}
```

