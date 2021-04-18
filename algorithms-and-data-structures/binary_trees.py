"""General classes of a binary tree"""
class Node:

    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None

"""Get max width of a layer in a binary tree"""
def getMaxWidth(root):
    maxWidth = 0
    h = height(root)
    # Get width of each level and compare to get the max
    for i in range(1, h+1):
        width = getWidth(root, i)

        if (width > maxWidth):
            maxWidth = width
        
        return maxWidth

def getWidth(root, level):
    """ Get width of a given level"""
    if root is None:
        return 0

    if level == 1:
        return 1

    return getWidth(root.left, level-1) + getWidth(root.right, level-1)

def height(node):
    if node is None:
        return 0

    lh = height(node.left)
    rh = height(node.right)

    return lh + 1 if lh > rh else rh + 1


"""Diameter of a binary tree"""
def diameter(root):

    if root is None:
        return 0

    lh = height(root.left)
    rh = height(root.right)

    ld = diameter(root.left)
    rd = diameter(root.right)

    return max(lh + rh + 1, max( ld, rd))

"""Check if a binary tree is a subtree of another binary tree"""
def contains(root1, root2):
    
    if root1 is None and root2 is None:
        return True

    if root1 is None or root2 is None:
        return False

    if root1.data != root2.data:
        return False

    lBranch = contains(root1.left, root2.left)
    rBranch = contains(root1.right, root2.right)

    return lBranch and rBranch


"""Traverse inorder"""
def inorder(root):

    if root is None:
        return
    inorder(root.left)
    print(root.data)
    inorder(root.right)

def preorder(root):

    if root is None:
        return
    print(root.data)
    preorder(root.left)
    preorder(root.right)

def postorder(root):

    if root is None:
        return
    lb = postorder(root.left)
    rb = postorder(root.right)

    print(lb)
    print(rb)

    return root.data

def order_successor(root, initial=True):
    if initial:
        initial=False
        order_succesor(root.right, initial)

    if node is None:
        return
    
    return min(order_succesor(root.left, initial), root.data)

