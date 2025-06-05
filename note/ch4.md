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


    
    


