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
    let mut all_levels = Vec::new();
    let mut safe_levels = Vec::new();

    for line in reader.lines() {
        let line = line?;

        let level: Vec<i32> = line.split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        all_levels.push(level);
    }

    for level in all_levels.iter() {
        if is_safe(level) {
            safe_levels.push(level);
        }
    }

    Ok(safe_levels.len() as i32)
}

fn is_safe(level: &Vec<i32>) -> bool {
    if level.len() <= 2 {
        return true;
    }
    let mut direction = level[1] - level[0];
    if direction.abs() < 1 || direction.abs() > 3 {
        return false;
    }
    for i in 2..level.len() {
        let diff = level[i] - level[i-1];
        if diff.abs() < 1 || diff.abs() > 3 {
            return false;
        }
        if (direction > 0 && diff < 0) || (direction < 0 && diff > 0) {
            return false;
        }
        direction = diff;
    }
    true
}

fn part2(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);
    let mut all_levels = Vec::new();
    let mut safe_levels = Vec::new();

    for line in reader.lines() {
        let line = line?;

        let level: Vec<i32> = line.split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        all_levels.push(level);
    }

    for level in all_levels.iter() {
        if is_safe2(level) {
            safe_levels.push(level);
        }
    }

    Ok(safe_levels.len() as i32)
}

fn is_safe2(level: &Vec<i32>) -> bool {
    if level.len() <= 2 {
        return true;
    }

    if is_safe(level) {
        return true;
    }

    for i in 0..level.len() {
        let mut modified_level = level.clone();
        modified_level.remove(i);
        if is_safe(&modified_level) {
            return true;
        }
    }

    false
}