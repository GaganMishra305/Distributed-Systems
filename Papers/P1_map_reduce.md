# Map Reduce
Paper: https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf

## 1. What?
```
Input1  ->  map() -> a1   b1   

Input2  ->  map() ->      b2   

Input3  ->  map() -> a3        c3
                      
                    reduce(a, 2)
                    reduce(b, 2)
                    reduce(c, 1)
```

## 2. Map()
```
map(k, v):
    split v into words
    for each word w
        emit(w, 1)
```

## 3. Reduce()
```
reduce(k, v):
    emit(len(v))
```

## 4. Example: 
count all occurences of all words in a large document.
```
map(String key, String value):
    // key: document name
    // value: document contents
    for each word w in value:
    EmitIntermediate(w, "1");

reduce(String key, Iterator values):
    // key: a word
    // values: a list of counts
    int result = 0;
    for each v in values:
    result += ParseInt(v);
    Emit(AsString(result));
```
