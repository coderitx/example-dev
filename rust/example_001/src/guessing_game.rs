// 导入用到的包到环境中
use std::io;
use rand::Rng;
use std::cmp::Ordering;

// 构建一个猜数字的游戏
#[allow(dead_code)]
pub fn guessing_game(){
    println!("Guess the number!");

    // 生成一个随机数
    let secret_number = rand::thread_rng().gen_range(1..=100);

    loop {
        println!("Please input your guess.");

        // 定义一个可变的字符串变量
        let mut guess = String::new();

        // 接收标准输入
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        // 将标准输入转换为uint32
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {guess}");
        // 比较输入的和生成的数字是否相等
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}

