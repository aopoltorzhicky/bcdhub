package bcd

import (
	"github.com/pkg/errors"
)

// This is the list of protocols BCD supports
// Every time new protocol is proposed we determine if everything works fine or implement a custom handler otherwise
// After that we append protocol to this list with a corresponding handler id (aka symlink)
var symLinks = map[string]string{
	"ProtoGenesisGenesisGenesisGenesisGenesisGenesk612im": SymLinkAlpha,
	"PrihK96nBAFSxVL1GLJTVhu9YnzkMFiBeuJRPA8NwuZVZCE1L6i": SymLinkAlpha,
	"PtBMwNZT94N7gXKw4i273CKcSaBrrBnqnt3RATExNKr9KNX2USV": SymLinkAlpha,
	"ProtoDemoNoopsDemoNoopsDemoNoopsDemoNoopsDemo6XBoYp": SymLinkAlpha,
	"PtYuensgYBb3G3x1hLLbCmcav8ue8Kyd2khADcL5LsT5R1hcXex": SymLinkAlpha,
	"Ps9mPmXaRzmzk35gbAYNCAw6UXdE2qoABTHbN2oEEc1qM7CwT9P": SymLinkAlpha,
	"PsYLVpVvgbLhAhoqAkMFUo6gudkJ9weNXhUYCiLDzcUpFpkk8Wt": SymLinkAlpha,
	"PsddFKi32cMJ2qPjf43Qv5GDWLDPZb3T3bF6fLKiF5HtvHNU7aP": SymLinkAlpha,
	"Pt24m4xiPbLDhVgVfABUjirbmda3yohdN82Sp9FeuAXJ4eV9otd": SymLinkAlpha,
	"PtCJ7pwoxe8JasnHY8YonnLYjcVHmhiARPJvqcC6VfHT5s8k8sY": SymLinkAlpha,
	"PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS": SymLinkBabylon,
	"PsBABY5HQTSkA4297zNHfsZNKtxULfL18y95qb3m53QJiXGmrbU": SymLinkBabylon,
	"PsCARTHAGazKbHtnKfLzQg3kms52kSRpgnDY982a9oYsSXRLQEb": SymLinkBabylon, // Carthagenet
	"PryLyZ8A11FXDr1tRE9zQ7Di6Y8zX48RfFCFpkjC8Pt9yCBLhtN": SymLinkBabylon, // Dalphanet
	"PsDELPH1Kxsxt8f9eWbxQeRxkjfbxoqM52jvs5Y5fBxWWh4ifpo": SymLinkBabylon, // Delphinet
	"PtEdoTezd3RHSC31mpxxo1npxFjoWWcFgQtxapi51Z8TLu6v6Uq": SymLinkBabylon, // Edonet 8.1
	"PtEdo2ZkT9oKpimTah6x2embF25oss54njMuPzkJTEi5RqfdZFA": SymLinkBabylon, // Edonet 8.2
	"PrrUA9dCzbqBzugjQyw65HLHKjhH3HMFSLLHLZjj5rkmkG13Fej": SymLinkBabylon, // Falphanet
	"PsrsRVg1Gycjn5LvMtoYSQah1znvYmGp8bHLxwYLBZaYFf2CEkV": SymLinkBabylon, // Falphanet
	"PsFLorenaUUuikDWvMDr6fGBRG8kt3e3D3fHoXK1j1BFRxeSH4i": SymLinkBabylon, // Florencenet (no baking accounts)
	"PtGRANADsDU8R9daYKAgWnQYAJ64omN1o3KMGVCykShA97vQbvV": SymLinkBabylon, // Granadanet
	"PtHangzHogokSuiMHemCuowEavgYTP8J5qQ9fQS793MHYFpCY3r": SymLinkBabylon, // Hangzhounet
	"PtHangz2aRngywmSRGGvrcTyMbbdpWdpFKuS4uMWxg2RaH9i1qx": SymLinkBabylon, // Hangzhounet 2
	"PsiThaCaT47Zboaw71QWScM8sXeMM7bbQFncK9FLqYc6EKdpjVP": SymLinkBabylon, // Itacanet
	"Psithaca2MLRFYargivpo7YvUr7wUDqyxrdhC5CQq78mRvimz6A": SymLinkBabylon, // Itacanet 2
}

// GetProtoSymLink -
func GetProtoSymLink(protocol string) (string, error) {
	if protoSymLink, ok := symLinks[protocol]; ok {
		return protoSymLink, nil
	}
	return "", errors.Errorf("Unknown protocol: %s", protocol)
}

// GetCurrentSymLink -
func GetCurrentSymLink() string {
	return SymLinkBabylon
}

// GetCurrentProtocol - returns last supported protocol
func GetCurrentProtocol() string {
	return "PtHangz2aRngywmSRGGvrcTyMbbdpWdpFKuS4uMWxg2RaH9i1qx"
}

// Symbolic links
const (
	SymLinkAlpha   = "alpha"
	SymLinkBabylon = "babylon"
)

var ChainID = map[string]string{
	"NetXdQprcVkpaWU": "mainnet",
	"NetXZSsxBpMQeAT": "hangzhou2net",
	"NetXnHfVqm9iesp": "ithacanet",
}
