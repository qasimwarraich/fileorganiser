use chrono::{DateTime, Local};
use std::{fs, io, time::SystemTime};

fn main() {
    let path = String::from("../test_dir");
    read_files(&path).expect("You supplied an invalid path");
}

fn read_files(path: &String) -> Result<String, io::Error> {
    let mut res = String::new();

    for file in fs::read_dir(path)? {
        let time = file?.metadata()?.accessed();

        let datetime = format_date(time?);
        res = datetime?;
        println!("Time: {:?}", res);
    }

    return Ok(res);
}

fn format_date(timestamp: SystemTime) -> Result<String, io::Error> {
    let datetime: DateTime<Local> = timestamp.into();
    let datetime = format!("{}", datetime.format("%Y-%m-%d"));

    return Ok(datetime);
}
