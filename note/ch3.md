# CH3

1. Lox is dynamically typed 
2. managing memory: reference counting, and tracing garbage collector

## Data Types

1. Booleans
2. Integer
3. double-precision floating point
4. String
5. Nil (no value)

## Expressions

1. Arithmetic
    - Binary operation
    - Infix 
    - mixfix ? true:false
    - Both infix and prefix like '-'
    - comparison and euality
    - no implicit conversins
    - logical operators 
    - Precedence and grouping

## Statement

> "expression's main job is to produc a value "
> "Statement job is to produce an effect"

1. expression Statement use -> ; like "high";
2. Block -> scope {}

### Variable 

1. Using var Statement
    - default value is nil

### Control flow

1. if (condition) {}else{}
2. while(true){}
3. for(variable,condition,operation){}

### Functions

1. func(variables:parameter){}
    func(argument);
2. is dynamically typed , so fully specifies the function  
    include it's body
3. without return , return nil but interpreter
4. first class, mean real values -> func1(func2)(1,2);
5. function wrap another function
6. closures

### class
1. core concept instances and classes
1. class or prototypes
    1. static dispatch
    2. dynamic dispatch
2. prototype there are only objects no class ,  
    objects contain state and method
3. you can create class , like c++ without func keyword
4. also first class , var t = myClass;
5. this create new instance var n = myClass();
6. you create field, auto create field if it didn't exist
7. contructor , call init()
8. inheritance , by using < less than operator class a < b 
9. every method define in superclass also available its subclass
10. super keyword call super class init , init(){super.init()}


     



