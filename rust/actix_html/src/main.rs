use actix_files::Files;
use actix_web::{middleware::Logger, App, HttpServer};

fn index() -> impl actix_web::dev::HttpServiceFactory {
    Files::new("/", "./static/").index_file("index.html")
}

fn create() -> impl actix_web::dev::HttpServiceFactory {
    Files::new("/create", "./static/").index_file("create.html")
}

fn view() -> impl actix_web::dev::HttpServiceFactory {
    Files::new("/view", "./static/").index_file("view.html")
}

fn update() -> impl actix_web::dev::HttpServiceFactory {
    Files::new("/update", "./static/").index_file("update.html")
}

fn delete() -> impl actix_web::dev::HttpServiceFactory {
    Files::new("/delete", "./static/").index_file("delete.html")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Starting the HTTP server at port: 8080");

    HttpServer::new(|| {
        App::new()
            .service(index())
            .service(create())
            .service(view())
            .service(update())
            .service(delete())
            .wrap(Logger::default())
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
