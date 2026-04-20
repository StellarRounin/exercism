pub fn is_armstrong_number(num: u32) -> bool {
    let digits = num.to_string();
    let digits_list = digits.chars();
    let mut sum = 0;

    for digit in digits_list {
        sum += digit.to_digit(10).unwrap().pow(digits.len() as u32);
    }

    if sum == num { true } else { false }
}
