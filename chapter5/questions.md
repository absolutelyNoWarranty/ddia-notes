Q: True or False? Keeping data geographically close to users can be achieved through replication.

Q: True or False? Replication is easy, even if your data is constantly changing.

Q: A node that stores a copy of the database is called a
1. clone
2. replicant 
3. replica
4. shadow

Q: True or False? In leader-based replication, the leaders processes only writes.

Q: The leader sends a message to the follower and waits for a response from the follower
before reporting success to the client. Thus the replication to the follower is
1. asynchronous
2. synchronous

Q: True or False? In most database systems, enabling synchronous replication means that ALL followers are
synchronous.

Q: True or False? Asynchronous replication is less durable than synchronous replication.

Q: The process of handling a failure of the leader is called:
1. handover
2. handoff
3. failover
4. fail-through

Q: True or False? It's OK to have unnecessary failovers True or False? It's OK to have unnecessary failovers.

Q: What are some issues with statement-based replication (choose all that apply)?
1. statements that have side effects may result in different side effects on each replica
2. statement-based replication is not compact
3. statements that depend on data in the db have to be executed in exactly the same order on each replica
4. deterministic functions start generating non-deterministic values

Q: MySQL supports both row-based and statement-based replication. True or False? Statement-based replication
is the older way to do replication.

Q: What do we call the log used for replication which is decoupled from the storage engine internals?

Q: Sending the log off to an external application such as a data warehouse is called ____.

Q: True or False. Eventual consistency means there is no limit to how far behind a replica can be.

Q: What is read-your-writes consistency?
1. The guarantee that the writer's changes will be visible sometime.
2. The guarantee that the writer's changes will be immediately visible to all.
3. The guarantee that the writer will see their own changes immediately.

Q: What is a way to implement read-after-write consistency for edits to a user profile?
1. make the user's page only editable by the user, and always read from the leader when
they ask for it
2. allow any user to edit the profile page, and always read from the leader
3. read from the follower or leader

Q: Which is the stronger guarantee? Monotonic reads or eventual consistency?

Q: If writes happen in a certain order, then anyone reading those writes will see
them in the same order. What is this guarantee called?
1. consistent ordering
2. consistent prefix reads
3. write-then-read consistency
4. consistent suffix reads

TODO: Solutions for Replication Lag ...

