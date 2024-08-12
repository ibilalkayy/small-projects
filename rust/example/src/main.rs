use actix_web::{web, App, HttpServer, HttpResponse, Responder};
use serde::Deserialize;
use std::fs;
use std::path::PathBuf;

#[derive(Deserialize)]
struct SubmitForm {
    action: String,
}

async fn redirect(form: web::Form<SubmitForm>) -> impl Responder {
    let (path, redirect_url) = match form.action.as_str() {
        "create" => (PathBuf::from("static/create.html"), "/create"),
        "update" => (PathBuf::from("static/update.html"), "/update"),
        _ => (PathBuf::from("static/error.html"), "/"),
    };

    let content = fs::read_to_string(path).unwrap_or_else(|_| "Error loading page.".into());
    
    HttpResponse::Ok()
        .content_type("text/html")
        .insert_header(("HX-Redirect", redirect_url))
        .body(content)
}

async fn serve_create() -> impl Responder {
    let path = PathBuf::from("static/create.html");
    let content = fs::read_to_string(path).unwrap_or_else(|_| "Error loading create page.".into());
    HttpResponse::Ok().content_type("text/html").body(content)
}

async fn serve_update() -> impl Responder {
    let path = PathBuf::from("static/update.html");
    let content = fs::read_to_string(path).unwrap_or_else(|_| "Error loading update page.".into());
    HttpResponse::Ok().content_type("text/html").body(content)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(|| async {
                let path = PathBuf::from("static/index.html");
                HttpResponse::Ok().body(fs::read_to_string(path).unwrap_or_else(|_| "Error loading index page.".into()))
            }))
            .route("/redirect", web::post().to(redirect))
            .route("/create", web::get().to(serve_create))
            .route("/update", web::get().to(serve_update))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}

