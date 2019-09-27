package picfightcoin

import (
	"fmt"
	"github.com/jfixby/bignum"
	"github.com/jfixby/coin"
	"testing"
)

func TestLinearDown(t *testing.T) {
	checkLinearDown(t, bignum.Float64Engine{}, int64(7), coin.FromFloat(5), 1, false)
	checkLinearDown(t, bignum.Float64Engine{}, int64(3), coin.FromFloat(3), 1, false)
	checkLinearDown(t, bignum.Float64Engine{}, int64(7), coin.FromFloat(7), 1, false)

	checkLinearDown(t, bignum.Float64Engine{}, int64(25), coin.FromFloat(100), 1, true)

	N := 44 * 365 * 24 * 60 / 5
	checkLinearDown(t, bignum.Float64Engine{}, int64(N), coin.FromFloat(7777777), N/10, false)
	//checkLinearDown(t, bignum.BigDecimalEngine{}, int64(N), coin.FromFloat(7777777), N/10, false)
}

func checkLinearDown(t *testing.T, engine bignum.BigNumEngine, generateTotalBlocks int64, generateTotalCoins coin.Amount, printIterations int, PrintChartData bool) {
	t.Log(fmt.Sprintf("Generating: %v in %v steps using %v",
		generateTotalCoins,
		generateTotalBlocks,
		engine,
	))

	expected := engine.NewBigNum(generateTotalCoins.ToCoins())
	generated := engine.NewBigNum(0)

	N := int(generateTotalBlocks)
	for i := 0; i < N; i++ {
		blockIndex := int64(i)
		generatedAtBlock_i := LinearDownGenerate(engine,
			generateTotalBlocks,
			generateTotalCoins,
			blockIndex,
		)
		generated = generated.Add(generated, generatedAtBlock_i)

		if int(blockIndex)%printIterations == 0 {
			blockIndexPad := fmt.Sprintf("%15v", blockIndex)
			genPad := fmt.Sprintf("%-30v", generatedAtBlock_i.ToFloat64())
			totalGenPad := fmt.Sprintf("%-30v", generated.ToFloat64())
			if PrintChartData {
				fmt.Println(fmt.Sprintf("%v	%v", i, generatedAtBlock_i.ToFloat64()))
			} else {
				t.Log(fmt.Sprintf("[%v] coins %v total %v", blockIndexPad, genPad, totalGenPad))
			}
		}
	}
	if PrintChartData {
		fmt.Println(fmt.Sprintf("total: %v", generated.ToFloat64()))
	} else {
		t.Log(fmt.Sprintf("Generated: %v coins",
			generated.ToFloat64(),
		))
	}

	if expected.ToFloat64() != generated.ToFloat64() {
		t.Fatalf("mismatched total subsidy -- \n got %v, \nwant %v",
			generated,
			expected,
		)
	}
}
