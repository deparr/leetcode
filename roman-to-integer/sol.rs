impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut res = 0;
        let mut prev = '_';
        for c in s.chars() {
            match (prev, c) {
                (_, 'I') => res += 1,
                ('I', 'V') => res += 3,
                (_, 'V') => res += 5,
                ('I', 'X') => res += 8,
                (_, 'X') => res += 10,
                ('X', 'L') => res += 30,
                (_, 'L') => res += 50,
                ('X', 'C') => res += 80,
                (_, 'C') => res += 100,
                ('C', 'D') => res += 300,
                (_, 'D') => res += 500,
                ('C', 'M') => res += 800,
                (_, 'M') => res += 1000,
                _ => panic!("unreachable"),
            }
            prev = c;
        }
        return res;
    }
}
