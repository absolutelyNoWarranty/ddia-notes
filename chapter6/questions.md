Q: The main reason for wanting to partition data is
1. scalability
2. availability
3. durability

Q: True or false? Partitioning is usually combined with replication.

Q: True or false? A node may only store one partition.

Q: What is an example of the hot spot concept?
1. Traffic increases by a lot and all partitions have higher load.
2. A partition stops receiving any traffic.
3. A partition start receiving much more traffic than any other partition.
4. Some partitions receiving slightly more traffic than others.

Q: True or false? Partitioning by key range is a good way to prevent hot spots.

Q: After going through a hash function skewed data ____.
1. becomes normally distributed.
2. becomes uniformly distributed.
3. stays the same.

Q: In range partitioning each partition is assigned a range of ____. In hash partitioning,
each partition is assigned a range of ____.
1. keys; values
2. hashes; values
3. values; nodes
4. keys; hashes

Q: True or false. Hash partitioning allows us to do efficient range queries.

Q: True or false. With hash partitioning, we don't have to worry about hot spots.

Q: A document-partitioned index is also known as a ____.
1. local index
2. global index
3. tag index

Q: The scatter/gather approach to querying a partitioned databased is ____.
1. where you need to query all partitions and combine the results you get back.
2. where you query a random number of partitions and combine the results you get back.
3. where you query a single random partition

Q: True or false. In practice, updates to global secondary indexes are often asynchronous.

TODO: Partitioning Secondary Indexes by Term
