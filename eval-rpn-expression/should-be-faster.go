func add(op1 int, op2 int) int { return op1 + op2 }
func sub(op1 int, op2 int) int { return op1 - op2 }
func mul(op1 int, op2 int) int { return op1 * op2 }
func div(op1 int, op2 int) int { return op1 / op2 }

func evalRPN(tokens []string) int {
    operands := make([]int, len(tokens))
    idx := 0
    var ( op1 int; op2 int; op func(int,int)int )
    for i := range tokens {
        val, err := strconv.ParseInt(tokens[i], 10, 32)
        if err != nil {
            switch tokens[i] {
                case `+`:
                op = add
                case `-`:
                op = sub
                case `*`:
                op = mul
                case `/`:
                op = div
            }
            op1 = operands[idx-1]
            op2 = operands[idx-2]
            idx--
            operands[idx-1] = op(op2, op1)
        } else {
            operands[idx] = int(val)
            idx++
        }
    }
    return operands[0]
}
