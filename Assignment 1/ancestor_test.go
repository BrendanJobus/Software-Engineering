package main

import "testing"

func TestEmptyTree(t *testing.T) {
  tree := &Tree{}

  emptyTree := findLCA(tree.Root, 2, 5)
  if emptyTree != -1 {
    t.Errorf("LCA failed, function does not work for empty tree")
  } else {
    t.Logf("LCA passed")
  }
}


func TestInsert(t *testing.T) {
  values := []int{10, 5, 13, 6, 4, 15, 2, 7, 35}					//				  10
                                                          //			  5	  13
	tree := &Tree{}													                //		  4  6	 15
	for i := 0; i < len(values); i++ {								      //		 2    7 	 35
		err := tree.Insert(values[i])
		if err != nil {
			t.Errorf("Error inserting value %v", values[i])
		} else {
      t.Logf("Value inserted '%v'", values[i])
    }
	}
}

func TestFirstChildren(t *testing.T) {
  values := []int{10, 5, 13}

  tree := &Tree{}
  for i := 0; i < len(values); i++ {
    tree.Insert(values[i])
  }

  result := findLCA(tree.Root, 5, 13)
  if result != 10 {
    t.Errorf("findLCA(tree.Root, 5, 13) failed, expected %v, got %v", 10, result)
  } else {
    t.Logf("findLCA(tree.Root, 5, 13) success, expected %v, got %v", 10, result)
  }
}

func TestSecondChildren(t *testing.T) {
  values := []int{10, 5, 13, 6, 4, 15}

  tree := &Tree{}
  for i := 0; i < len(values); i++ {
    tree.Insert(values[i])
  }

  result := findLCA(tree.Root, 6, 4)
  if result != 5 {
    t.Errorf("findLCA(tree.Root, 6, 4) failed, expected %v, got %v", 5, result)
  } else {
    t.Logf("findLCA(tree.Root, 6, 4) success, expected %v, got %v", 5, result)
  }

  result = findLCA(tree.Root, 6, 15)
  if result != 10 {
    t.Errorf("findLCA(tree.Root, 6, 15) failed, expected %v, got %v", 10, result)
  } else {
    t.Logf("findLCA(tree.Root, 6, 15) success, expected %v, got %v", 10, result)
  }

  result = findLCA(tree.Root, 4, 15)
  if result != 10 {
    t.Errorf("findLCA(tree.Root, 4, 15) failed, expected %v, got %v", 10, result)
  } else {
    t.Logf("findLCA(tree.Root, 4, 15) success, expected %v, got %v", 10, result)
  }
}

func TestBetweenRows(t *testing.T) {
  values := []int{10, 5, 13, 6, 4, 15, 2, 7, 35}

  tree := &Tree{}
  for i := 0; i < len(values); i++ {
    tree.Insert(values[i])
  }

  result := findLCA(tree.Root, 2, 6)
  if result != 5 {
    t.Errorf("findLCA(tree.Root, 4, 6) failed, expected %v, got %v", 5, result)
  } else {
    t.Logf("findLCA(tree.Root, 2, 6) success, expected %v, got %v", 5, result)
  }

  result = findLCA(tree.Root, 2, 15)
  if result != 10 {
    t.Errorf("findLCA(tree.Root, 2, 15) failed, expected %v, got %v", 10, result)
  } else {
    t.Logf("findLCA(tree.Root, 2, 15) success, expected %v, got %v", 10, result)
  }
}

func TestCallOnNonExistentNode(t *testing.T) {
  values := []int{10, 5, 13, 6, 4, 15, 2, 7, 35}

  tree := &Tree{}
  for i := 0; i < len(values); i++ {
    tree.Insert(values[i])
  }

  result := findLCA(tree.Root, 2, 9)
  if result != -1 {
    t.Errorf("findLCA(2, 9) on non-existent 9 failed, expected %v, got %v", -1, result)
  } else {
    t.Logf("findLCA(2, 9) on non-existent 9 success, expected %v, got %v", -1, result)
  }

  result = findLCA(tree.Root, 1, 6)
  if result != -1 {
    t.Errorf("findLCA(1, 6) on non-existent 1 failed, expected %v, got %v", -1, result)
  } else {
    t.Logf("findLCA(1, 6) on non-existent 1 success, expected %v, got %v", -1, result)
  }

  result = findLCA(tree.Root, 6, 39)
  if result != -1 {
    t.Errorf("findLCA(6, 39) on non-existent 39 failed, expected %v, got %v", -1, result)
  } else {
    t.Logf("findLCA(6, 39) on non-existent 39 success, expected %v, got %v", -1, result)
  }

  result = findLCA(tree.Root, 9, 39)
  if result != -1 {
    t.Errorf("findLCA(9, 39) on non-existent 9 and non-existent 39 failed, expected %v, got %v", -1, result)
  } else {
    t.Logf("findLCA(9, 39) on non-existent 9 and non-existent 39 success, expected %v, got %v", -1, result)
  }
}
