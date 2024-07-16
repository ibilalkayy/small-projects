use leptos::*;

fn main() {
    leptos::mount_to_body(|| view! { < App /> });
}

#[component]
fn App() -> impl IntoView {
    let (count, set_count) = create_signal(0);
    let double_count = move || count() * 2;

    view! {
        <button
            on:click=move |_| { set_count.update(|n| * n += 1); }
            class:red=move || count() % 2 == 1
        >
            "Click Me: "
            {move || count()}
        </button>
        <br/>
        
        <p> "Count: " {count} </p>
        <ProgressBar progress=count/>

        <p> "Double Count: " {double_count} </p>
        <ProgressBar progress=double_count/>
    }
}

#[component]
fn ProgressBar(#[prop(default = 10)] max: u16, progress: impl Fn() -> i32 + 'static) -> impl IntoView {
    view ! {
        <progress
            max=max
            value=progress
        />
    }
}

