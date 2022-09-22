vector<int> plusOne(vector<int>& digits) {
    for (size_t i = digits.size()-1; i < digits.size(); i--) {
        if (digits[i] != 9) {
            digits[i] += 1;
            break;
        } else {
            digits[i] = 0;
            if (i == 0) {
                digits.insert(digits.begin(), 1);
            }
        }
    }
    return digits;
}
