use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};

#[get("/")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello, World!")
}

#[post("/bye")]
async fn bye() -> impl Responder {
    HttpResponse::Ok().body("Bye world!")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(hello)
            .service(bye)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
