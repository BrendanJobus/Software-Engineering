import pytest
from ancestor import *

@pytest.fixture
def full_tree():
    return [10, 5, 13, 6, 4, 15, 2, 7, 35]

class TestAncestor:
    def test_emptyTree(fullTree):
        tree = Tree(None)
        result = findLCA(tree.root, 2, 6)
        assert result == -1

    def test_insert(fullTree):
        values =  [10, 5, 13, 6, 4, 15, 2, 7, 35]
        tree = Tree(None)

        for i in values:
            err = insertInTree(i, tree)
            if err != 0:
                assert False
        assert True

    def test_firstChildren(fullTree):
        values = [10, 5, 13]
        tree = Tree(None)

        for i in values:
            err = insertInTree(i, tree)

        result = findLCA(tree.root, 5, 13)
        assert result == 10

    def test_secondChildren(fullTree):
        values = [10, 5, 13, 6, 4, 15]
        tree = Tree(None)

        for i in values:
            err = insertInTree(i, tree)

        result = findLCA(tree.root, 6, 4)
        assert result == 5

        result = findLCA(tree.root, 6, 15)
        assert result == 10

    def test_betweenRows(fullTree):
        values = [10, 5, 13, 6, 4, 15, 2, 7, 35]
        tree = Tree(None)

        for i in values:
            err = insertInTree(i, tree)

        result = findLCA(tree.root, 2, 6)
        assert result == 5

        result = findLCA(tree.root, 2, 15)
        assert result == 10

    def test_callOnNonExistentNode(fullTree):
        values = [10, 5, 13, 6, 4, 15, 2, 7, 35]
        tree = Tree(None)

        for i in values:
            err = insertInTree(i, tree)

        result = findLCA(tree.root, 2, 9)
        assert result == -1

        result = findLCA(tree.root, 1, 6)
        assert result == -1

        result = findLCA(tree.root, 6, 39)
        assert result == -1

        result = findLCA(tree.root, 9, 39)
        assert result == -1
