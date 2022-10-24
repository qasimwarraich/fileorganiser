use chrono::{DateTime, Local};
use std::{fs, io, time};

fn main() {
    let path: String = "../test_dir".to_string();
    let spam = read_files(&path);
    println!("{}", spam.unwrap());
}

fn read_files(path: &String) -> Result<String, io::Error> {
    for file in fs::read_dir(path)? {
        let time = file?.metadata()?.accessed();

        let datetime = format_date(time?);
        println!("Time: {:?}", datetime.unwrap());
    }
    let s = String::from("Spam");
    return Ok(s);
}

fn format_date(timestamp: time::SystemTime) -> Result<String, io::Error> {
    let datetime: DateTime<Local> = timestamp.into();
    let datetime = format!("{}", datetime.format("%Y-%m-%d"));

    Ok(datetime)
}
