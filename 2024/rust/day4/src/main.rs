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

    let mut result: i32 = 0; 

    let mut input_as_vec = Vec::new();

    for line in reader.lines() {
        let line = line?;

        // Recreate input as a new dimension 
        let mut input_as_vec_line = Vec::new();

        let char_vec: Vec<char> = line.chars().collect();

        for c in char_vec {
            input_as_vec_line.push(c)
        }

        input_as_vec.push(input_as_vec_line);
    }

    for (y, line) in input_as_vec.iter().enumerate() {
        for (x, _) in line.iter().enumerate() {
            result += count_xmas_horizontal(x, y, &input_as_vec);
            result += count_xmas_vertical(x, y, &input_as_vec);
            result += count_xmas_diagonal(x, y, &input_as_vec);
        }
    }

    Ok(result)
}


fn count_xmas_horizontal(x: usize, y: usize, input: &Vec<Vec<char>>) -> i32 {
    // Find forwards and backwards horizontal
    // right: x+1
    // left: x-1
    let max = input[0].len();
    let mut result: i32 = 0;

    // Don't try to access afer the last item in the vec
    if x < max - 3 {
        // Left to right
        if input[y][x] == 'X' && input[y][x+1] == 'M' && input[y][x+2] == 'A' && input[y][x+3] == 'S' {
            result += 1;
        }
    }

    // Don't try to access an item before the vec
    if x > 2 {
        // Right to left
        if input[y][x] == 'X' && input[y][x-1] == 'M' && input[y][x-2] == 'A' && input[y][x-3] == 'S' {
            result += 1;
        }
    }
    
    result
}

fn count_xmas_vertical(x: usize, y: usize, input: &Vec<Vec<char>>) -> i32 {
    // Find forwards and backwards vertical 
    // down: y+1
    // up: y-1
    let max = input.len();
    let mut result: i32 = 0;
    
    // Don't try to access after the last line of the input
    if y < max - 3 {
        // Downwards
        if input[y][x] == 'X' && input[y+1][x] == 'M' && input[y+2][x] == 'A' && input[y+3][x] == 'S' {
            result += 1;
        }
    }


    // Don't try to access before the first line of the input
    if y > 2 {
        // Upwards
        if input[y][x] == 'X' && input[y-1][x] == 'M' && input[y-2][x] == 'A' && input[y-3][x] == 'S' {
            result += 1; 
        }
    }

    result 
}


fn count_xmas_diagonal(x: usize, y: usize, input: &Vec<Vec<char>>) -> i32 {
    // Find diagonal
    // diag right down: x+1 y+1
    // diag right up: x+1 y-1
    // diag left down: x-1 y+1
    // diag left up: x-1 y-1
    let max_x = input[0].len();
    let max_y = input.len();
    let mut result: i32 = 0;

    // Diag right down
    if x < max_x - 3 && y < max_y - 3 {
        if input[y][x] == 'X' && input[y+1][x+1] == 'M' && input[y+2][x+2] == 'A' && input[y+3][x+3] == 'S' {
            result += 1;
        }
    } 

    // Diag right up
    if x < max_x - 3 && y > 2 {
        if input[y][x] == 'X' && input[y-1][x+1] == 'M' && input[y-2][x+2] == 'A' && input[y-3][x+3] == 'S' {
            result += 1;
        }
    } 

    // Diag left down 
    if x > 2 && y < max_y - 3 {
        if input[y][x] == 'X' && input[y+1][x-1] == 'M' && input[y+2][x-2] == 'A' && input[y+3][x-3] == 'S' {
            result += 1;
        }
    } 

    // Diag left up
    if x > 2 && y > 2 {
        if input[y][x] == 'X' && input[y-1][x-1] == 'M' && input[y-2][x-2] == 'A' && input[y-3][x-3] == 'S' {
            result += 1;
        }
    } 

    result
}



fn part2(filename: &str) -> Result<i32, Error> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);

    let mut result: i32 = 0; 

    let mut input_as_vec = Vec::new();

    for line in reader.lines() {
        let line = line?;

        // Recreate input as a new dimension 
        let mut input_as_vec_line = Vec::new();

        let char_vec: Vec<char> = line.chars().collect();

        for c in char_vec {
            input_as_vec_line.push(c)
        }

        input_as_vec.push(input_as_vec_line);
    }

    for (y, line) in input_as_vec.iter().enumerate() {
        for (x, character) in line.iter().enumerate() {
            if character.to_string() == "A" {
                result += count_mas_diagonal(x, y, &input_as_vec);
            }
        }
    }

    Ok(result)
}


fn count_mas_diagonal(x: usize, y: usize, input: &Vec<Vec<char>>) -> i32 {
    // Find diagonal
    // diag right down: x+1 y+1
    // diag right up: x+1 y-1
    // diag left down: x-1 y+1
    // diag left up: x-1 y-1
    let max_x = input[0].len();
    let max_y = input.len();
    let mut result: i32 = 0;


    if y > 0 && y < max_y - 1 && x > 0 && x < max_x - 1 {
        // M.S
        // .A.
        // M.S
        if input[y-1][x-1] == 'M' && input[y+1][x-1] == 'M' && input[y-1][x+1] == 'S' && input[y+1][x+1] == 'S' {
            result += 1;
        }
        
        // M.M
        // .A.
        // S.S
        if input[y-1][x-1] == 'M' && input[y+1][x-1] == 'S' && input[y-1][x+1] == 'M' && input[y+1][x+1] == 'S' {
            result += 1;
        }

        // S.M
        // .A.
        // S.M
        if input[y-1][x-1] == 'S' && input[y+1][x-1] == 'S' && input[y-1][x+1] == 'M' && input[y+1][x+1] == 'M' {
            result += 1;
        }

        // S.S
        // .A.
        // M.M
        if input[y-1][x-1] == 'S' && input[y+1][x-1] == 'M' && input[y-1][x+1] == 'S' && input[y+1][x+1] == 'M' {
            result += 1;
        }
    }

    result
}