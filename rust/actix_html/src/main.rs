use actix_files::Files;
use actix_web::{middleware::Logger, App, HttpServer};

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    
    println!("Starting the HTTP server at port: 8080");

    HttpServer::new(|| {
        App::new()
            .service(Files::new("/", "./static/").index_file("index.html"))
            .wrap(Logger::default())
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
