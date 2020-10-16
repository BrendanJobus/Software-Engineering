from dataclasses import dataclass

# Binary tree node

@dataclass
class Node:
    key: int
    left: 'Node'
    right: 'Node'

@dataclass
class Tree:
    root: 'Node'

def main():
    values = [10, 5, 13, 6, 4, 15, 2, 7, 35]

    tree = Tree(None)

    for i in values:
        err = insertInTree(i, tree)
        if err != 0:
            print("Error inserting value \'{}\': {}".format(i, err))

    x = findLCA(tree.root, 2, 7)
    print(x)

"""
            10
      5             13
    4   6               15
  2       7                 35
"""

def findPath( root, path, k ):
    # Base case
    if root is None:
        return False

    path.append(root.key)

    if root.key == k :
        return True

    if (( root.left != None and findPath(root.left, path, k)) or (root.right != None and findPath(root.right, path, k))):
        return True

    path.pop()
    return False

def findLCA(root, n1, n2):
    path1 = []
    path2 = []

    if (not findPath(root, path1, n1) or not findPath(root, path2, n2)) :
        return -1

    i = 0
    while (i < len(path1) and i < len(path2)):
        if path1[i] != path2[i]:
            break
        i += 1
    if i > 0:
        return path1[i - 1]
    return path1[i]

def insertInTree(value, t):
    if t.root is None:
        t.root = Node(value, None, None)
        return 0

    return insert(value, t.root)

def insert(value, n):
    if n is None:
        return 1

    if value is n.key:
        return 0
    elif value < n.key:
        if n.left is None:
            n.left = Node(value, None, None)
            return 0
        return insert(value,n.left)
    elif value > n.key:
        if n.right is None:
            n.right = Node(value, None, None)
            return 0
        return insert(value, n.right)
    return 0

if __name__ == "__main__":
    main()
