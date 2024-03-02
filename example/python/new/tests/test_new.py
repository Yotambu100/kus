from example.python.new.new import add, minus


def test_add():
    assert add(2, 3) == 5
    assert add(-1, 1) == 0
    assert add(0, 0) == 0


def test_minus():
    assert minus(5, -3) == 8
    assert minus(10, -10) == 20
    assert minus(-5, -5) == 0
