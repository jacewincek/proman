use axum::{
    routing::get,
    response::IntoResponse,
    Router,
    Json,
    extract::{Query, Path, State},
};
use serde::{Deserialize, Serialize};
use sqlx::SqlitePool;
use std::env;

#[tokio::main]
async fn main() {
    let pool = SqlitePool::connect(&env::var("DATABASE_URL").unwrap()).await.unwrap();

    let app = Router::new()
        .route("/projects", get(projects_index).post(projects_create))
        .route("/projects/:id", get(projects_view).patch(projects_update).delete(projects_delete))
        .with_state(pool);

    axum::Server::bind(&"127.0.0.1:8080".parse().unwrap())
        .serve(app.into_make_service())
        .await
        .unwrap();

}

async fn projects_index(
    State(pool): State<SqlitePool>,
) -> impl IntoResponse {

}

async fn projects_create(
    State(pool): State<SqlitePool>,
) -> impl IntoResponse {
    
}

async fn projects_view(
    State(pool): State<SqlitePool>,
) -> impl IntoResponse {
    
}

async fn projects_update(
    State(pool): State<SqlitePool>,
) -> impl IntoResponse {
    
}

async fn projects_delete(
    State(pool): State<SqlitePool>,
) -> impl IntoResponse {
    
}
