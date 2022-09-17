int lengthOfLastWord(string s) {
    int count = 0;
    for (int i = s.size()-1; i > -1; i--) {
        if (s[i] != ' ')
            count++;
        else if (count > 0)
            break;
    }
    return count;
}
