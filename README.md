# split

split slide into chucks using generics
## install

> go get github.com/Philiphil/split
## usage

```
package main
import (     
      split "[github.com/philiphil/split](http://github.com/Philiphil/split)"  
 ) 

func main() {     
    var i = []int{1, 2, 3, 4, 5}      
    si := split.SplitToChunks(i, 1)    
}
```
