pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let mut song: String = String::from("");
    
    for i in (start_bottles+1-take_down..=start_bottles).rev() {
        let start_bottles_str: String = capitalize_first(&number_to_str(i));
        let remaining_bottles_str: String = number_to_str(i - 1);

        song += &(start_bottles_str.clone());

        if start_bottles_str == "One" {
            song += " green bottle hanging on the wall,\n";
        } else {
            song += " green bottles hanging on the wall,\n";
        };

        song += &(start_bottles_str.clone());

        if start_bottles_str == "One" {
            song += " green bottle hanging on the wall,\n";
        } else {
            song += " green bottles hanging on the wall,\n";
        };

        song += &(
            "And if one green bottle should accidentally fall,\nThere'll be ".to_owned() + &remaining_bottles_str);
        
        if remaining_bottles_str == "one" {
            song += " green bottle hanging on the wall.\n\n";
        } else {
            song += " green bottles hanging on the wall.\n\n";
        };
    }

    println!("{:?}", song);

    song
}

pub fn number_to_str(number: u32) -> String {
    match number {
        1 => "one".to_string(),
        2 => "two".to_string(),
        3 => "three".to_string(),
        4 => "four".to_string(),
        5 => "five".to_string(),
        6 => "six".to_string(),
        7 => "seven".to_string(),
        8 => "eight".to_string(),
        9 => "nine".to_string(),
        10 => "ten".to_string(),
        _ => "no".to_string(),
    }
}

pub fn capitalize_first(s: &str) -> String {
    let mut chars = s.chars();
    match chars.next() {
        None => String::new(),
        Some(f) => f.to_uppercase().collect::<String>() + chars.as_str(),
    }
}