Q:True or False. A change to the schema affects the data model so we never have to change the application code.

Q:Backward compatibility is when ____ code can read data that was written by ____ older code. Forward compatibility is when ____ code can read data that was written by newer code.
A. older; newer; newer; older
B. newer; older; older; newer

Q:True or False. The sequence-of-bytes representation used to write data to a file or send it over the network is usually quite similar to the in-memory representation.

Q: ____ is the process of translating from the in-memory representation to a byte sequence. ____ is the reverse.
A. encoding; deconding
B. decoding; encoding

Q:True or False. Python has a built-in encoding called pickle. We should use it for long-term data storage.

Q:True or False. What are some disadvantages of using built-in encodings?
1. bad security`
2. they are often not designed to be efficient
3. bad portability, you must commit to using the same language
4. usually not designed with backwards and forwards compatibility in mind

Q: Which of the following are true about JSON?
1. JSON distinguishes between strings and numbers.
2. JSON distinguisheds between integers and floating-point numbers
3. JSON supports Unicode character strings
4. JSON supports binary strings
5. Many JSON tools don't bother using the optional JSON schema support.

Q: What are some examples of binary encoding libraries?
1. Apache Thrift
2. JSON
3. Protocol Buffers
4. XML
5. Apache Avro
6. CSV

Q: Match the scenarios with their definitions.
A. Message-Passing Dataflow
B. Dataflow Through Services
C. Dataflow Through Databases

1. Sending a message to be read anytime in the future.
2. Sending a message and expecting a response ASAP.
3. Sending a message asynchronously and with low-latency.

Q: True or False. A distributed actor framework essentially integrates
a message broker and the actor programming model into a single framework.

Q: True or False. In the actor framework, message delivery is guaranteed.

Q: The ____ model tries to make a request to a remote service look the same as
calling a function or method in your programming language.
A. actor
B. message-passing
C. RPC
D. REST

Q: What are some features of the REST design philosophy?
1. using URLs for identifying actions
2. using XML features for cache control
3. emphasizes describing the API using a Description Language to support code generation
4. complex data formats

Q: What are some considerations when designing RPC APIs.
1. timeouts
2. idempotence
3. variable execution times
4. encoding objects
