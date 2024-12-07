use std::fs::File;
use std::env;
use std::io::{self, BufRead, BufReader, Error};

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Usage {} <filename>", args[0]);
        std::process::exit(1);
    }

    let filename = &args[1];

    let part1_ans = part1(filename)?;
    println!("Part1: {}", part1_ans);

    let part2_ans = part2(filename)?;
    println!("Part2: {}", part2_ans);

    Ok(())
}


fn part1(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let mut first_column = Vec::new();
    let mut second_column = Vec::new();
    let mut difference_values = Vec::new();

    for line in reader.lines() {
        let line = line?;
        let numbers: Vec<i32> = line.split_whitespace()
                                    .map(|s| s.parse().unwrap())
                                    .collect();

        if numbers.len() == 2 {
            first_column.push(numbers[0]);
            second_column.push(numbers[1]);
        }
    }

    first_column.sort();
    second_column.sort();

    for (idx, number) in first_column.iter().enumerate() {
        let difference = (number - second_column[idx]).abs();
        difference_values.push(difference);
    }

    Ok(difference_values.iter().sum())
}


fn part2(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let mut first_column = Vec::new();
    let mut second_column = Vec::new();
    let mut multiplied_differences = Vec::new();

    for line in reader.lines() {
        let line = line?;
        let numbers: Vec<i32> = line.split_whitespace()
                                    .map(|s| s.parse().unwrap())
                                    .collect();

        if numbers.len() == 2 {
            first_column.push(numbers[0]);
            second_column.push(numbers[1]);
        }
    }

    for number in first_column.iter() {
        let times_appeared = second_column.iter().filter(|&n| *n == *number).count();
        multiplied_differences.push(number * times_appeared as i32)
    }

    Ok(multiplied_differences.iter().sum())
}