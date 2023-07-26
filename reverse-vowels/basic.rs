impl Solution {
    pub fn reverse_words(s: String) -> String {
        let mut res = String::with_capacity(s.len());
        for val in s.split_ascii_whitespace().rev() {
            res.push_str(val);
            res.push(' ');
        }
        res.pop();
        return res;
    }
}

