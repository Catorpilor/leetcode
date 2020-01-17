package uf

type UnionFind interface {
    Union(i, j int)
    Find(i, j int) bool
}
