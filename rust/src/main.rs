use chrono::{DateTime, Local};
use std::{fs, io, time::SystemTime};

fn main() {
    let path = String::from("../test_dir");
    organise_files(&path).expect("You supplied an invalid path");
}

fn organise_files(path: &String) -> Result<(), io::Error> {
    let dir_path = fs::canonicalize(path)?;

    for file in fs::read_dir(path)? {
        let file = file?;
        let file_path = dir_path.join(file.file_name());

        let metadata = file.metadata()?;
        let datetime = format_date(metadata.modified()?)?;

        if !metadata.is_dir() {
            fs::create_dir_all(dir_path.join(&datetime))?;
            fs::rename(&file_path, &dir_path.join(&datetime).join(file.file_name()))?;
        }
    }
    return Ok(());
}

fn format_date(timestamp: SystemTime) -> Result<String, io::Error> {
    let datetime: DateTime<Local> = timestamp.into();
    let datetime = format!("{}", datetime.format("%Y-%m-%d"));
    return Ok(datetime);
}
