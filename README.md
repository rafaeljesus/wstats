When started, it will listen on port 5555 (but this may be configurable through a command-line flag).

Clients will be able to connect to this port and send arbitrary natural language over the wire.

The purpose of the application is to process the text, and store some stats about the different words that it sees.

The application will also expose an HTTP interface on port 8080 (configurable):

clients hitting the /stats endpoint will receive a JSON representation of the statistics about the words that the application has seen so far.

Specifically, the JSON response should look like:

```js
{
  "count": 42,
  "top_5_words": ["lorem", "ipsum", "dolor", "sit", "amet"],
  "top_5_letters": ["e", "t", "a", "o", "i"]
}
```

Where count represents the total number of words seen, top_5_words contains the 5 words that have been seen with the highest frequency,
and top_5_letters contains the 5 letters that have been seen with the highest frequency (you may choose to transform all letters to lowercase if you so wish).

A few things to look out for:

* The number of words to process may be large, although you may safely assume that they will fit within main memory.
* The application should support a high degree of concurrency, whereby many clients would be sending text at the same time.
* While we only expect three metrics over the collection of words, you should assume that a more fully-fledged version of the application would be collecting many more of these.
* We would like to see your approach to automated testing for this type of Go program.
