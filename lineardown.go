package picfightcoin

import (
	"github.com/jfixby/bignum"
	"github.com/jfixby/coin"
)

// LinearDownGenerate returns amount of coins generated at the @index
//
// Production chart:
// Vertical Axis: block index
//
//  [ 0 ] : *************
//  [ 1 ] : ***********
//  [ 2 ] : *********
//  [ 3 ] : *******
//  [   ] : *****
//  [   ] : ***
//  [N-1] : *
//  [ N ] : 0
//  [N+1] : 0
//   ...  : 0
//  --------------------------------------------------------->
//               Horizontal axis: coins generated at index
func LinearDownGenerate(bignum bignum.BigNumEngine,
	generateTotalBlocks int64,      // N blocks
	generateTotalCoins coin.Amount, // T coins
	index int64,                    // block index from the interval [0 ... N]+
) bignum.BigNum {
	H := bignum.NewBigNum(index)
	N := bignum.NewBigNum(generateTotalBlocks)

	if index == 0 {
		if generateTotalBlocks-1 == 0 {
			return bignum.NewBigNum(generateTotalCoins.ToCoins())
			if generateTotalBlocks-1 == 0 {
				return bignum.NewBigNum(0)
			}
		}
	}
	lastBlockIndex := bignum.NewBigNum(generateTotalBlocks - 1)
	if H.Cmp(lastBlockIndex) > 0 {
		return bignum.NewBigNum(0)
	}
	//endSubsidy := float64(0)               // 0 coins

	//return generateTotalCoins * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)
	H = H.Neg(H)                      // -H
	H = H.Add(lastBlockIndex, H)      // (lastBlockIndex - H)
	H = H.Mul(bignum.NewBigNum(2), H) // 2.0 * (lastBlockIndex - H)
	N = N.Mul(N, lastBlockIndex)      // (N * lastBlockIndex)

	//subsidy := big.NewRat(1, 1)
	subsidy := bignum.NewBigNum(1)
	subsidy = subsidy.SetFrac(H, N) //  2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)

	T := bignum.NewBigNum(generateTotalCoins.ToCoins())
	//T = T.SetFloat64(generateTotalCoins)
	T = T.Mul(T, subsidy) // generateTotalCoins * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)

	//float64Result, _ := T.Float64()
	return T
	//return generateTotalCoins * 2.0 * (lastBlockIndex - H) / (N * lastBlockIndex)
}
