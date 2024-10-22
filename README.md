100.0% coverage

# Graph Algorithms
This repository contains graph algorithms made in golang by Lars M Bek and Ida M Jensen<br>

## Looking for a main.go?
This code is open-source which means you can use it in your own code. If you want to test the software you can test it from the test file.<br>
There is main.go's multiple places in the repository, one for each graph algorithm. You can find one at "./dijkstra/example/main.go" and compile a binary from there.<br>

## Information about this project
In this project we will develop both abstracted graph algorithms and concrete (implemented cases) for graph algorithms. <br>

### The repository contains the following algorthms<br>
* Dijkstra (currently only n^2 time complexity version)

### Future plans
We are looking into added the following:
* Dijkstra will get its priority queue (algorithm will become around n*log(n) time complexity , and a few optimizations will be added
* Weights will be able to compare with multiple variables, so that variables such as distance, time, openDoors etc. Will all be able to get a biased weight, so it can put percentage importance on each of them.
* AStar will be implemented
* Other graph algorithms (such as flow algorithms - Maximum flow problem) will be added 

### Information about the code:<br>
* The code's test coverage is 100%
* The code style is interface driven
  * The Go code follows an interface-driven design, where interfaces define behavior and decouple implementations, promoting flexibility, testability, and modularity. This approach ensures that dependencies are injected via interfaces, enabling easy swapping of implementations for testing or future scalability.
 
## Contributors
Lars M Bek <br>
Ida M Jensen
