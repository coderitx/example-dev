

#[allow(dead_code)]
pub fn _function() {
    println!("this is a function");

    // 定义一个宏调用
    // 宏调用是一个表达式
    // x + 1 表达式的结尾没有分号。如果在表达式的结尾加上分号，它就变成了语句，而语句不会返回值
    let num = {
        let x = 10;
         x + 1
    };
    println!("The value of y is: {num}");
}


#[allow(dead_code)]
pub fn five(number: i32) -> i32{
    println!("this value is: {number}");

    number + 1
}