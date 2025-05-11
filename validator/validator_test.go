package validator

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// Benchmark expression parameters
	// These are the maximum values for the number of factors and the size of each factor
	// in the benchmark expression.cd ..
	benchmarkExpFactorSizeMax    = 5
	benchmarkExpFactorsAmountMax = 100

	// Random number generation parameters
	// These are the maximum values for the integer and decimal parts of the random numbers
	// used in the benchmark expression.
	intPartMax = 1000
	decPartMax = 1000
)

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	// Test case for bemchmark expression(initialized in init)
	benchmarkExp = ""

	benchmarkExpLarge      = buildRandomExpression(1000, 1)
	benchmarkExpDeepNested = buildRandomExpression(2, 1000)
)

func init() {
	benchmarkExpFactorSize := rnd.Intn(benchmarkExpFactorSizeMax) + 1
	benchmarkExpFactorsAmount := rnd.Intn(benchmarkExpFactorsAmountMax) + 1

	benchmarkExp = buildRandomExpression(benchmarkExpFactorSize, benchmarkExpFactorsAmount)
}

func TestValidateExpression_ValidExpressions(t *testing.T) {
	expressions := []string{
		"3.5 + (2 - 4.1)",
		" -3 + (-2.5) ",
		"(1.1 + (2.2 - (3.3 + 4.4)))",
		"8+767.635+495-275.115-17.231+483.737-62.046-718.404+162.371-900.815-(-580.982+696-526.634+462.047+381-(432-909+104.991-462.178-378.176-(-175-337.010-586+311.798+859-(652.268+346.714+663.331+161-617.723-(-935.993-308-824.761-744.853+179-(468-326+243-563.273+782.274-(-156-120.090-70-267.940+846.111-(191.376+407.730-881-959.969+303.097-(-255.417+5.068+306-345-989-(845+298.783+164+351.366+515+790.133-194.218-592.785+304.757-541.649+(-752+570+556.357+375-708.314-(662.690-163-24-984.448+533-(720.337-676.216-363.995-670-830-(-307.152-167-977.699+149+741-(757+222.913+204.969+223-963.091+715.217+911+796.872+901-669+493.005+358.194+442+814+728.406+795+928-465.246-650+426+(-79-845.949+452.126-66+310.272-(74+180+168+891-314-(-550-949.024+380.226+569.157-114.860+906+686-10.997-452.570-780.366-(954.275-76.006-233.620-691.386+729.626-(197-998.138-128-263.287+9.224-(450-850.474+469+922.634-795+(-532.903-391.845-971-164-939.546+(-349+565+481-362.224+493.937+430.410+752.566+595.917-145.390-827-(147.662+782.329-46-291+982+777+794+637.175+686.948-548.818-(-713.593-83+64.429+309+485.845-(-771.065-905+934+35+747-(-132-40+530+159.116+958.375-(-248.410+31.639+455-47-20+(-313.176+578.573+958.820+886+73.968-(240-340+465.574-800-343-(964.980-820-8.762-647-808)))))))))))))))))))))))))))))))",
		buildRandomExpression(4, 8),
	}
	for _, expr := range expressions {
		ok, err := ValidateExpression(expr)
		require.NoError(t, err)
		assert.True(t, ok, "Expected valid: %s", expr)
	}
}

func TestValidateExpression_InvalidExpressions(t *testing.T) {
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
		ok, err := ValidateExpression(expr)
		assert.Error(t, err)
		assert.False(t, ok, "Expected invalid: %s", expr)
	}
}

func BenchmarkValidateExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ok, err := ValidateExpression(benchmarkExp)
		require.NoError(b, err)
		require.True(b, ok, "Expected valid: %s", benchmarkExp)
	}
}

func BenchmarkValidateExpression_LargeExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ValidateExpression(benchmarkExpLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValidateExpression_DeepNestedExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ValidateExpression(benchmarkExpDeepNested)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Utility functions

// BuildNFactorExpression builds an expression of the form: "term1 op term2 op term3 ..."
// Each operator is randomly "+" or "-".
// The first term may optionally start with a unary "-".
func buildNFactorExpression(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	if n <= 0 {
		return ""
	}

	var b strings.Builder

	// First value with optional unary minus
	first := randomDecimal(rnd)
	if rnd.Intn(2) == 0 {
		b.WriteString("-")
	}
	b.WriteString(first)

	for i := 1; i < n; i++ {
		op := randomOp()
		val := randomDecimal(rnd)
		b.WriteString(op)
		b.WriteString(val)
	}

	return b.String()
}

// buildNestedExpression builds an expression like:
//
//	values[0] +/- (buildNestedExpression(values[1:]))  // if op == "-" or values[1] starts with "-"
//	values[0] + buildNestedExpression(values[1:])      // otherwise
func buildNestedExpression(values []string) string {
	if len(values) == 0 {
		return ""
	}
	if len(values) == 1 {
		return values[0]
	}

	var b strings.Builder
	left := values[0]
	rightExpr := buildNestedExpression(values[1:])

	op := randomOp()

	needsParens := op == "-" || strings.HasPrefix(rightExpr, "-")
	if needsParens {
		rightExpr = "(" + rightExpr + ")"
	}

	b.WriteString(left)
	b.WriteString(op)
	b.WriteString(rightExpr)

	// // Final nesting layer K times
	// result := b.String()
	// for i := 0; i < k; i++ {
	// 	result = "(" + result + ")"
	// }

	return b.String()
}

// randomOp returns either "+" or "-" randomly
func randomOp() string {
	if rand.Intn(2) == 0 {
		return "+"
	}
	return "-"
}

// randomDecimal produces a decimal like "123.456" or "78"
func randomDecimal(rnd *rand.Rand) string {
	intPart := rnd.Intn(intPartMax)
	if rnd.Intn(2) == 0 {
		// no decimal part
		return fmt.Sprintf("%d", intPart)
	}
	decPart := rnd.Intn(decPartMax)

	return fmt.Sprintf("%d.%03d", intPart, decPart)
}

func buildRandomExpression(factorSizeMax int, nestedFactorsAmountMax int) string {
	// Randomly generate the number of factors to create a benchmark expression.
	// The number of factors is between 1 and factorsAmountMax
	// and the size of each factor is between 1 and factorSizeMax.
	terms := make([]string, nestedFactorsAmountMax)
	for i := 0; i < nestedFactorsAmountMax; i++ {
		terms[i] = buildNFactorExpression(factorSizeMax)
	}

	return buildNestedExpression(terms)
}
