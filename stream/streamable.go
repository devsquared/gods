package stream

import "github.com/devsquared/gods/collection"

type Streamable[T any] interface {
	// Spliterate splits and iterates over the implementing collection.
	Spliterate(collection.Collectioner[T])

	// Apply allows for application of operations on all stream elements.
	Apply(func(T) T)

	// Combine is the terminal operation of a stream that collects the elements into a collection.
	Combine(func() collection.Collectioner[T])
}

// May need to rethink this interface. This is what a stream does - spliterate, apply..., and then combine. But does it make
// sense for a collection to implement something like this? Think about applications like streaming a list of ints and
// sorting them. How does Apply fit in on the list? why should I be able to say list.Apply?

// Or should we have "Stream" objects that surround collections? If you use a ListStream, does that make sense? I think
// that this is what we want. We can define a Stream[List] that will allow us to do these things.

// Streamable needs to tell us how to spliterate, apply, and combine from a defined collection.
