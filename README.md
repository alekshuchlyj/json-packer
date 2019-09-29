# json-packer
Packs and unpacks nesting Json values

Input: 
{"a":"1", "b":{"c":"2", "d":{"e":"3"}, "g":{"h":"5"}}, "f":"4"}
Output:
b.g.h: 5
b.c: 2
b.d.e: 3
f: 4
a: 1   
