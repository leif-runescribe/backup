use std::sync::{Arc, Mutex};
use std::collections::{HashSet, VecDeque};
use tokio;
use reqwest;
use scraper::{Html, Selector};
use url:Url;
use thiserror::Error;

#[derive(Error, Debug)]
enum CrawlerError {
    #[error("Request error: {0}")]
    RequestError(#[from] reqwest::Error),
    #[error("URL parse error: {0}")]
    UrlParseError(#[from] url::ParseError),
    #[error("Invalid URL: {0}")]
    InvalidUrl(String),
}