# Golang Collections [![CI][1]][1] [![Go Report Card][2]][2] [![codecov.io][3]][4] [![GoDoc][5]][6]

[1]: https://github.com/billryan/collections/workflows/CI/badge.svg
[2]: https://goreportcard.com/badge/github.com/billryan/collections
[3]: https://codecov.io/github/billryan/collections/coverage.svg?branch=master "Coverage badge"
[4]: https://codecov.io/github/billryan/collections?branch=master "Codecov Status"
[5]: https://godoc.org/github.com/billryan/collections?status.svg "GoDoc badge"
[6]: https://godoc.org/github.com/billryan/collections "GoDoc"

Maps and slices go a long way in Go, but sometimes you need more. This is a collection of collections that may be useful.

## Queue
A [queue](https://en.wikipedia.org/wiki/Queue_\(data_structure\)) is a first-in first-out data structure.

## Set
A [set](https://en.wikipedia.org/wiki/Set_\(computer_science\)) is an unordered collection of unique values typically used for testing membership.

## Skip list
A [skip list](https://en.wikipedia.org/wiki/Skip_list) is a data structure that stores nodes in a hierarchy of linked lists. It gives performance similar to binary search trees by using a random number of forward links to skip parts of the list.

## Splay Tree

A [splay tree](https://en.wikipedia.org/wiki/Splay_tree) is a type of binary search tree where every access to the tree results in the tree being rearranged so that the current node gets put on top.

## Stack
A [stack](https://en.wikipedia.org/wiki/Stack_\(abstract_data_type\)) is a last-in last-out data structure.

## Trie
A [trie](http://en.wikipedia.org/wiki/Trie) is a type of tree where each node represents one byte of a key.

## Ternary Search Tree

A [ternary search tree](http://en.wikipedia.org/wiki/Ternary_search_tree) is similar to a trie in that nodes store the letters of the key, but instead of either using a list or hash at each node a binary tree is used. Ternary search trees have the performance benefits of a trie without the usual memory costs.