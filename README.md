# match

Simple program to match two strings... 

usage:

./match pattern string

Match will check to see if the pattern matches the string using a few different Go libray 
calls:

regexp:
   Match()
   
path/filepath
   MatchString()
   
If there is a problem (error) in the patter, ./match will print the error... 
