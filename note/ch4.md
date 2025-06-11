# Ch4 Scanning

1. "The scanner takes in raw source code as a series  
    of characters and group it into a series of chunk  
    we call **token**"
2. lexical analysis is: scan through the list of character  
    group them together into samllest sequences that still  
    represent something. Each blobs of characteris called  
    lexeme var lang = "lox"; ->
        1. var
        2. lang
        3. =
        4. "lox"
        5. ;
3. The core of the scanner is a loop Starting at the first  
    character of the source cod, the scanner figure out what  
    lexeme the character belong, and following character,
    when it reach enf of character it emit token, and so on 
4. reserve word eg: keyword are identifier , but we need seperate  
    between user identifier , keyword, so Token conain  
    1. KeyWordType
    2. identifier
5. about semicolon -> implicit semicolon  
    so case like go [Effective go](https://go.dev/doc/effective_go#semicolons)
    "The rule is this. If the last token before a newline is an identifier  
    (which includes words like int and float64),  
    a basic literal such as a number or string constant, or one of the tokens "
    > break continue fallthrough return ++ -- ) }
    so yes
    ```go
    if(){
    }
    ```
    so no, because ) it trigger identifier
    ```go
    if()
    {
    }
    ```
    so probally just design no implicit semicolon
