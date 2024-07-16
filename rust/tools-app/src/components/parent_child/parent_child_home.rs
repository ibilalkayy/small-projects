use leptos::*;
use leptos_router::*;

pub fn ParentChildHome() -> impl IntoView {
    view! {
        <div>
            <h1> "Parent Child Home" </h1>
            <ul>
                <li>
                    <a href="/parent-child/write-signal"> "Write Signal" </a>
                </li>
            </ul>
            <Outlet/>
        </div>
    }
}
