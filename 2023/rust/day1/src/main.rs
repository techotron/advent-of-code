use std::fs::File;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    // Open the file in read-only mode
    let file_path = "input";
    let file = File::open(file_path)?;
    let reader = io::BufReader::new(file);

    // Initialize the sum
    let mut sum = 0;

    // Iterate over lines in the calibration document
    for line in reader.lines() {
        // Handle each line
        if let Ok(content) = line {
            // Extract the first and last digits
            if let (Some(first_digit), Some(last_digit)) = (content.chars().next(), content.chars().last()) {
                // Convert characters to digits
                if let (Some(first), Some(last)) = (first_digit.to_digit(10), last_digit.to_digit(10)) {
                    // Add the two digits to the sum
                    sum += first * 10 + last;
                }
            }
        }
    }

    // Print the sum of all calibration values
    println!("Sum of calibration values: {}", sum);

    Ok(())
}
