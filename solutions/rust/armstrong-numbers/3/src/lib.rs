pub fn is_armstrong_number(num: u32) -> bool {
    // let digits = num.to_string();
    // let digits_list = digits.chars();
    // let mut sum = 0;

    let digits: Vec<u32> = num.to_string()
        .chars()
        .map(|c| c.to_digit(10).unwrap())
        .collect();

    let digits_len = digits.len() as u32;

    let sum: u32 = digits.iter()
        .map(|d| d.pow(digits_len))
        .sum();

    sum == num
}
