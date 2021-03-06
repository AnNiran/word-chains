/*
package build contains functionality for building a word chain using starting and ending word (boundary)
using a unidirectional graph data structure

implemented as per the technical specification at:
http://codekata.com/kata/kata19-word-chains/

The path is built based Dijkstra's algorithm,
but with modifications since Dijkstra's algorithm involves measuring 
weights of the different edges, and in the case we have all edges equal 
=> this leads to the solution of using a breadth-first search

Also, the graph is not preliminary built ahead of time.
It is instead built dynamically when each node is visited - then on the fly the next words 
are generated

Functionality works in both directions and uses a real dictionary content in
cmd/dictionary/wordlist.txt with over 45 000 words
*/

package build