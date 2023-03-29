package stream

import "github.com/devsquared/gods/collection"

// A lot of scratch work in this file.

type Streamable[T any] interface {
	// Spliterate splits and iterates over the implementing collection.
	//
	// NOTES: how are we going to split collections and iterate through these? channels? how many?
	// are we returning a single chan? is this receiving from multiple channels?
	Spliterate(collection.Collectioner[T]) chan T

	// Apply allows for application of operations on all stream elements.
	//
	// NOTES: should we be accepting a channel here and returning a channel?
	// or should we just accept and return the elements themselves?
	Apply(func(chan T) chan T)

	// Combine is the terminal operation of a stream that collects the elements into a collection.
	Combine(func(chan T) collection.Collectioner[T])
}

// May need to rethink this interface. This is what a stream does - spliterate, apply..., and then combine. But does it make
// sense for a collection to implement something like this? Think about applications like streaming a list of ints and
// sorting them. How does Apply fit in on the list? why should I be able to say list.Apply?

// Or should we have "Stream" objects that surround collections? If you use a ListStream, does that make sense? I think
// that this is what we want. We can define a Stream[List] that will allow us to do these things.

// Streamable needs to tell us how to spliterate, apply, and combine from a defined collection. In go, we have channels
// which can act as the pipeline portion of streams.

// When we get to testing this, we need to understand what else could implement the interface. Is that acceptable?
