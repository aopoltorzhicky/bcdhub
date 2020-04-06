package unpack

import (
	"testing"
)

func TestPulickKey(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "ed25519 public key",
			input:  "004e4ca2abb4baeed702a0ac5b0de9b5607dd1fedb399c0ce25e15b3868f67269e",
			result: "edpkuEhzJqdFBCWMw6TU3deADRK2fq3GuwWFUphwyH7ero1Na4oGFP",
		},
		{
			name:   "secp256k1 public key",
			input:  "01030ed412d33412ab4b71df0aaba07df7ddd2a44eb55c87bf81868ba09a358bc0e0",
			result: "sppk7bMuoa8w2LSKz3XEuPsKx1WavsMLCWgbWG9CZNAsJg9eTmkXRPd",
		},
		{
			name:   "p256 public key",
			input:  "02031a3ad5ea94de6912f9bc83fd31de49816e90602c5252d77b5b233bfe711b0dd2",
			result: "p2pk66iTZwLmRPshQgUr2HE3RUzSFwAN5MNaBQ5rfduT1dGKXd25pNN",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := PublicKey(tt.input)
			if err != nil {
				t.Errorf("Error in PublicKey. Input: %v. Error: %v", tt.input, err)
			}
			if res != tt.result {
				t.Errorf("Error in PublicKey. Input: %v. Got: %v. Expected: %v.", tt.input, res, tt.result)
			}
		})
	}
}

func TestKeyHash(t *testing.T) {
	tests := []struct {
		input  string
		result string
	}{
		{
			input:  "0010fc2282886d9cf8a1eebdc2733e302c7b110f38",
			result: "tz1MBqYpcoGU93c1bePp5A6dmwKYjmHXRopS",
		},
		{
			input:  "003c8c2fe0f75ce212558df94c7a7306c2eeadd979",
			result: "tz1RABAzdLWVvxAFf1wpeUALAkp32mVhSGXX",
		},
		{
			input:  "004bf0acca4cc9e034b1d5f0f783c78e5ed44d866e",
			result: "tz1SZZgtvMVXaBKPcez4gfjKUsDz1gs6vg6X",
		},
		{
			input:  "0079e68d8f0a8d64ec856e193efc0a347ef4adf8ee",
			result: "tz1WkaeRycRr999GrVFepJd9Nqi1FWqGyGqq",
		},
		{
			input:  "01028562fb176188114cf437a757cdc75bc4aa8cae",
			result: "tz28YZoayJjVz2bRgGeVjxE8NonMiJ3r2Wdu",
		},
		{
			input:  "029d6a61cd3510193e257128da8f09a0b173bff695",
			result: "tz3agP9LGe2cXmKQyYn6T68BHKjjktDbbSWX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res, err := KeyHash(tt.input)

			if err != nil {
				t.Errorf("Error in KeyHash. Error: %v", err)
			}

			if res != tt.result {
				t.Errorf("Error in Keyhash. Got: %v. Expected: %v", res, tt.result)
			}
		})
	}
}

func TestAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		result  string
	}{
		{
			name:    "tz1 address",
			address: "00009e6ac2e529a49aedbcdd0ac9542d5c0f4ce76f77",
			result:  `tz1a5fMLLY5WCarCzH7RKTJHX9mJFN8eaaWG`,
		},
		{
			name:    "tz1 address",
			address: "000010fc2282886d9cf8a1eebdc2733e302c7b110f38",
			result:  "tz1MBqYpcoGU93c1bePp5A6dmwKYjmHXRopS",
		},
		{
			name:    "tz1 address",
			address: "000006a868bd80219eb1f6a25108d1bdaa98ae27b2d9",
			result:  "tz1LFEVYR7YRCxT6Nm3Zfjdnfj77xZqhbR5U",
		},
		{
			name:    "tz1 address",
			address: "00003c8c2fe0f75ce212558df94c7a7306c2eeadd979",
			result:  "tz1RABAzdLWVvxAFf1wpeUALAkp32mVhSGXX",
		},
		{
			name:    "tz1 address",
			address: "00004bf0acca4cc9e034b1d5f0f783c78e5ed44d866e",
			result:  "tz1SZZgtvMVXaBKPcez4gfjKUsDz1gs6vg6X",
		},
		{
			name:    "tz1 address",
			address: "000079e68d8f0a8d64ec856e193efc0a347ef4adf8ee",
			result:  "tz1WkaeRycRr999GrVFepJd9Nqi1FWqGyGqq",
		},
		{
			name:    "tz1 address",
			address: "0000d27fcbd31910d2226ba4c8f646d3d4c7b2f3a756",
			result:  "tz1eq3gqb2iZHjHVHoPJqV84gZdBF2TMQiH4",
		},
		{
			name:    "tz1 address",
			address: "0000107c4009f2bcfcc248d6952998af5b7203b8ff59",
			result:  "tz1M9CMEtsXm3QxA7FmMU2Qh7xzsuGXVbcDr",
		},
		{
			name:    "tz2 address",
			address: "0001028562fb176188114cf437a757cdc75bc4aa8cae",
			result:  "tz28YZoayJjVz2bRgGeVjxE8NonMiJ3r2Wdu",
		},
		{
			name:    "tz3 address",
			address: "00029d6a61cd3510193e257128da8f09a0b173bff695",
			result:  "tz3agP9LGe2cXmKQyYn6T68BHKjjktDbbSWX",
		},
		{
			name:    "KT address",
			address: "0168b709e887ddc34c3c9e468b5819b2f012b60ef700",
			result:  "KT1J8T7U6J1BAo9fJAxvedHsNErnejwvPyUH",
		},
		{
			name:    "KT address",
			address: "011fb03e3ff9fedaf3a2200ffc64d27812da734bba00",
			result:  "KT1BUKeJTemAaVBfRz6cqxeUBQGQqMxfG19A",
		},
		{
			name:    "KT address",
			address: "01d50e3f6f059dc86f5591455549313ce42d0c50f100",
			result:  "KT1U1JZaXoG4u1EPnhHL4R4otzkWc1L34q3c",
		},
		{
			name:    "KT address",
			address: "01f8f6c6a0af7c20251bc7df108f2a6e2879a06c9a00",
			result:  "KT1XHAmdRKugP1Q38CxDmpcRSxq143KpEiYx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.address, func(t *testing.T) {
			res, err := Address(tt.address)
			if err != nil {
				t.Errorf("Error in Address. Error: %v", err)
			}
			if res != tt.result {
				t.Errorf("Error in Address. Got %v, Expected: %v", res, tt.result)
			}
		})
	}
}

func TestSignature(t *testing.T) {
	validTestCases := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "sig",
			input:  "bdc36db614aaa6084549020d376bb2469b5ea888dca2f7afbe5a0095bcc45ca0d8b5f00a051969437fe092debbcfe19d66378fbb74104de7eb1ecd895a64a80a",
			result: `signpEFVQ1rW3TnVhc3PXf6SHRj7PvxwfJhBukWfB5X9rDhzpEk3ms5gRh763e922n52uQcjeqhqPdYi7WbFs2ERrNAPmCZJ`,
		},
		{
			name:   "sig",
			input:  "a04991b4e938cc42d6c01c42be3649a81a9f80d244d9b90e7ec4edf8e0a7b68b6c212da2fef076e48fed66802fa83442b960a36afdb3e60c3cf14d4010f41f03",
			result: `sigixZejtj1GfDpyiWAQAmvbtnNmCXKyADqVvCaXJH9xHyhSnYYV8696Z3kkns5DNV7oMnMPfNzo3qm84DfEx1XG6saZmHiA`,
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Signature(tc.input)
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}
			if res != tc.result {
				t.Errorf("\nInput: %v. \nGot: %v, \nexpected: %v.", tc.input, res, tc.result)
			}
		})
	}
}

func TestChainID(t *testing.T) {
	validTestCases := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "chainID/main",
			input:  "7a06a770",
			result: `NetXdQprcVkpaWU`,
		},
		{
			name:   "chainID/babylon",
			input:  "458aa837",
			result: `NetXUdfLh6Gm88t`,
		},
		{
			name:   "chainID/carthage",
			input:  "9caecab9",
			result: `NetXjD3HPJJjmcd`,
		},
		{
			name:   "chainID/zeronet",
			input:  "0f6f0310",
			result: `NetXKakFj1A7ouL`,
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ChainID(tc.input)
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}
			if res != tc.result {
				t.Errorf("\nInput: %v. \nGot: %v, \nexpected: %v.", tc.input, res, tc.result)
			}
		})
	}
}

