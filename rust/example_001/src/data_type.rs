/*
长度	有符号	无符号
8-bit	i8	u8
16-bit	i16	u16
32-bit	i32	u32
64-bit	i64	u64
128-bit	i128	u128
arch	isize	usize
*/

use std::io;

#[allow(dead_code)]
pub fn data_type(){
    // addition
    let sum = 5 + 10;
    println!("addition result: {sum}");

    // subtraction
    let difference = 95.5 - 4.3;
    println!("subtraction result: {difference}");

    // multiplication
    let product = 4 * 30;
    println!("multiplication result: {product}");

    // division
    let quotient = 56.7 / 32.2;
    let truncated = -5 / 3; // 结果为 -1
    println!("division result: {quotient} {truncated}");

    // remainder
    let remainder = 43 % 5;
    println!("remainder result: {remainder}");

    // tuple 类型
    let tup: (i32,f32,bool) = (88,88.8,false);
    // 元组解析
    let (x,y,z) = tup;
    println!("tup: {x} {y} {z}");

    // 元组索引访问
    let one: i32 = tup.0;
    let two: f32 = tup.1;
    let three: bool = tup.2;
    println!("tuple index result index=0: {one} index=1: {two} index=2: {three}");


    // 数组类型
    let arr = [1,2,4,5];
    let _arr1: [u32;2] = [100,200];
    // 长度为10，初始值为999
    let _arr2 = [999;10];

    let first = arr[0];
    let second = arr[1];
    println!("first: {first} second: {second}")
}

#[allow(dead_code)]
pub fn array_idnex(){
    let _arr = [1,2,4,5,6];

    println!("Please enter an array index.");

    let mut index = String::new();

    io::stdin().read_line(&mut index).expect("Failed to read line");
    let index: usize = index.trim().parse().expect("Index entered was not a number");
    
    let element = _arr[index];

    println!("this value of element at index {index} is : {element}");
}