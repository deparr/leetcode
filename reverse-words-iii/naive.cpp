class Solution {
public:
    string reverseWords(string s) {
        int begin = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != ' ' && i != s.size()-1)
                continue;
            int end = i == s.size()-1 ? i : i-1;
            while (end > begin) {
                char tmp = s[end];
                s[end--] = s[begin];
                s[begin++] = tmp;
            }
            begin = i + 1;
        }
        
        return s;
    }
};
