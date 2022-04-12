// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::fmt;
/// various log levels
#[derive(Clone, PartialEq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
    Debug,
}

impl fmt::Display for LogLevel {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            LogLevel::Info => write!(f, "INFO"),
            LogLevel::Warning => write!(f, "WARNING"),
            LogLevel::Error => write!(f, "ERROR"),
            LogLevel::Debug => write!(f, "DEBUG"),
        }
    }
}

/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
    let mut l = level.to_string();
    match level {
        LogLevel::Info => format!("[{}]: {}", l, message),
        LogLevel::Warning => format!("[{}]: {}", l, message),
        LogLevel::Error => format!("[{}]: {}", l, message),
        LogLevel::Debug => format!("[{}]: {}", l, message),
    }
}
pub fn info(message: &str) -> String {
    format!("[INFO]: {}", message)
}
pub fn warn(message: &str) -> String {
    format!("[WARNING]: {}", message)
}
pub fn error(message: &str) -> String {
    format!("[ERROR]: {}", message)
}
