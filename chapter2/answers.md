Q:The best known data-model today is probably SQL, which is based on the ____ model.
=> B. relational

Q:True or False? PostgreSQL is a relational db. It is impossible for it to support storing JSON documents directly.
=> False. It also support storing JSON values.

Q:SQL is a ____ query language.
=> A. declarative

Q:The awkward disconnect between the model of how the data is stored in a database and the data model the application 
code uses is sometimes called an ____ ____.
=> C. impedance mismatch

Q:True or False. It is inappropriate to use JSON to represent self-contained documents like a person's resume.
=> False.

Q:True or False. In terms of locality, a multiple-table representation is better than a JSON representation.
=> False. All data sould be stored in a JSON object, you don't have to query multiple times.

Q:What are some advantages of using the id of a string in a standarized list instead of free-text?
=> 1. Decreases duplication
2. Is meaningful to humans
=> 3. Better search 
=> 4. Easier to update
5. Reduces the amount of joining we have to do.

Q:True or False. The relational model cannot handle many-to-many relationships.
=> False.

Q:True or False. Document and graph databases do not enforce a schema on write. So they
have no schema.
=> False. Data will be checked with schema when reading.
