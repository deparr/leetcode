impl Solution {
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        for i in 0..matrix.len() {
            for j in i..matrix.len() {
                let temp = matrix[i][j];
                matrix[i][j] = matrix[j][i];
                matrix[j][i] = temp;
            }
        }
        for row in matrix.iter_mut() {
            row.reverse();
        }
    }
}
