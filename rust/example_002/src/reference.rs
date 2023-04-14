

#[allow(dead_code)]
pub fn param_reference(){
    let s1 = String::from("rust");
    let size = calculate_length(&s1);
    println!("s1 = {s1} size = {size}");
}

fn calculate_length(s: &String) -> usize{
    s.len()
}