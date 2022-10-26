func evalRPN(tokens []string) int {
    operands := make([]int, len(tokens))
    idx := 0
    var ( op1 int; op2 int )
    for _, token := range tokens {
        val, err := strconv.Atoi(token)
        // fmt.Println(token, idx, operands)
        if err != nil {
            switch token[0] {
                case '+':
                op1 = operands[idx-1]
                op2 = operands[idx-2]
                idx--
                operands[idx-1]= (op2 + op1)
                case '-':
                op1 = operands[idx-1]
                op2 = operands[idx-2]
                idx--
                operands[idx-1]= (op2 - op1)
                case '*':
                op1 = operands[idx-1]
                op2 = operands[idx-2]
                idx--
                operands[idx-1]= (op2 * op1)
                case '/':
                op1 = operands[idx-1]
                op2 = operands[idx-2]
                idx--
                operands[idx-1]= (op2 / op1)
            }
        } else {
            operands[idx] = val
            idx++
        }
    }
    return operands[0]
}
