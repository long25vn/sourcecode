create table posts with this form

    create table posts (
        id serial primary key not null,
        title text,
        alias text,
        intro_text text,
        full_text text,
        image text,
        published text,
        published_at timestamptz,
        categories text null,
        type text null,
        created_at timestamptz,
        created_by text null,
        modified_at timestamptz,
        modified_by text null,
        author_visible text null
    );


and command  ` go run main.go`