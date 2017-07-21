faircoin rpcclient
==================

A thin wrapper around the FairCoin rpc calls:

* Blockchain
  * [`getblockcount`][1]
* Cvn
  * [`getactivecvns`][2]

# Usage

```
package main

import (
    faircoin "github.com/mmoya/faircoin_rpcclient"
)

func main() {
    c := faircoin.NewClient("http://127.0.0.1:40405", faircoin.CookieCredential())

    activeCvns, err := c.GetActiveCVNs()
    if err != nil {
        log.Fatal("GetActiveCVNs: ", err)
    }

    fmt.Println(activeCvns)
}
```


[1]: https://github.com/faircoin/faircoin/blob/v2.0.0/src/rpc/blockchain.cpp#L171
[2]: https://github.com/faircoin/faircoin/blob/v2.0.0/src/rpc/cvn.cpp#L35
