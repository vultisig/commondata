package chainutil

import "testing"

func TestResolveChainAlias_EVMChains(t *testing.T) {
	evmAliases := []string{
		"Arbitrum", "Base", "Blast", "Optimism", "Zksync",
		"Mantle", "Avalanche", "CronosChain", "BSC", "Polygon",
		"Hyperliquid", "Sei",
	}
	for _, chain := range evmAliases {
		result := ResolveChainAlias(chain)
		if result != "Ethereum" {
			t.Errorf("ResolveChainAlias(%q) = %q, want %q", chain, result, "Ethereum")
		}
	}
}

func TestResolveChainAlias_THORChainGroup(t *testing.T) {
	result := ResolveChainAlias("MayaChain")
	if result != "THORChain" {
		t.Errorf("ResolveChainAlias(MayaChain) = %q, want THORChain", result)
	}
}

func TestResolveChainAlias_CosmosGroup(t *testing.T) {
	cosmosAliases := []string{"Osmosis", "Dydx", "Kujira", "Noble", "Akash", "QBTC"}
	for _, chain := range cosmosAliases {
		result := ResolveChainAlias(chain)
		if result != "Cosmos" {
			t.Errorf("ResolveChainAlias(%q) = %q, want %q", chain, result, "Cosmos")
		}
	}
}

func TestResolveChainAlias_TerraGroup(t *testing.T) {
	result := ResolveChainAlias("TerraClassic")
	if result != "Terra" {
		t.Errorf("ResolveChainAlias(TerraClassic) = %q, want Terra", result)
	}
}

func TestResolveChainAlias_CanonicalReturnsSelf(t *testing.T) {
	canonicals := []string{"Ethereum", "THORChain", "Cosmos", "Terra", "Bitcoin", "Solana"}
	for _, chain := range canonicals {
		result := ResolveChainAlias(chain)
		if result != chain {
			t.Errorf("ResolveChainAlias(%q) = %q, want %q", chain, result, chain)
		}
	}
}

func TestResolveChainAlias_UnknownChainPassthrough(t *testing.T) {
	unknowns := []string{"Bitcoin", "Litecoin", "Dogecoin", "Solana", "Ripple", "SomeNewChain"}
	for _, chain := range unknowns {
		result := ResolveChainAlias(chain)
		if result != chain {
			t.Errorf("ResolveChainAlias(%q) = %q, want %q", chain, result, chain)
		}
	}
}

func TestResolveChainAlias_CaseInsensitive(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"polygon", "Ethereum"},
		{"POLYGON", "Ethereum"},
		{"Polygon", "Ethereum"},
		{"mayachain", "THORChain"},
		{"MAYACHAIN", "THORChain"},
		{"osmosis", "Cosmos"},
		{"OSMOSIS", "Cosmos"},
		{"terraclassic", "Terra"},
		{"TERRACLASSIC", "Terra"},
	}
	for _, tc := range cases {
		result := ResolveChainAlias(tc.input)
		if result != tc.want {
			t.Errorf("ResolveChainAlias(%q) = %q, want %q", tc.input, result, tc.want)
		}
	}
}
