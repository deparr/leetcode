impl Solution {
    pub fn reverse_words(s: String) -> String {
        String::from(s.split_ascii_whitespace().rev().collect::<Vec<&str>>().join(" "))
    }
}

