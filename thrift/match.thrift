namespace cpp match

struct User {
    1: i32 id,
    2: string username,
    3: i32 score,
}

service Match{
    i32 add(1: User user),
    i32 remove(1: User user),
}
