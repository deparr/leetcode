string addStrings(string num1, string num2) {
    char carry = 0;
    string res = "";
    int xi = (int)num1.size()-1, yi = (int)num2.size()-1;
    while (xi > -1 || yi > -1) {
        char xc = xi > -1 ? num1[xi] & 0xf : 0x00;
        char yc = yi > -1 ? num2[yi] & 0xf : 0x00;
        char total = xc + yc + carry;
        carry = 0;
        if (total >= 0x0a) {
            carry = 0x1;
            total = 0x30 | (total - 0x0a);
        } else {
            total = 0x30 | total;
        }
        res.insert(res.begin(), total);
        xi--; yi--;
    }
    
    if (carry > 0) {
        res.insert(res.begin(), 0x31);
    }
    
    return res;
}