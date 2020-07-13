package kinds

// Fa1_2 -
type Fa1_2 struct{}

// GetName -
func (fa12 Fa1_2) GetName() string {
	return "fa12"
}

// GetJSON -
func (fa12 Fa1_2) GetJSON() string {
	return `
	[
	{
		"name": "approve",
		"prim": "pair",
		"args": [
		{
			"prim": "address"
		},
		{
			"prim": "nat"
		}
		]
	},
	{
		"name": "getAllowance",
		"prim": "pair",
		"args": [
		{
			"args": [
			{
				"prim": "address"
			},
			{
				"prim": "address"
			}
			],
			"prim": "pair"
		},
		{
			"parameter": {
			"prim": "nat"
			},
			"prim": "contract"
		}
		]
	},
	{
		"name": "getBalance",
		"prim": "pair",
		"args": [
		{
			"prim": "address"
		},
		{
			"parameter": {
			"prim": "nat"
			},
			"prim": "contract"
		}
		]
	},
	{
		"name": "getTotalSupply",
		"prim": "pair",
		"args": [
		{
			"prim": "unit"
		},
		{
			"parameter": {
			"prim": "nat"
			},
			"prim": "contract"
		}
		]
	},
	{
		"name": "transfer",
		"prim": "pair",
		"args": [
		{
			"prim": "address"
		},
		{
			"args": [
			{
				"prim": "address"
			},
			{
				"prim": "nat"
			}
			],
			"prim": "pair"
		}
		]
	}
	]`
}
