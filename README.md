# D2iQ Take Home Challenge

## Cache with eviction policy

**Problem Statement**: Create an in-memory cache in Go for storing String values based on a String key. Expect that the cache will continually have previously unseen keys added to it, and that some keys will be fetched more frequently than others.

Your cache should have the following signature (in pseudocode):

```
type Cache {
    // Inserts the provided key/value pair into the cache, making it
    // available for future get() calls.
    void put(String key, String value);
    
    // Returns a value previously provided via put(). An empty Option
    // may be returned if the requested data was never inserted or is
    // no longer available.
    Option[String] get(String key);
}
```

The choice of eviction policy is up to you, but you could consider least recently used or time to live policies, a combination of the two, or any other eviction policy you choose.

The code should at least contain:

- A `README.md` explaining how to build the project and how to use the library (and any other details you think are important, e.g. limitations)
- A `Makefile`
- Tests 🙂

Good luck! 🎉
