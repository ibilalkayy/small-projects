use leptos::*;
use leptos_router::*;

pub fn ParentChildHome() -> impl IntoView {
    view! {
        <div>
            <h1> "Parent Child Home" </h1>
            <Outlet/>
        </div>
    }
}
