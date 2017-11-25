holdemhandrank - Texas Holdem hand rank evaluator for Go
===========================

# holdemhandrank

Package `holdemhandrank` is a [Go](http://golang.org) client library providing ranking of 7 cards hand.

The Two Plus Two alogorithm and lookup table are used. 

The lookup table HandRanks.dat is included in the module which should be placed into the root folder of the app.


# Installation

## Development

```
go get -v github.com/mahinko/holdemhandrank/
```

# Usage
```go
rank := holdemhandrank.Eval([]string{ "ac", "kc", "qc", "jc", "tc", "2h", "3s" })

//{ HandType: 9,
//  HandRank: 10,
//  Value: 36874,
//  HandName: 'Straight Flush' }
```

# Links
The logic was taken from the following project - https://github.com/chenosaurus/poker-evaluator
