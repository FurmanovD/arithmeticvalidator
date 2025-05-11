package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateLinear_ValidExpressions(t *testing.T) {
	expressions := []string{
		"3.5 + (2 - 4.1)",
		" -3 + (-2.5) ",
		"(1.1 + (2.2 - (3.3 + 4.4)))",
		"8+767.635+495-275.115-17.231+483.737-62.046-718.404+162.371-900.815-(-580.982+696-526.634+462.047+381-(432-909+104.991-462.178-378.176-(-175-337.010-586+311.798+859-(652.268+346.714+663.331+161-617.723-(-935.993-308-824.761-744.853+179-(468-326+243-563.273+782.274-(-156-120.090-70-267.940+846.111-(191.376+407.730-881-959.969+303.097-(-255.417+5.068+306-345-989-(845+298.783+164+351.366+515+790.133-194.218-592.785+304.757-541.649+(-752+570+556.357+375-708.314-(662.690-163-24-984.448+533-(720.337-676.216-363.995-670-830-(-307.152-167-977.699+149+741-(757+222.913+204.969+223-963.091+715.217+911+796.872+901-669+493.005+358.194+442+814+728.406+795+928-465.246-650+426+(-79-845.949+452.126-66+310.272-(74+180+168+891-314-(-550-949.024+380.226+569.157-114.860+906+686-10.997-452.570-780.366-(954.275-76.006-233.620-691.386+729.626-(197-998.138-128-263.287+9.224-(450-850.474+469+922.634-795+(-532.903-391.845-971-164-939.546+(-349+565+481-362.224+493.937+430.410+752.566+595.917-145.390-827-(147.662+782.329-46-291+982+777+794+637.175+686.948-548.818-(-713.593-83+64.429+309+485.845-(-771.065-905+934+35+747-(-132-40+530+159.116+958.375-(-248.410+31.639+455-47-20+(-313.176+578.573+958.820+886+73.968-(240-340+465.574-800-343-(964.980-820-8.762-647-808)))))))))))))))))))))))))))))))",
		buildRandomExpression(4, 8),
	}
	for _, expr := range expressions {
		ok, err := ValidateLinear(expr)
		require.NoErrorf(t, err, "Expected valid: %s", expr)
		assert.True(t, ok, "Expected valid: %s", expr)
	}
}

func TestValidateLnear_InvalidExpressions(t *testing.T) {
	expressions := []string{
		"3 + + 2",
		"4..5 - 2",
		"(2 + 3",
		"abc + 1",
		"1 + 2)",
		"1 + (2 - 3))",
		"1 + (2 - 3) +",
		"1 + (2 - 3) + (4 - 5",
		"1 + (2 - 3) + (4 - 5))",
		"1 + (2 - 3) + (4 - 5)) + 6",
	}
	for _, expr := range expressions {
		ok, err := ValidateLinear(expr)
		assert.Error(t, err)
		assert.False(t, ok, "Expected invalid: %s", expr)
	}
}

func BenchmarkValidateLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ok, err := ValidateLinear(benchmarkExp)
		require.NoErrorf(b, err, "Expected valid: %s", benchmarkExp)
		require.True(b, ok, "Expected valid: %s", benchmarkExp)
	}
}

func BenchmarkValidateLinear_LargeExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ValidateLinear(benchmarkExpLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateLinear_DeepNestedExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ValidateLinear(benchmarkExpDeepNested)
		if err != nil {
			b.Fatal(err)
		}
	}
}
