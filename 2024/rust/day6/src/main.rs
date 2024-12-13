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

    // let part2_ans = part2(filename)?;
    // println!("Part2: {}", part2_ans);

    Ok(())
}


fn part1(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line?;
        println!("{}", line);
    }

    Ok(0)
}


fn part2(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    for line in reader.lines() {
        let line = line?;
        println!("{}", line);
    }

    Ok(0)
}