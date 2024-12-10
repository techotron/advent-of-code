use std::fs::File;
use std::env;
use std::io::{self, BufRead, BufReader, Error};
use regex::Regex;

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        eprintln!("Usage {} <filename>", args[0]);
        std::process::exit(1);
    }

    let filename = &args[1];

    // let part1_ans = part1(filename)?;
    // println!("Part1: {}", part1_ans);

    let part2_ans = part2(filename)?;
    println!("Part2: {}", part2_ans);

    Ok(())
}


fn part1(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let re = Regex::new(r"mul\(\d{1,3},\d{1,3}\)").unwrap();
    let mut calculations= Vec::new();

    for line in reader.lines() {
        let line = line?;
        for cap in re.captures_iter(&line) {
            if let Some(calculation) = cap.get(0) {
                calculations.push(calculation.as_str().to_string());
            }
        }
    }

    Ok(calculate_string(calculations))
}

fn calculate_string(calculations: Vec<String>) -> i32 {
    let mut results: Vec<i32> = Vec::new();
    for calc in calculations.iter() {
        let calc = calc.replace("mul", "");
        let calc = calc.replace("(", "");
        let calc = calc.replace(")", "");
        let parts = calc.split(",");
        let numbers = parts.collect::<Vec<&str>>();
        results.push(numbers[0].parse::<i32>().unwrap() * numbers[1].parse::<i32>().unwrap());
    }

    results.iter().sum()
}

fn part2(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let re = Regex::new(r"do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)").unwrap();
    let mut calculations= Vec::new();
    
    let mut skip = false;

    for line in reader.lines() {
        let line = line?;
        for cap in re.captures_iter(&line) {
            let matched_str = cap.get(0).unwrap().as_str();
            if matched_str == "don't()" {
                skip = true;
            } else if matched_str == "do()" {
                skip = false;
            } else if !skip {
                calculations.push(matched_str.to_string());
            }
        }
    }

    Ok(calculate_string(calculations))
}