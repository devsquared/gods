# Gods - Go Data Structures

For experimentation and learning, I wanted to start implementing data structures in go.
This repo houses those structures and are open to use and changes. 

## Collections
Collection is any simple group of data.
### List
- List is an unordered collection of data. It is backed by a simple golang slice. Basic add here will append the data.

## Heap
This repo contains ["array" implementation of heaps](https://www.geeksforgeeks.org/array-representation-of-binary-heap/). 
- [Max Heap](https://www.digitalocean.com/community/tutorials/max-heap-java)
  - The [max heap](https://github.com/devsquared/gods/blob/main/heap/max_heap.go) is a complete binary tree that has the max nodes at the top. This is nice for when you want the popped value to be the highest in the tree.

## Queue
- [Ring Queue](https://en.wikipedia.org/wiki/Circular_buffer) or ring buffer 
  - This implementation is quick and cheap in regard to performance and memory. The ring queue here utilizes [bit masking](https://www.scaler.com/topics/data-structures/bit-masking/) and some bitwise magic to speed things up.
- [Priority Queue](https://www.programiz.com/dsa/priority-queue)
  - Backed by our max heap, this priority queue allows for quickly popping off the highest priority element in the queue. 

## TODO
- [ ] Update README with outline of what is in the repo. Add outline as you add structures.
- [ ] Collections
  - [ ] List
  - [ ] Set
  - [ ] Map