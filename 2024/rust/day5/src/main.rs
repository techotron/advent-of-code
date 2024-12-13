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

    let part1_ans = part1(filename)?;
    println!("Part1: {}", part1_ans);

    let part2_ans = part2(filename)?;
    println!("Part2: {}", part2_ans);

    Ok(())
}

fn parse_input(reader: BufReader<File>) -> (Vec<Vec<i32>>, Vec<Vec<i32>>) {
    let mut rules = Vec::new();
    let mut updates= Vec::new();

    let re_rule = Regex::new(r"^(\d+)\|(\d+)$").unwrap();
    let re_update = Regex::new(r"^\d+(,\d+)*$").unwrap();

    for line in reader.lines() {
        let line = line.unwrap();
        let mut rule = Vec::new();
        let mut update = Vec::new();

        if let Some(caps) = re_rule.captures(&line) {
            let before = caps.get(1).unwrap().as_str().parse::<i32>().unwrap();
            let after = caps.get(2).unwrap().as_str().parse::<i32>().unwrap();
            rule.push(before);
            rule.push(after);
        }

        if re_update.is_match(&line) {
            for num in line.split(',') {
                update.push(num.parse::<i32>().unwrap());
            }
        }

        if rule.len() > 1 {
            rules.push(rule);
        }

        if update.len() > 0 {
            updates.push(update);
        }
    }

    (rules, updates)
}


fn part1(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);
    let mut updates_in_correct_order = Vec::new();
    let mut result = 0;

    let (rules, updates) = parse_input(reader);

    for update in updates.iter() {
        if is_correct_order(update, &rules) {
            updates_in_correct_order.push(update);
        }
    }

    for update in updates_in_correct_order {
        result += update[update.len() / 2];
    }
    Ok(result)
}

fn is_correct_order(update: &Vec<i32>, rules: &Vec<Vec<i32>>) -> bool {
    // let mut all_rules = Vec::new();
    for (idx, page) in update.iter().enumerate() {
        for rule in rules {
            let before = rule[0];
            let after = rule[1];

            if page == &after {
                // The first number of the rule is found, so split the update and check if the second number of the rule is found in the before part of the vec and return false if so
                let pages_before: Vec<i32> = Vec::from_iter(update[0..idx].iter().cloned());
                let pages_after: Vec<i32> = Vec::from_iter(update[idx + 1..update.len()].iter().cloned());

                for p in pages_before {
                    if p == after {
                        return false
                    }
                }

                for p in pages_after {
                    if p == before {
                        return false
                    }
                }
            }
        }
    }
    true 
}

fn part2(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);
    let mut fixed_updates= Vec::new();
    let mut result = 0;

    let (rules, updates) = parse_input(reader);

    for update in updates.iter() {
        if !is_correct_order(update, &rules) {
            fixed_updates.push(fix_incorrect_updates(update ,&rules));
        }
    }

    for update in fixed_updates {
        result += update[update.len() / 2];
    }
    Ok(result)
}

fn fix_incorrect_updates(update: &Vec<i32>, rules: &Vec<Vec<i32>>) -> Vec<i32> {
    let mut fixed_update = update.clone();
    let mut rule = &rules[0];
    let mut keep_running = true;
    let mut counter: usize = 0;

    while keep_running {
        let updated = fixed_update.clone();
        let mut before_idx: usize = 0;
        let mut after_idx: usize = 0;

        for (idx, page) in updated.iter().enumerate() {
            if *page == rule[0] {
                before_idx = idx
            }
            if *page == rule[1] {
                after_idx = idx
            }
        }

        if need_to_apply_rule(&fixed_update, &rule) {
            fixed_update.swap(before_idx, after_idx);
            counter = 0;
        } else {
            counter += 1;
        }
        if counter == rules.len() {
            keep_running = false;
        } else {
            rule = &rules[counter];
        }

    }
    fixed_update
}

fn need_to_apply_rule(update: &Vec<i32>, rule: &Vec<i32>) -> bool {
    for (idx, page) in update.iter().enumerate() {
        let before = rule[0];
        let after = rule[1];

        if page == &after {
            // The first number of the rule is found, so split the update and check if the second number of the rule is found in the before part of the vec and return false if so
            let pages_before: Vec<i32> = Vec::from_iter(update[0..idx].iter().cloned());
            let pages_after: Vec<i32> = Vec::from_iter(update[idx + 1..update.len()].iter().cloned());

            for p in pages_before {
                if p == after {
                    return true
                }
            }

            for p in pages_after {
                if p == before {
                    return true
                }
            }
        }
    }
    false 
}