Q: Many databases use a log internally, which is an ____ data file. 
=> C. append-only

Q: An ____ is a type of data structure for efficiently finding a particular key in a database. It is a ____ structure derived from ____. 
=> B. index; additional; primary data 

Q: Storage engines fall into which two broad categories?
=> OLTP & OLAP

Q: True or False. OLTP systems are typically user-facing.
=> True. OLAP are typically for analysting data.

Q: The bottleneck for OLTP systems is usually ____. The bottleneck for OLAP is usually ____. 
=> A. disk seek time; disk bandwidth

Q: True or False. Both OLTP and OLAP systems typically must serve a huge volume of requsets.
=> False.

Exercise: Implement column compression. See bitmap_encoding.go/bitmap_encoding_test.
Will finish this in a near future 😬