func TestContract(t *testing.T) {
	validTestCases := []struct {
		name   string
		input  string
		result string
	}{
		{
			name:   "contract/tz3",
			input:  "0002358cbffa97149631cfb999fa47f0035fb1ea8636",
			result: `tz3RDC3Jdn4j15J7bBHZd29EUee9gVB1CxD9`,
		},
		{
			name:   "contract/KT",
			input:  "016f516588d2ee560385e386708a13bd63da907cf300",
			result: `KT1JjN5bTE9yayzYHiBm6ruktwEWSHRF8aDm`,
		},
		{
			name:   "contract/KT%",
			input:  "01e5bae183211979a662665319a0900df3542e65ba00646f",
			result: `KT1VXUBQbYMt58yoKhNo73Zf8HTMfAd8Fqge%do`,
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Contract(tc.input)
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}
			if res != tc.result {
				t.Errorf("\nInput: %v. \nGot: %v, \nexpected: %v.", tc.input, res, tc.result)
			}
		})
	}
}

func TestBytes(t *testing.T) {
	validTestCases := []struct {
		name    string
		input   string
		jsonstr string
		result  string
	}{
		{
			name:    "Small int",
			input:   "050006",
			jsonstr: `{ "int": "6" }`,
			result:  `6`,
		},
		{
			name:    "Negative small int",
			input:   "050046",
			jsonstr: `{ "int": "-6" }`,
			result:  `-6`,
		},
		{
			name:    "Medium int",
			input:   "0500840e",
			jsonstr: `{ "int": "900" }`,
			result:  `900`,
		},
		{
			name:    "Negative medium int",
			input:   "0500c40e",
			jsonstr: `{ "int": "-900" }`,
			result:  `-900`,
		},
		{
			name:    "Large int",
			input:   "0500ba9af7ea06",
			jsonstr: `{ "int": "917431994" }`,
			result:  `917431994`,
		},
		{
			name:    "Negative large int",
			input:   "0500c0f9b9d4c723",
			jsonstr: `{ "int": "-610913435200" }`,
			result:  `-610913435200`,
		},
		{
			name:    "String",
			input:   "0501000000096d696368656c696e65",
			jsonstr: `{ "string": "micheline" }`,
			result:  `"micheline"`,
		},
		{
			name:    "Empty string",
			input:   "050100000000",
			jsonstr: `{ "string": "" }`,
			result:  `""`,
		},
		{
			name:    "Bytes",
			input:   "050a000000080123456789abcdef",
			jsonstr: `{ "bytes": "0123456789abcdef" }`,
			result:  "0x0123456789abcdef",
		},
		{
			name:    "Mixed literal array",
			input:   "0502000000210061010000000574657a6f730100000000010000000b63727970746f6e6f6d6963",
			jsonstr: `[ { "int": "-33" }, { "string": "tezos" }, { "string": "" }, { "string": "cryptonomic" } ]`,
			result:  `{ -33 ; "tezos" ; "" ; "cryptonomic" }`,
		},
		{
			name:    "Single primitive",
			input:   "050343",
			jsonstr: `{ "prim": "PUSH" }`,
			result:  "PUSH",
		},
		{
			name:    "Single primitive with a single annotation",
			input:   "0504430000000440636261",
			jsonstr: `{ "prim": "PUSH", "annots": [ "@cba" ] }`,
			result:  "PUSH @cba",
		},
		{
			name:    "Single primitive with a single argument",
			input:   "05053d036d",
			jsonstr: `{ "prim": "NIL", "args": [ { "prim": "operation" } ] }`,
			result:  "NIL operation",
		},
		{
			name:    "Single primitive with a single argument and annotation",
			input:   "05063d036d0000000440636261",
			jsonstr: `{ "prim": "NIL", "args": [ { "prim": "operation" } ], "annots": [ "@cba" ] }`,
			result:  "NIL @cba operation",
		},
		{
			name:    "Single primitive with two arguments",
			input:   "05073d036d036d",
			jsonstr: `{ "prim": "NIL", "args": [ { "prim": "operation" }, { "prim": "operation" } ] }`,
			result:  "NIL operation operation",
		},
		{
			name:    "Single primitive with two arguments and annotation",
			input:   "05083d036d036d0000000440636261",
			jsonstr: `{ "prim": "NIL", "args": [ { "prim": "operation" }, { "prim": "operation" } ], "annots": [ "@cba" ] }`,
			result:  "NIL @cba operation operation",
		},
		{
			name:    "Single primitive with more than two arguments and no annotations",
			input:   "05093d00000006036d036d036d00000000",
			jsonstr: `{ "prim": "NIL", "args": [ { "prim": "operation" }, { "prim": "operation" }, { "prim": "operation" } ] }`,
			result:  "NIL operation operation operation",
		},
		{
			name:    "Single primitive with more than two arguments and multiple annotations",
			input:   "05093d00000006036d036d036d00000011407265642040677265656e2040626c7565",
			jsonstr: `{ "prim": "NIL", "args": [ { "prim": "operation" }, { "prim": "operation" }, { "prim": "operation" } ], "annots": [ "@red", "@green", "@blue" ] }`,
			result:  "NIL @red @green @blue operation operation operation",
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "05010000000561646d696e",
			jsonstr: `{ "string": "admin" }`,
			result:  `"admin"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4 decode address from bytes #1",
			input:   "050a00000016000006a868bd80219eb1f6a25108d1bdaa98ae27b2d9",
			jsonstr: `{ "bytes": "000006a868bd80219eb1f6a25108d1bdaa98ae27b2d9" }`,
			result:  `"tz1LFEVYR7YRCxT6Nm3Zfjdnfj77xZqhbR5U"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050100000006706175736564",
			jsonstr: `{ "string": "paused" }`,
			result:  `"paused"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050303",
			jsonstr: `{ "prim": "False" }`,
			result:  "False",
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "05010000000866616c6c6261636b",
			jsonstr: `{ "string": "fallback" }`,
			result:  `"fallback"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "0502000000270316031607430368010000001655706172616d4e6f53756368456e747279506f696e7403420327",
			jsonstr: `[ { "prim": "CAR" }, { "prim": "CAR" }, { "prim": "PUSH", "args": [ { "prim": "string" }, { "string": "UparamNoSuchEntryPoint" } ] }, { "prim": "PAIR" }, { "prim": "FAILWITH" } ]`,
			result:  `{ CAR ; CAR ; PUSH string "UparamNoSuchEntryPoint" ; PAIR ; FAILWITH }`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "0501000000086e65774f776e6572",
			jsonstr: `{ "string": "newOwner" }`,
			result:  `"newOwner"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050306",
			jsonstr: `{ "prim": "None" }`,
			result:  "None",
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "0501000000096f70657261746f7273",
			jsonstr: `{ "string": "operators" }`,
			result:  `"operators"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050200000000",
			jsonstr: `[]`,
			result:  `{}`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050100000009746f6b656e636f6465",
			jsonstr: `{ "string": "tokencode" }`,
			result:  `"tokencode"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050100000005545a425443",
			jsonstr: `{ "string": "TZBTC" }`,
			result:  `"TZBTC"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050100000009746f6b656e6e616d65",
			jsonstr: `{ "string": "tokenname" }`,
			result:  `"tokenname"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "05010000000b746f74616c4275726e6564",
			jsonstr: `{ "string": "totalBurned" }`,
			result:  `"totalBurned"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "050000",
			jsonstr: `{ "int": "0" }`,
			result:  "0",
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "05010000000b746f74616c4d696e746564",
			jsonstr: `{ "string": "totalMinted" }`,
			result:  `"totalMinted"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "05010000000b746f74616c537570706c79",
			jsonstr: `{ "string": "totalSupply" }`,
			result:  `"totalSupply"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4",
			input:   "05010000000d72656465656d41646472657373",
			jsonstr: `{ "string": "redeemAddress" }`,
			result:  `"redeemAddress"`,
		},
		{
			name:    "KT1FgscaMyhxoVLbVirJVVKpRXgiSGtDG9Z4 decode address from bytes #2",
			input:   "050a00000016000006a868bd80219eb1f6a25108d1bdaa98ae27b2d9",
			jsonstr: `{ "bytes": "000006a868bd80219eb1f6a25108d1bdaa98ae27b2d9" }`,
			result:  `"tz1LFEVYR7YRCxT6Nm3Zfjdnfj77xZqhbR5U"`,
		},
		{
			name:    "decode keyhash from bytes",
			input:   "050a000000150044c6f8bc6088cd3b64f0bca87f812634c3f0ed30",
			jsonstr: ``,
			result:  `"tz1RugzxKA8NwuymbGcy2wkSTvfRJpckfmDF"`,
		},
		// {
		// 	name:    "if it`s",
		// 	input:   "0502000003f103210316051f020000000203170321051f02000001d10321051f02000000f0051f02000000020321051f020000000203160743036200020342030c0329072f0200000004053e0362020000005c072e020000001207430368010000000742616443746f720327020000003c072e020000001e072e0200000000020000001207430368010000000742616443746f720327020000001207430368010000000742616443746f7203270346072f02000000060743036200000200000000051f020000000a03210317031703170316034c034b0356072f0200000031074303680100000026496e7465726e616c3a204e65676174697665206e6577206469766964656e6420706f696e747303270200000000034c051f0200000088051f02000000020321051f020000000203160743036200000342030c0329072f020000000c053e076503620760036e0362020000003e072e020000001e072e0200000000020000001207430368010000000742616443746f720327020000001207430368010000000742616443746f7203270346072f020000000607430362000002000000020316033a07430362009a9c01034c0322072f020000002b074303680100000020496e7465726e616c3a204e6567617469766520707265636973696f6e4d756c74032702000000020316034c032107430362000003190325072c02000000040320032002000001ec034c0321051f0200000148034c051f02000000020321034c051f02000000a3051f020000009c051f02000000020321051f020000000203160743036200030342030c0329072f0200000004053e0362020000005c072e020000001207430368010000000742616443746f720327020000003c072e020000001e072e020000001207430368010000000742616443746f7203270200000000020000001207430368010000000742616443746f7203270346072f02000000060743036200000200000000051f020000000403120346051f020000000b051f0200000004032103160743036200030342030c051f0200000042072f020000001c053e07640764076503620760036e0362036c07640764036203620369020000001a054403620533036905440764076503620760036e0362036c03460350051f020000000d0321051f020000000203170316051f020000000203200342051f020000000c032103170317031703160346051f020000000b051f0200000004032103160743036200020342030c051f0200000042072f020000001c053e07640764076503620760036e0362036c07640764036203620369020000001a053303620533036905440764076503620760036e0362036c03460350051f020000000d0321051f020000000203170316051f020000000203200342",
		// 	jsonstr: ``,
		// 	result:  "",
		// },
		// {
		// 	name:    "something",
		// 	input:   "05020000037103210316051f02000000020317050d0765036e0362072f0200000029034f07430368010000001a55706172616d417267756d656e74556e7061636b4661696c6564034203270200000000034203210316051f0200000002031703480342051f02000000b40321074303690a0000000c0501000000067061757365640329072f02000000220743036801000000175553746f72653a206e6f206669656c642070617573656403270200000000050d0359072f020000002a07430368010000001f5553746f72653a206661696c656420746f20756e7061636b2070617573656403270200000000072c0200000027034f074303680100000018546f6b656e4f7065726174696f6e73417265506175736564034203270200000000051f02000000020321034c051f02000000020321034c0321051f020000008703160743036801000000066c65646765720342030c0329072f020000000c053e076503620760036e03620200000044050d076503620760036e0362072f020000002a07430368010000001f5553746f72653a206661696c656420746f20756e7061636b206c6564676572032702000000000346072f02000000060723036e036202000000020317031703160329072f02000000060743036200000200000000032103300325072c020000000203200200000043051f02000000020321034c0317031703300325072c020000000203200200000022074303680100000015556e73616665416c6c6f77616e63654368616e676503420327051f02000000020321034c051f020000000403210316034c0743036801000000066c65646765720342030c0329072f020000000c053e076503620760036e03620200000044050d076503620760036e0362072f020000002a07430368010000001f5553746f72653a206661696c656420746f20756e7061636b206c6564676572032702000000000346072f02000000140723036e036207430362000003420723036e0362020000000403210317071f0002020000000203210570000203170317032103300325072c02000000060320053e036202000000020346071f00030200000002032105700003031703160350051f020000000d0321051f020000000203160317034c0320034c0342034c03160743036801000000066c65646765720342030c051f0200000004030c03460350053d036d034203210316051f020000000203170342",
		// 	jsonstr: ``,
		// 	result:  "",
		// },
		// {
		// 	name:    "strange",
		// 	input:   "05020000055403210316051f02000000020317050d0362072f0200000029034f07430368010000001a55706172616d417267756d656e74556e7061636b4661696c6564034203270200000000034203210316051f02000000020317051f02000000c20321074303690a0000000f0501000000096f70657261746f72730329072f020000002507430368010000001a5553746f72653a206e6f206669656c64206f70657261746f727303270200000000050d0566036e072f020000002d0743036801000000225553746f72653a206661696c656420746f20756e7061636b206f70657261746f72730327020000000003480339072c02000000000200000026074307650368036c0707010000001353656e64657249734e6f744f70657261746f72030b0327051f02000000960321074303690a0000001305010000000d72656465656d416464726573730329072f020000002907430368010000001e5553746f72653a206e6f206669656c642072656465656d4164647265737303270200000000050d036e072f02000000310743036801000000265553746f72653a206661696c656420746f20756e7061636b2072656465656d4164647265737303270200000000034203210316051f020000000203170321051f02000002b3034c0342051f02000000020321034c051f02000000020321034c03160743036801000000066c65646765720342030c0329072f020000000c053e076503620760036e03620200000044050d076503620760036e0362072f020000002a07430368010000001f5553746f72653a206661696c656420746f20756e7061636b206c6564676572032702000000000346072f02000000290317074303620000034c03420743036801000000104e6f74456e6f75676842616c616e636503420327020000000003210316071f000202000000020321057000020317034c034b0356072f020000002e0316051f02000000020321034c031703420743036801000000104e6f74456e6f75676842616c616e6365034203270200000000051f020000000d0321051f020000000203170316034c03200342051f02000000020321034c051f02000000700321031603300325072c020000002603210317034503300325072c020000000e0320053e076503620760036e03620200000002034602000000020346034c03160743036801000000066c65646765720342030c051f0200000014072f0200000004053e03690200000004030c034603500321051f02000000f50317033b051f02000000900321074303690a0000001105010000000b746f74616c537570706c790329072f020000002707430368010000001c5553746f72653a206e6f206669656c6420746f74616c537570706c7903270200000000050d0362072f020000002f0743036801000000245553746f72653a206661696c656420746f20756e7061636b20746f74616c537570706c790327020000000003120356072f020000002a07430368010000001f496e7465726e616c3a204e6567617469766520746f74616c20737570706c7903270200000000030c0346074303690a0000001105010000000b746f74616c537570706c7903500320051f02000000900321074303690a0000001105010000000b746f74616c4275726e65640329072f020000002707430368010000001c5553746f72653a206e6f206669656c6420746f74616c4275726e656403270200000000050d0362072f020000002f0743036801000000245553746f72653a206661696c656420746f20756e7061636b20746f74616c4275726e6564032702000000000312030c0346074303690a0000001105010000000b746f74616c4275726e65640350053d036d034203210316051f020000000203170342",
		// 	jsonstr: ``,
		// 	result:  "",
		// },
		{
			name:    "Test ascii string",
			input:   "68747470733a2f2f6d6f6e6579657665727933646179732e636f6d2f6d6f6e6579657665727933646179732e6a736f6e",
			jsonstr: ``,
			result:  "https://moneyevery3days.com/moneyevery3days.json",
		},
		{
			name:    "Test non-ascii string",
			input:   "ff68747470733a2f2f6d6f6e6579657665727933646179732e636f6d2f6d6f6e6579657665727933646179732e6a736f6e",
			jsonstr: ``,
			result:  "ff68747470733a2f2f6d6f6e6579657665727933646179732e636f6d2f6d6f6e6579657665727933646179732e6a736f6e",
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Bytes(tc.input)
			if res != tc.result {
				t.Errorf("\nInput: %v. \nGot: %v\nexpected: %v.", tc.input, res, tc.result)
			}
		})
	}
}
