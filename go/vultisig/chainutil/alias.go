package chainutil

import "strings"

var chainAliases = map[string]string{
	// EVM — all share m/44'/60'/0'/0/0
	"arbitrum":    "Ethereum",
	"base":        "Ethereum",
	"blast":       "Ethereum",
	"optimism":    "Ethereum",
	"zksync":      "Ethereum",
	"mantle":      "Ethereum",
	"avalanche":   "Ethereum",
	"cronoschain": "Ethereum",
	"bsc":         "Ethereum",
	"polygon":     "Ethereum",
	"hyperliquid": "Ethereum",
	"sei":         "Ethereum",

	// THORChain group — m/44'/931'/0'/0/0
	"mayachain": "THORChain",

	// Cosmos-118 — m/44'/118'/0'/0/0
	"osmosis": "Cosmos",
	"dydx":    "Cosmos",
	"kujira":  "Cosmos",
	"noble":   "Cosmos",
	"akash":   "Cosmos",
	"qbtc":    "Cosmos",

	// Terra-330 — m/44'/330'/0'/0/0
	"terraclassic": "Terra",
}

func ResolveChainAlias(chain string) string {
	canonical, ok := chainAliases[strings.ToLower(chain)]
	if ok {
		return canonical
	}
	return chain
}
