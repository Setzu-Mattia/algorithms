
# Credits:
# https://www.hackerearth.com/practice/notes/binary-indexed-tree-or-fenwick-tree/
# http://www.geeksforgeeks.org/binary-indexed-tree-or-fenwick-tree-2/
# http://citeseerx.ist.psu.edu/viewdoc/download;jsessionid=A6CAC51B4339EAF134D34B47FBCB1987?doi=10.1.1.14.8917&rep=rep1&type=pdf

def down(x):
    return x - (x & (-x))

def up(x):
    return x + (x & (-x))

n = 10 # input('Array size? ')
bit = [0] * (n + 1)

def update(i, deltaValue): # O(logi)
    assert i > 0 and i <= n
    while i <= n:
        bit[i] += deltaValue
        i = up(i)

def query(i): # O(logn)
    result = 0
    while i > 0:
        result += bit[i]
        i = down(i)
    return result

def sum(i, j): # O(2*logn)
    assert i > 0 and i <= j and j <= n
    return query(j) - query(i-1)

update(1, 10)
update(2, 2)
update(3, -5)
update(8, -5)
print sum(1, 4)
