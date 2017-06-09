# search-mapper-and-reducer
distributed mapper and reducer model for searching

this project searches for keyword in files in certain directory on server.

the project is implementation for distributed systems so there is 3 mappers and a reducer.

each mapper search for the word in a part of the files in the directory.

after that the mapper save its result in a file saved in a certian place on server.

then the reducer collects the results from the files saved by the mappers.

the keyword , jobid and mapperid are passed as a parameter to the program.
